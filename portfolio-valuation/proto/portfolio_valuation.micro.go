// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/portfolio_valuation.proto

package portfolio_valuation

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

// Client API for PortfolioValuation service

type PortfolioValuationService interface {
	GetPortfolio(ctx context.Context, in *Portfolio, opts ...client.CallOption) (*Portfolio, error)
}

type portfolioValuationService struct {
	c    client.Client
	name string
}

func NewPortfolioValuationService(name string, c client.Client) PortfolioValuationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "portfoliovaluation"
	}
	return &portfolioValuationService{
		c:    c,
		name: name,
	}
}

func (c *portfolioValuationService) GetPortfolio(ctx context.Context, in *Portfolio, opts ...client.CallOption) (*Portfolio, error) {
	req := c.c.NewRequest(c.name, "PortfolioValuation.GetPortfolio", in)
	out := new(Portfolio)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PortfolioValuation service

type PortfolioValuationHandler interface {
	GetPortfolio(context.Context, *Portfolio, *Portfolio) error
}

func RegisterPortfolioValuationHandler(s server.Server, hdlr PortfolioValuationHandler, opts ...server.HandlerOption) error {
	type portfolioValuation interface {
		GetPortfolio(ctx context.Context, in *Portfolio, out *Portfolio) error
	}
	type PortfolioValuation struct {
		portfolioValuation
	}
	h := &portfolioValuationHandler{hdlr}
	return s.Handle(s.NewHandler(&PortfolioValuation{h}, opts...))
}

type portfolioValuationHandler struct {
	PortfolioValuationHandler
}

func (h *portfolioValuationHandler) GetPortfolio(ctx context.Context, in *Portfolio, out *Portfolio) error {
	return h.PortfolioValuationHandler.GetPortfolio(ctx, in, out)
}
