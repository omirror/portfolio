// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/users.proto

package users

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

// Client API for Users service

type UsersService interface {
	Count(ctx context.Context, in *CountRequest, opts ...client.CallOption) (*CountResponse, error)
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	Find(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	Update(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*ListResponse, error)
	All(ctx context.Context, in *AllRequest, opts ...client.CallOption) (*ListResponse, error)
	ValidatePassword(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
}

type usersService struct {
	c    client.Client
	name string
}

func NewUsersService(name string, c client.Client) UsersService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "users"
	}
	return &usersService{
		c:    c,
		name: name,
	}
}

func (c *usersService) Count(ctx context.Context, in *CountRequest, opts ...client.CallOption) (*CountResponse, error) {
	req := c.c.NewRequest(c.name, "Users.Count", in)
	out := new(CountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Create(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Users.Create", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Find(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Users.Find", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Update(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Users.Update", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Users.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Users.Search", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) All(ctx context.Context, in *AllRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Users.All", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) ValidatePassword(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Users.ValidatePassword", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersHandler interface {
	Count(context.Context, *CountRequest, *CountResponse) error
	Create(context.Context, *User, *User) error
	Find(context.Context, *User, *User) error
	Update(context.Context, *User, *User) error
	List(context.Context, *ListRequest, *ListResponse) error
	Search(context.Context, *SearchRequest, *ListResponse) error
	All(context.Context, *AllRequest, *ListResponse) error
	ValidatePassword(context.Context, *User, *User) error
}

func RegisterUsersHandler(s server.Server, hdlr UsersHandler, opts ...server.HandlerOption) error {
	type users interface {
		Count(ctx context.Context, in *CountRequest, out *CountResponse) error
		Create(ctx context.Context, in *User, out *User) error
		Find(ctx context.Context, in *User, out *User) error
		Update(ctx context.Context, in *User, out *User) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Search(ctx context.Context, in *SearchRequest, out *ListResponse) error
		All(ctx context.Context, in *AllRequest, out *ListResponse) error
		ValidatePassword(ctx context.Context, in *User, out *User) error
	}
	type Users struct {
		users
	}
	h := &usersHandler{hdlr}
	return s.Handle(s.NewHandler(&Users{h}, opts...))
}

type usersHandler struct {
	UsersHandler
}

func (h *usersHandler) Count(ctx context.Context, in *CountRequest, out *CountResponse) error {
	return h.UsersHandler.Count(ctx, in, out)
}

func (h *usersHandler) Create(ctx context.Context, in *User, out *User) error {
	return h.UsersHandler.Create(ctx, in, out)
}

func (h *usersHandler) Find(ctx context.Context, in *User, out *User) error {
	return h.UsersHandler.Find(ctx, in, out)
}

func (h *usersHandler) Update(ctx context.Context, in *User, out *User) error {
	return h.UsersHandler.Update(ctx, in, out)
}

func (h *usersHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.UsersHandler.List(ctx, in, out)
}

func (h *usersHandler) Search(ctx context.Context, in *SearchRequest, out *ListResponse) error {
	return h.UsersHandler.Search(ctx, in, out)
}

func (h *usersHandler) All(ctx context.Context, in *AllRequest, out *ListResponse) error {
	return h.UsersHandler.All(ctx, in, out)
}

func (h *usersHandler) ValidatePassword(ctx context.Context, in *User, out *User) error {
	return h.UsersHandler.ValidatePassword(ctx, in, out)
}
