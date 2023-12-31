// Code generated by goa v3.11.3, DO NOT EDIT.
//
// diary endpoints
//
// Command:
// $ goa gen github.com/shibayu36/go-diary-api/design

package diary

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "diary" service endpoints.
type Endpoints struct {
	UserSignup  goa.Endpoint
	Signin      goa.Endpoint
	CreateDiary goa.Endpoint
}

// NewEndpoints wraps the methods of the "diary" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		UserSignup:  NewUserSignupEndpoint(s),
		Signin:      NewSigninEndpoint(s),
		CreateDiary: NewCreateDiaryEndpoint(s, a.APIKeyAuth),
	}
}

// Use applies the given middleware to all the "diary" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.UserSignup = m(e.UserSignup)
	e.Signin = m(e.Signin)
	e.CreateDiary = m(e.CreateDiary)
}

// NewUserSignupEndpoint returns an endpoint function that calls the method
// "UserSignup" of service "diary".
func NewUserSignupEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UserSignupPayload)
		return nil, s.UserSignup(ctx, p)
	}
}

// NewSigninEndpoint returns an endpoint function that calls the method
// "Signin" of service "diary".
func NewSigninEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*SigninPayload)
		return s.Signin(ctx, p)
	}
}

// NewCreateDiaryEndpoint returns an endpoint function that calls the method
// "CreateDiary" of service "diary".
func NewCreateDiaryEndpoint(s Service, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreateDiaryPayload)
		var err error
		sc := security.APIKeyScheme{
			Name:           "api_key",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var key string
		if p.Key != nil {
			key = *p.Key
		}
		ctx, err = authAPIKeyFn(ctx, key, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.CreateDiary(ctx, p)
	}
}
