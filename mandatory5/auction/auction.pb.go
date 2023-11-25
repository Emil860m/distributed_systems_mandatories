// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: auction.proto

package auction

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

type AccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerNodeId string `protobuf:"bytes,1,opt,name=serverNodeId,proto3" json:"serverNodeId,omitempty"`
	BidTimestamp int32  `protobuf:"varint,2,opt,name=bidTimestamp,proto3" json:"bidTimestamp,omitempty"`
}

func (x *AccessRequest) Reset() {
	*x = AccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequest) ProtoMessage() {}

func (x *AccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequest.ProtoReflect.Descriptor instead.
func (*AccessRequest) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{0}
}

func (x *AccessRequest) GetServerNodeId() string {
	if x != nil {
		return x.ServerNodeId
	}
	return ""
}

func (x *AccessRequest) GetBidTimestamp() int32 {
	if x != nil {
		return x.BidTimestamp
	}
	return 0
}

type NewHighestBid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount    int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Timestamp int32  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *NewHighestBid) Reset() {
	*x = NewHighestBid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewHighestBid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewHighestBid) ProtoMessage() {}

func (x *NewHighestBid) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewHighestBid.ProtoReflect.Descriptor instead.
func (*NewHighestBid) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{1}
}

func (x *NewHighestBid) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *NewHighestBid) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewHighestBid) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ongoing bool   `protobuf:"varint,1,opt,name=ongoing,proto3" json:"ongoing,omitempty"`
	Amount  int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{2}
}

func (x *Outcome) GetOngoing() bool {
	if x != nil {
		return x.Ongoing
	}
	return false
}

func (x *Outcome) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Outcome) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BidMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int32  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BidMessage) Reset() {
	*x = BidMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidMessage) ProtoMessage() {}

func (x *BidMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidMessage.ProtoReflect.Descriptor instead.
func (*BidMessage) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{3}
}

func (x *BidMessage) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BidMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outcome string `protobuf:"bytes,1,opt,name=outcome,proto3" json:"outcome,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{4}
}

func (x *Ack) GetOutcome() string {
	if x != nil {
		return x.Outcome
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[5]
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
	return file_auction_proto_rawDescGZIP(), []int{5}
}

var File_auction_proto protoreflect.FileDescriptor

var file_auction_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x57, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x69,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x62, 0x69, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x59,
	0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x4f, 0x0a, 0x07, 0x6f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x0a, 0x62, 0x69,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x75,
	0x74, 0x63, 0x6f, 0x6d, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xc1,
	0x01, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a,
	0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x13,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x36, 0x0a, 0x12, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4e, 0x65, 0x77, 0x48, 0x69, 0x67, 0x68,
	0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x12, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4e, 0x65,
	0x77, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x1a, 0x0b, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12,
	0x10, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x62, 0x69, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x09, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x61, 0x63, 0x6b, 0x12, 0x24, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x6f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auction_proto_rawDescOnce sync.Once
	file_auction_proto_rawDescData = file_auction_proto_rawDesc
)

func file_auction_proto_rawDescGZIP() []byte {
	file_auction_proto_rawDescOnce.Do(func() {
		file_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_auction_proto_rawDescData)
	})
	return file_auction_proto_rawDescData
}

var file_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_auction_proto_goTypes = []interface{}{
	(*AccessRequest)(nil), // 0: chat.AccessRequest
	(*NewHighestBid)(nil), // 1: chat.NewHighestBid
	(*Outcome)(nil),       // 2: chat.outcome
	(*BidMessage)(nil),    // 3: chat.bidMessage
	(*Ack)(nil),           // 4: chat.ack
	(*Empty)(nil),         // 5: chat.Empty
}
var file_auction_proto_depIdxs = []int32{
	0, // 0: chat.serverNode.RequestAccess:input_type -> chat.AccessRequest
	1, // 1: chat.serverNode.ShareNewHighestBid:input_type -> chat.NewHighestBid
	3, // 2: chat.serverNode.Bid:input_type -> chat.bidMessage
	5, // 3: chat.serverNode.Result:input_type -> chat.Empty
	5, // 4: chat.serverNode.RequestAccess:output_type -> chat.Empty
	5, // 5: chat.serverNode.ShareNewHighestBid:output_type -> chat.Empty
	4, // 6: chat.serverNode.Bid:output_type -> chat.ack
	2, // 7: chat.serverNode.Result:output_type -> chat.outcome
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auction_proto_init() }
func file_auction_proto_init() {
	if File_auction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessRequest); i {
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
		file_auction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewHighestBid); i {
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
		file_auction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outcome); i {
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
		file_auction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidMessage); i {
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
		file_auction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_auction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auction_proto_goTypes,
		DependencyIndexes: file_auction_proto_depIdxs,
		MessageInfos:      file_auction_proto_msgTypes,
	}.Build()
	File_auction_proto = out.File
	file_auction_proto_rawDesc = nil
	file_auction_proto_goTypes = nil
	file_auction_proto_depIdxs = nil
}