// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/tradier.proto

package tests

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

type MarketeStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarketeStateRequest) Reset() {
	*x = MarketeStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tradier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketeStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketeStateRequest) ProtoMessage() {}

func (x *MarketeStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tradier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketeStateRequest.ProtoReflect.Descriptor instead.
func (*MarketeStateRequest) Descriptor() ([]byte, []int) {
	return file_proto_tradier_proto_rawDescGZIP(), []int{0}
}

type MarketeStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time        string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	State       string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Nextstate   string `protobuf:"bytes,4,opt,name=nextstate,proto3" json:"nextstate,omitempty"`
}

func (x *MarketeStateResponse) Reset() {
	*x = MarketeStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tradier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketeStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketeStateResponse) ProtoMessage() {}

func (x *MarketeStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tradier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketeStateResponse.ProtoReflect.Descriptor instead.
func (*MarketeStateResponse) Descriptor() ([]byte, []int) {
	return file_proto_tradier_proto_rawDescGZIP(), []int{1}
}

func (x *MarketeStateResponse) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *MarketeStateResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *MarketeStateResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MarketeStateResponse) GetNextstate() string {
	if x != nil {
		return x.Nextstate
	}
	return ""
}

type QuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *QuoteRequest) Reset() {
	*x = QuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tradier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteRequest) ProtoMessage() {}

func (x *QuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tradier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteRequest.ProtoReflect.Descriptor instead.
func (*QuoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_tradier_proto_rawDescGZIP(), []int{2}
}

func (x *QuoteRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type QuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote []*QuoteInfo `protobuf:"bytes,1,rep,name=quote,proto3" json:"quote,omitempty"`
}

func (x *QuoteResponse) Reset() {
	*x = QuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tradier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteResponse) ProtoMessage() {}

func (x *QuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tradier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteResponse.ProtoReflect.Descriptor instead.
func (*QuoteResponse) Descriptor() ([]byte, []int) {
	return file_proto_tradier_proto_rawDescGZIP(), []int{3}
}

func (x *QuoteResponse) GetQuote() []*QuoteInfo {
	if x != nil {
		return x.Quote
	}
	return nil
}

type QuoteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol   string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Last     float64 `protobuf:"fixed64,2,opt,name=last,proto3" json:"last,omitempty"`
	AskPrice float64 `protobuf:"fixed64,3,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	BidPrice float64 `protobuf:"fixed64,4,opt,name=bid_price,json=bidPrice,proto3" json:"bid_price,omitempty"`
	AskSize  int32   `protobuf:"varint,5,opt,name=ask_size,json=askSize,proto3" json:"ask_size,omitempty"`
	BidSize  int32   `protobuf:"varint,6,opt,name=bid_size,json=bidSize,proto3" json:"bid_size,omitempty"`
}

func (x *QuoteInfo) Reset() {
	*x = QuoteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tradier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteInfo) ProtoMessage() {}

func (x *QuoteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tradier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteInfo.ProtoReflect.Descriptor instead.
func (*QuoteInfo) Descriptor() ([]byte, []int) {
	return file_proto_tradier_proto_rawDescGZIP(), []int{4}
}

func (x *QuoteInfo) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *QuoteInfo) GetLast() float64 {
	if x != nil {
		return x.Last
	}
	return 0
}

func (x *QuoteInfo) GetAskPrice() float64 {
	if x != nil {
		return x.AskPrice
	}
	return 0
}

func (x *QuoteInfo) GetBidPrice() float64 {
	if x != nil {
		return x.BidPrice
	}
	return 0
}

func (x *QuoteInfo) GetAskSize() int32 {
	if x != nil {
		return x.AskSize
	}
	return 0
}

func (x *QuoteInfo) GetBidSize() int32 {
	if x != nil {
		return x.BidSize
	}
	return 0
}

var File_proto_tradier_proto protoreflect.FileDescriptor

var file_proto_tradier_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x69, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x72, 0x61, 0x64, 0x69, 0x65, 0x72, 0x22, 0x15,
	0x0a, 0x13, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x14, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65,
	0x78, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x65, 0x78, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x26, 0x0a, 0x0c, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x22, 0x39, 0x0a, 0x0d, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x69, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x09,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x04, 0x6c, 0x61, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x62, 0x69, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x61, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69,
	0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x69,
	0x64, 0x53, 0x69, 0x7a, 0x65, 0x32, 0x94, 0x01, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x64, 0x69, 0x65,
	0x72, 0x12, 0x38, 0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x74, 0x72, 0x61,
	0x64, 0x69, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x69, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e,
	0x74, 0x72, 0x61, 0x64, 0x69, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x72,
	0x61, 0x64, 0x69, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_tradier_proto_rawDescOnce sync.Once
	file_proto_tradier_proto_rawDescData = file_proto_tradier_proto_rawDesc
)

func file_proto_tradier_proto_rawDescGZIP() []byte {
	file_proto_tradier_proto_rawDescOnce.Do(func() {
		file_proto_tradier_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_tradier_proto_rawDescData)
	})
	return file_proto_tradier_proto_rawDescData
}

var file_proto_tradier_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_tradier_proto_goTypes = []interface{}{
	(*MarketeStateRequest)(nil),  // 0: tradier.MarketeStateRequest
	(*MarketeStateResponse)(nil), // 1: tradier.MarketeStateResponse
	(*QuoteRequest)(nil),         // 2: tradier.QuoteRequest
	(*QuoteResponse)(nil),        // 3: tradier.QuoteResponse
	(*QuoteInfo)(nil),            // 4: tradier.QuoteInfo
}
var file_proto_tradier_proto_depIdxs = []int32{
	4, // 0: tradier.QuoteResponse.quote:type_name -> tradier.QuoteInfo
	2, // 1: tradier.Tradier.Quote:input_type -> tradier.QuoteRequest
	0, // 2: tradier.Tradier.GetMarketState:input_type -> tradier.MarketeStateRequest
	3, // 3: tradier.Tradier.Quote:output_type -> tradier.QuoteResponse
	1, // 4: tradier.Tradier.GetMarketState:output_type -> tradier.MarketeStateResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_tradier_proto_init() }
func file_proto_tradier_proto_init() {
	if File_proto_tradier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_tradier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketeStateRequest); i {
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
		file_proto_tradier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketeStateResponse); i {
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
		file_proto_tradier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteRequest); i {
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
		file_proto_tradier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteResponse); i {
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
		file_proto_tradier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteInfo); i {
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
			RawDescriptor: file_proto_tradier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_tradier_proto_goTypes,
		DependencyIndexes: file_proto_tradier_proto_depIdxs,
		MessageInfos:      file_proto_tradier_proto_msgTypes,
	}.Build()
	File_proto_tradier_proto = out.File
	file_proto_tradier_proto_rawDesc = nil
	file_proto_tradier_proto_goTypes = nil
	file_proto_tradier_proto_depIdxs = nil
}
