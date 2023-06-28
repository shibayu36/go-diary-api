// Code generated by goa v3.11.3, DO NOT EDIT.
//
// diary HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/shibayu36/go-diary-api/design

package server

import (
	"context"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeUserSignupResponse returns an encoder for responses returned by the
// diary UserSignup endpoint.
func EncodeUserSignupResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// DecodeUserSignupRequest returns a decoder for requests sent to the diary
// UserSignup endpoint.
func DecodeUserSignupRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body UserSignupRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUserSignupRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewUserSignupPayload(&body)

		return payload, nil
	}
}

// EncodeSigninResponse returns an encoder for responses returned by the diary
// Signin endpoint.
func EncodeSigninResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSigninRequest returns a decoder for requests sent to the diary Signin
// endpoint.
func DecodeSigninRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body SigninRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateSigninRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewSigninPayload(&body)

		return payload, nil
	}
}
