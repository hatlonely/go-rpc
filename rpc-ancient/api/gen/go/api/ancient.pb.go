// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: api/ancient.proto

package api

import (
	context "context"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Ancient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author  string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Dynasty string `protobuf:"bytes,4,opt,name=dynasty,proto3" json:"dynasty,omitempty"`
	Content string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Ancient) Reset() {
	*x = Ancient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ancient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ancient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ancient) ProtoMessage() {}

func (x *Ancient) ProtoReflect() protoreflect.Message {
	mi := &file_api_ancient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ancient.ProtoReflect.Descriptor instead.
func (*Ancient) Descriptor() ([]byte, []int) {
	return file_api_ancient_proto_rawDescGZIP(), []int{0}
}

func (x *Ancient) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Ancient) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Ancient) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Ancient) GetDynasty() string {
	if x != nil {
		return x.Dynasty
	}
	return ""
}

func (x *Ancient) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type GetAncientReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAncientReq) Reset() {
	*x = GetAncientReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ancient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAncientReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAncientReq) ProtoMessage() {}

func (x *GetAncientReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_ancient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAncientReq.ProtoReflect.Descriptor instead.
func (*GetAncientReq) Descriptor() ([]byte, []int) {
	return file_api_ancient_proto_rawDescGZIP(), []int{1}
}

func (x *GetAncientReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PutAncientReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ancient *Ancient `protobuf:"bytes,1,opt,name=ancient,proto3" json:"ancient,omitempty"`
}

func (x *PutAncientReq) Reset() {
	*x = PutAncientReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ancient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutAncientReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutAncientReq) ProtoMessage() {}

func (x *PutAncientReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_ancient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutAncientReq.ProtoReflect.Descriptor instead.
func (*PutAncientReq) Descriptor() ([]byte, []int) {
	return file_api_ancient_proto_rawDescGZIP(), []int{2}
}

func (x *PutAncientReq) GetAncient() *Ancient {
	if x != nil {
		return x.Ancient
	}
	return nil
}

type UpdateAncientReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ancient    *Ancient              `protobuf:"bytes,1,opt,name=ancient,proto3" json:"ancient,omitempty"`
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateAncientReq) Reset() {
	*x = UpdateAncientReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ancient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAncientReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAncientReq) ProtoMessage() {}

func (x *UpdateAncientReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_ancient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAncientReq.ProtoReflect.Descriptor instead.
func (*UpdateAncientReq) Descriptor() ([]byte, []int) {
	return file_api_ancient_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAncientReq) GetAncient() *Ancient {
	if x != nil {
		return x.Ancient
	}
	return nil
}

func (x *UpdateAncientReq) GetUpdateMask() *field_mask.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_api_ancient_proto protoreflect.FileDescriptor

var file_api_ancient_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x07, 0x41,
	0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x79, 0x6e, 0x61, 0x73, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x79, 0x6e, 0x61, 0x73, 0x74, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41,
	0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x0d, 0x50, 0x75, 0x74,
	0x41, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x07, 0x61, 0x6e,
	0x63, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x41, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x6e, 0x63, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0x77, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x63, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x07, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6e,
	0x63, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x3b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x32, 0xbd, 0x02, 0x0a, 0x0e,
	0x41, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x63, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x50, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x41,
	0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x74,
	0x41, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x8e, 0x01, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x48, 0x1a, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x7b, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x69, 0x64, 0x7d, 0x3a, 0x07, 0x61,
	0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x5a, 0x23, 0x32, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e,
	0x63, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x69,
	0x64, 0x7d, 0x3a, 0x07, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x74, 0x6c, 0x6f, 0x6e,
	0x65, 0x6c, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2d, 0x61,
	0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ancient_proto_rawDescOnce sync.Once
	file_api_ancient_proto_rawDescData = file_api_ancient_proto_rawDesc
)

func file_api_ancient_proto_rawDescGZIP() []byte {
	file_api_ancient_proto_rawDescOnce.Do(func() {
		file_api_ancient_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ancient_proto_rawDescData)
	})
	return file_api_ancient_proto_rawDescData
}

