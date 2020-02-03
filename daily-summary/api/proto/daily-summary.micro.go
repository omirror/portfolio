// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/daily-summary.proto

package daily_summary

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DailySummary service

type DailySummaryService interface {
	Morning(ctx context.Context, in *Request, opts ...client.CallOption) (*MorningSummary, error)
	Evening(ctx context.Context, in *Request, opts ...client.CallOption) (*EveningSummary, error)
}

type dailySummaryService struct {
	c    client.Client
	name string
}

func NewDailySummaryService(name string, c client.Client) DailySummaryService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "dailysummary"
	}
	return &dailySummaryService{
		c:    c,
		name: name,
	}
}

func (c *dailySummaryService) Morning(ctx context.Context, in *Request, opts ...client.CallOption) (*MorningSummary, error) {
	req := c.c.NewRequest(c.name, "DailySummary.Morning", in)
	out := new(MorningSummary)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dailySummaryService) Evening(ctx context.Context, in *Request, opts ...client.CallOption) (*EveningSummary, error) {
	req := c.c.NewRequest(c.name, "DailySummary.Evening", in)
	out := new(EveningSummary)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DailySummary service

type DailySummaryHandler interface {
	Morning(context.Context, *Request, *MorningSummary) error
	Evening(context.Context, *Request, *EveningSummary) error
}

func RegisterDailySummaryHandler(s server.Server, hdlr DailySummaryHandler, opts ...server.HandlerOption) error {
	type dailySummary interface {
		Morning(ctx context.Context, in *Request, out *MorningSummary) error
		Evening(ctx context.Context, in *Request, out *EveningSummary) error
	}
	type DailySummary struct {
		dailySummary
	}
	h := &dailySummaryHandler{hdlr}
	return s.Handle(s.NewHandler(&DailySummary{h}, opts...))
}

type dailySummaryHandler struct {
	DailySummaryHandler
}

func (h *dailySummaryHandler) Morning(ctx context.Context, in *Request, out *MorningSummary) error {
	return h.DailySummaryHandler.Morning(ctx, in, out)
}

func (h *dailySummaryHandler) Evening(ctx context.Context, in *Request, out *EveningSummary) error {
	return h.DailySummaryHandler.Evening(ctx, in, out)
}