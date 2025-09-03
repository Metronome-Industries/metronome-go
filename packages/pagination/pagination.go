// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type CursorPage[T any] struct {
	// Cursor to fetch the next page
	NextPage string `json:"next_page"`
	// Items of the page
	Data []T `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPage    respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorPage[T]) GetNextPage() (res *CursorPage[T], err error) {
	next := r.NextPage
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("next_page", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorPageAutoPager[T any] struct {
	page *CursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorPageAutoPager[T any](page *CursorPage[T], err error) *CursorPageAutoPager[T] {
	return &CursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorPageAutoPager[T]) Next() bool {
	if r.page == nil {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	// if the API returned empty data then keep iterating
	// until we either get more data or there are no more pages
	// to fetch
	for len(r.page.Data) == 0 {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorPageAutoPager[T]) Index() int {
	return r.run
}

type BodyCursorPage[T any] struct {
	// Cursor to fetch the next page
	NextPage string `json:"next_page"`
	// Items of the page
	Data []T `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPage    respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r BodyCursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *BodyCursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *BodyCursorPage[T]) GetNextPage() (res *BodyCursorPage[T], err error) {
	next := r.NextPage
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("next_page", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *BodyCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &BodyCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type BodyCursorPageAutoPager[T any] struct {
	page *BodyCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewBodyCursorPageAutoPager[T any](page *BodyCursorPage[T], err error) *BodyCursorPageAutoPager[T] {
	return &BodyCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *BodyCursorPageAutoPager[T]) Next() bool {
	if r.page == nil {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	// if the API returned empty data then keep iterating
	// until we either get more data or there are no more pages
	// to fetch
	for len(r.page.Data) == 0 {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *BodyCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *BodyCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *BodyCursorPageAutoPager[T]) Index() int {
	return r.run
}

type CursorPageWithoutLimit[T any] struct {
	// Cursor to fetch the next page
	NextPage string `json:"next_page"`
	// Items of the page
	Data []T `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPage    respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorPageWithoutLimit[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorPageWithoutLimit[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorPageWithoutLimit[T]) GetNextPage() (res *CursorPageWithoutLimit[T], err error) {
	next := r.NextPage
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("next_page", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorPageWithoutLimit[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorPageWithoutLimit[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorPageWithoutLimitAutoPager[T any] struct {
	page *CursorPageWithoutLimit[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorPageWithoutLimitAutoPager[T any](page *CursorPageWithoutLimit[T], err error) *CursorPageWithoutLimitAutoPager[T] {
	return &CursorPageWithoutLimitAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorPageWithoutLimitAutoPager[T]) Next() bool {
	if r.page == nil {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	// if the API returned empty data then keep iterating
	// until we either get more data or there are no more pages
	// to fetch
	for len(r.page.Data) == 0 {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorPageWithoutLimitAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorPageWithoutLimitAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorPageWithoutLimitAutoPager[T]) Index() int {
	return r.run
}
