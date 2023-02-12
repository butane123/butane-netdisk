// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: repository.proto

package repository

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

type RepositoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepositoryId int64 `protobuf:"varint,1,opt,name=repositoryId,proto3" json:"repositoryId,omitempty"`
}

func (x *RepositoryReq) Reset() {
	*x = RepositoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryReq) ProtoMessage() {}

func (x *RepositoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryReq.ProtoReflect.Descriptor instead.
func (*RepositoryReq) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{0}
}

func (x *RepositoryReq) GetRepositoryId() int64 {
	if x != nil {
		return x.RepositoryId
	}
	return 0
}

type RepositoryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ext  string `protobuf:"bytes,1,opt,name=ext,proto3" json:"ext,omitempty"`
	Size int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RepositoryResp) Reset() {
	*x = RepositoryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryResp) ProtoMessage() {}

func (x *RepositoryResp) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryResp.ProtoReflect.Descriptor instead.
func (*RepositoryResp) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{1}
}

func (x *RepositoryResp) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

func (x *RepositoryResp) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *RepositoryResp) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RepositoryResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepositoryId int64 `protobuf:"varint,1,opt,name=repositoryId,proto3" json:"repositoryId,omitempty"`
}

func (x *DeleteByIdReq) Reset() {
	*x = DeleteByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByIdReq) ProtoMessage() {}

func (x *DeleteByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteByIdReq.ProtoReflect.Descriptor instead.
func (*DeleteByIdReq) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteByIdReq) GetRepositoryId() int64 {
	if x != nil {
		return x.RepositoryId
	}
	return 0
}

type DeleteByIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *DeleteByIdResp) Reset() {
	*x = DeleteByIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteByIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByIdResp) ProtoMessage() {}

func (x *DeleteByIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteByIdResp.ProtoReflect.Descriptor instead.
func (*DeleteByIdResp) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteByIdResp) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_repository_proto protoreflect.FileDescriptor

var file_repository_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x33,
	0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x32, 0xab,
	0x01, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x58, 0x0a,
	0x1f, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x6f,
	0x6f, 0x6c, 0x42, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x12, 0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0e, 0x5a, 0x0c,
	0x2e, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_repository_proto_rawDescOnce sync.Once
	file_repository_proto_rawDescData = file_repository_proto_rawDesc
)

func file_repository_proto_rawDescGZIP() []byte {
	file_repository_proto_rawDescOnce.Do(func() {
		file_repository_proto_rawDescData = protoimpl.X.CompressGZIP(file_repository_proto_rawDescData)
	})
	return file_repository_proto_rawDescData
}

var file_repository_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_repository_proto_goTypes = []interface{}{
	(*RepositoryReq)(nil),  // 0: repository.RepositoryReq
	(*RepositoryResp)(nil), // 1: repository.RepositoryResp
	(*DeleteByIdReq)(nil),  // 2: repository.DeleteByIdReq
	(*DeleteByIdResp)(nil), // 3: repository.DeleteByIdResp
}
var file_repository_proto_depIdxs = []int32{
	0, // 0: repository.repository.getRepositoryPoolByRepositoryId:input_type -> repository.RepositoryReq
	2, // 1: repository.repository.deleteById:input_type -> repository.DeleteByIdReq
	1, // 2: repository.repository.getRepositoryPoolByRepositoryId:output_type -> repository.RepositoryResp
	3, // 3: repository.repository.deleteById:output_type -> repository.DeleteByIdResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_repository_proto_init() }
func file_repository_proto_init() {
	if File_repository_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_repository_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryReq); i {
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
		file_repository_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryResp); i {
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
		file_repository_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteByIdReq); i {
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
		file_repository_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteByIdResp); i {
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
			RawDescriptor: file_repository_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_repository_proto_goTypes,
		DependencyIndexes: file_repository_proto_depIdxs,
		MessageInfos:      file_repository_proto_msgTypes,
	}.Build()
	File_repository_proto = out.File
	file_repository_proto_rawDesc = nil
	file_repository_proto_goTypes = nil
	file_repository_proto_depIdxs = nil
}