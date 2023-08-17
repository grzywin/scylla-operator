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

	"github.com/scylladb/scylladb-swagger-go-client/scylladb/gen/v1/models"
)

// StorageProxyMetricsCasReadConditionNotMetGetReader is a Reader for the StorageProxyMetricsCasReadConditionNotMetGet structure.
type StorageProxyMetricsCasReadConditionNotMetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsCasReadConditionNotMetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsCasReadConditionNotMetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsCasReadConditionNotMetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsCasReadConditionNotMetGetOK creates a StorageProxyMetricsCasReadConditionNotMetGetOK with default headers values
func NewStorageProxyMetricsCasReadConditionNotMetGetOK() *StorageProxyMetricsCasReadConditionNotMetGetOK {
	return &StorageProxyMetricsCasReadConditionNotMetGetOK{}
}

/*
StorageProxyMetricsCasReadConditionNotMetGetOK handles this case with default header values.

StorageProxyMetricsCasReadConditionNotMetGetOK storage proxy metrics cas read condition not met get o k
*/
type StorageProxyMetricsCasReadConditionNotMetGetOK struct {
	Payload int32
}

func (o *StorageProxyMetricsCasReadConditionNotMetGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageProxyMetricsCasReadConditionNotMetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyMetricsCasReadConditionNotMetGetDefault creates a StorageProxyMetricsCasReadConditionNotMetGetDefault with default headers values
func NewStorageProxyMetricsCasReadConditionNotMetGetDefault(code int) *StorageProxyMetricsCasReadConditionNotMetGetDefault {
	return &StorageProxyMetricsCasReadConditionNotMetGetDefault{
		_statusCode: code,
	}
}

/*
StorageProxyMetricsCasReadConditionNotMetGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsCasReadConditionNotMetGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics cas read condition not met get default response
func (o *StorageProxyMetricsCasReadConditionNotMetGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsCasReadConditionNotMetGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsCasReadConditionNotMetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsCasReadConditionNotMetGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}