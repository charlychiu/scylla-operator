// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylladb-swagger-go-client/scylladb/gen/v2/models"
)

// FindConfigHintedHandoffThrottleInKbReader is a Reader for the FindConfigHintedHandoffThrottleInKb structure.
type FindConfigHintedHandoffThrottleInKbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigHintedHandoffThrottleInKbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigHintedHandoffThrottleInKbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigHintedHandoffThrottleInKbDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigHintedHandoffThrottleInKbOK creates a FindConfigHintedHandoffThrottleInKbOK with default headers values
func NewFindConfigHintedHandoffThrottleInKbOK() *FindConfigHintedHandoffThrottleInKbOK {
	return &FindConfigHintedHandoffThrottleInKbOK{}
}

/*
FindConfigHintedHandoffThrottleInKbOK handles this case with default header values.

Config value
*/
type FindConfigHintedHandoffThrottleInKbOK struct {
	Payload int64
}

func (o *FindConfigHintedHandoffThrottleInKbOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigHintedHandoffThrottleInKbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigHintedHandoffThrottleInKbDefault creates a FindConfigHintedHandoffThrottleInKbDefault with default headers values
func NewFindConfigHintedHandoffThrottleInKbDefault(code int) *FindConfigHintedHandoffThrottleInKbDefault {
	return &FindConfigHintedHandoffThrottleInKbDefault{
		_statusCode: code,
	}
}

/*
FindConfigHintedHandoffThrottleInKbDefault handles this case with default header values.

unexpected error
*/
type FindConfigHintedHandoffThrottleInKbDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config hinted handoff throttle in kb default response
func (o *FindConfigHintedHandoffThrottleInKbDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigHintedHandoffThrottleInKbDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigHintedHandoffThrottleInKbDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigHintedHandoffThrottleInKbDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}