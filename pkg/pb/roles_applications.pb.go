// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: roles_applications.proto

package pb

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

type RoleApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity        *Entity `protobuf:"bytes,1,opt,name=Entity,proto3" json:"Entity,omitempty"`
	IdRole        string  `protobuf:"bytes,2,opt,name=id_role,json=idRole,proto3" json:"id_role,omitempty"`
	IdApplication string  `protobuf:"bytes,3,opt,name=id_application,json=idApplication,proto3" json:"id_application,omitempty"`
}

func (x *RoleApplication) Reset() {
	*x = RoleApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roles_applications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleApplication) ProtoMessage() {}

func (x *RoleApplication) ProtoReflect() protoreflect.Message {
	mi := &file_roles_applications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleApplication.ProtoReflect.Descriptor instead.
func (*RoleApplication) Descriptor() ([]byte, []int) {
	return file_roles_applications_proto_rawDescGZIP(), []int{0}
}

func (x *RoleApplication) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *RoleApplication) GetIdRole() string {
	if x != nil {
		return x.IdRole
	}
	return ""
}

func (x *RoleApplication) GetIdApplication() string {
	if x != nil {
		return x.IdApplication
	}
	return ""
}

type RoleApplicationListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  *RoleApplication `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Metadata *Metadata        `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *RoleApplicationListRequest) Reset() {
	*x = RoleApplicationListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roles_applications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleApplicationListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleApplicationListRequest) ProtoMessage() {}

func (x *RoleApplicationListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_roles_applications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleApplicationListRequest.ProtoReflect.Descriptor instead.
func (*RoleApplicationListRequest) Descriptor() ([]byte, []int) {
	return file_roles_applications_proto_rawDescGZIP(), []int{1}
}

func (x *RoleApplicationListRequest) GetPayload() *RoleApplication {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *RoleApplicationListRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type RoleApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *RoleApplication `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *RoleApplicationResponse) Reset() {
	*x = RoleApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roles_applications_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleApplicationResponse) ProtoMessage() {}

func (x *RoleApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_roles_applications_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleApplicationResponse.ProtoReflect.Descriptor instead.
func (*RoleApplicationResponse) Descriptor() ([]byte, []int) {
	return file_roles_applications_proto_rawDescGZIP(), []int{2}
}

func (x *RoleApplicationResponse) GetResponse() *RoleApplication {
	if x != nil {
		return x.Response
	}
	return nil
}

type RoleApplicationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []*RoleApplication `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *RoleApplicationListResponse) Reset() {
	*x = RoleApplicationListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roles_applications_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleApplicationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleApplicationListResponse) ProtoMessage() {}

func (x *RoleApplicationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_roles_applications_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleApplicationListResponse.ProtoReflect.Descriptor instead.
func (*RoleApplicationListResponse) Descriptor() ([]byte, []int) {
	return file_roles_applications_proto_rawDescGZIP(), []int{3}
}

func (x *RoleApplicationListResponse) GetResponse() []*RoleApplication {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_roles_applications_proto protoreflect.FileDescriptor

var file_roles_applications_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79,
	0x0a, 0x0f, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x26, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x64, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7d, 0x0a, 0x1a, 0x52, 0x6f, 0x6c,
	0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x17, 0x52, 0x6f, 0x6c, 0x65,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x0a, 0x1b, 0x52, 0x6f, 0x6c, 0x65,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf4, 0x01, 0x0a,
	0x16, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x17, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x52, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x22, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_roles_applications_proto_rawDescOnce sync.Once
	file_roles_applications_proto_rawDescData = file_roles_applications_proto_rawDesc
)

func file_roles_applications_proto_rawDescGZIP() []byte {
	file_roles_applications_proto_rawDescOnce.Do(func() {
		file_roles_applications_proto_rawDescData = protoimpl.X.CompressGZIP(file_roles_applications_proto_rawDescData)
	})
	return file_roles_applications_proto_rawDescData
}

var file_roles_applications_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_roles_applications_proto_goTypes = []any{
	(*RoleApplication)(nil),             // 0: domain.RoleApplication
	(*RoleApplicationListRequest)(nil),  // 1: domain.RoleApplicationListRequest
	(*RoleApplicationResponse)(nil),     // 2: domain.RoleApplicationResponse
	(*RoleApplicationListResponse)(nil), // 3: domain.RoleApplicationListResponse
	(*Entity)(nil),                      // 4: domain.Entity
	(*Metadata)(nil),                    // 5: domain.Metadata
}
var file_roles_applications_proto_depIdxs = []int32{
	4, // 0: domain.RoleApplication.Entity:type_name -> domain.Entity
	0, // 1: domain.RoleApplicationListRequest.payload:type_name -> domain.RoleApplication
	5, // 2: domain.RoleApplicationListRequest.metadata:type_name -> domain.Metadata
	0, // 3: domain.RoleApplicationResponse.response:type_name -> domain.RoleApplication
	0, // 4: domain.RoleApplicationListResponse.response:type_name -> domain.RoleApplication
	0, // 5: domain.RoleApplicationService.Create:input_type -> domain.RoleApplication
	0, // 6: domain.RoleApplicationService.Update:input_type -> domain.RoleApplication
	1, // 7: domain.RoleApplicationService.FindAll:input_type -> domain.RoleApplicationListRequest
	2, // 8: domain.RoleApplicationService.Create:output_type -> domain.RoleApplicationResponse
	2, // 9: domain.RoleApplicationService.Update:output_type -> domain.RoleApplicationResponse
	3, // 10: domain.RoleApplicationService.FindAll:output_type -> domain.RoleApplicationListResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_roles_applications_proto_init() }
func file_roles_applications_proto_init() {
	if File_roles_applications_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_roles_applications_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RoleApplication); i {
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
		file_roles_applications_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RoleApplicationListRequest); i {
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
		file_roles_applications_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RoleApplicationResponse); i {
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
		file_roles_applications_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RoleApplicationListResponse); i {
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
			RawDescriptor: file_roles_applications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_roles_applications_proto_goTypes,
		DependencyIndexes: file_roles_applications_proto_depIdxs,
		MessageInfos:      file_roles_applications_proto_msgTypes,
	}.Build()
	File_roles_applications_proto = out.File
	file_roles_applications_proto_rawDesc = nil
	file_roles_applications_proto_goTypes = nil
	file_roles_applications_proto_depIdxs = nil
}
