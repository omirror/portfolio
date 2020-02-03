// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/portfolios.proto

package portfolios

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

type ListRequest struct {
	UserUuids            []string `protobuf:"bytes,1,rep,name=user_uuids,json=userUuids,proto3" json:"user_uuids,omitempty"`
	Uuids                []string `protobuf:"bytes,2,rep,name=uuids,proto3" json:"uuids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fac9e0553d4fb04e, []int{0}
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

func (m *ListRequest) GetUserUuids() []string {
	if m != nil {
		return m.UserUuids
	}
	return nil
}

func (m *ListRequest) GetUuids() []string {
	if m != nil {
		return m.Uuids
	}
	return nil
}

type ListResponse struct {
	Portfolios           []*Portfolio `protobuf:"bytes,1,rep,name=portfolios,proto3" json:"portfolios,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fac9e0553d4fb04e, []int{1}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetPortfolios() []*Portfolio {
	if m != nil {
		return m.Portfolios
	}
	return nil
}

type AllRequest struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllRequest) Reset()         { *m = AllRequest{} }
func (m *AllRequest) String() string { return proto.CompactTextString(m) }
func (*AllRequest) ProtoMessage()    {}
func (*AllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fac9e0553d4fb04e, []int{2}
}

func (m *AllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllRequest.Unmarshal(m, b)
}
func (m *AllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllRequest.Marshal(b, m, deterministic)
}
func (m *AllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllRequest.Merge(m, src)
}
func (m *AllRequest) XXX_Size() int {
	return xxx_messageInfo_AllRequest.Size(m)
}
func (m *AllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllRequest proto.InternalMessageInfo

func (m *AllRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *AllRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type AllResponse struct {
	Portfolios           []*Portfolio `protobuf:"bytes,1,rep,name=portfolios,proto3" json:"portfolios,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AllResponse) Reset()         { *m = AllResponse{} }
func (m *AllResponse) String() string { return proto.CompactTextString(m) }
func (*AllResponse) ProtoMessage()    {}
func (*AllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fac9e0553d4fb04e, []int{3}
}

func (m *AllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllResponse.Unmarshal(m, b)
}
func (m *AllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllResponse.Marshal(b, m, deterministic)
}
func (m *AllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllResponse.Merge(m, src)
}
func (m *AllResponse) XXX_Size() int {
	return xxx_messageInfo_AllResponse.Size(m)
}
func (m *AllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllResponse proto.InternalMessageInfo

func (m *AllResponse) GetPortfolios() []*Portfolio {
	if m != nil {
		return m.Portfolios
	}
	return nil
}

type Portfolio struct {
	Uuid                                string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid                            string   `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	AssetClassTargetStocks              float32  `protobuf:"fixed32,3,opt,name=asset_class_target_stocks,json=assetClassTargetStocks,proto3" json:"asset_class_target_stocks,omitempty"`
	AssetClassTargetCash                float32  `protobuf:"fixed32,4,opt,name=asset_class_target_cash,json=assetClassTargetCash,proto3" json:"asset_class_target_cash,omitempty"`
	IndustryTargetInformationTechnology float32  `protobuf:"fixed32,5,opt,name=industry_target_information_technology,json=industryTargetInformationTechnology,proto3" json:"industry_target_information_technology,omitempty"`
	IndustryTargetFinancials            float32  `protobuf:"fixed32,6,opt,name=industry_target_financials,json=industryTargetFinancials,proto3" json:"industry_target_financials,omitempty"`
	IndustryTargetEnergy                float32  `protobuf:"fixed32,7,opt,name=industry_target_energy,json=industryTargetEnergy,proto3" json:"industry_target_energy,omitempty"`
	IndustryTargetHealthCare            float32  `protobuf:"fixed32,8,opt,name=industry_target_healthCare,json=industryTargetHealthCare,proto3" json:"industry_target_healthCare,omitempty"`
	IndustryTargetMaterials             float32  `protobuf:"fixed32,9,opt,name=industry_target_materials,json=industryTargetMaterials,proto3" json:"industry_target_materials,omitempty"`
	IndustryTargetUtilities             float32  `protobuf:"fixed32,10,opt,name=industry_target_utilities,json=industryTargetUtilities,proto3" json:"industry_target_utilities,omitempty"`
	IndustryTargetRealEstate            float32  `protobuf:"fixed32,11,opt,name=industry_target_real_estate,json=industryTargetRealEstate,proto3" json:"industry_target_real_estate,omitempty"`
	IndustryTargetConsumerDiscretionary float32  `protobuf:"fixed32,12,opt,name=industry_target_consumer_discretionary,json=industryTargetConsumerDiscretionary,proto3" json:"industry_target_consumer_discretionary,omitempty"`
	IndustryTargetConsumerStaples       float32  `protobuf:"fixed32,13,opt,name=industry_target_consumer_staples,json=industryTargetConsumerStaples,proto3" json:"industry_target_consumer_staples,omitempty"`
	IndustryTargetCommunicationServices float32  `protobuf:"fixed32,14,opt,name=industry_target_communication_services,json=industryTargetCommunicationServices,proto3" json:"industry_target_communication_services,omitempty"`
	IndustryTargetIndustrials           float32  `protobuf:"fixed32,15,opt,name=industry_target_industrials,json=industryTargetIndustrials,proto3" json:"industry_target_industrials,omitempty"`
	IndustryTargetMiscellaneous         float32  `protobuf:"fixed32,16,opt,name=industry_target_miscellaneous,json=industryTargetMiscellaneous,proto3" json:"industry_target_miscellaneous,omitempty"`
	XXX_NoUnkeyedLiteral                struct{} `json:"-"`
	XXX_unrecognized                    []byte   `json:"-"`
	XXX_sizecache                       int32    `json:"-"`
}

func (m *Portfolio) Reset()         { *m = Portfolio{} }
func (m *Portfolio) String() string { return proto.CompactTextString(m) }
func (*Portfolio) ProtoMessage()    {}
func (*Portfolio) Descriptor() ([]byte, []int) {
	return fileDescriptor_fac9e0553d4fb04e, []int{4}
}

func (m *Portfolio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Portfolio.Unmarshal(m, b)
}
func (m *Portfolio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Portfolio.Marshal(b, m, deterministic)
}
func (m *Portfolio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Portfolio.Merge(m, src)
}
func (m *Portfolio) XXX_Size() int {
	return xxx_messageInfo_Portfolio.Size(m)
}
func (m *Portfolio) XXX_DiscardUnknown() {
	xxx_messageInfo_Portfolio.DiscardUnknown(m)
}

var xxx_messageInfo_Portfolio proto.InternalMessageInfo

func (m *Portfolio) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Portfolio) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *Portfolio) GetAssetClassTargetStocks() float32 {
	if m != nil {
		return m.AssetClassTargetStocks
	}
	return 0
}

