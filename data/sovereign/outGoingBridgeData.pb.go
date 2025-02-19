// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: outGoingBridgeData.proto

package sovereign

import (
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

type BridgeOperations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*BridgeOutGoingData `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *BridgeOperations) Reset() {
	*x = BridgeOperations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outGoingBridgeData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BridgeOperations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BridgeOperations) ProtoMessage() {}

func (x *BridgeOperations) ProtoReflect() protoreflect.Message {
	mi := &file_outGoingBridgeData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BridgeOperations.ProtoReflect.Descriptor instead.
func (*BridgeOperations) Descriptor() ([]byte, []int) {
	return file_outGoingBridgeData_proto_rawDescGZIP(), []int{0}
}

func (x *BridgeOperations) GetData() []*BridgeOutGoingData {
	if x != nil {
		return x.Data
	}
	return nil
}

type BridgeOutGoingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                int32                `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Hash                []byte               `protobuf:"bytes,2,opt,name=Hash,proto3" json:"Hash,omitempty"`
	OutGoingOperations  []*OutGoingOperation `protobuf:"bytes,3,rep,name=OutGoingOperations,proto3" json:"OutGoingOperations,omitempty"`
	AggregatedSignature []byte               `protobuf:"bytes,4,opt,name=AggregatedSignature,proto3" json:"AggregatedSignature,omitempty"`
	LeaderSignature     []byte               `protobuf:"bytes,5,opt,name=LeaderSignature,proto3" json:"LeaderSignature,omitempty"`
	PubKeysBitmap       []byte               `protobuf:"bytes,6,opt,name=PubKeysBitmap,proto3" json:"PubKeysBitmap,omitempty"`
	Epoch               uint32               `protobuf:"varint,7,opt,name=Epoch,proto3" json:"Epoch,omitempty"`
}

func (x *BridgeOutGoingData) Reset() {
	*x = BridgeOutGoingData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outGoingBridgeData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BridgeOutGoingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BridgeOutGoingData) ProtoMessage() {}

func (x *BridgeOutGoingData) ProtoReflect() protoreflect.Message {
	mi := &file_outGoingBridgeData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BridgeOutGoingData.ProtoReflect.Descriptor instead.
func (*BridgeOutGoingData) Descriptor() ([]byte, []int) {
	return file_outGoingBridgeData_proto_rawDescGZIP(), []int{1}
}

func (x *BridgeOutGoingData) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *BridgeOutGoingData) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *BridgeOutGoingData) GetOutGoingOperations() []*OutGoingOperation {
	if x != nil {
		return x.OutGoingOperations
	}
	return nil
}

func (x *BridgeOutGoingData) GetAggregatedSignature() []byte {
	if x != nil {
		return x.AggregatedSignature
	}
	return nil
}

func (x *BridgeOutGoingData) GetLeaderSignature() []byte {
	if x != nil {
		return x.LeaderSignature
	}
	return nil
}

func (x *BridgeOutGoingData) GetPubKeysBitmap() []byte {
	if x != nil {
		return x.PubKeysBitmap
	}
	return nil
}

func (x *BridgeOutGoingData) GetEpoch() uint32 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

type OutGoingOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *OutGoingOperation) Reset() {
	*x = OutGoingOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outGoingBridgeData_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutGoingOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutGoingOperation) ProtoMessage() {}

func (x *OutGoingOperation) ProtoReflect() protoreflect.Message {
	mi := &file_outGoingBridgeData_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutGoingOperation.ProtoReflect.Descriptor instead.
func (*OutGoingOperation) Descriptor() ([]byte, []int) {
	return file_outGoingBridgeData_proto_rawDescGZIP(), []int{2}
}

func (x *OutGoingOperation) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *OutGoingOperation) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BridgeOutGoingDataValidatorSetChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKeyIDs [][]byte `protobuf:"bytes,1,rep,name=PubKeyIDs,proto3" json:"PubKeyIDs,omitempty"`
}

func (x *BridgeOutGoingDataValidatorSetChange) Reset() {
	*x = BridgeOutGoingDataValidatorSetChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outGoingBridgeData_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BridgeOutGoingDataValidatorSetChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BridgeOutGoingDataValidatorSetChange) ProtoMessage() {}

func (x *BridgeOutGoingDataValidatorSetChange) ProtoReflect() protoreflect.Message {
	mi := &file_outGoingBridgeData_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BridgeOutGoingDataValidatorSetChange.ProtoReflect.Descriptor instead.
func (*BridgeOutGoingDataValidatorSetChange) Descriptor() ([]byte, []int) {
	return file_outGoingBridgeData_proto_rawDescGZIP(), []int{3}
}

func (x *BridgeOutGoingDataValidatorSetChange) GetPubKeyIDs() [][]byte {
	if x != nil {
		return x.PubKeyIDs
	}
	return nil
}

type BridgeOperationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHashes []string `protobuf:"bytes,1,rep,name=TxHashes,proto3" json:"TxHashes,omitempty"`
}

func (x *BridgeOperationsResponse) Reset() {
	*x = BridgeOperationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outGoingBridgeData_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BridgeOperationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BridgeOperationsResponse) ProtoMessage() {}

