// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/param"
)

type ID struct {
	ID   string `json:"id,required" format:"uuid"`
	JSON idJSON `json:"-"`
}

// idJSON contains the JSON metadata for the struct [ID]
type idJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idJSON) RawJSON() string {
	return r.raw
}

type IDParam struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r IDParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
