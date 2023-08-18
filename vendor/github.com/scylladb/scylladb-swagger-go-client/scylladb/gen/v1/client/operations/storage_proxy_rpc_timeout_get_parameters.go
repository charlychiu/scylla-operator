// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewStorageProxyRPCTimeoutGetParams creates a new StorageProxyRPCTimeoutGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageProxyRPCTimeoutGetParams() *StorageProxyRPCTimeoutGetParams {
	return &StorageProxyRPCTimeoutGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyRPCTimeoutGetParamsWithTimeout creates a new StorageProxyRPCTimeoutGetParams object
// with the ability to set a timeout on a request.
func NewStorageProxyRPCTimeoutGetParamsWithTimeout(timeout time.Duration) *StorageProxyRPCTimeoutGetParams {
	return &StorageProxyRPCTimeoutGetParams{
		timeout: timeout,
	}
}

// NewStorageProxyRPCTimeoutGetParamsWithContext creates a new StorageProxyRPCTimeoutGetParams object
// with the ability to set a context for a request.
func NewStorageProxyRPCTimeoutGetParamsWithContext(ctx context.Context) *StorageProxyRPCTimeoutGetParams {
	return &StorageProxyRPCTimeoutGetParams{
		Context: ctx,
	}
}

// NewStorageProxyRPCTimeoutGetParamsWithHTTPClient creates a new StorageProxyRPCTimeoutGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageProxyRPCTimeoutGetParamsWithHTTPClient(client *http.Client) *StorageProxyRPCTimeoutGetParams {
	return &StorageProxyRPCTimeoutGetParams{
		HTTPClient: client,
	}
}

/*
StorageProxyRPCTimeoutGetParams contains all the parameters to send to the API endpoint

	for the storage proxy Rpc timeout get operation.

	Typically these are written to a http.Request.
*/
type StorageProxyRPCTimeoutGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage proxy Rpc timeout get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageProxyRPCTimeoutGetParams) WithDefaults() *StorageProxyRPCTimeoutGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage proxy Rpc timeout get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageProxyRPCTimeoutGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage proxy Rpc timeout get params
func (o *StorageProxyRPCTimeoutGetParams) WithTimeout(timeout time.Duration) *StorageProxyRPCTimeoutGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy Rpc timeout get params
func (o *StorageProxyRPCTimeoutGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy Rpc timeout get params
func (o *StorageProxyRPCTimeoutGetParams) WithContext(ctx context.Context) *StorageProxyRPCTimeoutGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy Rpc timeout get params
func (o *StorageProxyRPCTimeoutGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy Rpc timeout get params
func (o *StorageProxyRPCTimeoutGetParams) WithHTTPClient(client *http.Client) *StorageProxyRPCTimeoutGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy Rpc timeout get params
func (o *StorageProxyRPCTimeoutGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyRPCTimeoutGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}