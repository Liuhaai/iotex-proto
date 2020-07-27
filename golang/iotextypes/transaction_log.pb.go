// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/types/transaction_log.proto

package iotextypes

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

type TransactionLog_Type int32

const (
	TransactionLog_IN_CONTRACT_TRANSFER       TransactionLog_Type = 0
	TransactionLog_BUCKET_WITHDRAW_AMOUNT     TransactionLog_Type = 1
	TransactionLog_BUCKET_CREATE_AMOUNT       TransactionLog_Type = 2
	TransactionLog_BUCKET_DEPOSIT_AMOUNT      TransactionLog_Type = 3
	TransactionLog_CANDIDATE_SELF_STAKE       TransactionLog_Type = 4
	TransactionLog_CANDIDATE_REGISTRATION_FEE TransactionLog_Type = 5
	TransactionLog_GAS_FEE                    TransactionLog_Type = 6
	TransactionLog_NATIVE_TRANSFER            TransactionLog_Type = 7
	TransactionLog_DEPOSIT_TO_REWARDING_FUND  TransactionLog_Type = 8
	TransactionLog_CLAIM_FROM_REWARDING_FUND  TransactionLog_Type = 9
)

var TransactionLog_Type_name = map[int32]string{
	0: "IN_CONTRACT_TRANSFER",
	1: "BUCKET_WITHDRAW_AMOUNT",
	2: "BUCKET_CREATE_AMOUNT",
	3: "BUCKET_DEPOSIT_AMOUNT",
	4: "CANDIDATE_SELF_STAKE",
	5: "CANDIDATE_REGISTRATION_FEE",
	6: "GAS_FEE",
	7: "NATIVE_TRANSFER",
	8: "DEPOSIT_TO_REWARDING_FUND",
	9: "CLAIM_FROM_REWARDING_FUND",
}

var TransactionLog_Type_value = map[string]int32{
	"IN_CONTRACT_TRANSFER":       0,
	"BUCKET_WITHDRAW_AMOUNT":     1,
	"BUCKET_CREATE_AMOUNT":       2,
	"BUCKET_DEPOSIT_AMOUNT":      3,
	"CANDIDATE_SELF_STAKE":       4,
	"CANDIDATE_REGISTRATION_FEE": 5,
	"GAS_FEE":                    6,
	"NATIVE_TRANSFER":            7,
	"DEPOSIT_TO_REWARDING_FUND":  8,
	"CLAIM_FROM_REWARDING_FUND":  9,
}

func (x TransactionLog_Type) String() string {
	return proto.EnumName(TransactionLog_Type_name, int32(x))
}

func (TransactionLog_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce17c7385702f675, []int{0, 0}
}

type TransactionLog_TokenSymbol int32

const (
	TransactionLog_IOTX TransactionLog_TokenSymbol = 0
)

var TransactionLog_TokenSymbol_name = map[int32]string{
	0: "IOTX",
}

var TransactionLog_TokenSymbol_value = map[string]int32{
	"IOTX": 0,
}

func (x TransactionLog_TokenSymbol) String() string {
	return proto.EnumName(TransactionLog_TokenSymbol_name, int32(x))
}

func (TransactionLog_TokenSymbol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce17c7385702f675, []int{0, 1}
}

type TransactionLog struct {
	ActionHash           []byte                        `protobuf:"bytes,1,opt,name=actionHash,proto3" json:"actionHash,omitempty"`
	NumTransactions      uint64                        `protobuf:"varint,2,opt,name=numTransactions,proto3" json:"numTransactions,omitempty"`
	Transactions         []*TransactionLog_Transaction `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *TransactionLog) Reset()         { *m = TransactionLog{} }
func (m *TransactionLog) String() string { return proto.CompactTextString(m) }
func (*TransactionLog) ProtoMessage()    {}
func (*TransactionLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce17c7385702f675, []int{0}
}

func (m *TransactionLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionLog.Unmarshal(m, b)
}
func (m *TransactionLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionLog.Marshal(b, m, deterministic)
}
func (m *TransactionLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionLog.Merge(m, src)
}
func (m *TransactionLog) XXX_Size() int {
	return xxx_messageInfo_TransactionLog.Size(m)
}
func (m *TransactionLog) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionLog.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionLog proto.InternalMessageInfo

func (m *TransactionLog) GetActionHash() []byte {
	if m != nil {
		return m.ActionHash
	}
	return nil
}

func (m *TransactionLog) GetNumTransactions() uint64 {
	if m != nil {
		return m.NumTransactions
	}
	return 0
}

func (m *TransactionLog) GetTransactions() []*TransactionLog_Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type TransactionLog_Transaction struct {
	Topic                []byte                     `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Amount               string                     `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Sender               string                     `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            string                     `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Symbol               TransactionLog_TokenSymbol `protobuf:"varint,5,opt,name=symbol,proto3,enum=iotextypes.TransactionLog_TokenSymbol" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TransactionLog_Transaction) Reset()         { *m = TransactionLog_Transaction{} }
