package metronome_test

import (
	"context"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/option"
)

const (
	pathTestBaseURL = "https://proxy.example/metronome/"
	pathTestUUID    = "13117714-3f05-48e5-a6e9-a66093f13b4d"
	pathTestEnum    = "stripe"
)

// Expected requests for valid path parameters, with {id} standing for
// pathTestUUID and the base URL removed.
var pathTestFixtures = map[string]string{
	"V1BillableMetricService.Get":                          "GET v1/billable-metrics/{id}",
	"V1CustomerBillingConfigService.Delete":                "DELETE v1/customers/{id}/billing-config/stripe",
	"V1CustomerBillingConfigService.Get":                   "GET v1/customers/{id}/billing-config/stripe",
	"V1CustomerBillingConfigService.New":                   "POST v1/customers/{id}/billing-config/stripe",
	"V1CustomerInvoiceService.AddCharge":                   "POST v1/customers/{id}/addCharge",
	"V1CustomerInvoiceService.Get":                         "GET v1/customers/{id}/invoices/{id}",
	"V1CustomerInvoiceService.GetPdf":                      "GET v1/customers/{id}/invoices/{id}/pdf",
	"V1CustomerInvoiceService.List":                        "GET v1/customers/{id}/invoices",
	"V1CustomerInvoiceService.ListAutoPaging":              "GET v1/customers/{id}/invoices",
	"V1CustomerInvoiceService.ListBreakdowns":              "GET v1/customers/{id}/invoices/breakdowns?ending_before=0001-01-01T00%3A00%3A00Z&starting_on=0001-01-01T00%3A00%3A00Z",
	"V1CustomerInvoiceService.ListBreakdownsAutoPaging":    "GET v1/customers/{id}/invoices/breakdowns?ending_before=0001-01-01T00%3A00%3A00Z&starting_on=0001-01-01T00%3A00%3A00Z",
	"V1CustomerPlanService.Add":                            "POST v1/customers/{id}/plans/add",
	"V1CustomerPlanService.End":                            "POST v1/customers/{id}/plans/{id}/end",
	"V1CustomerPlanService.List":                           "GET v1/customers/{id}/plans",
	"V1CustomerPlanService.ListAutoPaging":                 "GET v1/customers/{id}/plans",
	"V1CustomerPlanService.ListPriceAdjustments":           "GET v1/customers/{id}/plans/{id}/priceAdjustments",
	"V1CustomerPlanService.ListPriceAdjustmentsAutoPaging": "GET v1/customers/{id}/plans/{id}/priceAdjustments",
	"V1CustomerService.Get":                                "GET v1/customers/{id}",
	"V1CustomerService.ListBillableMetrics":                "GET v1/customers/{id}/billable-metrics",
	"V1CustomerService.ListBillableMetricsAutoPaging":      "GET v1/customers/{id}/billable-metrics",
	"V1CustomerService.ListCosts":                          "GET v1/customers/{id}/costs?ending_before=0001-01-01T00%3A00%3A00Z&starting_on=0001-01-01T00%3A00%3A00Z",
	"V1CustomerService.ListCostsAutoPaging":                "GET v1/customers/{id}/costs?ending_before=0001-01-01T00%3A00%3A00Z&starting_on=0001-01-01T00%3A00%3A00Z",
	"V1CustomerService.PreviewEvents":                      "POST v1/customers/{id}/previewEvents",
	"V1CustomerService.SetIngestAliases":                   "POST v1/customers/{id}/setIngestAliases",
	"V1CustomerService.SetName":                            "POST v1/customers/{id}/setName",
	"V1CustomerService.UpdateConfig":                       "POST v1/customers/{id}/updateConfig",
	"V1PlanService.GetDetails":                             "GET v1/planDetails/{id}",
	"V1PlanService.ListCharges":                            "GET v1/planDetails/{id}/charges",
	"V1PlanService.ListChargesAutoPaging":                  "GET v1/planDetails/{id}/charges",
	"V1PlanService.ListCustomers":                          "GET v1/planDetails/{id}/customers",
	"V1PlanService.ListCustomersAutoPaging":                "GET v1/planDetails/{id}/customers",
}

type pathMethod struct {
	name      string
	fn        reflect.Value
	params    reflect.Type
	pathField []int
}

