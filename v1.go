// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"github.com/Metronome-Industries/metronome-go/v3/option"
)

// V1Service contains methods and other services that help with interacting with
// the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1Service] method instead.
type V1Service struct {
	Options []option.RequestOption
	// [Alerts](https://docs.metronome.com/connecting-metronome/alerts/) monitor
	// customer spending, balances, and other billing factors. Use these endpoints to
	// create, retrieve, and archive customer alerts. To view sample alert payloads by
	// alert type, navigate
	// [here.](https://docs.metronome.com/manage-product-access/create-manage-alerts/#webhook-notifications)
	Alerts V1AlertService
	// [Plans](https://docs.metronome.com/pricing-and-packaging/create-plans/)
	// determine the base pricing for a customer. Use these endpoints to add a plan to
	// a customer, end a customer plan, retrieve plans, and retrieve plan details.
	// Create plans in the [Metronome app](https://app.metronome.com/plans).
	Plans V1PlanService
	// [Credit grants](https://docs.metronome.com/invoicing/how-billing-works/manage-credits/)
	// adjust a customer balance for prepayments, reimbursements, promotions, and so
	// on. Use these endpoints to create, retrieve, update, and delete credit grants.
	CreditGrants V1CreditGrantService
	// Use these endpoints to configure a billing API key, a webhook secret, or invoice
	// finalization behavior.
	PricingUnits V1PricingUnitService
	Customers    V1CustomerService
	// [Customers](https://docs.metronome.com/provisioning/create-customers/) in
	// Metronome represent your users for all billing and reporting. Use these
	// endpoints to create, retrieve, update, and archive customers and their billing
	// configuration.
	Dashboards V1DashboardService
	// [Usage events](https://docs.metronome.com/connecting-metronome/send-usage-data/)
	// are the basis for billable metrics. Use these endpoints to send usage events to
	// Metronome and retrieve aggregated event data.
	Usage V1UsageService
	// [Security](https://docs.metronome.com/developer-resources/security/) endpoints
	// allow you to retrieve security-related data.
	AuditLogs V1AuditLogService
	// [Custom fields](https://docs.metronome.com/integrations/custom-fields/) enable
	// adding additional data to Metronome entities. Use these endpoints to create,
	// retrieve, update, and delete custom fields.
	CustomFields V1CustomFieldService
	// [Billable metrics](https://docs.metronome.com/understanding-metronome/how-metronome-works#billable-metrics)
	// in Metronome represent the various consumption components that Metronome meters
	// and aggregates.
	BillableMetrics V1BillableMetricService
	// [Security](https://docs.metronome.com/developer-resources/security/) endpoints
	// allow you to retrieve security-related data.
	Services V1ServiceService
	// [Invoices](https://docs.metronome.com/invoicing/) reflect how much a customer
	// spent during a period, which is the basis for billing. Metronome automatically
	// generates invoices based upon your pricing, packaging, and usage events. Use
	// these endpoints to retrieve invoices.
	Invoices  V1InvoiceService
	Contracts V1ContractService
	Packages  V1PackageService
	Payments  V1PaymentService
	// Use these endpoints to configure a billing API key, a webhook secret, or invoice
	// finalization behavior.
	Settings V1SettingService
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
	r.Packages = NewV1PackageService(opts...)
	r.Payments = NewV1PaymentService(opts...)
	r.Settings = NewV1SettingService(opts...)
	return
}