func (m *TransactionLog_Transaction) String() string { return proto.CompactTextString(m) }
func (*TransactionLog_Transaction) ProtoMessage()    {}
func (*TransactionLog_Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce17c7385702f675, []int{0, 0}
}

func (m *TransactionLog_Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionLog_Transaction.Unmarshal(m, b)
}
func (m *TransactionLog_Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionLog_Transaction.Marshal(b, m, deterministic)
}
func (m *TransactionLog_Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionLog_Transaction.Merge(m, src)
}
func (m *TransactionLog_Transaction) XXX_Size() int {
	return xxx_messageInfo_TransactionLog_Transaction.Size(m)
}
func (m *TransactionLog_Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionLog_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionLog_Transaction proto.InternalMessageInfo

func (m *TransactionLog_Transaction) GetTopic() []byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *TransactionLog_Transaction) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *TransactionLog_Transaction) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *TransactionLog_Transaction) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *TransactionLog_Transaction) GetSymbol() TransactionLog_TokenSymbol {
	if m != nil {
		return m.Symbol
	}
	return TransactionLog_IOTX
}

type BlockTransactionLog struct {
	NumTransactions      uint64            `protobuf:"varint,1,opt,name=numTransactions,proto3" json:"numTransactions,omitempty"`
	TransactionLog       []*TransactionLog `protobuf:"bytes,2,rep,name=transactionLog,proto3" json:"transactionLog,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BlockTransactionLog) Reset()         { *m = BlockTransactionLog{} }
func (m *BlockTransactionLog) String() string { return proto.CompactTextString(m) }
func (*BlockTransactionLog) ProtoMessage()    {}
func (*BlockTransactionLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce17c7385702f675, []int{1}
}

func (m *BlockTransactionLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockTransactionLog.Unmarshal(m, b)
}
func (m *BlockTransactionLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockTransactionLog.Marshal(b, m, deterministic)
}
func (m *BlockTransactionLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockTransactionLog.Merge(m, src)
}
func (m *BlockTransactionLog) XXX_Size() int {
	return xxx_messageInfo_BlockTransactionLog.Size(m)
}
func (m *BlockTransactionLog) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockTransactionLog.DiscardUnknown(m)
}

var xxx_messageInfo_BlockTransactionLog proto.InternalMessageInfo

func (m *BlockTransactionLog) GetNumTransactions() uint64 {
	if m != nil {
		return m.NumTransactions
	}
	return 0
}

func (m *BlockTransactionLog) GetTransactionLog() []*TransactionLog {
	if m != nil {
		return m.TransactionLog
	}
	return nil
}

func init() {
	proto.RegisterEnum("iotextypes.TransactionLog_Type", TransactionLog_Type_name, TransactionLog_Type_value)
	proto.RegisterEnum("iotextypes.TransactionLog_TokenSymbol", TransactionLog_TokenSymbol_name, TransactionLog_TokenSymbol_value)
	proto.RegisterType((*TransactionLog)(nil), "iotextypes.TransactionLog")
	proto.RegisterType((*TransactionLog_Transaction)(nil), "iotextypes.TransactionLog.Transaction")
	proto.RegisterType((*BlockTransactionLog)(nil), "iotextypes.BlockTransactionLog")
}

func init() { proto.RegisterFile("proto/types/transaction_log.proto", fileDescriptor_ce17c7385702f675) }

var fileDescriptor_ce17c7385702f675 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5f, 0x6f, 0x93, 0x50,
	0x18, 0xc6, 0x47, 0xff, 0x6d, 0x7b, 0xbb, 0x74, 0xe4, 0x6c, 0x4e, 0xd6, 0xe8, 0x52, 0x7b, 0x61,
	0x7a, 0x23, 0x4d, 0x34, 0xde, 0x9a, 0x9c, 0xc2, 0x69, 0x87, 0x6b, 0x61, 0x39, 0x9c, 0x5a, 0x63,
	0x62, 0x48, 0x8b, 0x84, 0xe1, 0x5a, 0x0e, 0x81, 0xd3, 0xc4, 0x5e, 0x1a, 0x2f, 0xfd, 0x34, 0x7e,
	0x43, 0x53, 0xa0, 0x81, 0x36, 0xd3, 0xcb, 0xe7, 0xf9, 0x3d, 0x0f, 0xef, 0x1b, 0x78, 0x81, 0x57,
	0x51, 0xcc, 0x05, 0xef, 0x8b, 0x4d, 0xe4, 0x25, 0x7d, 0x11, 0xcf, 0xc3, 0x64, 0xee, 0x8a, 0x80,
	0x87, 0xce, 0x92, 0xfb, 0x6a, 0xca, 0x10, 0x04, 0x5c, 0x78, 0x3f, 0xd2, 0x44, 0xf7, 0x67, 0x1d,
	0x5a, 0xac, 0x48, 0x8d, 0xb9, 0x8f, 0x6e, 0x00, 0x32, 0x71, 0x3b, 0x4f, 0x1e, 0x14, 0xa9, 0x23,
	0xf5, 0xce, 0x68, 0xc9, 0x41, 0x3d, 0x38, 0x0f, 0xd7, 0xab, 0x52, 0x29, 0x51, 0x2a, 0x1d, 0xa9,
	0x57, 0xa3, 0x87, 0x36, 0xfa, 0x08, 0x67, 0xa2, 0x1c, 0xab, 0x76, 0xaa, 0xbd, 0xe6, 0xdb, 0xd7,
	0x6a, 0x31, 0x5f, 0xdd, 0x9f, 0x5d, 0x96, 0x74, 0xaf, 0xdb, 0xfe, 0x23, 0x41, 0xb3, 0x44, 0xd1,
	0x25, 0xd4, 0x05, 0x8f, 0x02, 0x37, 0x5f, 0x30, 0x13, 0xe8, 0x0a, 0x1a, 0xf3, 0x15, 0x5f, 0x87,
	0x22, 0x5d, 0xe9, 0x94, 0xe6, 0x6a, 0xeb, 0x27, 0x5e, 0xf8, 0xcd, 0x8b, 0x95, 0x6a, 0xe6, 0x67,
	0x0a, 0xbd, 0x80, 0xd3, 0xd8, 0x73, 0x83, 0x28, 0xf0, 0x42, 0xa1, 0xd4, 0x52, 0x54, 0x18, 0xe8,
	0x03, 0x34, 0x92, 0xcd, 0x6a, 0xc1, 0x97, 0x4a, 0xbd, 0x23, 0xf5, 0x5a, 0xff, 0xdf, 0x9c, 0x3f,
	0x7a, 0xa1, 0x9d, 0xa6, 0x69, 0xde, 0xea, 0xfe, 0xae, 0x40, 0x8d, 0x6d, 0x22, 0x0f, 0x29, 0x70,
	0x69, 0x98, 0x8e, 0x66, 0x99, 0x8c, 0x62, 0x8d, 0x39, 0x8c, 0x62, 0xd3, 0x1e, 0x12, 0x2a, 0x1f,
	0xa1, 0x36, 0x5c, 0x0d, 0xa6, 0xda, 0x1d, 0x61, 0xce, 0xcc, 0x60, 0xb7, 0x3a, 0xc5, 0x33, 0x07,
	0x4f, 0xac, 0xa9, 0xc9, 0x64, 0x69, 0xdb, 0xca, 0x99, 0x46, 0x09, 0x66, 0x64, 0x47, 0x2a, 0xe8,
	0x1a, 0x9e, 0xe5, 0x44, 0x27, 0xf7, 0x96, 0x6d, 0xb0, 0x1d, 0xaa, 0x6e, 0x4b, 0x1a, 0x36, 0x75,
	0x43, 0xdf, 0x16, 0x6c, 0x32, 0x1e, 0x3a, 0x36, 0xc3, 0x77, 0x44, 0xae, 0xa1, 0x1b, 0x68, 0x17,
	0x84, 0x92, 0x91, 0x61, 0x33, 0x8a, 0x99, 0x61, 0x99, 0xce, 0x90, 0x10, 0xb9, 0x8e, 0x9a, 0x70,
	0x3c, 0xc2, 0x76, 0x2a, 0x1a, 0xe8, 0x02, 0xce, 0x4d, 0xcc, 0x8c, 0x4f, 0xa4, 0x58, 0xf6, 0x18,
	0xbd, 0x84, 0xeb, 0xdd, 0x3c, 0x66, 0x39, 0x94, 0xcc, 0x30, 0xd5, 0x0d, 0x73, 0xe4, 0x0c, 0xa7,
	0xa6, 0x2e, 0x9f, 0x6c, 0xb1, 0x36, 0xc6, 0xc6, 0xc4, 0x19, 0x52, 0x6b, 0x72, 0x88, 0x4f, 0xbb,
	0xcf, 0xa1, 0x59, 0x7a, 0x49, 0xe8, 0x04, 0x6a, 0x86, 0xc5, 0x3e, 0xcb, 0x47, 0xdd, 0x5f, 0x12,
	0x5c, 0x0c, 0x96, 0xdc, 0x7d, 0x3c, 0x38, 0xc4, 0x27, 0x0e, 0x4d, 0x7a, 0xfa, 0xd0, 0x06, 0xd0,
	0x12, 0x7b, 0x5d, 0xa5, 0x92, 0x9e, 0x5a, 0xfb, 0xdf, 0x1f, 0x8c, 0x1e, 0x34, 0x06, 0x5f, 0xa1,
	0xeb, 0xf2, 0x95, 0xea, 0x07, 0xe2, 0x61, 0xbd, 0xc8, 0x7a, 0x51, 0xcc, 0xbf, 0x7b, 0xae, 0x50,
	0xfd, 0x38, 0x72, 0xd5, 0xf4, 0x21, 0xf7, 0xd2, 0x97, 0xf7, 0x79, 0xc2, 0xe5, 0xab, 0x7e, 0x39,
	0x95, 0x89, 0x37, 0xd9, 0xef, 0xe7, 0xf3, 0xe5, 0x3c, 0xf4, 0xfb, 0xc5, 0xf4, 0x45, 0x23, 0x05,
	0xef, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x1d, 0x67, 0x81, 0xa0, 0x03, 0x00, 0x00,
}
