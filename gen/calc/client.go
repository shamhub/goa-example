// Code generated by goa v3.14.6, DO NOT EDIT.
//
// calc client
//
// Command:
// $ goa gen github.com/shamhub/goa-example/api

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "calc" service client.
type Client struct {
	DivideEndpoint goa.Endpoint
	AddEndpoint    goa.Endpoint
}

// NewClient initializes a "calc" service client given the endpoints.
func NewClient(divide, add goa.Endpoint) *Client {
	return &Client{
		DivideEndpoint: divide,
		AddEndpoint:    add,
	}
}

// Divide calls the "divide" endpoint of the "calc" service.
// Divide may return the following errors:
//   - "div_by_zero" (type *DivByZero): division by 0
//   - "timeout" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) Divide(ctx context.Context, p *DividePayload) (res *DivideResult, err error) {
	var ires any
	ires, err = c.DivideEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*DivideResult), nil
}

// Add calls the "add" endpoint of the "calc" service.
// Add may return the following errors:
//   - "timeout" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) Add(ctx context.Context, p *AddPayload) (res int, err error) {
	var ires any
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}