// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"github.com/Metronome-Industries/metronome-go/option"
)

// V1Service contains methods and other services that help with interacting with
// the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1Service] method instead.
type V1Service struct {
	Options         []option.RequestOption
	Alerts          V1AlertService
	Plans           V1PlanService
	CreditGrants    V1CreditGrantService
	PricingUnits    V1PricingUnitService
	Customers       V1CustomerService
	Dashboards      V1DashboardService
	Usage           V1UsageService
	AuditLogs       V1AuditLogService
	CustomFields    V1CustomFieldService
	BillableMetrics V1BillableMetricService
	Services        V1ServiceService
	Invoices        V1InvoiceService
	Contracts       V1ContractService
	Payments        V1PaymentService
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r V1Service) {
	r = V1Service{}
	r.Options = opts
	r.Alerts = NewV1AlertService(opts...)
	r.Plans = NewV1PlanService(opts...)
	r.CreditGrants = NewV1CreditGrantService(opts...)
	r.PricingUnits = NewV1PricingUnitService(opts...)
	r.Customers = NewV1CustomerService(opts...)
	r.Dashboards = NewV1DashboardService(opts...)
	r.Usage = NewV1UsageService(opts...)
	r.AuditLogs = NewV1AuditLogService(opts...)
	r.CustomFields = NewV1CustomFieldService(opts...)
	r.BillableMetrics = NewV1BillableMetricService(opts...)
	r.Services = NewV1ServiceService(opts...)
	r.Invoices = NewV1InvoiceService(opts...)
	r.Contracts = NewV1ContractService(opts...)
	r.Payments = NewV1PaymentService(opts...)
	return
}
