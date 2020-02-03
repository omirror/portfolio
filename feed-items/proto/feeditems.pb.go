// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/feeditems.proto

package feeditems

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type BulkDeleteRequest struct {
	FeedType             string   `protobuf:"bytes,1,opt,name=feed_type,json=feedType,proto3" json:"feed_type,omitempty"`
	FeedUuid             string   `protobuf:"bytes,2,opt,name=feed_uuid,json=feedUuid,proto3" json:"feed_uuid,omitempty"`
	PostUuids            []string `protobuf:"bytes,3,rep,name=post_uuids,json=postUuids,proto3" json:"post_uuids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BulkDeleteRequest) Reset()         { *m = BulkDeleteRequest{} }
func (m *BulkDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*BulkDeleteRequest) ProtoMessage()    {}
func (*BulkDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6742ade9c041eff, []int{0}
}

func (m *BulkDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkDeleteRequest.Unmarshal(m, b)
}
func (m *BulkDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkDeleteRequest.Marshal(b, m, deterministic)
}
func (m *BulkDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkDeleteRequest.Merge(m, src)
}
func (m *BulkDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_BulkDeleteRequest.Size(m)
}
func (m *BulkDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BulkDeleteRequest proto.InternalMessageInfo

func (m *BulkDeleteRequest) GetFeedType() string {
	if m != nil {
		return m.FeedType
	}
	return ""
}

func (m *BulkDeleteRequest) GetFeedUuid() string {
	if m != nil {
		return m.FeedUuid
	}
	return ""
}

func (m *BulkDeleteRequest) GetPostUuids() []string {
	if m != nil {
		return m.PostUuids
	}
	return nil
}

type FeedItem struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	FeedType             string   `protobuf:"bytes,2,opt,name=feed_type,json=feedType,proto3" json:"feed_type,omitempty"`
	FeedUuid             string   `protobuf:"bytes,3,opt,name=feed_uuid,json=feedUuid,proto3" json:"feed_uuid,omitempty"`
	Tag                  string   `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	PostUuid             string   `protobuf:"bytes,5,opt,name=post_uuid,json=postUuid,proto3" json:"post_uuid,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            int64    `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItem) Reset()         { *m = FeedItem{} }
func (m *FeedItem) String() string { return proto.CompactTextString(m) }
func (*FeedItem) ProtoMessage()    {}
func (*FeedItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6742ade9c041eff, []int{1}
}

func (m *FeedItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItem.Unmarshal(m, b)
}
func (m *FeedItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItem.Marshal(b, m, deterministic)
}
func (m *FeedItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItem.Merge(m, src)
}
func (m *FeedItem) XXX_Size() int {
	return xxx_messageInfo_FeedItem.Size(m)
}
func (m *FeedItem) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItem.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItem proto.InternalMessageInfo

func (m *FeedItem) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *FeedItem) GetFeedType() string {
	if m != nil {
		return m.FeedType
	}
	return ""
}

func (m *FeedItem) GetFeedUuid() string {
	if m != nil {
		return m.FeedUuid
	}
	return ""
}

func (m *FeedItem) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *FeedItem) GetPostUuid() string {
	if m != nil {
		return m.PostUuid
	}
	return ""
}

func (m *FeedItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FeedItem) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6742ade9c041eff, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Response struct {
	Error                *Error    `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Item                 *FeedItem `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6742ade9c041eff, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetItem() *FeedItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type GetFeedRequest struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Page                 int32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedRequest) Reset()         { *m = GetFeedRequest{} }
func (m *GetFeedRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedRequest) ProtoMessage()    {}
func (*GetFeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6742ade9c041eff, []int{4}
}

func (m *GetFeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedRequest.Unmarshal(m, b)
}
func (m *GetFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedRequest.Merge(m, src)
}
func (m *GetFeedRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedRequest.Size(m)
}
func (m *GetFeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedRequest proto.InternalMessageInfo

func (m *GetFeedRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetFeedRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *GetFeedRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetFeedRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListRequest struct {
	Uuids                []string `protobuf:"bytes,1,rep,name=uuids,proto3" json:"uuids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6742ade9c041eff, []int{5}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetUuids() []string {
	if m != nil {
		return m.Uuids
	}
	return nil
}

type GetFeedResponse struct {
	Error                *Error      `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Items                []*FeedItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetFeedResponse) Reset()         { *m = GetFeedResponse{} }
func (m *GetFeedResponse) String() string { return proto.CompactTextString(m) }
func (*GetFeedResponse) ProtoMessage()    {}
func (*GetFeedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6742ade9c041eff, []int{6}
}

func (m *GetFeedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedResponse.Unmarshal(m, b)
}
func (m *GetFeedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedResponse.Marshal(b, m, deterministic)
}
func (m *GetFeedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedResponse.Merge(m, src)
}
func (m *GetFeedResponse) XXX_Size() int {
	return xxx_messageInfo_GetFeedResponse.Size(m)
}
func (m *GetFeedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedResponse proto.InternalMessageInfo

func (m *GetFeedResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetFeedResponse) GetItems() []*FeedItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*BulkDeleteRequest)(nil), "BulkDeleteRequest")
	proto.RegisterType((*FeedItem)(nil), "FeedItem")
	proto.RegisterType((*Error)(nil), "Error")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*GetFeedRequest)(nil), "GetFeedRequest")
	proto.RegisterType((*ListRequest)(nil), "ListRequest")
	proto.RegisterType((*GetFeedResponse)(nil), "GetFeedResponse")
}

func init() { proto.RegisterFile("proto/feeditems.proto", fileDescriptor_b6742ade9c041eff) }

