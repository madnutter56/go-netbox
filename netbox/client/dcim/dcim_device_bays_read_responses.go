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

// DcimDeviceBaysReadReader is a Reader for the DcimDeviceBaysRead structure.
type DcimDeviceBaysReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBaysReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceBaysReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceBaysReadOK creates a DcimDeviceBaysReadOK with default headers values
func NewDcimDeviceBaysReadOK() *DcimDeviceBaysReadOK {
	return &DcimDeviceBaysReadOK{}
}

/*
DcimDeviceBaysReadOK describes a response with status code 200, with default header values.

DcimDeviceBaysReadOK dcim device bays read o k
*/
type DcimDeviceBaysReadOK struct {
	Payload *models.DeviceBay
}

// IsSuccess returns true when this dcim device bays read o k response has a 2xx status code
func (o *DcimDeviceBaysReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bays read o k response has a 3xx status code
func (o *DcimDeviceBaysReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bays read o k response has a 4xx status code
func (o *DcimDeviceBaysReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bays read o k response has a 5xx status code
func (o *DcimDeviceBaysReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bays read o k response a status code equal to that given
func (o *DcimDeviceBaysReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDeviceBaysReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/device-bays/{id}/][%d] dcimDeviceBaysReadOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBaysReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/device-bays/{id}/][%d] dcimDeviceBaysReadOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBaysReadOK) GetPayload() *models.DeviceBay {
	return o.Payload
}

func (o *DcimDeviceBaysReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceBaysReadDefault creates a DcimDeviceBaysReadDefault with default headers values
func NewDcimDeviceBaysReadDefault(code int) *DcimDeviceBaysReadDefault {
	return &DcimDeviceBaysReadDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceBaysReadDefault describes a response with status code -1, with default header values.

DcimDeviceBaysReadDefault dcim device bays read default
*/
type DcimDeviceBaysReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device bays read default response
func (o *DcimDeviceBaysReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device bays read default response has a 2xx status code
func (o *DcimDeviceBaysReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device bays read default response has a 3xx status code
func (o *DcimDeviceBaysReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device bays read default response has a 4xx status code
func (o *DcimDeviceBaysReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device bays read default response has a 5xx status code
func (o *DcimDeviceBaysReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device bays read default response a status code equal to that given
func (o *DcimDeviceBaysReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceBaysReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/device-bays/{id}/][%d] dcim_device-bays_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBaysReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/device-bays/{id}/][%d] dcim_device-bays_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBaysReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBaysReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
