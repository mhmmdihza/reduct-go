// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostAPIV1BBucketNameEntryNameBatchParams creates a new PostAPIV1BBucketNameEntryNameBatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV1BBucketNameEntryNameBatchParams() *PostAPIV1BBucketNameEntryNameBatchParams {
	return &PostAPIV1BBucketNameEntryNameBatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV1BBucketNameEntryNameBatchParamsWithTimeout creates a new PostAPIV1BBucketNameEntryNameBatchParams object
// with the ability to set a timeout on a request.
func NewPostAPIV1BBucketNameEntryNameBatchParamsWithTimeout(timeout time.Duration) *PostAPIV1BBucketNameEntryNameBatchParams {
	return &PostAPIV1BBucketNameEntryNameBatchParams{
		timeout: timeout,
	}
}

// NewPostAPIV1BBucketNameEntryNameBatchParamsWithContext creates a new PostAPIV1BBucketNameEntryNameBatchParams object
// with the ability to set a context for a request.
func NewPostAPIV1BBucketNameEntryNameBatchParamsWithContext(ctx context.Context) *PostAPIV1BBucketNameEntryNameBatchParams {
	return &PostAPIV1BBucketNameEntryNameBatchParams{
		Context: ctx,
	}
}

// NewPostAPIV1BBucketNameEntryNameBatchParamsWithHTTPClient creates a new PostAPIV1BBucketNameEntryNameBatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV1BBucketNameEntryNameBatchParamsWithHTTPClient(client *http.Client) *PostAPIV1BBucketNameEntryNameBatchParams {
	return &PostAPIV1BBucketNameEntryNameBatchParams{
		HTTPClient: client,
	}
}

/*
PostAPIV1BBucketNameEntryNameBatchParams contains all the parameters to send to the API endpoint

	for the post API v1 b bucket name entry name batch operation.

	Typically these are written to a http.Request.
*/
type PostAPIV1BBucketNameEntryNameBatchParams struct {

	/* Body.

	   Batch of records payload

	   Format: binary
	*/
	Body io.ReadCloser

	/* BucketName.

	   Name of bucket
	*/
	BucketName string

	/* EntryName.

	   Name of entry
	*/
	EntryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API v1 b bucket name entry name batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV1BBucketNameEntryNameBatchParams) WithDefaults() *PostAPIV1BBucketNameEntryNameBatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v1 b bucket name entry name batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV1BBucketNameEntryNameBatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v1 b bucket name entry name batch params
func (o *PostAPIV1BBucketNameEntryNameBatchParams) WithTimeout(timeout time.Duration) *PostAPIV1BBucketNameEntryNameBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v1 b bucket name entry name batch params
func (o *PostAPIV1BBucketNameEntryNameBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v1 b bucket name entry name batch params
func (o *PostAPIV1BBucketNameEntryNameBatchParams) WithContext(ctx context.Context) *PostAPIV1BBucketNameEntryNameBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v1 b bucket name entry name batch params
func (o *PostAPIV1BBucketNameEntryNameBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v1 b bucket name entry name batch params
func (o *PostAPIV1BBucketNameEntryNameBatchParams) WithHTTPClient(client *http.Client) *PostAPIV1BBucketNameEntryNameBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v1 b bucket name entry name batch params
func (o *PostAPIV1BBucketNameEntryNameBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post API v1 b bucket name entry name batch params
func (o *PostAPIV1BBucketNameEntryNameBatchParams) WithBody(body io.ReadCloser) *PostAPIV1BBucketNameEntryNameBatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API v1 b bucket name entry name batch params
func (o *PostAPIV1BBucketNameEntryNameBatchParams) SetBody(body io.ReadCloser) {
	o.Body = body
}

// WithBucketName adds the bucketName to the post API v1 b bucket name entry name batch params
func (o *PostAPIV1BBucketNameEntryNameBatchParams) WithBucketName(bucketName string) *PostAPIV1BBucketNameEntryNameBatchParams {
	o.SetBucketName(bucketName)
	return o
}

// SetBucketName adds the bucketName to the post API v1 b bucket name entry name batch params
func (o *PostAPIV1BBucketNameEntryNameBatchParams) SetBucketName(bucketName string) {
	o.BucketName = bucketName
}

// WithEntryName adds the entryName to the post API v1 b bucket name entry name batch params
func (o *PostAPIV1BBucketNameEntryNameBatchParams) WithEntryName(entryName string) *PostAPIV1BBucketNameEntryNameBatchParams {
	o.SetEntryName(entryName)
	return o
}

// SetEntryName adds the entryName to the post API v1 b bucket name entry name batch params
func (o *PostAPIV1BBucketNameEntryNameBatchParams) SetEntryName(entryName string) {
	o.EntryName = entryName
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV1BBucketNameEntryNameBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param bucket_name
	if err := r.SetPathParam("bucket_name", o.BucketName); err != nil {
		return err
	}

	// path param entry_name
	if err := r.SetPathParam("entry_name", o.EntryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
