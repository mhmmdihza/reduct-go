// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPIV1BBucketNameReader is a Reader for the DeleteAPIV1BBucketName structure.
type DeleteAPIV1BBucketNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIV1BBucketNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIV1BBucketNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAPIV1BBucketNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAPIV1BBucketNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPIV1BBucketNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v1/b/{bucket_name}] DeleteAPIV1BBucketName", response, response.Code())
	}
}

// NewDeleteAPIV1BBucketNameOK creates a DeleteAPIV1BBucketNameOK with default headers values
func NewDeleteAPIV1BBucketNameOK() *DeleteAPIV1BBucketNameOK {
	return &DeleteAPIV1BBucketNameOK{}
}

/*
DeleteAPIV1BBucketNameOK describes a response with status code 200, with default header values.

The bucket is deleted
*/
type DeleteAPIV1BBucketNameOK struct {
}

// IsSuccess returns true when this delete Api v1 b bucket name o k response has a 2xx status code
func (o *DeleteAPIV1BBucketNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Api v1 b bucket name o k response has a 3xx status code
func (o *DeleteAPIV1BBucketNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api v1 b bucket name o k response has a 4xx status code
func (o *DeleteAPIV1BBucketNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api v1 b bucket name o k response has a 5xx status code
func (o *DeleteAPIV1BBucketNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api v1 b bucket name o k response a status code equal to that given
func (o *DeleteAPIV1BBucketNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete Api v1 b bucket name o k response
func (o *DeleteAPIV1BBucketNameOK) Code() int {
	return 200
}

func (o *DeleteAPIV1BBucketNameOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/b/{bucket_name}][%d] deleteApiV1BBucketNameOK", 200)
}

func (o *DeleteAPIV1BBucketNameOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/b/{bucket_name}][%d] deleteApiV1BBucketNameOK", 200)
}

func (o *DeleteAPIV1BBucketNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIV1BBucketNameUnauthorized creates a DeleteAPIV1BBucketNameUnauthorized with default headers values
func NewDeleteAPIV1BBucketNameUnauthorized() *DeleteAPIV1BBucketNameUnauthorized {
	return &DeleteAPIV1BBucketNameUnauthorized{}
}

/*
DeleteAPIV1BBucketNameUnauthorized describes a response with status code 401, with default header values.

Access token is invalid or empty
*/
type DeleteAPIV1BBucketNameUnauthorized struct {
}

// IsSuccess returns true when this delete Api v1 b bucket name unauthorized response has a 2xx status code
func (o *DeleteAPIV1BBucketNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api v1 b bucket name unauthorized response has a 3xx status code
func (o *DeleteAPIV1BBucketNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api v1 b bucket name unauthorized response has a 4xx status code
func (o *DeleteAPIV1BBucketNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api v1 b bucket name unauthorized response has a 5xx status code
func (o *DeleteAPIV1BBucketNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api v1 b bucket name unauthorized response a status code equal to that given
func (o *DeleteAPIV1BBucketNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete Api v1 b bucket name unauthorized response
func (o *DeleteAPIV1BBucketNameUnauthorized) Code() int {
	return 401
}

func (o *DeleteAPIV1BBucketNameUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/b/{bucket_name}][%d] deleteApiV1BBucketNameUnauthorized", 401)
}

func (o *DeleteAPIV1BBucketNameUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/b/{bucket_name}][%d] deleteApiV1BBucketNameUnauthorized", 401)
}

func (o *DeleteAPIV1BBucketNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIV1BBucketNameForbidden creates a DeleteAPIV1BBucketNameForbidden with default headers values
func NewDeleteAPIV1BBucketNameForbidden() *DeleteAPIV1BBucketNameForbidden {
	return &DeleteAPIV1BBucketNameForbidden{}
}

/*
DeleteAPIV1BBucketNameForbidden describes a response with status code 403, with default header values.

Access token doesn't have enough permissions
*/
type DeleteAPIV1BBucketNameForbidden struct {
}

// IsSuccess returns true when this delete Api v1 b bucket name forbidden response has a 2xx status code
func (o *DeleteAPIV1BBucketNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api v1 b bucket name forbidden response has a 3xx status code
func (o *DeleteAPIV1BBucketNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api v1 b bucket name forbidden response has a 4xx status code
func (o *DeleteAPIV1BBucketNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api v1 b bucket name forbidden response has a 5xx status code
func (o *DeleteAPIV1BBucketNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api v1 b bucket name forbidden response a status code equal to that given
func (o *DeleteAPIV1BBucketNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete Api v1 b bucket name forbidden response
func (o *DeleteAPIV1BBucketNameForbidden) Code() int {
	return 403
}

func (o *DeleteAPIV1BBucketNameForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/b/{bucket_name}][%d] deleteApiV1BBucketNameForbidden", 403)
}

func (o *DeleteAPIV1BBucketNameForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/b/{bucket_name}][%d] deleteApiV1BBucketNameForbidden", 403)
}

func (o *DeleteAPIV1BBucketNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIV1BBucketNameNotFound creates a DeleteAPIV1BBucketNameNotFound with default headers values
func NewDeleteAPIV1BBucketNameNotFound() *DeleteAPIV1BBucketNameNotFound {
	return &DeleteAPIV1BBucketNameNotFound{}
}

/*
DeleteAPIV1BBucketNameNotFound describes a response with status code 404, with default header values.

Bucket doesn't exist
*/
type DeleteAPIV1BBucketNameNotFound struct {
}

// IsSuccess returns true when this delete Api v1 b bucket name not found response has a 2xx status code
func (o *DeleteAPIV1BBucketNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api v1 b bucket name not found response has a 3xx status code
func (o *DeleteAPIV1BBucketNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api v1 b bucket name not found response has a 4xx status code
func (o *DeleteAPIV1BBucketNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api v1 b bucket name not found response has a 5xx status code
func (o *DeleteAPIV1BBucketNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api v1 b bucket name not found response a status code equal to that given
func (o *DeleteAPIV1BBucketNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Api v1 b bucket name not found response
func (o *DeleteAPIV1BBucketNameNotFound) Code() int {
	return 404
}

func (o *DeleteAPIV1BBucketNameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/b/{bucket_name}][%d] deleteApiV1BBucketNameNotFound", 404)
}

func (o *DeleteAPIV1BBucketNameNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/b/{bucket_name}][%d] deleteApiV1BBucketNameNotFound", 404)
}

func (o *DeleteAPIV1BBucketNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
