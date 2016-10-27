package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListTasksParams creates a new ListTasksParams object
// with the default values initialized.
func NewListTasksParams() *ListTasksParams {

	return &ListTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTasksParamsWithTimeout creates a new ListTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTasksParamsWithTimeout(timeout time.Duration) *ListTasksParams {

	return &ListTasksParams{

		timeout: timeout,
	}
}

// NewListTasksParamsWithContext creates a new ListTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTasksParamsWithContext(ctx context.Context) *ListTasksParams {

	return &ListTasksParams{

		Context: ctx,
	}
}

/*ListTasksParams contains all the parameters to send to the API endpoint
for the list tasks operation typically these are written to a http.Request
*/
type ListTasksParams struct {
	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the list tasks params
func (o *ListTasksParams) WithTimeout(timeout time.Duration) *ListTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tasks params
func (o *ListTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tasks params
func (o *ListTasksParams) WithContext(ctx context.Context) *ListTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tasks params
func (o *ListTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *ListTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
