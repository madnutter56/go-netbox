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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// ExtrasTagsBulkUpdateReader is a Reader for the ExtrasTagsBulkUpdate structure.
type ExtrasTagsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasTagsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasTagsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasTagsBulkUpdateOK creates a ExtrasTagsBulkUpdateOK with default headers values
func NewExtrasTagsBulkUpdateOK() *ExtrasTagsBulkUpdateOK {
	return &ExtrasTagsBulkUpdateOK{}
}

/*
ExtrasTagsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasTagsBulkUpdateOK extras tags bulk update o k
*/
type ExtrasTagsBulkUpdateOK struct {
	Payload *models.Tag
}

// IsSuccess returns true when this extras tags bulk update o k response has a 2xx status code
func (o *ExtrasTagsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras tags bulk update o k response has a 3xx status code
func (o *ExtrasTagsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras tags bulk update o k response has a 4xx status code
func (o *ExtrasTagsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras tags bulk update o k response has a 5xx status code
func (o *ExtrasTagsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras tags bulk update o k response a status code equal to that given
func (o *ExtrasTagsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasTagsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/tags/][%d] extrasTagsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasTagsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/tags/][%d] extrasTagsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasTagsBulkUpdateOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *ExtrasTagsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasTagsBulkUpdateDefault creates a ExtrasTagsBulkUpdateDefault with default headers values
func NewExtrasTagsBulkUpdateDefault(code int) *ExtrasTagsBulkUpdateDefault {
	return &ExtrasTagsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasTagsBulkUpdateDefault describes a response with status code -1, with default header values.

ExtrasTagsBulkUpdateDefault extras tags bulk update default
*/
type ExtrasTagsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras tags bulk update default response
func (o *ExtrasTagsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras tags bulk update default response has a 2xx status code
func (o *ExtrasTagsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras tags bulk update default response has a 3xx status code
func (o *ExtrasTagsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras tags bulk update default response has a 4xx status code
func (o *ExtrasTagsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras tags bulk update default response has a 5xx status code
func (o *ExtrasTagsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras tags bulk update default response a status code equal to that given
func (o *ExtrasTagsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasTagsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/tags/][%d] extras_tags_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasTagsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/tags/][%d] extras_tags_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasTagsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasTagsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
