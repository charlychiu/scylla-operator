// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v1/models"
)

// StorageProxyHintedHandoffEnabledByDcPostReader is a Reader for the StorageProxyHintedHandoffEnabledByDcPost structure.
type StorageProxyHintedHandoffEnabledByDcPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyHintedHandoffEnabledByDcPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyHintedHandoffEnabledByDcPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyHintedHandoffEnabledByDcPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyHintedHandoffEnabledByDcPostOK creates a StorageProxyHintedHandoffEnabledByDcPostOK with default headers values
func NewStorageProxyHintedHandoffEnabledByDcPostOK() *StorageProxyHintedHandoffEnabledByDcPostOK {
	return &StorageProxyHintedHandoffEnabledByDcPostOK{}
}

/*
StorageProxyHintedHandoffEnabledByDcPostOK handles this case with default header values.

Success
*/
type StorageProxyHintedHandoffEnabledByDcPostOK struct {
}

func (o *StorageProxyHintedHandoffEnabledByDcPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageProxyHintedHandoffEnabledByDcPostDefault creates a StorageProxyHintedHandoffEnabledByDcPostDefault with default headers values
func NewStorageProxyHintedHandoffEnabledByDcPostDefault(code int) *StorageProxyHintedHandoffEnabledByDcPostDefault {
	return &StorageProxyHintedHandoffEnabledByDcPostDefault{
		_statusCode: code,
	}
}

/*
StorageProxyHintedHandoffEnabledByDcPostDefault handles this case with default header values.

internal server error
*/
type StorageProxyHintedHandoffEnabledByDcPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy hinted handoff enabled by dc post default response
func (o *StorageProxyHintedHandoffEnabledByDcPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyHintedHandoffEnabledByDcPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyHintedHandoffEnabledByDcPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyHintedHandoffEnabledByDcPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}