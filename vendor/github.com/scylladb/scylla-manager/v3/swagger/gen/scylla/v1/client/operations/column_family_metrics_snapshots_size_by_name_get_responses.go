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

// ColumnFamilyMetricsSnapshotsSizeByNameGetReader is a Reader for the ColumnFamilyMetricsSnapshotsSizeByNameGet structure.
type ColumnFamilyMetricsSnapshotsSizeByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsSnapshotsSizeByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsSnapshotsSizeByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsSnapshotsSizeByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsSnapshotsSizeByNameGetOK creates a ColumnFamilyMetricsSnapshotsSizeByNameGetOK with default headers values
func NewColumnFamilyMetricsSnapshotsSizeByNameGetOK() *ColumnFamilyMetricsSnapshotsSizeByNameGetOK {
	return &ColumnFamilyMetricsSnapshotsSizeByNameGetOK{}
}

/*
ColumnFamilyMetricsSnapshotsSizeByNameGetOK handles this case with default header values.

Success
*/
type ColumnFamilyMetricsSnapshotsSizeByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsSnapshotsSizeByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsSnapshotsSizeByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsSnapshotsSizeByNameGetDefault creates a ColumnFamilyMetricsSnapshotsSizeByNameGetDefault with default headers values
func NewColumnFamilyMetricsSnapshotsSizeByNameGetDefault(code int) *ColumnFamilyMetricsSnapshotsSizeByNameGetDefault {
	return &ColumnFamilyMetricsSnapshotsSizeByNameGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsSnapshotsSizeByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsSnapshotsSizeByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics snapshots size by name get default response
func (o *ColumnFamilyMetricsSnapshotsSizeByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsSnapshotsSizeByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsSnapshotsSizeByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsSnapshotsSizeByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}