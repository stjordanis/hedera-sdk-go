// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/TransactionReceipt.proto

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

// The consensus result for a transaction, which might not be currently known, or may  succeed or fail.
type TransactionReceipt struct {
	Status       ResponseCodeEnum `protobuf:"varint,1,opt,name=status,proto3,enum=proto.ResponseCodeEnum" json:"status,omitempty"`
	AccountID    *AccountID       `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	FileID       *FileID          `protobuf:"bytes,3,opt,name=fileID,proto3" json:"fileID,omitempty"`
	ContractID   *ContractID      `protobuf:"bytes,4,opt,name=contractID,proto3" json:"contractID,omitempty"`
	ExchangeRate *ExchangeRateSet `protobuf:"bytes,5,opt,name=exchangeRate,proto3" json:"exchangeRate,omitempty"`
	TopicID      *TopicID         `protobuf:"bytes,6,opt,name=topicID,proto3" json:"topicID,omitempty"`
	// Updated sequence number for a consensus service topic. The result of a ConsensusSubmitMessage.
	TopicSequenceNumber uint64 `protobuf:"varint,7,opt,name=topicSequenceNumber,proto3" json:"topicSequenceNumber,omitempty"`
	// Updated running hash for a consensus service topic. The result of a ConsensusSubmitMessage.
	TopicRunningHash     []byte   `protobuf:"bytes,8,opt,name=topicRunningHash,proto3" json:"topicRunningHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionReceipt) Reset()         { *m = TransactionReceipt{} }
func (m *TransactionReceipt) String() string { return proto.CompactTextString(m) }
func (*TransactionReceipt) ProtoMessage()    {}
func (*TransactionReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_01c1490b2e797be2, []int{0}
}

func (m *TransactionReceipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionReceipt.Unmarshal(m, b)
}
func (m *TransactionReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionReceipt.Marshal(b, m, deterministic)
}
func (m *TransactionReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionReceipt.Merge(m, src)
}
func (m *TransactionReceipt) XXX_Size() int {
	return xxx_messageInfo_TransactionReceipt.Size(m)
}
func (m *TransactionReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionReceipt proto.InternalMessageInfo

func (m *TransactionReceipt) GetStatus() ResponseCodeEnum {
	if m != nil {
		return m.Status
	}
	return ResponseCodeEnum_OK
}

func (m *TransactionReceipt) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *TransactionReceipt) GetFileID() *FileID {
	if m != nil {
		return m.FileID
	}
	return nil
}

func (m *TransactionReceipt) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

func (m *TransactionReceipt) GetExchangeRate() *ExchangeRateSet {
	if m != nil {
		return m.ExchangeRate
	}
	return nil
}

func (m *TransactionReceipt) GetTopicID() *TopicID {
	if m != nil {
		return m.TopicID
	}
	return nil
}

func (m *TransactionReceipt) GetTopicSequenceNumber() uint64 {
	if m != nil {
		return m.TopicSequenceNumber
	}
	return 0
}

func (m *TransactionReceipt) GetTopicRunningHash() []byte {
	if m != nil {
		return m.TopicRunningHash
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionReceipt)(nil), "proto.TransactionReceipt")
}

func init() { proto.RegisterFile("proto/TransactionReceipt.proto", fileDescriptor_01c1490b2e797be2) }

