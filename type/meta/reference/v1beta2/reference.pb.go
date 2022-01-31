// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: type/meta/reference/v1beta2/reference.proto

package reference

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

type ReferenceType int32

const (
	ReferenceType_ReferenceType_UNSPECIFIED ReferenceType = 0
	ReferenceType_Internal                  ReferenceType = 1
	ReferenceType_External                  ReferenceType = 2
	ReferenceType_Owner                     ReferenceType = 3
)

// Enum value maps for ReferenceType.
var (
	ReferenceType_name = map[int32]string{
		0: "ReferenceType_UNSPECIFIED",
		1: "Internal",
		2: "External",
		3: "Owner",
	}
	ReferenceType_value = map[string]int32{
		"ReferenceType_UNSPECIFIED": 0,
		"Internal":                  1,
		"External":                  2,
		"Owner":                     3,
	}
)

func (x ReferenceType) Enum() *ReferenceType {
	p := new(ReferenceType)
	*p = x
	return p
}

func (x ReferenceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReferenceType) Descriptor() protoreflect.EnumDescriptor {
	return file_type_meta_reference_v1beta2_reference_proto_enumTypes[0].Descriptor()
}

func (ReferenceType) Type() protoreflect.EnumType {
	return &file_type_meta_reference_v1beta2_reference_proto_enumTypes[0]
}

func (x ReferenceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReferenceType.Descriptor instead.
func (ReferenceType) EnumDescriptor() ([]byte, []int) {
	return file_type_meta_reference_v1beta2_reference_proto_rawDescGZIP(), []int{0}
}

type ReferenceMatchType int32

const (
	ReferenceMatchType_ReferenceMatchType_UNSPECIFIED ReferenceMatchType = 0
	ReferenceMatchType_ALL                            ReferenceMatchType = 1
	ReferenceMatchType_ANY                            ReferenceMatchType = 2
)

// Enum value maps for ReferenceMatchType.
var (
	ReferenceMatchType_name = map[int32]string{
		0: "ReferenceMatchType_UNSPECIFIED",
		1: "ALL",
		2: "ANY",
	}
	ReferenceMatchType_value = map[string]int32{
		"ReferenceMatchType_UNSPECIFIED": 0,
		"ALL":                            1,
		"ANY":                            2,
	}
)

func (x ReferenceMatchType) Enum() *ReferenceMatchType {
	p := new(ReferenceMatchType)
	*p = x
	return p
}

func (x ReferenceMatchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReferenceMatchType) Descriptor() protoreflect.EnumDescriptor {
	return file_type_meta_reference_v1beta2_reference_proto_enumTypes[1].Descriptor()
}

func (ReferenceMatchType) Type() protoreflect.EnumType {
	return &file_type_meta_reference_v1beta2_reference_proto_enumTypes[1]
}

func (x ReferenceMatchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReferenceMatchType.Descriptor instead.
func (ReferenceMatchType) EnumDescriptor() ([]byte, []int) {
	return file_type_meta_reference_v1beta2_reference_proto_rawDescGZIP(), []int{1}
}

type SingleReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ReferenceType `protobuf:"varint,1,opt,name=type,proto3,enum=innovationupstream.type.meta.reference.v1beta2.ReferenceType" json:"type,omitempty"`
	Id   string        `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SingleReference) Reset() {
	*x = SingleReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_meta_reference_v1beta2_reference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleReference) ProtoMessage() {}

func (x *SingleReference) ProtoReflect() protoreflect.Message {
	mi := &file_type_meta_reference_v1beta2_reference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleReference.ProtoReflect.Descriptor instead.
func (*SingleReference) Descriptor() ([]byte, []int) {
	return file_type_meta_reference_v1beta2_reference_proto_rawDescGZIP(), []int{0}
}

func (x *SingleReference) GetType() ReferenceType {
	if x != nil {
		return x.Type
	}
	return ReferenceType_ReferenceType_UNSPECIFIED
}

func (x *SingleReference) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LabelledSingleReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label     string           `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Reference *SingleReference `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
}

func (x *LabelledSingleReference) Reset() {
	*x = LabelledSingleReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_meta_reference_v1beta2_reference_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelledSingleReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelledSingleReference) ProtoMessage() {}

func (x *LabelledSingleReference) ProtoReflect() protoreflect.Message {
	mi := &file_type_meta_reference_v1beta2_reference_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelledSingleReference.ProtoReflect.Descriptor instead.
func (*LabelledSingleReference) Descriptor() ([]byte, []int) {
	return file_type_meta_reference_v1beta2_reference_proto_rawDescGZIP(), []int{1}
}

func (x *LabelledSingleReference) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *LabelledSingleReference) GetReference() *SingleReference {
	if x != nil {
		return x.Reference
	}
	return nil
}

type ReferenceQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchType ReferenceMatchType `protobuf:"varint,1,opt,name=match_type,json=matchType,proto3,enum=innovationupstream.type.meta.reference.v1beta2.ReferenceMatchType" json:"match_type,omitempty"`
	Ids       []string           `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	Type      ReferenceType      `protobuf:"varint,3,opt,name=type,proto3,enum=innovationupstream.type.meta.reference.v1beta2.ReferenceType" json:"type,omitempty"`
}

func (x *ReferenceQuery) Reset() {
	*x = ReferenceQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_meta_reference_v1beta2_reference_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferenceQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferenceQuery) ProtoMessage() {}

func (x *ReferenceQuery) ProtoReflect() protoreflect.Message {
	mi := &file_type_meta_reference_v1beta2_reference_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferenceQuery.ProtoReflect.Descriptor instead.
func (*ReferenceQuery) Descriptor() ([]byte, []int) {
	return file_type_meta_reference_v1beta2_reference_proto_rawDescGZIP(), []int{2}
}

func (x *ReferenceQuery) GetMatchType() ReferenceMatchType {
	if x != nil {
		return x.MatchType
	}
	return ReferenceMatchType_ReferenceMatchType_UNSPECIFIED
}

func (x *ReferenceQuery) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ReferenceQuery) GetType() ReferenceType {
	if x != nil {
		return x.Type
	}
	return ReferenceType_ReferenceType_UNSPECIFIED
}

var File_type_meta_reference_v1beta2_reference_proto protoreflect.FileDescriptor

var file_type_meta_reference_v1beta2_reference_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x69,
	0x6e, 0x6e, 0x6f, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x22, 0x74, 0x0a,
	0x0f, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x51, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d,
	0x2e, 0x69, 0x6e, 0x6e, 0x6f, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x75, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x17, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x6c, 0x65, 0x64,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x5d, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x69, 0x6e, 0x6e, 0x6f, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x22, 0xd8, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x61, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x69, 0x6e,
	0x6e, 0x6f, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x51, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x69, 0x6e, 0x6e,
	0x6f, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a,
	0x55, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x19, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x10, 0x03, 0x2a, 0x4a, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x1e,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4e, 0x59,
	0x10, 0x02, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x6f, 0x2e, 0x69, 0x6e, 0x6e, 0x6f, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x32, 0x3b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_type_meta_reference_v1beta2_reference_proto_rawDescOnce sync.Once
	file_type_meta_reference_v1beta2_reference_proto_rawDescData = file_type_meta_reference_v1beta2_reference_proto_rawDesc
)

