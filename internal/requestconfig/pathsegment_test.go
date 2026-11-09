package requestconfig

import "testing"

func TestPathSegmentEscapes(t *testing.T) {
	for in, want := range map[string]string{
		"13117714-3f05-48e5-a6e9-a66093f13b4d": "13117714-3f05-48e5-a6e9-a66093f13b4d",
		"aws_marketplace":                      "aws_marketplace",
		"a/b":                                  "a%2Fb",
		"a?b":                                  "a%3Fb",
		"a#b":                                  "a%23b",
		"a%2Fb":                                "a%252Fb",
		"../x":                                 "..%2Fx",
		"...":                                  "...",
	} {
		got, err := PathSegment(in)
		if err != nil {
			t.Errorf("PathSegment(%q): unexpected error %v", in, err)
		}
		if got != want {
			t.Errorf("PathSegment(%q) = %q, expected %q", in, got, want)
		}
	}
}

func TestPathSegmentRejectsDotSegments(t *testing.T) {
	for _, in := range []string{"", ".", "..", "%2e", "%2E", "%2e%2e", "%2E%2E", ".%2e", "%2E.", ".%2E"} {
		if _, err := PathSegment(in); err == nil {
			t.Errorf("PathSegment(%q): expected an error", in)
		}
	}
}
