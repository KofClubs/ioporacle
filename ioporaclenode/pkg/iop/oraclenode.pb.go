// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: oraclenode.proto

package iop

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

type SendDealRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deal *Deal `protobuf:"bytes,1,opt,name=deal,proto3" json:"deal,omitempty"`
}

func (x *SendDealRequest) Reset() {
	*x = SendDealRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oraclenode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDealRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDealRequest) ProtoMessage() {}

func (x *SendDealRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oraclenode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDealRequest.ProtoReflect.Descriptor instead.
func (*SendDealRequest) Descriptor() ([]byte, []int) {
	return file_oraclenode_proto_rawDescGZIP(), []int{0}
}

func (x *SendDealRequest) GetDeal() *Deal {
	if x != nil {
		return x.Deal
	}
	return nil
}

type SendDealResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendDealResponse) Reset() {
	*x = SendDealResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oraclenode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDealResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDealResponse) ProtoMessage() {}

func (x *SendDealResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oraclenode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDealResponse.ProtoReflect.Descriptor instead.
func (*SendDealResponse) Descriptor() ([]byte, []int) {
	return file_oraclenode_proto_rawDescGZIP(), []int{1}
}

type ValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oraclenode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oraclenode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_oraclenode_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateRequest) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type Deal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index     uint32         `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Deal      *EncryptedDeal `protobuf:"bytes,2,opt,name=deal,proto3" json:"deal,omitempty"`
	Signature []byte         `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Deal) Reset() {
	*x = Deal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oraclenode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deal) ProtoMessage() {}

func (x *Deal) ProtoReflect() protoreflect.Message {
	mi := &file_oraclenode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deal.ProtoReflect.Descriptor instead.
func (*Deal) Descriptor() ([]byte, []int) {
	return file_oraclenode_proto_rawDescGZIP(), []int{3}
}

func (x *Deal) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Deal) GetDeal() *EncryptedDeal {
	if x != nil {
		return x.Deal
	}
	return nil
}

func (x *Deal) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type EncryptedDeal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DhKey     []byte `protobuf:"bytes,1,opt,name=dhKey,proto3" json:"dhKey,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Nonce     []byte `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Cipher    []byte `protobuf:"bytes,4,opt,name=cipher,proto3" json:"cipher,omitempty"`
}

func (x *EncryptedDeal) Reset() {
	*x = EncryptedDeal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oraclenode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedDeal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedDeal) ProtoMessage() {}

func (x *EncryptedDeal) ProtoReflect() protoreflect.Message {
	mi := &file_oraclenode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedDeal.ProtoReflect.Descriptor instead.
func (*EncryptedDeal) Descriptor() ([]byte, []int) {
	return file_oraclenode_proto_rawDescGZIP(), []int{4}
}

func (x *EncryptedDeal) GetDhKey() []byte {
	if x != nil {
		return x.DhKey
	}
	return nil
}

func (x *EncryptedDeal) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *EncryptedDeal) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *EncryptedDeal) GetCipher() []byte {
	if x != nil {
		return x.Cipher
	}
	return nil
}

type ValidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash        []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Valid       bool   `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	BlockNumber int64  `protobuf:"varint,3,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	Signature   []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oraclenode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oraclenode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_oraclenode_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateResponse) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *ValidateResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *ValidateResponse) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *ValidateResponse) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_oraclenode_proto protoreflect.FileDescriptor

var file_oraclenode_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x69, 0x6f, 0x70, 0x22, 0x30, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x44,
	0x65, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x65,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x69, 0x6f, 0x70, 0x2e, 0x44,
	0x65, 0x61, 0x6c, 0x52, 0x04, 0x64, 0x65, 0x61, 0x6c, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x6e,
	0x64, 0x44, 0x65, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a,
	0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x22, 0x62, 0x0a, 0x04, 0x44, 0x65, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x65, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x69, 0x6f, 0x70, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x44, 0x65, 0x61, 0x6c, 0x52, 0x04, 0x64, 0x65, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x71, 0x0a, 0x0d, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x65, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x68, 0x4b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x64, 0x68, 0x4b, 0x65, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x22, 0x7c, 0x0a, 0x10, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0xcd, 0x01, 0x0a, 0x0a, 0x4f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64,
	0x44, 0x65, 0x61, 0x6c, 0x12, 0x14, 0x2e, 0x69, 0x6f, 0x70, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x44,
	0x65, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x69, 0x6f, 0x70,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x69, 0x6f, 0x70,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x69, 0x6f, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0d, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x2e, 0x69, 0x6f, 0x70,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x69, 0x6f, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x69,
	0x6f, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oraclenode_proto_rawDescOnce sync.Once
	file_oraclenode_proto_rawDescData = file_oraclenode_proto_rawDesc
)

func file_oraclenode_proto_rawDescGZIP() []byte {
	file_oraclenode_proto_rawDescOnce.Do(func() {
		file_oraclenode_proto_rawDescData = protoimpl.X.CompressGZIP(file_oraclenode_proto_rawDescData)
	})
	return file_oraclenode_proto_rawDescData
}

var file_oraclenode_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_oraclenode_proto_goTypes = []interface{}{
	(*SendDealRequest)(nil),  // 0: iop.SendDealRequest
	(*SendDealResponse)(nil), // 1: iop.SendDealResponse
	(*ValidateRequest)(nil),  // 2: iop.ValidateRequest
	(*Deal)(nil),             // 3: iop.Deal
	(*EncryptedDeal)(nil),    // 4: iop.EncryptedDeal
	(*ValidateResponse)(nil), // 5: iop.ValidateResponse
}
var file_oraclenode_proto_depIdxs = []int32{
	3, // 0: iop.SendDealRequest.deal:type_name -> iop.Deal
	4, // 1: iop.Deal.deal:type_name -> iop.EncryptedDeal
	0, // 2: iop.OracleNode.SendDeal:input_type -> iop.SendDealRequest
	2, // 3: iop.OracleNode.ValidateTransaction:input_type -> iop.ValidateRequest
	2, // 4: iop.OracleNode.ValidateBlock:input_type -> iop.ValidateRequest
	1, // 5: iop.OracleNode.SendDeal:output_type -> iop.SendDealResponse
	5, // 6: iop.OracleNode.ValidateTransaction:output_type -> iop.ValidateResponse
	5, // 7: iop.OracleNode.ValidateBlock:output_type -> iop.ValidateResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_oraclenode_proto_init() }
func file_oraclenode_proto_init() {
	if File_oraclenode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oraclenode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDealRequest); i {
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
		file_oraclenode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDealResponse); i {
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
		file_oraclenode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequest); i {
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
		file_oraclenode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deal); i {
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
		file_oraclenode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedDeal); i {
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
		file_oraclenode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateResponse); i {
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
			RawDescriptor: file_oraclenode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oraclenode_proto_goTypes,
		DependencyIndexes: file_oraclenode_proto_depIdxs,
		MessageInfos:      file_oraclenode_proto_msgTypes,
	}.Build()
	File_oraclenode_proto = out.File
	file_oraclenode_proto_rawDesc = nil
	file_oraclenode_proto_goTypes = nil
	file_oraclenode_proto_depIdxs = nil
}
