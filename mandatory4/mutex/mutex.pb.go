// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: mutex.proto

package mutex

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId  string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Timestamp int32  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mutex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_mutex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_mutex_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Request) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mutex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_mutex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_mutex_proto_rawDescGZIP(), []int{1}
}

type LamportTimestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int32 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *LamportTimestamp) Reset() {
	*x = LamportTimestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mutex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LamportTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LamportTimestamp) ProtoMessage() {}

func (x *LamportTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_mutex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LamportTimestamp.ProtoReflect.Descriptor instead.
func (*LamportTimestamp) Descriptor() ([]byte, []int) {
	return file_mutex_proto_rawDescGZIP(), []int{2}
}

func (x *LamportTimestamp) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type PeerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerList []string `protobuf:"bytes,1,rep,name=peerList,proto3" json:"peerList,omitempty"`
}

func (x *PeerList) Reset() {
	*x = PeerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mutex_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerList) ProtoMessage() {}

func (x *PeerList) ProtoReflect() protoreflect.Message {
	mi := &file_mutex_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerList.ProtoReflect.Descriptor instead.
func (*PeerList) Descriptor() ([]byte, []int) {
	return file_mutex_proto_rawDescGZIP(), []int{3}
}

func (x *PeerList) GetPeerList() []string {
	if x != nil {
		return x.PeerList
	}
	return nil
}

var File_mutex_proto protoreflect.FileDescriptor

var file_mutex_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63,
	0x68, 0x61, 0x74, 0x22, 0x43, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x30, 0x0a, 0x10, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x26, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x65, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xd9, 0x01, 0x0a, 0x05,
	0x4d, 0x75, 0x74, 0x65, 0x78, 0x12, 0x2b, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x40, 0x0a, 0x17, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0d, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x30, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x65, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x65,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x11, 0x4c, 0x65, 0x74, 0x50, 0x65, 0x65,
	0x72, 0x4b, 0x6e, 0x6f, 0x77, 0x49, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6d, 0x75, 0x74, 0x65,
	0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mutex_proto_rawDescOnce sync.Once
	file_mutex_proto_rawDescData = file_mutex_proto_rawDesc
)

func file_mutex_proto_rawDescGZIP() []byte {
	file_mutex_proto_rawDescOnce.Do(func() {
		file_mutex_proto_rawDescData = protoimpl.X.CompressGZIP(file_mutex_proto_rawDescData)
	})
	return file_mutex_proto_rawDescData
}

var file_mutex_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mutex_proto_goTypes = []interface{}{
	(*Request)(nil),          // 0: chat.Request
	(*Empty)(nil),            // 1: chat.Empty
	(*LamportTimestamp)(nil), // 2: chat.LamportTimestamp
	(*PeerList)(nil),         // 3: chat.PeerList
}
var file_mutex_proto_depIdxs = []int32{
	0, // 0: chat.Mutex.RequestAccess:input_type -> chat.Request
	0, // 1: chat.Mutex.RequestLamportTimestamp:input_type -> chat.Request
	0, // 2: chat.Mutex.RequestPeerList:input_type -> chat.Request
	0, // 3: chat.Mutex.LetPeerKnowIExist:input_type -> chat.Request
	1, // 4: chat.Mutex.RequestAccess:output_type -> chat.Empty
	2, // 5: chat.Mutex.RequestLamportTimestamp:output_type -> chat.LamportTimestamp
	3, // 6: chat.Mutex.RequestPeerList:output_type -> chat.PeerList
	1, // 7: chat.Mutex.LetPeerKnowIExist:output_type -> chat.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mutex_proto_init() }
func file_mutex_proto_init() {
	if File_mutex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mutex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_mutex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_mutex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LamportTimestamp); i {
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
		file_mutex_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerList); i {
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
			RawDescriptor: file_mutex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mutex_proto_goTypes,
		DependencyIndexes: file_mutex_proto_depIdxs,
		MessageInfos:      file_mutex_proto_msgTypes,
	}.Build()
	File_mutex_proto = out.File
	file_mutex_proto_rawDesc = nil
	file_mutex_proto_goTypes = nil
	file_mutex_proto_depIdxs = nil
}
