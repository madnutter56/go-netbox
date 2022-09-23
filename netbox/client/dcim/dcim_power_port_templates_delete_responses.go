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
)

// DcimPowerPortTemplatesDeleteReader is a Reader for the DcimPowerPortTemplatesDelete structure.
type DcimPowerPortTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerPortTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortTemplatesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortTemplatesDeleteNoContent creates a DcimPowerPortTemplatesDeleteNoContent with default headers values
func NewDcimPowerPortTemplatesDeleteNoContent() *DcimPowerPortTemplatesDeleteNoContent {
	return &DcimPowerPortTemplatesDeleteNoContent{}
}

/*
DcimPowerPortTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerPortTemplatesDeleteNoContent dcim power port templates delete no content
*/
type DcimPowerPortTemplatesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim power port templates delete no content response has a 2xx status code
func (o *DcimPowerPortTemplatesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power port templates delete no content response has a 3xx status code
func (o *DcimPowerPortTemplatesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power port templates delete no content response has a 4xx status code
func (o *DcimPowerPortTemplatesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power port templates delete no content response has a 5xx status code
func (o *DcimPowerPortTemplatesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power port templates delete no content response a status code equal to that given
func (o *DcimPowerPortTemplatesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimPowerPortTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesDeleteNoContent ", 204)
}

func (o *DcimPowerPortTemplatesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesDeleteNoContent ", 204)
}

func (o *DcimPowerPortTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimPowerPortTemplatesDeleteDefault creates a DcimPowerPortTemplatesDeleteDefault with default headers values
func NewDcimPowerPortTemplatesDeleteDefault(code int) *DcimPowerPortTemplatesDeleteDefault {
	return &DcimPowerPortTemplatesDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortTemplatesDeleteDefault describes a response with status code -1, with default header values.

DcimPowerPortTemplatesDeleteDefault dcim power port templates delete default
*/
type DcimPowerPortTemplatesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power port templates delete default response
func (o *DcimPowerPortTemplatesDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power port templates delete default response has a 2xx status code
func (o *DcimPowerPortTemplatesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power port templates delete default response has a 3xx status code
func (o *DcimPowerPortTemplatesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power port templates delete default response has a 4xx status code
func (o *DcimPowerPortTemplatesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power port templates delete default response has a 5xx status code
func (o *DcimPowerPortTemplatesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power port templates delete default response a status code equal to that given
func (o *DcimPowerPortTemplatesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerPortTemplatesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-port-templates/{id}/][%d] dcim_power-port-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-port-templates/{id}/][%d] dcim_power-port-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortTemplatesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
