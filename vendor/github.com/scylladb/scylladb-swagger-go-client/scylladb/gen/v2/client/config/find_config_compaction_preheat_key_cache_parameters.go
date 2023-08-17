// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewFindConfigCompactionPreheatKeyCacheParams creates a new FindConfigCompactionPreheatKeyCacheParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindConfigCompactionPreheatKeyCacheParams() *FindConfigCompactionPreheatKeyCacheParams {
	return &FindConfigCompactionPreheatKeyCacheParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigCompactionPreheatKeyCacheParamsWithTimeout creates a new FindConfigCompactionPreheatKeyCacheParams object
// with the ability to set a timeout on a request.
func NewFindConfigCompactionPreheatKeyCacheParamsWithTimeout(timeout time.Duration) *FindConfigCompactionPreheatKeyCacheParams {
	return &FindConfigCompactionPreheatKeyCacheParams{
		timeout: timeout,
	}
}

// NewFindConfigCompactionPreheatKeyCacheParamsWithContext creates a new FindConfigCompactionPreheatKeyCacheParams object
// with the ability to set a context for a request.
func NewFindConfigCompactionPreheatKeyCacheParamsWithContext(ctx context.Context) *FindConfigCompactionPreheatKeyCacheParams {
	return &FindConfigCompactionPreheatKeyCacheParams{
		Context: ctx,
	}
}

// NewFindConfigCompactionPreheatKeyCacheParamsWithHTTPClient creates a new FindConfigCompactionPreheatKeyCacheParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindConfigCompactionPreheatKeyCacheParamsWithHTTPClient(client *http.Client) *FindConfigCompactionPreheatKeyCacheParams {
	return &FindConfigCompactionPreheatKeyCacheParams{
		HTTPClient: client,
	}
}

/*
FindConfigCompactionPreheatKeyCacheParams contains all the parameters to send to the API endpoint

	for the find config compaction preheat key cache operation.

	Typically these are written to a http.Request.
*/
type FindConfigCompactionPreheatKeyCacheParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find config compaction preheat key cache params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigCompactionPreheatKeyCacheParams) WithDefaults() *FindConfigCompactionPreheatKeyCacheParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find config compaction preheat key cache params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigCompactionPreheatKeyCacheParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find config compaction preheat key cache params
func (o *FindConfigCompactionPreheatKeyCacheParams) WithTimeout(timeout time.Duration) *FindConfigCompactionPreheatKeyCacheParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config compaction preheat key cache params
func (o *FindConfigCompactionPreheatKeyCacheParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config compaction preheat key cache params
func (o *FindConfigCompactionPreheatKeyCacheParams) WithContext(ctx context.Context) *FindConfigCompactionPreheatKeyCacheParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config compaction preheat key cache params
func (o *FindConfigCompactionPreheatKeyCacheParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config compaction preheat key cache params
func (o *FindConfigCompactionPreheatKeyCacheParams) WithHTTPClient(client *http.Client) *FindConfigCompactionPreheatKeyCacheParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config compaction preheat key cache params
func (o *FindConfigCompactionPreheatKeyCacheParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigCompactionPreheatKeyCacheParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}