var fileDescriptor_01c1490b2e797be2 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x6f, 0xda, 0x30,
	0x14, 0xc7, 0x95, 0x01, 0x61, 0xf3, 0x18, 0x62, 0x9e, 0xc6, 0x2c, 0x0e, 0x53, 0x34, 0x69, 0x52,
	0x34, 0x89, 0x64, 0xa5, 0xb7, 0xde, 0xca, 0x8f, 0x0a, 0x2e, 0x55, 0x65, 0x38, 0xf5, 0x66, 0xcc,
	0x6b, 0x62, 0xb5, 0xb1, 0xd3, 0xd8, 0x91, 0xda, 0xbf, 0xa3, 0xff, 0x70, 0x55, 0xdb, 0xd0, 0x20,
	0x7a, 0x7a, 0xd1, 0xf7, 0xf3, 0xf9, 0xe6, 0xbd, 0x04, 0xfd, 0x2e, 0x2b, 0x65, 0x54, 0xba, 0xa9,
	0x98, 0xd4, 0x8c, 0x1b, 0xa1, 0x24, 0x05, 0x0e, 0xa2, 0x34, 0x89, 0x05, 0xb8, 0x63, 0xc7, 0x68,
	0xe8, 0xb4, 0x29, 0xd3, 0x82, 0x6f, 0x9e, 0x4b, 0xd0, 0x0e, 0x8f, 0x88, 0xcb, 0x29, 0xe8, 0x52,
	0x49, 0x0d, 0x33, 0xb5, 0x83, 0x63, 0xb2, 0x78, 0xe2, 0x39, 0x93, 0x19, 0x50, 0x66, 0xf6, 0xe4,
	0xa7, 0x5f, 0x29, 0x0a, 0xd0, 0x86, 0x15, 0xa5, 0x8b, 0xff, 0xbc, 0xb4, 0x10, 0x3e, 0x3d, 0x03,
	0xa7, 0x28, 0xd4, 0x86, 0x99, 0x5a, 0x93, 0x20, 0x0a, 0xe2, 0xfe, 0xe4, 0x97, 0xd3, 0x93, 0xe6,
	0xca, 0x85, 0xac, 0x0b, 0xea, 0x35, 0x9c, 0xa0, 0x2f, 0x8c, 0x73, 0x55, 0x4b, 0xb3, 0x9a, 0x93,
	0x4f, 0x51, 0x10, 0x7f, 0x9d, 0x0c, 0x7c, 0xe7, 0x72, 0x9f, 0xd3, 0x77, 0x05, 0xff, 0x45, 0xe1,
	0x9d, 0x78, 0x80, 0xd5, 0x9c, 0xb4, 0xac, 0xfc, 0xcd, 0xcb, 0x57, 0x36, 0xa4, 0x1e, 0xe2, 0x33,
	0x84, 0xb8, 0x92, 0xa6, 0x62, 0xfc, 0xed, 0xbd, 0x6d, 0xab, 0x7e, 0xf7, 0xea, 0xec, 0x00, 0x68,
	0x43, 0xc2, 0x17, 0xa8, 0x07, 0x8d, 0xcf, 0x27, 0x1d, 0x5b, 0x1a, 0xfa, 0x52, 0xf3, 0xcf, 0xac,
	0xc1, 0xd0, 0x23, 0x17, 0xc7, 0xa8, 0x6b, 0x54, 0x29, 0xf8, 0x6a, 0x4e, 0x42, 0x5b, 0xeb, 0xfb,
	0xda, 0xc6, 0xa5, 0x74, 0x8f, 0xf1, 0x7f, 0xf4, 0xc3, 0x3e, 0xae, 0xe1, 0xb1, 0x06, 0xc9, 0xe1,
	0xba, 0x2e, 0xb6, 0x50, 0x91, 0x6e, 0x14, 0xc4, 0x6d, 0xfa, 0x11, 0xc2, 0xff, 0xd0, 0xc0, 0xc6,
	0xb4, 0x96, 0x52, 0xc8, 0x6c, 0xc9, 0x74, 0x4e, 0x3e, 0x47, 0x41, 0xdc, 0xa3, 0x27, 0xf9, 0x74,
	0x89, 0x46, 0x5c, 0x15, 0x49, 0x0e, 0x3b, 0xa8, 0x58, 0x92, 0x33, 0x9d, 0x67, 0x15, 0x2b, 0x73,
	0x77, 0xcc, 0x4d, 0x70, 0x1b, 0x67, 0xc2, 0xe4, 0xf5, 0x36, 0xe1, 0xaa, 0x48, 0x0f, 0x34, 0x75,
	0xfa, 0x58, 0xef, 0xee, 0xc7, 0x99, 0x4a, 0xad, 0xbb, 0x0d, 0xed, 0x38, 0x7f, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xb3, 0x8c, 0xf6, 0x79, 0x72, 0x02, 0x00, 0x00,
}
