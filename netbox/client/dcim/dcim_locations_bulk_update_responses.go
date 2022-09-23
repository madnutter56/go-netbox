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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimLocationsBulkUpdateReader is a Reader for the DcimLocationsBulkUpdate structure.
type DcimLocationsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimLocationsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimLocationsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimLocationsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimLocationsBulkUpdateOK creates a DcimLocationsBulkUpdateOK with default headers values
func NewDcimLocationsBulkUpdateOK() *DcimLocationsBulkUpdateOK {
	return &DcimLocationsBulkUpdateOK{}
}

/*
DcimLocationsBulkUpdateOK describes a response with status code 200, with default header values.

DcimLocationsBulkUpdateOK dcim locations bulk update o k
*/
type DcimLocationsBulkUpdateOK struct {
	Payload *models.Location
}

// IsSuccess returns true when this dcim locations bulk update o k response has a 2xx status code
func (o *DcimLocationsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim locations bulk update o k response has a 3xx status code
func (o *DcimLocationsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim locations bulk update o k response has a 4xx status code
func (o *DcimLocationsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim locations bulk update o k response has a 5xx status code
func (o *DcimLocationsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim locations bulk update o k response a status code equal to that given
func (o *DcimLocationsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimLocationsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/locations/][%d] dcimLocationsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimLocationsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/locations/][%d] dcimLocationsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimLocationsBulkUpdateOK) GetPayload() *models.Location {
	return o.Payload
}

func (o *DcimLocationsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Location)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimLocationsBulkUpdateDefault creates a DcimLocationsBulkUpdateDefault with default headers values
func NewDcimLocationsBulkUpdateDefault(code int) *DcimLocationsBulkUpdateDefault {
	return &DcimLocationsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimLocationsBulkUpdateDefault describes a response with status code -1, with default header values.

DcimLocationsBulkUpdateDefault dcim locations bulk update default
*/
type DcimLocationsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim locations bulk update default response
func (o *DcimLocationsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim locations bulk update default response has a 2xx status code
func (o *DcimLocationsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim locations bulk update default response has a 3xx status code
func (o *DcimLocationsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim locations bulk update default response has a 4xx status code
func (o *DcimLocationsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim locations bulk update default response has a 5xx status code
func (o *DcimLocationsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim locations bulk update default response a status code equal to that given
func (o *DcimLocationsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimLocationsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/locations/][%d] dcim_locations_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/locations/][%d] dcim_locations_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimLocationsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