func (m *Portfolio) GetAssetClassTargetCash() float32 {
	if m != nil {
		return m.AssetClassTargetCash
	}
	return 0
}

func (m *Portfolio) GetIndustryTargetInformationTechnology() float32 {
	if m != nil {
		return m.IndustryTargetInformationTechnology
	}
	return 0
}

func (m *Portfolio) GetIndustryTargetFinancials() float32 {
	if m != nil {
		return m.IndustryTargetFinancials
	}
	return 0
}

func (m *Portfolio) GetIndustryTargetEnergy() float32 {
	if m != nil {
		return m.IndustryTargetEnergy
	}
	return 0
}

func (m *Portfolio) GetIndustryTargetHealthCare() float32 {
	if m != nil {
		return m.IndustryTargetHealthCare
	}
	return 0
}

func (m *Portfolio) GetIndustryTargetMaterials() float32 {
	if m != nil {
		return m.IndustryTargetMaterials
	}
	return 0
}

func (m *Portfolio) GetIndustryTargetUtilities() float32 {
	if m != nil {
		return m.IndustryTargetUtilities
	}
	return 0
}

func (m *Portfolio) GetIndustryTargetRealEstate() float32 {
	if m != nil {
		return m.IndustryTargetRealEstate
	}
	return 0
}

func (m *Portfolio) GetIndustryTargetConsumerDiscretionary() float32 {
	if m != nil {
		return m.IndustryTargetConsumerDiscretionary
	}
	return 0
}

func (m *Portfolio) GetIndustryTargetConsumerStaples() float32 {
	if m != nil {
		return m.IndustryTargetConsumerStaples
	}
	return 0
}

func (m *Portfolio) GetIndustryTargetCommunicationServices() float32 {
	if m != nil {
		return m.IndustryTargetCommunicationServices
	}
	return 0
}

func (m *Portfolio) GetIndustryTargetIndustrials() float32 {
	if m != nil {
		return m.IndustryTargetIndustrials
	}
	return 0
}

