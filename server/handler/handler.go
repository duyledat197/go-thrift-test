package handler

import (
	"context"

	"github.com/duyledat197/test-thrift/thr/shared"
	"github.com/duyledat197/test-thrift/thr/tutorial"
)

// CaculatorHandler ...
type CaculatorHandler struct {
}

//NewCaculatorHandler ...
func NewCaculatorHandler() tutorial.Calculator {
	return &CaculatorHandler{}
}

// Add ...
func (c *CaculatorHandler) Add(ctx context.Context, req *tutorial.AddRequest) (r *tutorial.AddResponse, err error) {
	return &tutorial.AddResponse{
		Result_: req.GetA() + req.GetB(),
	}, nil
}

// GetStruct ...
func (c *CaculatorHandler) GetStruct(ctx context.Context, key int32) (r *shared.SharedStruct, err error) {
	return &shared.SharedStruct{
		Key:   key,
		Value: "hello",
	}, nil
}