func (x *BridgeOperationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_outGoingBridgeData_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BridgeOperationsResponse.ProtoReflect.Descriptor instead.
func (*BridgeOperationsResponse) Descriptor() ([]byte, []int) {
	return file_outGoingBridgeData_proto_rawDescGZIP(), []int{4}
}

func (x *BridgeOperationsResponse) GetTxHashes() []string {
	if x != nil {
		return x.TxHashes
	}
	return nil
}

var File_outGoingBridgeData_proto protoreflect.FileDescriptor

var file_outGoingBridgeData_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6f, 0x75, 0x74, 0x47, 0x6f, 0x69, 0x6e, 0x67, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x6f, 0x76, 0x65,
	0x72, 0x65, 0x69, 0x67, 0x6e, 0x22, 0x45, 0x0a, 0x10, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x6f, 0x76, 0x65, 0x72, 0x65,
	0x69, 0x67, 0x6e, 0x2e, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x47, 0x6f, 0x69,
	0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0xa2, 0x02, 0x0a,
	0x12, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x47, 0x6f, 0x69, 0x6e, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x4c, 0x0a, 0x12, 0x4f,
	0x75, 0x74, 0x47, 0x6f, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6f, 0x76, 0x65, 0x72, 0x65,
	0x69, 0x67, 0x6e, 0x2e, 0x4f, 0x75, 0x74, 0x47, 0x6f, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x4f, 0x75, 0x74, 0x47, 0x6f, 0x69, 0x6e, 0x67, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x73,
	0x42, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x50, 0x75,
	0x62, 0x4b, 0x65, 0x79, 0x73, 0x42, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x70, 0x6f, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x45, 0x70, 0x6f, 0x63,
	0x68, 0x22, 0x3b, 0x0a, 0x11, 0x4f, 0x75, 0x74, 0x47, 0x6f, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x44,
	0x0a, 0x24, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x47, 0x6f, 0x69, 0x6e, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79,
	0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x50, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x49, 0x44, 0x73, 0x22, 0x36, 0x0a, 0x18, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x32, 0x5c, 0x0a, 0x0e,
	0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x54, 0x78, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x4a,
	0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1b, 0x2e, 0x73, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x69,
	0x67, 0x6e, 0x2e, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x23, 0x2e, 0x73, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x2e,
	0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x78, 0x2f, 0x6d, 0x78, 0x2d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2d, 0x63, 0x6f, 0x72,
	0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73, 0x6f, 0x76, 0x65, 0x72, 0x65,
	0x69, 0x67, 0x6e, 0x3b, 0x73, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_outGoingBridgeData_proto_rawDescOnce sync.Once
	file_outGoingBridgeData_proto_rawDescData = file_outGoingBridgeData_proto_rawDesc
)

func file_outGoingBridgeData_proto_rawDescGZIP() []byte {
	file_outGoingBridgeData_proto_rawDescOnce.Do(func() {
		file_outGoingBridgeData_proto_rawDescData = protoimpl.X.CompressGZIP(file_outGoingBridgeData_proto_rawDescData)
	})
	return file_outGoingBridgeData_proto_rawDescData
}

var file_outGoingBridgeData_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_outGoingBridgeData_proto_goTypes = []interface{}{
	(*BridgeOperations)(nil),                     // 0: sovereign.BridgeOperations
	(*BridgeOutGoingData)(nil),                   // 1: sovereign.BridgeOutGoingData
	(*OutGoingOperation)(nil),                    // 2: sovereign.OutGoingOperation
	(*BridgeOutGoingDataValidatorSetChange)(nil), // 3: sovereign.BridgeOutGoingDataValidatorSetChange
	(*BridgeOperationsResponse)(nil),             // 4: sovereign.BridgeOperationsResponse
}
var file_outGoingBridgeData_proto_depIdxs = []int32{
	1, // 0: sovereign.BridgeOperations.Data:type_name -> sovereign.BridgeOutGoingData
	2, // 1: sovereign.BridgeOutGoingData.OutGoingOperations:type_name -> sovereign.OutGoingOperation
	0, // 2: sovereign.BridgeTxSender.Send:input_type -> sovereign.BridgeOperations
	4, // 3: sovereign.BridgeTxSender.Send:output_type -> sovereign.BridgeOperationsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_outGoingBridgeData_proto_init() }
func file_outGoingBridgeData_proto_init() {
	if File_outGoingBridgeData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_outGoingBridgeData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BridgeOperations); i {
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
		file_outGoingBridgeData_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BridgeOutGoingData); i {
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
		file_outGoingBridgeData_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutGoingOperation); i {
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
		file_outGoingBridgeData_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BridgeOutGoingDataValidatorSetChange); i {
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
		file_outGoingBridgeData_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BridgeOperationsResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_outGoingBridgeData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_outGoingBridgeData_proto_goTypes,
		DependencyIndexes: file_outGoingBridgeData_proto_depIdxs,
		MessageInfos:      file_outGoingBridgeData_proto_msgTypes,
	}.Build()
	File_outGoingBridgeData_proto = out.File
	file_outGoingBridgeData_proto_rawDesc = nil
	file_outGoingBridgeData_proto_goTypes = nil
	file_outGoingBridgeData_proto_depIdxs = nil
}
