// Code generated by goa v3.11.3, DO NOT EDIT.
//
// diary HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/shibayu36/go-playground/diary/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	diary "github.com/shibayu36/go-playground/diary/gen/diary"
	goahttp "goa.design/goa/v3/http"
)

// BuildUserSignupRequest instantiates a HTTP request object with method and
// path set to call the "diary" service "UserSignup" endpoint
func (c *Client) BuildUserSignupRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UserSignupDiaryPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("diary", "UserSignup", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUserSignupRequest returns an encoder for requests sent to the diary
// UserSignup server.
func EncodeUserSignupRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*diary.UserSignupPayload)
		if !ok {
			return goahttp.ErrInvalidType("diary", "UserSignup", "*diary.UserSignupPayload", v)
		}
		body := NewUserSignupRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("diary", "UserSignup", err)
		}
		return nil
	}
}

// DecodeUserSignupResponse returns a decoder for responses returned by the
// diary UserSignup endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeUserSignupResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("diary", "UserSignup", resp.StatusCode, string(body))
		}
	}
}