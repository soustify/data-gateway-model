// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: partners_categories.proto

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

type PartnerCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity     *Entity `protobuf:"bytes,1,opt,name=Entity,proto3" json:"Entity,omitempty"`
	IdPartner  string  `protobuf:"bytes,2,opt,name=id_partner,json=idPartner,proto3" json:"id_partner,omitempty"`
	IdCategory string  `protobuf:"bytes,3,opt,name=id_category,json=idCategory,proto3" json:"id_category,omitempty"`
}

func (x *PartnerCategory) Reset() {
	*x = PartnerCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partners_categories_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartnerCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartnerCategory) ProtoMessage() {}

func (x *PartnerCategory) ProtoReflect() protoreflect.Message {
	mi := &file_partners_categories_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartnerCategory.ProtoReflect.Descriptor instead.
func (*PartnerCategory) Descriptor() ([]byte, []int) {
	return file_partners_categories_proto_rawDescGZIP(), []int{0}
}

func (x *PartnerCategory) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PartnerCategory) GetIdPartner() string {
	if x != nil {
		return x.IdPartner
	}
	return ""
}

func (x *PartnerCategory) GetIdCategory() string {
	if x != nil {
		return x.IdCategory
	}
	return ""
}

type PartnerCategoryListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  *PartnerCategory `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Metadata *Metadata        `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *PartnerCategoryListRequest) Reset() {
	*x = PartnerCategoryListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partners_categories_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartnerCategoryListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartnerCategoryListRequest) ProtoMessage() {}

func (x *PartnerCategoryListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_partners_categories_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartnerCategoryListRequest.ProtoReflect.Descriptor instead.
func (*PartnerCategoryListRequest) Descriptor() ([]byte, []int) {
	return file_partners_categories_proto_rawDescGZIP(), []int{1}
}

func (x *PartnerCategoryListRequest) GetPayload() *PartnerCategory {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *PartnerCategoryListRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type PartnerCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *PartnerCategory `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *PartnerCategoryResponse) Reset() {
	*x = PartnerCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partners_categories_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartnerCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartnerCategoryResponse) ProtoMessage() {}

func (x *PartnerCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_partners_categories_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartnerCategoryResponse.ProtoReflect.Descriptor instead.
func (*PartnerCategoryResponse) Descriptor() ([]byte, []int) {
	return file_partners_categories_proto_rawDescGZIP(), []int{2}
}

func (x *PartnerCategoryResponse) GetResponse() *PartnerCategory {
	if x != nil {
		return x.Response
	}
	return nil
}

type PartnerCategoryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []*PartnerCategory `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *PartnerCategoryListResponse) Reset() {
	*x = PartnerCategoryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_partners_categories_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartnerCategoryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartnerCategoryListResponse) ProtoMessage() {}

func (x *PartnerCategoryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_partners_categories_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartnerCategoryListResponse.ProtoReflect.Descriptor instead.
func (*PartnerCategoryListResponse) Descriptor() ([]byte, []int) {
	return file_partners_categories_proto_rawDescGZIP(), []int{3}
}

func (x *PartnerCategoryListResponse) GetResponse() []*PartnerCategory {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_partners_categories_proto protoreflect.FileDescriptor

var file_partners_categories_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x79, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x26, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x52, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x64, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x64, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x69, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x7d, 0x0a, 0x1a, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x17, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x0a, 0x1b, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf4, 0x01,
	0x0a, 0x16, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x1f, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a,
	0x1f, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x52, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x22, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_partners_categories_proto_rawDescOnce sync.Once
	file_partners_categories_proto_rawDescData = file_partners_categories_proto_rawDesc
)

func file_partners_categories_proto_rawDescGZIP() []byte {
	file_partners_categories_proto_rawDescOnce.Do(func() {
		file_partners_categories_proto_rawDescData = protoimpl.X.CompressGZIP(file_partners_categories_proto_rawDescData)
	})
	return file_partners_categories_proto_rawDescData
}

var file_partners_categories_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_partners_categories_proto_goTypes = []any{
	(*PartnerCategory)(nil),             // 0: domain.PartnerCategory
	(*PartnerCategoryListRequest)(nil),  // 1: domain.PartnerCategoryListRequest
	(*PartnerCategoryResponse)(nil),     // 2: domain.PartnerCategoryResponse
	(*PartnerCategoryListResponse)(nil), // 3: domain.PartnerCategoryListResponse
	(*Entity)(nil),                      // 4: domain.Entity
	(*Metadata)(nil),                    // 5: domain.Metadata
}
var file_partners_categories_proto_depIdxs = []int32{
	4, // 0: domain.PartnerCategory.Entity:type_name -> domain.Entity
	0, // 1: domain.PartnerCategoryListRequest.payload:type_name -> domain.PartnerCategory
	5, // 2: domain.PartnerCategoryListRequest.metadata:type_name -> domain.Metadata
	0, // 3: domain.PartnerCategoryResponse.response:type_name -> domain.PartnerCategory
	0, // 4: domain.PartnerCategoryListResponse.response:type_name -> domain.PartnerCategory
	0, // 5: domain.PartnerCategoryService.Create:input_type -> domain.PartnerCategory
	0, // 6: domain.PartnerCategoryService.Update:input_type -> domain.PartnerCategory
	1, // 7: domain.PartnerCategoryService.FindAll:input_type -> domain.PartnerCategoryListRequest
	2, // 8: domain.PartnerCategoryService.Create:output_type -> domain.PartnerCategoryResponse
	2, // 9: domain.PartnerCategoryService.Update:output_type -> domain.PartnerCategoryResponse
	3, // 10: domain.PartnerCategoryService.FindAll:output_type -> domain.PartnerCategoryListResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_partners_categories_proto_init() }
func file_partners_categories_proto_init() {
	if File_partners_categories_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_partners_categories_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PartnerCategory); i {
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
		file_partners_categories_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PartnerCategoryListRequest); i {
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
		file_partners_categories_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PartnerCategoryResponse); i {
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
		file_partners_categories_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PartnerCategoryListResponse); i {
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
			RawDescriptor: file_partners_categories_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_partners_categories_proto_goTypes,
		DependencyIndexes: file_partners_categories_proto_depIdxs,
		MessageInfos:      file_partners_categories_proto_msgTypes,
	}.Build()
	File_partners_categories_proto = out.File
	file_partners_categories_proto_rawDesc = nil
	file_partners_categories_proto_goTypes = nil
	file_partners_categories_proto_depIdxs = nil
}
