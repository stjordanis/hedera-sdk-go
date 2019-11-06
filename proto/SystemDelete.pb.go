// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SystemDelete.proto

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

// Delete a file or smart contract - can only be done with a Hedera admin multisig. When it is deleted, it immediately disappears from the system as seen by the user, but is still stored internally until the expiration time, at which time it is truly and permanently deleted. Until that time, it can be undeleted by the Hedera admin multisig. When a smart contract is deleted, the cryptocurrency account within it continues to exist, and is not affected by the expiration time here.
type SystemDeleteTransactionBody struct {
	// Types that are valid to be assigned to Id:
	//	*SystemDeleteTransactionBody_FileID
	//	*SystemDeleteTransactionBody_ContractID
	Id                   isSystemDeleteTransactionBody_Id `protobuf_oneof:"id"`
	ExpirationTime       *TimestampSeconds                `protobuf:"bytes,3,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *SystemDeleteTransactionBody) Reset()         { *m = SystemDeleteTransactionBody{} }
func (m *SystemDeleteTransactionBody) String() string { return proto.CompactTextString(m) }
func (*SystemDeleteTransactionBody) ProtoMessage()    {}
func (*SystemDeleteTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_860c88b39ffc5dcc, []int{0}
}

func (m *SystemDeleteTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemDeleteTransactionBody.Unmarshal(m, b)
}
func (m *SystemDeleteTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemDeleteTransactionBody.Marshal(b, m, deterministic)
}
func (m *SystemDeleteTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemDeleteTransactionBody.Merge(m, src)
}
func (m *SystemDeleteTransactionBody) XXX_Size() int {
	return xxx_messageInfo_SystemDeleteTransactionBody.Size(m)
}
func (m *SystemDeleteTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemDeleteTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_SystemDeleteTransactionBody proto.InternalMessageInfo

type isSystemDeleteTransactionBody_Id interface {
	isSystemDeleteTransactionBody_Id()
}

type SystemDeleteTransactionBody_FileID struct {
	FileID *FileID `protobuf:"bytes,1,opt,name=fileID,proto3,oneof"`
}

type SystemDeleteTransactionBody_ContractID struct {
	ContractID *ContractID `protobuf:"bytes,2,opt,name=contractID,proto3,oneof"`
}

func (*SystemDeleteTransactionBody_FileID) isSystemDeleteTransactionBody_Id() {}

func (*SystemDeleteTransactionBody_ContractID) isSystemDeleteTransactionBody_Id() {}

func (m *SystemDeleteTransactionBody) GetId() isSystemDeleteTransactionBody_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SystemDeleteTransactionBody) GetFileID() *FileID {
	if x, ok := m.GetId().(*SystemDeleteTransactionBody_FileID); ok {
		return x.FileID
	}
	return nil
}

func (m *SystemDeleteTransactionBody) GetContractID() *ContractID {
	if x, ok := m.GetId().(*SystemDeleteTransactionBody_ContractID); ok {
		return x.ContractID
	}
	return nil
}

func (m *SystemDeleteTransactionBody) GetExpirationTime() *TimestampSeconds {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SystemDeleteTransactionBody) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SystemDeleteTransactionBody_FileID)(nil),
		(*SystemDeleteTransactionBody_ContractID)(nil),
	}
}

func init() {
	proto.RegisterType((*SystemDeleteTransactionBody)(nil), "proto.SystemDeleteTransactionBody")
}

func init() { proto.RegisterFile("SystemDelete.proto", fileDescriptor_860c88b39ffc5dcc) }

var fileDescriptor_860c88b39ffc5dcc = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x3b, 0x55, 0xbb, 0xb8, 0xe2, 0x5f, 0x36, 0x0e, 0x75, 0x23, 0x5d, 0xa8, 0xab, 0x59,
	0xd8, 0x07, 0x10, 0xe2, 0x20, 0xed, 0x4e, 0xa6, 0xf3, 0x02, 0xd7, 0xe4, 0x6a, 0x22, 0xcd, 0x0f,
	0x49, 0x10, 0xe7, 0xf5, 0x7c, 0x32, 0x69, 0x26, 0x94, 0xd2, 0x55, 0xc8, 0x39, 0xdf, 0x39, 0xf7,
	0x00, 0xdb, 0x0c, 0x31, 0x91, 0x69, 0x69, 0x4b, 0x89, 0x1a, 0x1f, 0x5c, 0x72, 0xec, 0x2c, 0x3f,
	0xf3, 0x6b, 0x8e, 0x51, 0x8b, 0x7e, 0xf0, 0x14, 0x47, 0x63, 0x7e, 0xd5, 0x6b, 0x43, 0x31, 0xa1,
	0xf1, 0xa3, 0xb0, 0xf8, 0xab, 0xe0, 0xee, 0xb0, 0xa0, 0x0f, 0x68, 0x23, 0x8a, 0xa4, 0x9d, 0xe5,
	0x4e, 0x0e, 0xec, 0x11, 0x66, 0x9f, 0x7a, 0x4b, 0xeb, 0xb6, 0xae, 0xee, 0xab, 0xa7, 0xf3, 0xe7,
	0x8b, 0x31, 0xd7, 0xbc, 0x65, 0x71, 0x35, 0xe9, 0x8a, 0xcd, 0x96, 0x00, 0xc2, 0xd9, 0x14, 0x50,
	0xa4, 0x75, 0x5b, 0x4f, 0x33, 0x7c, 0x53, 0xe0, 0xd7, 0xbd, 0xb1, 0x9a, 0x74, 0x07, 0x18, 0x7b,
	0x81, 0x4b, 0xfa, 0xf5, 0x3a, 0xe0, 0xee, 0xde, 0x6e, 0x5a, 0x7d, 0x92, 0x83, 0xb7, 0x25, 0xb8,
	0x5f, 0xbb, 0x21, 0xe1, 0xac, 0x8c, 0xdd, 0x11, 0xce, 0x4f, 0x61, 0xaa, 0x25, 0x7f, 0x80, 0x85,
	0x70, 0xa6, 0x51, 0x24, 0x29, 0xa0, 0xc2, 0xa8, 0xbe, 0x02, 0x7a, 0xd5, 0xa0, 0xd7, 0xa5, 0xe7,
	0x1b, 0x7f, 0xf0, 0xbd, 0xfa, 0x98, 0xe5, 0xdf, 0xf2, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xca, 0xcc,
	0x45, 0x9b, 0x33, 0x01, 0x00, 0x00,
}
