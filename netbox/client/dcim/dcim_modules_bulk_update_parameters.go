// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// NewDcimModulesBulkUpdateParams creates a new DcimModulesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModulesBulkUpdateParams() *DcimModulesBulkUpdateParams {
	return &DcimModulesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModulesBulkUpdateParamsWithTimeout creates a new DcimModulesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimModulesBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimModulesBulkUpdateParams {
	return &DcimModulesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimModulesBulkUpdateParamsWithContext creates a new DcimModulesBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimModulesBulkUpdateParamsWithContext(ctx context.Context) *DcimModulesBulkUpdateParams {
	return &DcimModulesBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimModulesBulkUpdateParamsWithHTTPClient creates a new DcimModulesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModulesBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimModulesBulkUpdateParams {
	return &DcimModulesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimModulesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim modules bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimModulesBulkUpdateParams struct {

	// Data.
	Data *models.WritableModule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim modules bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModulesBulkUpdateParams) WithDefaults() *DcimModulesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim modules bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModulesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim modules bulk update params
func (o *DcimModulesBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimModulesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim modules bulk update params
func (o *DcimModulesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim modules bulk update params
func (o *DcimModulesBulkUpdateParams) WithContext(ctx context.Context) *DcimModulesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim modules bulk update params
func (o *DcimModulesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim modules bulk update params
func (o *DcimModulesBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimModulesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim modules bulk update params
func (o *DcimModulesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim modules bulk update params
func (o *DcimModulesBulkUpdateParams) WithData(data *models.WritableModule) *DcimModulesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim modules bulk update params
func (o *DcimModulesBulkUpdateParams) SetData(data *models.WritableModule) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModulesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
