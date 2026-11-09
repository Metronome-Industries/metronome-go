package requestconfig

import (
	"errors"
	"net/url"
	"strings"
)

var errDotSegment = errors.New("path parameter must not be empty, '.' or '..'")

// PathSegment returns v escaped for use as exactly one URL path segment. It
// rejects values that URL resolution would treat as a dot segment.
func PathSegment(v string) (string, error) {
	switch strings.ToLower(v) {
	case "", ".", "..", "%2e", "%2e%2e", ".%2e", "%2e.":
		return "", errDotSegment
	}
	return url.PathEscape(v), nil
}
