// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.1
// source: nakji/data/market/v1beta1/pair.proto

package market

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

// Represents a pairing of tokens, for example trading markets
type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // address or other id [index = true, unique = true,
	// unmarshal.encoding = HEX]
	Token0 []byte `protobuf:"bytes,2,opt,name=token0,proto3" json:"token0,omitempty"` //	byte representation of token address exc. 0x (e.g.
	//\xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2), typically
	// corresponds with nakji.data.Token.id
	Token1 []byte `protobuf:"bytes,3,opt,name=token1,proto3" json:"token1,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nakji_data_market_v1beta1_pair_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_nakji_data_market_v1beta1_pair_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_nakji_data_market_v1beta1_pair_proto_rawDescGZIP(), []int{0}
}

func (x *Pair) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Pair) GetToken0() []byte {
	if x != nil {
		return x.Token0
	}
	return nil
}

func (x *Pair) GetToken1() []byte {
	if x != nil {
		return x.Token1
	}
	return nil
}

var File_nakji_data_market_v1beta1_pair_proto protoreflect.FileDescriptor

var file_nakji_data_market_v1beta1_pair_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6e, 0x61, 0x6b, 0x6a, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x69, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6e, 0x61, 0x6b, 0x6a, 0x69, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x22, 0x46, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x30, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x30, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x31, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x6b, 0x6a, 0x69, 0x2d, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6e, 0x61, 0x6b, 0x6a, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nakji_data_market_v1beta1_pair_proto_rawDescOnce sync.Once
	file_nakji_data_market_v1beta1_pair_proto_rawDescData = file_nakji_data_market_v1beta1_pair_proto_rawDesc
)

func file_nakji_data_market_v1beta1_pair_proto_rawDescGZIP() []byte {
	file_nakji_data_market_v1beta1_pair_proto_rawDescOnce.Do(func() {
		file_nakji_data_market_v1beta1_pair_proto_rawDescData = protoimpl.X.CompressGZIP(file_nakji_data_market_v1beta1_pair_proto_rawDescData)
	})
	return file_nakji_data_market_v1beta1_pair_proto_rawDescData
}

var file_nakji_data_market_v1beta1_pair_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nakji_data_market_v1beta1_pair_proto_goTypes = []interface{}{
	(*Pair)(nil), // 0: nakji.data.market.v1beta1.Pair
}
var file_nakji_data_market_v1beta1_pair_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nakji_data_market_v1beta1_pair_proto_init() }
func file_nakji_data_market_v1beta1_pair_proto_init() {
	if File_nakji_data_market_v1beta1_pair_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nakji_data_market_v1beta1_pair_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
			RawDescriptor: file_nakji_data_market_v1beta1_pair_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nakji_data_market_v1beta1_pair_proto_goTypes,
		DependencyIndexes: file_nakji_data_market_v1beta1_pair_proto_depIdxs,
		MessageInfos:      file_nakji_data_market_v1beta1_pair_proto_msgTypes,
	}.Build()
	File_nakji_data_market_v1beta1_pair_proto = out.File
	file_nakji_data_market_v1beta1_pair_proto_rawDesc = nil
	file_nakji_data_market_v1beta1_pair_proto_goTypes = nil
	file_nakji_data_market_v1beta1_pair_proto_depIdxs = nil
}
