package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
FilterInstances Filter instances
*/
func (a *Client) FilterInstances(params *FilterInstancesParams) (*FilterInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterInstancesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FilterInstances",
		Method:             "GET",
		PathPattern:        "/instances/filter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FilterInstancesReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FilterInstancesOK), nil

}

/*
FilterTasks Filter tasks
*/
func (a *Client) FilterTasks(params *FilterTasksParams) (*FilterTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FilterTasks",
		Method:             "GET",
		PathPattern:        "/tasks/filter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FilterTasksReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FilterTasksOK), nil

}

/*
GetInstance Get instance based on a single ARN
*/
func (a *Client) GetInstance(params *GetInstanceParams) (*GetInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInstanceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInstance",
		Method:             "GET",
		PathPattern:        "/instance/{arn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInstanceReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInstanceOK), nil

}

/*
GetTask Get task based on a single ARN
*/
func (a *Client) GetTask(params *GetTaskParams) (*GetTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTask",
		Method:             "GET",
		PathPattern:        "/task/{arn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTaskReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTaskOK), nil

}

/*
ListInstances Lists all instances
*/
func (a *Client) ListInstances(params *ListInstancesParams) (*ListInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListInstancesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListInstances",
		Method:             "GET",
		PathPattern:        "/instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListInstancesReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListInstancesOK), nil

}

/*
ListTasks Lists all tasks
*/
func (a *Client) ListTasks(params *ListTasksParams) (*ListTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListTasks",
		Method:             "GET",
		PathPattern:        "/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListTasksReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListTasksOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
