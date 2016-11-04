// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the License). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the license file accompanying this file. This file is distributed
// on an AS IS BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/aws/amazon-ecs-event-stream-handler/internal/client/operations"
)

// Default amazon ecs esh HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new amazon ecs esh HTTP client.
func NewHTTPClient(formats strfmt.Registry) *AmazonEcsEsh {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost:3000", "/v1", []string{"http"})
	return New(transport, formats)
}

// New creates a new amazon ecs esh client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *AmazonEcsEsh {
	cli := new(AmazonEcsEsh)
	cli.Transport = transport

	cli.Operations = operations.New(transport, formats)

	return cli
}

// AmazonEcsEsh is a client for amazon ecs esh
type AmazonEcsEsh struct {
	Operations *operations.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *AmazonEcsEsh) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Operations.SetTransport(transport)

}