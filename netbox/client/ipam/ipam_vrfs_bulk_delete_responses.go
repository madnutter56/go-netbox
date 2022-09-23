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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamVrfsBulkDeleteReader is a Reader for the IpamVrfsBulkDelete structure.
type IpamVrfsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVrfsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVrfsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVrfsBulkDeleteNoContent creates a IpamVrfsBulkDeleteNoContent with default headers values
func NewIpamVrfsBulkDeleteNoContent() *IpamVrfsBulkDeleteNoContent {
	return &IpamVrfsBulkDeleteNoContent{}
}

/*
IpamVrfsBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamVrfsBulkDeleteNoContent ipam vrfs bulk delete no content
*/
type IpamVrfsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this ipam vrfs bulk delete no content response has a 2xx status code
func (o *IpamVrfsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vrfs bulk delete no content response has a 3xx status code
func (o *IpamVrfsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vrfs bulk delete no content response has a 4xx status code
func (o *IpamVrfsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vrfs bulk delete no content response has a 5xx status code
func (o *IpamVrfsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vrfs bulk delete no content response a status code equal to that given
func (o *IpamVrfsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamVrfsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vrfs/][%d] ipamVrfsBulkDeleteNoContent ", 204)
}

func (o *IpamVrfsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/vrfs/][%d] ipamVrfsBulkDeleteNoContent ", 204)
}

func (o *IpamVrfsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamVrfsBulkDeleteDefault creates a IpamVrfsBulkDeleteDefault with default headers values
func NewIpamVrfsBulkDeleteDefault(code int) *IpamVrfsBulkDeleteDefault {
	return &IpamVrfsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
IpamVrfsBulkDeleteDefault describes a response with status code -1, with default header values.

IpamVrfsBulkDeleteDefault ipam vrfs bulk delete default
*/
type IpamVrfsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vrfs bulk delete default response
func (o *IpamVrfsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam vrfs bulk delete default response has a 2xx status code
func (o *IpamVrfsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vrfs bulk delete default response has a 3xx status code
func (o *IpamVrfsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vrfs bulk delete default response has a 4xx status code
func (o *IpamVrfsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vrfs bulk delete default response has a 5xx status code
func (o *IpamVrfsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vrfs bulk delete default response a status code equal to that given
func (o *IpamVrfsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamVrfsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vrfs/][%d] ipam_vrfs_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVrfsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /ipam/vrfs/][%d] ipam_vrfs_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVrfsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVrfsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