var fileDescriptor_b6742ade9c041eff = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xe3, 0x8f, 0xc4, 0x13, 0x89, 0x96, 0x51, 0x91, 0x56, 0x85, 0x08, 0x6b, 0xb9, 0xe4,
	0x82, 0x91, 0x82, 0xf8, 0x01, 0x7c, 0x46, 0x48, 0x1c, 0xd0, 0x0a, 0xce, 0x55, 0x1a, 0x0f, 0x95,
	0x45, 0x52, 0x1b, 0xef, 0xf8, 0xd0, 0x7f, 0xc8, 0x85, 0xff, 0x84, 0x76, 0xb6, 0x9b, 0xc4, 0x45,
	0x41, 0xbd, 0xcd, 0xbc, 0x19, 0xcf, 0xcc, 0x7b, 0xcf, 0x0b, 0x4f, 0xda, 0xae, 0xe1, 0xe6, 0xd5,
	0x0f, 0xa2, 0xaa, 0x66, 0xda, 0xda, 0x52, 0x72, 0xbd, 0x81, 0xc7, 0xef, 0xfa, 0xcd, 0xcf, 0x0f,
	0xb4, 0x21, 0x26, 0x43, 0xbf, 0x7a, 0xb2, 0x8c, 0x4f, 0x21, 0x77, 0x7d, 0x97, 0x7c, 0xdb, 0x92,
	0x8a, 0x8a, 0x68, 0x9e, 0x9b, 0x89, 0x03, 0xbe, 0xdd, 0xb6, 0xb4, 0x2b, 0xf6, 0x7d, 0x5d, 0xa9,
	0xd1, 0xbe, 0xf8, 0xbd, 0xaf, 0x2b, 0x9c, 0x01, 0xb4, 0x8d, 0x65, 0x29, 0x5a, 0x15, 0x17, 0xf1,
	0x3c, 0x37, 0xb9, 0x43, 0x5c, 0xd5, 0xea, 0xdf, 0x11, 0x4c, 0x3e, 0x11, 0x55, 0x9f, 0x99, 0xb6,
	0x88, 0x90, 0xc8, 0x0c, 0xbf, 0x40, 0xe2, 0xe1, 0xe6, 0xd1, 0xff, 0x36, 0xc7, 0xf7, 0x36, 0x9f,
	0x41, 0xcc, 0xab, 0x6b, 0x95, 0x08, 0xec, 0x42, 0xd7, 0xbe, 0xbb, 0x45, 0xa5, 0xbe, 0x3d, 0x9c,
	0x82, 0x05, 0x4c, 0x2b, 0xb2, 0xeb, 0xae, 0x6e, 0xb9, 0x6e, 0x6e, 0x54, 0x26, 0xe5, 0x43, 0xc8,
	0x51, 0x59, 0x77, 0xb4, 0x62, 0xaa, 0x2e, 0x57, 0xac, 0xc6, 0x45, 0x34, 0x8f, 0x4d, 0x7e, 0x87,
	0xbc, 0x65, 0xfd, 0x06, 0xd2, 0x8f, 0x5d, 0xd7, 0x74, 0x8e, 0xc6, 0xba, 0xa9, 0xbc, 0x4e, 0xa9,
	0x91, 0x18, 0x15, 0x8c, 0xb7, 0x64, 0xed, 0xea, 0x3a, 0x90, 0x08, 0xa9, 0x5e, 0xc2, 0xc4, 0x90,
	0x6d, 0x9b, 0x1b, 0x4b, 0xf8, 0x0c, 0x52, 0x72, 0x23, 0xe4, 0xd3, 0xe9, 0x22, 0x2b, 0x65, 0xa0,
	0xf1, 0x20, 0xce, 0x20, 0x71, 0x46, 0xc9, 0x80, 0xe9, 0x22, 0x2f, 0x83, 0x6e, 0x46, 0x60, 0x7d,
	0x05, 0x8f, 0x96, 0xc4, 0x0e, 0x0c, 0xae, 0x21, 0x24, 0x07, 0x86, 0x49, 0xbc, 0xd3, 0x78, 0x74,
	0xa0, 0x31, 0x42, 0xd2, 0xba, 0xcb, 0x62, 0x7f, 0xb0, 0x8b, 0xf1, 0x1c, 0xd2, 0x4d, 0xbd, 0xad,
	0x59, 0xf4, 0x4b, 0x8d, 0x4f, 0xf4, 0x0b, 0x98, 0x7e, 0xa9, 0x2d, 0x87, 0x05, 0xe7, 0x90, 0x7a,
	0x5f, 0x23, 0xf1, 0xd5, 0x27, 0xfa, 0x2b, 0x9c, 0xee, 0x0e, 0x79, 0x10, 0xb1, 0xe7, 0x90, 0xca,
	0x1f, 0xa8, 0x46, 0x45, 0x3c, 0x64, 0xe6, 0xf1, 0xc5, 0x9f, 0x08, 0xf2, 0x80, 0x59, 0x2c, 0x20,
	0x7b, 0x2f, 0xaa, 0xe3, 0xbe, 0xf3, 0x22, 0x2f, 0xc3, 0x32, 0x7d, 0x82, 0x33, 0x88, 0x97, 0xc4,
	0x47, 0xcb, 0x05, 0x64, 0xfe, 0xf7, 0x3e, 0xda, 0xf1, 0x12, 0x60, 0xff, 0x08, 0x10, 0xcb, 0x7f,
	0x5e, 0xc4, 0xb0, 0xbd, 0x84, 0xf1, 0x1d, 0x63, 0x3c, 0x2d, 0x87, 0x26, 0x5c, 0x9c, 0x95, 0xf7,
	0xc4, 0xd0, 0x27, 0x57, 0x99, 0x3c, 0xb5, 0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x0e,
	0x16, 0x5b, 0x83, 0x03, 0x00, 0x00,
}
