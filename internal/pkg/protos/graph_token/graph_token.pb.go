// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0--rc3
// source: protos/graph_token/graph_token.proto

package graph_token

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

type CreateTokensPairRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Roles  []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *CreateTokensPairRequest) Reset() {
	*x = CreateTokensPairRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_graph_token_graph_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokensPairRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokensPairRequest) ProtoMessage() {}

func (x *CreateTokensPairRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_graph_token_graph_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokensPairRequest.ProtoReflect.Descriptor instead.
func (*CreateTokensPairRequest) Descriptor() ([]byte, []int) {
	return file_protos_graph_token_graph_token_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTokensPairRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateTokensPairRequest) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type CreateTokensPairResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *CreateTokensPairResponce) Reset() {
	*x = CreateTokensPairResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_graph_token_graph_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokensPairResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokensPairResponce) ProtoMessage() {}

func (x *CreateTokensPairResponce) ProtoReflect() protoreflect.Message {
	mi := &file_protos_graph_token_graph_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokensPairResponce.ProtoReflect.Descriptor instead.
func (*CreateTokensPairResponce) Descriptor() ([]byte, []int) {
	return file_protos_graph_token_graph_token_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTokensPairResponce) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CreateTokensPairResponce) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

var File_protos_graph_token_graph_token_proto protoreflect.FileDescriptor

var file_protos_graph_token_graph_token_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x47, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x60, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x7a,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x50, 0x61,
	0x69, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x50, 0x61, 0x69, 0x72, 0x12, 0x24, 0x2e,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f,
	0x3b, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_graph_token_graph_token_proto_rawDescOnce sync.Once
	file_protos_graph_token_graph_token_proto_rawDescData = file_protos_graph_token_graph_token_proto_rawDesc
)

func file_protos_graph_token_graph_token_proto_rawDescGZIP() []byte {
	file_protos_graph_token_graph_token_proto_rawDescOnce.Do(func() {
		file_protos_graph_token_graph_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_graph_token_graph_token_proto_rawDescData)
	})
	return file_protos_graph_token_graph_token_proto_rawDescData
}

var file_protos_graph_token_graph_token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_graph_token_graph_token_proto_goTypes = []any{
	(*CreateTokensPairRequest)(nil),  // 0: graph_token.CreateTokensPairRequest
	(*CreateTokensPairResponce)(nil), // 1: graph_token.CreateTokensPairResponce
}
var file_protos_graph_token_graph_token_proto_depIdxs = []int32{
	0, // 0: graph_token.CreateTokensPairService.CreateTokensPair:input_type -> graph_token.CreateTokensPairRequest
	1, // 1: graph_token.CreateTokensPairService.CreateTokensPair:output_type -> graph_token.CreateTokensPairResponce
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_graph_token_graph_token_proto_init() }
func file_protos_graph_token_graph_token_proto_init() {
	if File_protos_graph_token_graph_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_graph_token_graph_token_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTokensPairRequest); i {
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
		file_protos_graph_token_graph_token_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTokensPairResponce); i {
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
			RawDescriptor: file_protos_graph_token_graph_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_graph_token_graph_token_proto_goTypes,
		DependencyIndexes: file_protos_graph_token_graph_token_proto_depIdxs,
		MessageInfos:      file_protos_graph_token_graph_token_proto_msgTypes,
	}.Build()
	File_protos_graph_token_graph_token_proto = out.File
	file_protos_graph_token_graph_token_proto_rawDesc = nil
	file_protos_graph_token_graph_token_proto_goTypes = nil
	file_protos_graph_token_graph_token_proto_depIdxs = nil
}
