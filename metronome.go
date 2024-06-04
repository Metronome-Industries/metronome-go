// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/param"
)

type IngestParams struct {
	Body []IngestParamsBody `json:"body,required"`
}

func (r IngestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type IngestParamsBody struct {
	CustomerID param.Field[string] `json:"customer_id,required"`
	EventType  param.Field[string] `json:"event_type,required"`
	// RFC 3339 formatted
	Timestamp     param.Field[string]                 `json:"timestamp,required"`
	TransactionID param.Field[string]                 `json:"transaction_id,required"`
	Properties    param.Field[map[string]interface{}] `json:"properties"`
}

func (r IngestParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
