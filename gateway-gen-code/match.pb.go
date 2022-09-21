// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: match.proto

package __

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

type S2CMatch_Msg int32

const (
	S2CMatch_MatchSuccess  S2CMatch_Msg = 0
	S2CMatch_CancelSuccess S2CMatch_Msg = 1
)

// Enum value maps for S2CMatch_Msg.
var (
	S2CMatch_Msg_name = map[int32]string{
		0: "MatchSuccess",
		1: "CancelSuccess",
	}
	S2CMatch_Msg_value = map[string]int32{
		"MatchSuccess":  0,
		"CancelSuccess": 1,
	}
)

func (x S2CMatch_Msg) Enum() *S2CMatch_Msg {
	p := new(S2CMatch_Msg)
	*p = x
	return p
}

func (x S2CMatch_Msg) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (S2CMatch_Msg) Descriptor() protoreflect.EnumDescriptor {
	return file_match_proto_enumTypes[0].Descriptor()
}

func (S2CMatch_Msg) Type() protoreflect.EnumType {
	return &file_match_proto_enumTypes[0]
}

func (x S2CMatch_Msg) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use S2CMatch_Msg.Descriptor instead.
func (S2CMatch_Msg) EnumDescriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{3, 0}
}

type C2SMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2SMatch) Reset() {
	*x = C2SMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SMatch) ProtoMessage() {}

func (x *C2SMatch) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SMatch.ProtoReflect.Descriptor instead.
func (*C2SMatch) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{0}
}

type C2SCancelMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2SCancelMatch) Reset() {
	*x = C2SCancelMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SCancelMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SCancelMatch) ProtoMessage() {}

func (x *C2SCancelMatch) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SCancelMatch.ProtoReflect.Descriptor instead.
func (*C2SCancelMatch) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{1}
}

type S2CCancelMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *S2CCancelMatch) Reset() {
	*x = S2CCancelMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CCancelMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CCancelMatch) ProtoMessage() {}

func (x *S2CCancelMatch) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CCancelMatch.ProtoReflect.Descriptor instead.
func (*S2CCancelMatch) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{2}
}

type S2CMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg        S2CMatch_Msg `protobuf:"varint,1,opt,name=msg,proto3,enum=match.S2CMatch_Msg" json:"msg,omitempty"`
	ServerInfo *ServerInfo  `protobuf:"bytes,2,opt,name=server_info,json=serverInfo,proto3" json:"server_info,omitempty"`
}

func (x *S2CMatch) Reset() {
	*x = S2CMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CMatch) ProtoMessage() {}

func (x *S2CMatch) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CMatch.ProtoReflect.Descriptor instead.
func (*S2CMatch) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{3}
}

func (x *S2CMatch) GetMsg() S2CMatch_Msg {
	if x != nil {
		return x.Msg
	}
	return S2CMatch_MatchSuccess
}

func (x *S2CMatch) GetServerInfo() *ServerInfo {
	if x != nil {
		return x.ServerInfo
	}
	return nil
}

type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip    string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port  int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{4}
}

func (x *ServerInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ServerInfo) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ServerInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_match_proto protoreflect.FileDescriptor

var file_match_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x22, 0x0a, 0x0a, 0x08, 0x43, 0x32, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x22, 0x10, 0x0a, 0x0e, 0x43, 0x32, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x32, 0x63, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x22, 0x91, 0x01, 0x0a, 0x08, 0x53, 0x32, 0x63, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x53, 0x32, 0x63, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x4d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x32, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2a, 0x0a, 0x03,
	0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x01, 0x22, 0x46, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x32, 0x75, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2d, 0x0a, 0x05, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x0f, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x43, 0x32, 0x73, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x1a, 0x0f, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x53, 0x32, 0x63, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3d, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x43, 0x32, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x15,
	0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x53, 0x32, 0x63, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x01, 0x2e, 0xaa, 0x02, 0x09, 0x47,
	0x72, 0x70, 0x63, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_match_proto_rawDescOnce sync.Once
	file_match_proto_rawDescData = file_match_proto_rawDesc
)

func file_match_proto_rawDescGZIP() []byte {
	file_match_proto_rawDescOnce.Do(func() {
		file_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_match_proto_rawDescData)
	})
	return file_match_proto_rawDescData
}

var file_match_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_match_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_match_proto_goTypes = []interface{}{
	(S2CMatch_Msg)(0),      // 0: match.S2cMatch.Msg
	(*C2SMatch)(nil),       // 1: match.C2sMatch
	(*C2SCancelMatch)(nil), // 2: match.C2sCancelMatch
	(*S2CCancelMatch)(nil), // 3: match.S2cCancelMatch
	(*S2CMatch)(nil),       // 4: match.S2cMatch
	(*ServerInfo)(nil),     // 5: match.ServerInfo
}
var file_match_proto_depIdxs = []int32{
	0, // 0: match.S2cMatch.msg:type_name -> match.S2cMatch.Msg
	5, // 1: match.S2cMatch.server_info:type_name -> match.ServerInfo
	1, // 2: match.Match.Match:input_type -> match.C2sMatch
	2, // 3: match.Match.CancelMatch:input_type -> match.C2sCancelMatch
	4, // 4: match.Match.Match:output_type -> match.S2cMatch
	3, // 5: match.Match.CancelMatch:output_type -> match.S2cCancelMatch
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_match_proto_init() }
func file_match_proto_init() {
	if File_match_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SMatch); i {
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
		file_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SCancelMatch); i {
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
		file_match_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2CCancelMatch); i {
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
		file_match_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2CMatch); i {
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
		file_match_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
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
			RawDescriptor: file_match_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_match_proto_goTypes,
		DependencyIndexes: file_match_proto_depIdxs,
		EnumInfos:         file_match_proto_enumTypes,
		MessageInfos:      file_match_proto_msgTypes,
	}.Build()
	File_match_proto = out.File
	file_match_proto_rawDesc = nil
	file_match_proto_goTypes = nil
	file_match_proto_depIdxs = nil
}