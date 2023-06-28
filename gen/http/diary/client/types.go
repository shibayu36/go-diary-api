// Code generated by goa v3.11.3, DO NOT EDIT.
//
// diary HTTP client types
//
// Command:
// $ goa gen github.com/shibayu36/go-diary-api/design

package client

import (
	diary "github.com/shibayu36/go-diary-api/gen/diary"
)

// UserSignupRequestBody is the type of the "diary" service "UserSignup"
// endpoint HTTP request body.
type UserSignupRequestBody struct {
	// User name
	Name string `form:"name" json:"name" xml:"name"`
	// User email
	Email string `form:"email" json:"email" xml:"email"`
}

// SigninRequestBody is the type of the "diary" service "Signin" endpoint HTTP
// request body.
type SigninRequestBody struct {
	// User email
	Email string `form:"email" json:"email" xml:"email"`
}

// NewUserSignupRequestBody builds the HTTP request body from the payload of
// the "UserSignup" endpoint of the "diary" service.
func NewUserSignupRequestBody(p *diary.UserSignupPayload) *UserSignupRequestBody {
	body := &UserSignupRequestBody{
		Name:  p.Name,
		Email: p.Email,
	}
	return body
}

// NewSigninRequestBody builds the HTTP request body from the payload of the
// "Signin" endpoint of the "diary" service.
func NewSigninRequestBody(p *diary.SigninPayload) *SigninRequestBody {
	body := &SigninRequestBody{
		Email: p.Email,
	}
	return body
}