func (m *Portfolio) GetIndustryTargetMiscellaneous() float32 {
	if m != nil {
		return m.IndustryTargetMiscellaneous
	}
	return 0
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "ListRequest")
	proto.RegisterType((*ListResponse)(nil), "ListResponse")
	proto.RegisterType((*AllRequest)(nil), "AllRequest")
	proto.RegisterType((*AllResponse)(nil), "AllResponse")
	proto.RegisterType((*Portfolio)(nil), "Portfolio")
}

func init() { proto.RegisterFile("proto/portfolios.proto", fileDescriptor_fac9e0553d4fb04e) }

var fileDescriptor_fac9e0553d4fb04e = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5f, 0x6f, 0xd3, 0x30,
	0x10, 0x5f, 0xd7, 0xad, 0x5b, 0xae, 0x1d, 0x20, 0x6b, 0xda, 0xbc, 0x4e, 0x13, 0x55, 0x10, 0x68,
	0xe2, 0xa1, 0x48, 0xe3, 0x8f, 0xb4, 0x09, 0x90, 0xb6, 0x32, 0x06, 0x12, 0x48, 0x28, 0xdd, 0x9e,
	0x23, 0x93, 0xde, 0x5a, 0x0b, 0x37, 0x0e, 0x3e, 0x07, 0xa9, 0xdf, 0x8e, 0xef, 0xc2, 0x17, 0x41,
	0x71, 0x96, 0x2c, 0xcd, 0xc2, 0x1e, 0x78, 0xb3, 0xef, 0xf7, 0xcf, 0xf1, 0x39, 0x07, 0x3b, 0x89,
	0xd1, 0x56, 0xbf, 0x48, 0xb4, 0xb1, 0xd7, 0x5a, 0x49, 0x4d, 0x43, 0x57, 0xf0, 0xcf, 0xa0, 0xfb,
	0x45, 0x92, 0x0d, 0xf0, 0x67, 0x8a, 0x64, 0xd9, 0x01, 0x40, 0x4a, 0x68, 0xc2, 0x34, 0x95, 0x13,
	0xe2, 0xad, 0x41, 0xfb, 0xd0, 0x0b, 0xbc, 0xac, 0x72, 0x95, 0x15, 0xd8, 0x36, 0xac, 0xe7, 0xc8,
	0xaa, 0x43, 0xf2, 0x8d, 0x7f, 0x02, 0xbd, 0xdc, 0x83, 0x12, 0x1d, 0x13, 0xb2, 0xe7, 0x00, 0xb7,
	0x39, 0xce, 0xa4, 0x7b, 0x04, 0xc3, 0x6f, 0x45, 0x29, 0xa8, 0xa0, 0xfe, 0x1b, 0x80, 0x53, 0xa5,
	0x8a, 0x78, 0x06, 0x6b, 0x89, 0x98, 0x22, 0x6f, 0x0d, 0x5a, 0x87, 0xed, 0xc0, 0xad, 0xb3, 0x4c,
	0x25, 0xe7, 0xd2, 0xf2, 0x55, 0x57, 0xcc, 0x37, 0xfe, 0x31, 0x74, 0x9d, 0xee, 0x3f, 0x22, 0xff,
	0x6c, 0x80, 0x57, 0x22, 0x59, 0x64, 0xf6, 0x15, 0x2e, 0xd2, 0x0b, 0xdc, 0x9a, 0xed, 0x83, 0x57,
	0xde, 0x82, 0x8b, 0xf5, 0x82, 0xcd, 0xe2, 0x12, 0xd8, 0x31, 0xec, 0x09, 0x22, 0xb4, 0x61, 0xa4,
	0x04, 0x51, 0x68, 0x85, 0x99, 0xa2, 0x0d, 0xc9, 0xea, 0xe8, 0x07, 0xf1, 0xf6, 0xa0, 0x75, 0xb8,
	0x1a, 0xec, 0x38, 0xc2, 0x28, 0xc3, 0x2f, 0x1d, 0x3c, 0x76, 0x28, 0x7b, 0x0d, 0xbb, 0x0d, 0xd2,
	0x48, 0xd0, 0x8c, 0xaf, 0x39, 0xe1, 0x76, 0x5d, 0x38, 0x12, 0x34, 0x63, 0x63, 0x78, 0x26, 0xe3,
	0x49, 0x4a, 0xd6, 0x2c, 0x0a, 0x8d, 0x8c, 0xaf, 0xb5, 0x99, 0x0b, 0x2b, 0x75, 0x1c, 0x5a, 0x8c,
	0x66, 0xb1, 0x56, 0x7a, 0xba, 0xe0, 0xeb, 0xce, 0xe5, 0x49, 0xc1, 0xce, 0x3d, 0x3e, 0xdf, 0x72,
	0x2f, 0x4b, 0x2a, 0x7b, 0x0b, 0xfd, 0xba, 0xe9, 0xb5, 0x8c, 0x45, 0x1c, 0x49, 0xa1, 0x88, 0x77,
	0x9c, 0x11, 0x5f, 0x36, 0xfa, 0x58, 0xe2, 0xec, 0x15, 0xec, 0xd4, 0xd5, 0x18, 0xa3, 0x99, 0x2e,
	0xf8, 0x46, 0xfe, 0x21, 0xcb, 0xca, 0x73, 0x87, 0x35, 0x65, 0xce, 0x50, 0x28, 0x3b, 0x1b, 0x09,
	0x83, 0x7c, 0xb3, 0x29, 0xf3, 0x53, 0x89, 0xb3, 0x13, 0xd8, 0xab, 0xab, 0xe7, 0xc2, 0xa2, 0x71,
	0x07, 0xf6, 0x9c, 0x78, 0x77, 0x59, 0xfc, 0xb5, 0x80, 0x9b, 0xb4, 0xa9, 0x95, 0x4a, 0x5a, 0x89,
	0xc4, 0xa1, 0x49, 0x7b, 0x55, 0xc0, 0xec, 0x1d, 0xec, 0xd7, 0xb5, 0x06, 0x85, 0x0a, 0x91, 0xac,
	0xb0, 0xc8, 0xbb, 0x4d, 0xc7, 0x0e, 0x50, 0xa8, 0x73, 0x87, 0x37, 0x75, 0x2f, 0xd2, 0x31, 0xa5,
	0x73, 0x34, 0xe1, 0x44, 0x52, 0x64, 0x30, 0xeb, 0x8c, 0x30, 0x0b, 0xde, 0x6b, 0xea, 0xde, 0xe8,
	0x86, 0xfb, 0xa1, 0x4a, 0x65, 0x17, 0x30, 0xf8, 0xa7, 0x29, 0x59, 0x91, 0x28, 0x24, 0xbe, 0xe5,
	0xec, 0x0e, 0x9a, 0xed, 0xc6, 0x39, 0xa9, 0xf9, 0x74, 0xf3, 0x79, 0x1a, 0xcb, 0x28, 0x7f, 0x5d,
	0x84, 0xe6, 0x97, 0x8c, 0x90, 0xf8, 0x83, 0xe6, 0xd3, 0x55, 0xb8, 0xe3, 0x1b, 0x2a, 0x7b, 0x7f,
	0xf7, 0xc6, 0x6e, 0xf6, 0xae, 0x57, 0x0f, 0x9d, 0xd3, 0x5e, 0xfd, 0x95, 0x96, 0x04, 0x76, 0x06,
	0x07, 0x77, 0x3a, 0x2d, 0x29, 0x42, 0xa5, 0x44, 0x8c, 0x3a, 0x25, 0xfe, 0xc8, 0x39, 0xec, 0xd7,
	0xba, 0x5d, 0xa5, 0x1c, 0xfd, 0x6e, 0x01, 0x94, 0x7f, 0x39, 0x31, 0x1f, 0x3a, 0x23, 0x83, 0x59,
	0x3f, 0x2a, 0x63, 0xa1, 0x5f, 0x59, 0xfb, 0x2b, 0xec, 0x31, 0xb4, 0x2f, 0xd0, 0xde, 0x43, 0xf0,
	0xa1, 0x73, 0x95, 0x4c, 0xee, 0x37, 0xf1, 0xa1, 0x7d, 0xaa, 0x14, 0xeb, 0x0e, 0x6f, 0xc7, 0x5a,
	0xbf, 0x37, 0xac, 0xcc, 0x2a, 0x7f, 0x85, 0x3d, 0x85, 0xb5, 0x6c, 0x60, 0xb2, 0xde, 0xb0, 0x32,
	0x7b, 0xfb, 0x5b, 0xc3, 0xea, 0x14, 0xf5, 0x57, 0xbe, 0x77, 0xdc, 0x88, 0x7e, 0xf9, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x9e, 0x43, 0x08, 0x24, 0xbc, 0x05, 0x00, 0x00,
}
