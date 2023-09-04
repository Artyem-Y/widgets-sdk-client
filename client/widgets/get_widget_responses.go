// Code generated by go-swagger; DO NOT EDIT.

package widgets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"terraform-provider-hashicups-pf/models"
)

// GetWidgetReader is a Reader for the GetWidget structure.
type GetWidgetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWidgetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWidgetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWidgetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWidgetOK creates a GetWidgetOK with default headers values
func NewGetWidgetOK() *GetWidgetOK {
	return &GetWidgetOK{}
}

/*
GetWidgetOK describes a response with status code 200, with default header values.

Widget details
*/
type GetWidgetOK struct {
	Payload *models.Widget
}

// IsSuccess returns true when this get widget o k response has a 2xx status code
func (o *GetWidgetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get widget o k response has a 3xx status code
func (o *GetWidgetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get widget o k response has a 4xx status code
func (o *GetWidgetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get widget o k response has a 5xx status code
func (o *GetWidgetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get widget o k response a status code equal to that given
func (o *GetWidgetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get widget o k response
func (o *GetWidgetOK) Code() int {
	return 200
}

func (o *GetWidgetOK) Error() string {
	return fmt.Sprintf("[GET /widgets/{id}][%d] getWidgetOK  %+v", 200, o.Payload)
}

func (o *GetWidgetOK) String() string {
	return fmt.Sprintf("[GET /widgets/{id}][%d] getWidgetOK  %+v", 200, o.Payload)
}

func (o *GetWidgetOK) GetPayload() *models.Widget {
	return o.Payload
}

func (o *GetWidgetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Widget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetDefault creates a GetWidgetDefault with default headers values
func NewGetWidgetDefault(code int) *GetWidgetDefault {
	return &GetWidgetDefault{
		_statusCode: code,
	}
}

/*
GetWidgetDefault describes a response with status code -1, with default header values.

Error response
*/
type GetWidgetDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get widget default response has a 2xx status code
func (o *GetWidgetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get widget default response has a 3xx status code
func (o *GetWidgetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get widget default response has a 4xx status code
func (o *GetWidgetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get widget default response has a 5xx status code
func (o *GetWidgetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get widget default response a status code equal to that given
func (o *GetWidgetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get widget default response
func (o *GetWidgetDefault) Code() int {
	return o._statusCode
}

func (o *GetWidgetDefault) Error() string {
	return fmt.Sprintf("[GET /widgets/{id}][%d] getWidget default  %+v", o._statusCode, o.Payload)
}

func (o *GetWidgetDefault) String() string {
	return fmt.Sprintf("[GET /widgets/{id}][%d] getWidget default  %+v", o._statusCode, o.Payload)
}

func (o *GetWidgetDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWidgetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
