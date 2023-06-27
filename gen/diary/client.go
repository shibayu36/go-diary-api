// Code generated by goa v3.11.3, DO NOT EDIT.
//
// diary client
//
// Command:
// $ goa gen github.com/shibayu36/go-diary-api/design

package diary

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "diary" service client.
type Client struct {
	UserSignupEndpoint goa.Endpoint
}

// NewClient initializes a "diary" service client given the endpoints.
func NewClient(userSignup goa.Endpoint) *Client {
	return &Client{
		UserSignupEndpoint: userSignup,
	}
}

// UserSignup calls the "UserSignup" endpoint of the "diary" service.
// UserSignup may return the following errors:
//   - "user_validation_error" (type *goa.ServiceError)
//   - "user_duplication_error" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) UserSignup(ctx context.Context, p *UserSignupPayload) (err error) {
	_, err = c.UserSignupEndpoint(ctx, p)
	return
}