var file_api_ancient_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_ancient_proto_goTypes = []interface{}{
	(*Ancient)(nil),              // 0: api.Ancient
	(*GetAncientReq)(nil),        // 1: api.GetAncientReq
	(*PutAncientReq)(nil),        // 2: api.PutAncientReq
	(*UpdateAncientReq)(nil),     // 3: api.UpdateAncientReq
	(*field_mask.FieldMask)(nil), // 4: google.protobuf.FieldMask
	(*empty.Empty)(nil),          // 5: google.protobuf.Empty
}
var file_api_ancient_proto_depIdxs = []int32{
	0, // 0: api.PutAncientReq.ancient:type_name -> api.Ancient
	0, // 1: api.UpdateAncientReq.ancient:type_name -> api.Ancient
	4, // 2: api.UpdateAncientReq.update_mask:type_name -> google.protobuf.FieldMask
	1, // 3: api.AncientService.GetAncient:input_type -> api.GetAncientReq
	2, // 4: api.AncientService.PutAncient:input_type -> api.PutAncientReq
	3, // 5: api.AncientService.UpdateAncient:input_type -> api.UpdateAncientReq
	0, // 6: api.AncientService.GetAncient:output_type -> api.Ancient
	5, // 7: api.AncientService.PutAncient:output_type -> google.protobuf.Empty
	5, // 8: api.AncientService.UpdateAncient:output_type -> google.protobuf.Empty
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_ancient_proto_init() }
func file_api_ancient_proto_init() {
	if File_api_ancient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ancient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ancient); i {
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
		file_api_ancient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAncientReq); i {
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
		file_api_ancient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutAncientReq); i {
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
		file_api_ancient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAncientReq); i {
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
			RawDescriptor: file_api_ancient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ancient_proto_goTypes,
		DependencyIndexes: file_api_ancient_proto_depIdxs,
		MessageInfos:      file_api_ancient_proto_msgTypes,
	}.Build()
	File_api_ancient_proto = out.File
	file_api_ancient_proto_rawDesc = nil
	file_api_ancient_proto_goTypes = nil
	file_api_ancient_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AncientServiceClient is the client API for AncientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AncientServiceClient interface {
	GetAncient(ctx context.Context, in *GetAncientReq, opts ...grpc.CallOption) (*Ancient, error)
	PutAncient(ctx context.Context, in *PutAncientReq, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateAncient(ctx context.Context, in *UpdateAncientReq, opts ...grpc.CallOption) (*empty.Empty, error)
}

type ancientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAncientServiceClient(cc grpc.ClientConnInterface) AncientServiceClient {
	return &ancientServiceClient{cc}
}

func (c *ancientServiceClient) GetAncient(ctx context.Context, in *GetAncientReq, opts ...grpc.CallOption) (*Ancient, error) {
	out := new(Ancient)
	err := c.cc.Invoke(ctx, "/api.AncientService/GetAncient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ancientServiceClient) PutAncient(ctx context.Context, in *PutAncientReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.AncientService/PutAncient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ancientServiceClient) UpdateAncient(ctx context.Context, in *UpdateAncientReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.AncientService/UpdateAncient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AncientServiceServer is the server API for AncientService service.
type AncientServiceServer interface {
	GetAncient(context.Context, *GetAncientReq) (*Ancient, error)
	PutAncient(context.Context, *PutAncientReq) (*empty.Empty, error)
	UpdateAncient(context.Context, *UpdateAncientReq) (*empty.Empty, error)
}

// UnimplementedAncientServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAncientServiceServer struct {
}

func (*UnimplementedAncientServiceServer) GetAncient(context.Context, *GetAncientReq) (*Ancient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAncient not implemented")
}
func (*UnimplementedAncientServiceServer) PutAncient(context.Context, *PutAncientReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutAncient not implemented")
}
func (*UnimplementedAncientServiceServer) UpdateAncient(context.Context, *UpdateAncientReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAncient not implemented")
}

func RegisterAncientServiceServer(s *grpc.Server, srv AncientServiceServer) {
	s.RegisterService(&_AncientService_serviceDesc, srv)
}

func _AncientService_GetAncient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAncientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AncientServiceServer).GetAncient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AncientService/GetAncient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AncientServiceServer).GetAncient(ctx, req.(*GetAncientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AncientService_PutAncient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutAncientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AncientServiceServer).PutAncient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AncientService/PutAncient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AncientServiceServer).PutAncient(ctx, req.(*PutAncientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AncientService_UpdateAncient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAncientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AncientServiceServer).UpdateAncient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AncientService/UpdateAncient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AncientServiceServer).UpdateAncient(ctx, req.(*UpdateAncientReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AncientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AncientService",
	HandlerType: (*AncientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAncient",
			Handler:    _AncientService_GetAncient_Handler,
		},
		{
			MethodName: "PutAncient",
			Handler:    _AncientService_PutAncient_Handler,
		},
		{
			MethodName: "UpdateAncient",
			Handler:    _AncientService_UpdateAncient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ancient.proto",
}
