// Code generated by goa v3.11.3, DO NOT EDIT.
//
// calc gRPC client types
//
// Command:
// $ goa gen github.com/shibayu36/go-playground/diary/design

package client

import (
	calc "github.com/shibayu36/go-playground/diary/gen/calc"
	calcpb "github.com/shibayu36/go-playground/diary/gen/grpc/calc/pb"
)

// NewProtoAddRequest builds the gRPC request type from the payload of the
// "add" endpoint of the "calc" service.
func NewProtoAddRequest(payload *calc.AddPayload) *calcpb.AddRequest {
	message := &calcpb.AddRequest{
		A: int32(payload.A),
		B: int32(payload.B),
	}
	return message
}

// NewAddResult builds the result type of the "add" endpoint of the "calc"
// service from the gRPC response type.
func NewAddResult(message *calcpb.AddResponse) int {
	result := int(message.Field)
	return result
}