func file_type_meta_reference_v1beta2_reference_proto_rawDescGZIP() []byte {
	file_type_meta_reference_v1beta2_reference_proto_rawDescOnce.Do(func() {
		file_type_meta_reference_v1beta2_reference_proto_rawDescData = protoimpl.X.CompressGZIP(file_type_meta_reference_v1beta2_reference_proto_rawDescData)
	})
	return file_type_meta_reference_v1beta2_reference_proto_rawDescData
}

var file_type_meta_reference_v1beta2_reference_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_type_meta_reference_v1beta2_reference_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_type_meta_reference_v1beta2_reference_proto_goTypes = []interface{}{
	(ReferenceType)(0),              // 0: innovationupstream.type.meta.reference.v1beta2.ReferenceType
	(ReferenceMatchType)(0),         // 1: innovationupstream.type.meta.reference.v1beta2.ReferenceMatchType
	(*SingleReference)(nil),         // 2: innovationupstream.type.meta.reference.v1beta2.SingleReference
	(*LabelledSingleReference)(nil), // 3: innovationupstream.type.meta.reference.v1beta2.LabelledSingleReference
	(*ReferenceQuery)(nil),          // 4: innovationupstream.type.meta.reference.v1beta2.ReferenceQuery
}
var file_type_meta_reference_v1beta2_reference_proto_depIdxs = []int32{
	0, // 0: innovationupstream.type.meta.reference.v1beta2.SingleReference.type:type_name -> innovationupstream.type.meta.reference.v1beta2.ReferenceType
	2, // 1: innovationupstream.type.meta.reference.v1beta2.LabelledSingleReference.reference:type_name -> innovationupstream.type.meta.reference.v1beta2.SingleReference
	1, // 2: innovationupstream.type.meta.reference.v1beta2.ReferenceQuery.match_type:type_name -> innovationupstream.type.meta.reference.v1beta2.ReferenceMatchType
	0, // 3: innovationupstream.type.meta.reference.v1beta2.ReferenceQuery.type:type_name -> innovationupstream.type.meta.reference.v1beta2.ReferenceType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_type_meta_reference_v1beta2_reference_proto_init() }
func file_type_meta_reference_v1beta2_reference_proto_init() {
	if File_type_meta_reference_v1beta2_reference_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_type_meta_reference_v1beta2_reference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleReference); i {
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
		file_type_meta_reference_v1beta2_reference_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelledSingleReference); i {
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
		file_type_meta_reference_v1beta2_reference_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferenceQuery); i {
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
			RawDescriptor: file_type_meta_reference_v1beta2_reference_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_type_meta_reference_v1beta2_reference_proto_goTypes,
		DependencyIndexes: file_type_meta_reference_v1beta2_reference_proto_depIdxs,
		EnumInfos:         file_type_meta_reference_v1beta2_reference_proto_enumTypes,
		MessageInfos:      file_type_meta_reference_v1beta2_reference_proto_msgTypes,
	}.Build()
	File_type_meta_reference_v1beta2_reference_proto = out.File
	file_type_meta_reference_v1beta2_reference_proto_rawDesc = nil
	file_type_meta_reference_v1beta2_reference_proto_goTypes = nil
	file_type_meta_reference_v1beta2_reference_proto_depIdxs = nil
}
