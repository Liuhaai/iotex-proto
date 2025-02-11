// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

// To compile the proto, run:
//      protoc --go_out=plugins=grpc:$GOPATH/src *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/types/consensus.proto

package iotextypes

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ConsensusVote_Topic int32

const (
	ConsensusVote_PROPOSAL ConsensusVote_Topic = 0
	ConsensusVote_LOCK     ConsensusVote_Topic = 1
	ConsensusVote_COMMIT   ConsensusVote_Topic = 2
)

// Enum value maps for ConsensusVote_Topic.
var (
	ConsensusVote_Topic_name = map[int32]string{
		0: "PROPOSAL",
		1: "LOCK",
		2: "COMMIT",
	}
	ConsensusVote_Topic_value = map[string]int32{
		"PROPOSAL": 0,
		"LOCK":     1,
		"COMMIT":   2,
	}
)

func (x ConsensusVote_Topic) Enum() *ConsensusVote_Topic {
	p := new(ConsensusVote_Topic)
	*p = x
	return p
}

func (x ConsensusVote_Topic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConsensusVote_Topic) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_consensus_proto_enumTypes[0].Descriptor()
}

func (ConsensusVote_Topic) Type() protoreflect.EnumType {
	return &file_proto_types_consensus_proto_enumTypes[0]
}

func (x ConsensusVote_Topic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConsensusVote_Topic.Descriptor instead.
func (ConsensusVote_Topic) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_consensus_proto_rawDescGZIP(), []int{1, 0}
}

type BlockProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block        *Block         `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Endorsements []*Endorsement `protobuf:"bytes,2,rep,name=endorsements,proto3" json:"endorsements,omitempty"`
}

func (x *BlockProposal) Reset() {
	*x = BlockProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_consensus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockProposal) ProtoMessage() {}

func (x *BlockProposal) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_consensus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockProposal.ProtoReflect.Descriptor instead.
func (*BlockProposal) Descriptor() ([]byte, []int) {
	return file_proto_types_consensus_proto_rawDescGZIP(), []int{0}
}

func (x *BlockProposal) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *BlockProposal) GetEndorsements() []*Endorsement {
	if x != nil {
		return x.Endorsements
	}
	return nil
}

type ConsensusVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash []byte              `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	Topic     ConsensusVote_Topic `protobuf:"varint,2,opt,name=topic,proto3,enum=iotextypes.ConsensusVote_Topic" json:"topic,omitempty"`
}

func (x *ConsensusVote) Reset() {
	*x = ConsensusVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_consensus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusVote) ProtoMessage() {}

func (x *ConsensusVote) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_consensus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusVote.ProtoReflect.Descriptor instead.
func (*ConsensusVote) Descriptor() ([]byte, []int) {
	return file_proto_types_consensus_proto_rawDescGZIP(), []int{1}
}

func (x *ConsensusVote) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *ConsensusVote) GetTopic() ConsensusVote_Topic {
	if x != nil {
		return x.Topic
	}
	return ConsensusVote_PROPOSAL
}

type ConsensusMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height      uint64       `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Endorsement *Endorsement `protobuf:"bytes,2,opt,name=endorsement,proto3" json:"endorsement,omitempty"`
	// Types that are assignable to Msg:
	//	*ConsensusMessage_BlockProposal
	//	*ConsensusMessage_Vote
	Msg isConsensusMessage_Msg `protobuf_oneof:"msg"`
}

func (x *ConsensusMessage) Reset() {
	*x = ConsensusMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_consensus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusMessage) ProtoMessage() {}

func (x *ConsensusMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_consensus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusMessage.ProtoReflect.Descriptor instead.
func (*ConsensusMessage) Descriptor() ([]byte, []int) {
	return file_proto_types_consensus_proto_rawDescGZIP(), []int{2}
}

func (x *ConsensusMessage) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ConsensusMessage) GetEndorsement() *Endorsement {
	if x != nil {
		return x.Endorsement
	}
	return nil
}

func (m *ConsensusMessage) GetMsg() isConsensusMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *ConsensusMessage) GetBlockProposal() *BlockProposal {
	if x, ok := x.GetMsg().(*ConsensusMessage_BlockProposal); ok {
		return x.BlockProposal
	}
	return nil
}

func (x *ConsensusMessage) GetVote() *ConsensusVote {
	if x, ok := x.GetMsg().(*ConsensusMessage_Vote); ok {
		return x.Vote
	}
	return nil
}

type isConsensusMessage_Msg interface {
	isConsensusMessage_Msg()
}

type ConsensusMessage_BlockProposal struct {
	BlockProposal *BlockProposal `protobuf:"bytes,100,opt,name=blockProposal,proto3,oneof"`
}

type ConsensusMessage_Vote struct {
	Vote *ConsensusVote `protobuf:"bytes,101,opt,name=vote,proto3,oneof"`
}

func (*ConsensusMessage_BlockProposal) isConsensusMessage_Msg() {}

func (*ConsensusMessage_Vote) isConsensusMessage_Msg() {}

var File_proto_types_consensus_proto protoreflect.FileDescriptor

var file_proto_types_consensus_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69,
	0x6f, 0x74, 0x65, 0x78, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x3b, 0x0a, 0x0c, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0c, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x91, 0x01,
	0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x56, 0x6f, 0x74, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x35, 0x0a,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x69,
	0x6f, 0x74, 0x65, 0x78, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x56, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x22, 0x2b, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0c, 0x0a,
	0x08, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c,
	0x4f, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10,
	0x02, 0x22, 0xe0, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x39,
	0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x0d, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0d, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x04,
	0x76, 0x6f, 0x74, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6f, 0x74,
	0x65, 0x78, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x56, 0x6f, 0x74, 0x65, 0x48, 0x00, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x42, 0x05, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x42, 0x5d, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_consensus_proto_rawDescOnce sync.Once
	file_proto_types_consensus_proto_rawDescData = file_proto_types_consensus_proto_rawDesc
)

func file_proto_types_consensus_proto_rawDescGZIP() []byte {
	file_proto_types_consensus_proto_rawDescOnce.Do(func() {
		file_proto_types_consensus_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_consensus_proto_rawDescData)
	})
	return file_proto_types_consensus_proto_rawDescData
}

var file_proto_types_consensus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_types_consensus_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_types_consensus_proto_goTypes = []interface{}{
	(ConsensusVote_Topic)(0), // 0: iotextypes.ConsensusVote.Topic
	(*BlockProposal)(nil),    // 1: iotextypes.BlockProposal
	(*ConsensusVote)(nil),    // 2: iotextypes.ConsensusVote
	(*ConsensusMessage)(nil), // 3: iotextypes.ConsensusMessage
	(*Block)(nil),            // 4: iotextypes.Block
	(*Endorsement)(nil),      // 5: iotextypes.Endorsement
}
var file_proto_types_consensus_proto_depIdxs = []int32{
	4, // 0: iotextypes.BlockProposal.block:type_name -> iotextypes.Block
	5, // 1: iotextypes.BlockProposal.endorsements:type_name -> iotextypes.Endorsement
	0, // 2: iotextypes.ConsensusVote.topic:type_name -> iotextypes.ConsensusVote.Topic
	5, // 3: iotextypes.ConsensusMessage.endorsement:type_name -> iotextypes.Endorsement
	1, // 4: iotextypes.ConsensusMessage.blockProposal:type_name -> iotextypes.BlockProposal
	2, // 5: iotextypes.ConsensusMessage.vote:type_name -> iotextypes.ConsensusVote
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_types_consensus_proto_init() }
func file_proto_types_consensus_proto_init() {
	if File_proto_types_consensus_proto != nil {
		return
	}
	file_proto_types_blockchain_proto_init()
	file_proto_types_endorsement_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_types_consensus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockProposal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_types_consensus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusVote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_types_consensus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_types_consensus_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ConsensusMessage_BlockProposal)(nil),
		(*ConsensusMessage_Vote)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_types_consensus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_consensus_proto_goTypes,
		DependencyIndexes: file_proto_types_consensus_proto_depIdxs,
		EnumInfos:         file_proto_types_consensus_proto_enumTypes,
		MessageInfos:      file_proto_types_consensus_proto_msgTypes,
	}.Build()
	File_proto_types_consensus_proto = out.File
	file_proto_types_consensus_proto_rawDesc = nil
	file_proto_types_consensus_proto_goTypes = nil
	file_proto_types_consensus_proto_depIdxs = nil
}
