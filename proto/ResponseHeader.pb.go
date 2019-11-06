// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ResponseHeader.proto

package proto

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

// Every query receives a response containing the QueryResponseHeader. Either or both of the cost and stateProof fields may be blank, if the responseType didn't ask for the cost or stateProof.
type ResponseHeader struct {
	NodeTransactionPrecheckCode ResponseCodeEnum `protobuf:"varint,1,opt,name=nodeTransactionPrecheckCode,proto3,enum=proto.ResponseCodeEnum" json:"nodeTransactionPrecheckCode,omitempty"`
	ResponseType                ResponseType     `protobuf:"varint,2,opt,name=responseType,proto3,enum=proto.ResponseType" json:"responseType,omitempty"`
	Cost                        uint64           `protobuf:"varint,3,opt,name=cost,proto3" json:"cost,omitempty"`
	StateProof                  []byte           `protobuf:"bytes,4,opt,name=stateProof,proto3" json:"stateProof,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}         `json:"-"`
	XXX_unrecognized            []byte           `json:"-"`
	XXX_sizecache               int32            `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_b707db0c2015df73, []int{0}
}

func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetNodeTransactionPrecheckCode() ResponseCodeEnum {
	if m != nil {
		return m.NodeTransactionPrecheckCode
	}
	return ResponseCodeEnum_OK
}

func (m *ResponseHeader) GetResponseType() ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return ResponseType_ANSWER_ONLY
}

func (m *ResponseHeader) GetCost() uint64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *ResponseHeader) GetStateProof() []byte {
	if m != nil {
		return m.StateProof
	}
	return nil
}

func init() {
	proto.RegisterType((*ResponseHeader)(nil), "proto.ResponseHeader")
}

func init() { proto.RegisterFile("ResponseHeader.proto", fileDescriptor_b707db0c2015df73) }

var fileDescriptor_b707db0c2015df73 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8d, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x89, 0xae, 0x1e, 0x86, 0x52, 0x30, 0x0a, 0x2e, 0x15, 0x64, 0xe9, 0x41, 0xf6, 0xb4,
	0x07, 0x3d, 0x78, 0x57, 0x04, 0x8f, 0x6b, 0xe8, 0xc5, 0xe3, 0x98, 0x8c, 0xa6, 0x4a, 0x33, 0x61,
	0x92, 0x0a, 0xfd, 0xa7, 0xfe, 0x1c, 0x69, 0x5a, 0x61, 0xd7, 0x43, 0x4f, 0x49, 0xde, 0x7b, 0xdf,
	0x17, 0xb8, 0x30, 0x94, 0x22, 0x87, 0x44, 0xcf, 0x84, 0x8e, 0xa4, 0x8b, 0xc2, 0x99, 0xf5, 0x49,
	0x39, 0x66, 0x67, 0x2f, 0x6b, 0x92, 0xcd, 0xb0, 0x99, 0xe9, 0xbf, 0xfd, 0x23, 0x3b, 0xda, 0x65,
	0xf3, 0x1f, 0x05, 0xd3, 0xb1, 0x46, 0xbf, 0xc2, 0x55, 0x60, 0x47, 0x0b, 0xc1, 0x90, 0xd0, 0xe6,
	0x25, 0x87, 0x5e, 0xc8, 0x7a, 0xb2, 0x5f, 0x5b, 0xae, 0x56, 0x8d, 0x6a, 0xa7, 0xb7, 0x97, 0x3b,
	0xbe, 0x1b, 0x2a, 0x9f, 0xc2, 0x7a, 0x65, 0x0e, 0xb1, 0xfa, 0x1e, 0x26, 0xb2, 0x07, 0x16, 0x9b,
	0x48, 0xf5, 0x51, 0x71, 0x9d, 0xff, 0x73, 0x6d, 0x2b, 0x33, 0x1a, 0x6a, 0x0d, 0x95, 0xe5, 0x94,
	0xeb, 0xe3, 0x46, 0xb5, 0x95, 0x29, 0x77, 0x7d, 0x0d, 0x90, 0x32, 0x66, 0xea, 0x85, 0xf9, 0xbd,
	0xae, 0x1a, 0xd5, 0x4e, 0xcc, 0x20, 0x79, 0xb8, 0x81, 0xb9, 0xe5, 0x55, 0xe7, 0xc9, 0x91, 0xa0,
	0xc7, 0xe4, 0x3f, 0x04, 0xa3, 0xef, 0x30, 0x2e, 0xf7, 0xff, 0x7d, 0xe2, 0x37, 0xf6, 0xea, 0xed,
	0xb4, 0xbc, 0xee, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x67, 0x8e, 0x5c, 0x93, 0x4f, 0x01, 0x00,
	0x00,
}
