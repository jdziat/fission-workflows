// Code generated by go-swagger; DO NOT EDIT.

package workflow_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fission/fission-workflows/cmd/wfcli/swagger-client/models"
)

// WfGetReader is a Reader for the WfGet structure.
type WfGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WfGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWfGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWfGetOK creates a WfGetOK with default headers values
func NewWfGetOK() *WfGetOK {
	return &WfGetOK{}
}

/*WfGetOK handles this case with default header values.

WfGetOK wf get o k
*/
type WfGetOK struct {
	Payload *models.Workflow
}

func (o *WfGetOK) Error() string {
	return fmt.Sprintf("[GET /workflow/{id}][%d] wfGetOK  %+v", 200, o.Payload)
}

func (o *WfGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Workflow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}