// call invokes the method with every path parameter set to value, or to
// enumValue for enum-typed parameters, and returns the method's error (for
// auto-pagers, the pager's error).
func (m pathMethod) call(recorded *[]*http.Request, value string, enumValue string) error {
	*recorded = nil
	params := reflect.New(m.params).Elem()
	for _, i := range m.pathField {
		field := params.Field(i)
		if field.Type() == reflect.TypeOf("") {
			field.SetString(value)
		} else {
			field.SetString(enumValue)
		}
	}
	out := m.fn.Call([]reflect.Value{reflect.ValueOf(context.Background()), params})
	switch last := out[len(out)-1].Interface().(type) {
	case error:
		return last
	case interface{ Err() error }:
		return last.Err()
	}
	return nil
}

func newPathTestClient(baseURL string, recorded *[]*http.Request) *metronome.Client {
	client := metronome.NewClient(
		option.WithBearerToken("My Bearer Token"),
		option.WithBaseURL(baseURL),
		option.WithMaxRetries(0),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					*recorded = append(*recorded, req)
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader("{}")),
					}, nil
				},
			},
		}),
	)
	return &client
}

// pathMethods returns every service method whose params struct has path
// parameters, so endpoints added later are covered without edits here.
func pathMethods(client *metronome.Client) []pathMethod {
	ctxType := reflect.TypeOf((*context.Context)(nil)).Elem()
	var methods []pathMethod
	var walk func(v reflect.Value)
	walk = func(v reflect.Value) {
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if field.Kind() != reflect.Struct || !strings.HasSuffix(field.Type().Name(), "Service") {
				continue
			}
			service := field.Addr()
			for j := 0; j < service.NumMethod(); j++ {
				fn := service.Method(j)
				ft := fn.Type()
				if ft.NumIn() != 3 || ft.In(0) != ctxType || ft.In(1).Kind() != reflect.Struct || !ft.IsVariadic() {
					continue
				}
				var pathField []int
				for k := 0; k < ft.In(1).NumField(); k++ {
					if _, ok := ft.In(1).Field(k).Tag.Lookup("path"); ok {
						pathField = append(pathField, k)
					}
				}
				if len(pathField) == 0 {
					continue
				}
				methods = append(methods, pathMethod{
					name:      field.Type().Name() + "." + service.Type().Method(j).Name,
					fn:        fn,
					params:    ft.In(1),
					pathField: pathField,
				})
			}
			walk(field)
		}
	}
	walk(reflect.ValueOf(client).Elem())
	sort.Slice(methods, func(a, b int) bool { return methods[a].name < methods[b].name })
	return methods
}

func TestPathParamsValidValuesUnchanged(t *testing.T) {
	var recorded []*http.Request
	methods := pathMethods(newPathTestClient(pathTestBaseURL, &recorded))
	seen := map[string]bool{}
	for _, m := range methods {
		want, ok := pathTestFixtures[m.name]
		if !ok {
			continue
		}
		seen[m.name] = true
		_ = m.call(&recorded, pathTestUUID, pathTestEnum)
		if len(recorded) != 1 {
			t.Fatalf("%s: expected 1 request, got %d", m.name, len(recorded))
		}
		method, path, _ := strings.Cut(want, " ")
		want = method + " " + pathTestBaseURL + strings.ReplaceAll(path, "{id}", pathTestUUID)
		if got := recorded[0].Method + " " + recorded[0].URL.String(); got != want {
			t.Errorf("%s: expected %q, got %q", m.name, want, got)
		}
	}
	for name := range pathTestFixtures {
		if !seen[name] {
			t.Errorf("%s: method with path parameters not found", name)
		}
	}
}

func TestPathParamsStayInOneSegment(t *testing.T) {
	var recorded []*http.Request
	methods := pathMethods(newPathTestClient(pathTestBaseURL, &recorded))
	for _, value := range []string{"../x", "../../../other", "a/b", "a?b=c", "a#b", "a%2Fb", "a b"} {
		for _, m := range methods {
			_ = m.call(&recorded, pathTestUUID, pathTestEnum)
			if len(recorded) != 1 {
				t.Fatalf("%s: expected 1 request, got %d", m.name, len(recorded))
			}
			wantQuery := recorded[0].URL.RawQuery

			_ = m.call(&recorded, value, value)
			if len(recorded) != 1 {
				t.Fatalf("%s(%q): expected 1 request, got %d", m.name, value, len(recorded))
			}
			u := recorded[0].URL
			if !strings.HasPrefix(u.EscapedPath(), "/metronome/v1/") && !strings.HasPrefix(u.EscapedPath(), "/metronome/v2/") {
				t.Errorf("%s(%q): path %q left the base URL prefix", m.name, value, u.EscapedPath())
			}
			if n := strings.Count(u.EscapedPath(), url.PathEscape(value)); n != len(m.pathField) {
				t.Errorf("%s(%q): expected %d escaped segments in %q, got %d", m.name, value, len(m.pathField), u.EscapedPath(), n)
			}
			if u.RawQuery != wantQuery || u.Fragment != "" {
				t.Errorf("%s(%q): query %q fragment %q, expected query %q and no fragment", m.name, value, u.RawQuery, u.Fragment, wantQuery)
			}
		}
	}
}

