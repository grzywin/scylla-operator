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

// StorageProxyMetricsRangeEstimatedHistogramGetReader is a Reader for the StorageProxyMetricsRangeEstimatedHistogramGet structure.
type StorageProxyMetricsRangeEstimatedHistogramGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsRangeEstimatedHistogramGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsRangeEstimatedHistogramGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsRangeEstimatedHistogramGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsRangeEstimatedHistogramGetOK creates a StorageProxyMetricsRangeEstimatedHistogramGetOK with default headers values
func NewStorageProxyMetricsRangeEstimatedHistogramGetOK() *StorageProxyMetricsRangeEstimatedHistogramGetOK {
	return &StorageProxyMetricsRangeEstimatedHistogramGetOK{}
}

/*
StorageProxyMetricsRangeEstimatedHistogramGetOK handles this case with default header values.

StorageProxyMetricsRangeEstimatedHistogramGetOK storage proxy metrics range estimated histogram get o k
*/
type StorageProxyMetricsRangeEstimatedHistogramGetOK struct {
}

func (o *StorageProxyMetricsRangeEstimatedHistogramGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageProxyMetricsRangeEstimatedHistogramGetDefault creates a StorageProxyMetricsRangeEstimatedHistogramGetDefault with default headers values
func NewStorageProxyMetricsRangeEstimatedHistogramGetDefault(code int) *StorageProxyMetricsRangeEstimatedHistogramGetDefault {
	return &StorageProxyMetricsRangeEstimatedHistogramGetDefault{
		_statusCode: code,
	}
}

/*
StorageProxyMetricsRangeEstimatedHistogramGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsRangeEstimatedHistogramGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics range estimated histogram get default response
func (o *StorageProxyMetricsRangeEstimatedHistogramGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsRangeEstimatedHistogramGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsRangeEstimatedHistogramGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsRangeEstimatedHistogramGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}