func TestPathParamsRejectDotSegments(t *testing.T) {
	var recorded []*http.Request
	methods := pathMethods(newPathTestClient(pathTestBaseURL, &recorded))
	for _, value := range []string{".", "..", "%2e", "%2E%2E", ".%2e", "%2e."} {
		for _, m := range methods {
			if err := m.call(&recorded, value, value); err == nil {
				t.Errorf("%s(%q): expected an error", m.name, value)
			}
			if len(recorded) != 0 {
				t.Errorf("%s(%q): expected no request, got %d", m.name, value, len(recorded))
			}
		}
	}
}

func TestInvoiceGetEscapesPathParams(t *testing.T) {
	var recorded []*http.Request
	client := newPathTestClient("https://api.example/", &recorded)
	for customerID, wantPath := range map[string]string{
		"a/b":  "/v1/customers/a%2Fb/invoices/inv_001",
		"100%": "/v1/customers/100%25/invoices/inv_001",
	} {
		recorded = nil
		_, _ = client.V1.Customers.Invoices.Get(context.Background(), metronome.V1CustomerInvoiceGetParams{
			CustomerID: customerID,
			InvoiceID:  "inv_001",
		})
		if len(recorded) != 1 {
			t.Fatalf("expected 1 request, got %d", len(recorded))
		}
		if got := recorded[0].URL.EscapedPath(); got != wantPath {
			t.Errorf("expected path %q, got %q", wantPath, got)
		}
		if recorded[0].URL.RawQuery != "" {
			t.Errorf("expected an empty query, got %q", recorded[0].URL.RawQuery)
		}
	}
}

// unescapedPathArgs returns the path template arguments in file that do not
// come from requestconfig.PathSegment in the same function.
func unescapedPathArgs(fset *token.FileSet, file *ast.File) []string {
	var offenders []string
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}
		escaped := map[string]bool{}
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			assign, ok := n.(*ast.AssignStmt)
			if !ok || len(assign.Rhs) != 1 || !isCall(assign.Rhs[0], "requestconfig", "PathSegment") {
				return true
			}
			if ident, ok := assign.Lhs[0].(*ast.Ident); ok {
				escaped[ident.Name] = true
			}
			return true
		})
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok || !isCall(call, "fmt", "Sprintf") || len(call.Args) < 2 {
				return true
			}
			template, ok := call.Args[0].(*ast.BasicLit)
			if !ok || !(strings.HasPrefix(template.Value, `"v1/`) || strings.HasPrefix(template.Value, `"v2/`)) {
				return true
			}
			for _, arg := range call.Args[1:] {
				if ident, ok := arg.(*ast.Ident); ok && escaped[ident.Name] {
					continue
				}
				offenders = append(offenders, fset.Position(arg.Pos()).String())
			}
			return true
		})
	}
	return offenders
}

func isCall(expr ast.Expr, pkg string, name string) bool {
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		return false
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != name {
		return false
	}
	ident, ok := sel.X.(*ast.Ident)
	return ok && ident.Name == pkg
}

func TestPathTemplatesUseEscapedSegments(t *testing.T) {
	fset := token.NewFileSet()
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return err
		}
		file, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return err
		}
		for _, offender := range unescapedPathArgs(fset, file) {
			t.Errorf("%s: path parameter is not escaped with requestconfig.PathSegment", offender)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnescapedPathArgsFindsRawParams(t *testing.T) {
	const src = `package metronome

func (r *S) Get(params P) {
	id, err := requestconfig.PathSegment(params.ID)
	_ = fmt.Sprintf("v1/things/%s", id)
	_ = fmt.Sprintf("v1/things/%s/%s", id, params.OtherID)
	_ = fmt.Sprintf("v1/things/%v", params.Kind)
	_ = err
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "fixture.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	if got := unescapedPathArgs(fset, file); len(got) != 2 {
		t.Errorf("expected 2 unescaped arguments, got %v", got)
	}
}
