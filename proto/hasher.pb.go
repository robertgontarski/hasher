// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: proto/hasher.proto

package proto

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

type HashEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *HashEmailRequest) Reset() {
	*x = HashEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hasher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashEmailRequest) ProtoMessage() {}

func (x *HashEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hasher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashEmailRequest.ProtoReflect.Descriptor instead.
func (*HashEmailRequest) Descriptor() ([]byte, []int) {
	return file_proto_hasher_proto_rawDescGZIP(), []int{0}
}

func (x *HashEmailRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HashEmailRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type HashPhoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Number      string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	CountryCode string `protobuf:"bytes,3,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
}

func (x *HashPhoneRequest) Reset() {
	*x = HashPhoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hasher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashPhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashPhoneRequest) ProtoMessage() {}

func (x *HashPhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hasher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashPhoneRequest.ProtoReflect.Descriptor instead.
func (*HashPhoneRequest) Descriptor() ([]byte, []int) {
	return file_proto_hasher_proto_rawDescGZIP(), []int{1}
}

func (x *HashPhoneRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HashPhoneRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *HashPhoneRequest) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

type HashNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
}

func (x *HashNameRequest) Reset() {
	*x = HashNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hasher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashNameRequest) ProtoMessage() {}

func (x *HashNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hasher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashNameRequest.ProtoReflect.Descriptor instead.
func (*HashNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_hasher_proto_rawDescGZIP(), []int{2}
}

func (x *HashNameRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HashNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HashNameRequest) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

type HashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *HashResponse) Reset() {
	*x = HashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hasher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashResponse) ProtoMessage() {}

func (x *HashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hasher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashResponse.ProtoReflect.Descriptor instead.
func (*HashResponse) Descriptor() ([]byte, []int) {
	return file_proto_hasher_proto_rawDescGZIP(), []int{3}
}

func (x *HashResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *HashResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HashResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_hasher_proto protoreflect.FileDescriptor

var file_proto_hasher_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x10,
	0x48, 0x61, 0x73, 0x68, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5d, 0x0a, 0x10, 0x48, 0x61,
	0x73, 0x68, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x4f, 0x0a, 0x0f, 0x48, 0x61, 0x73,
	0x68, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x0c, 0x48, 0x61,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0xca, 0x01, 0x0a, 0x0d, 0x48, 0x61, 0x73, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x48, 0x61, 0x73, 0x68, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x18, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x65, 0x72, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x09, 0x48, 0x61, 0x73, 0x68, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x65,
	0x72, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x08, 0x48, 0x61, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x62, 0x65,
	0x72, 0x74, 0x67, 0x6f, 0x6e, 0x74, 0x61, 0x72, 0x73, 0x6b, 0x69, 0x2f, 0x68, 0x61, 0x73, 0x68,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hasher_proto_rawDescOnce sync.Once
	file_proto_hasher_proto_rawDescData = file_proto_hasher_proto_rawDesc
)

func file_proto_hasher_proto_rawDescGZIP() []byte {
	file_proto_hasher_proto_rawDescOnce.Do(func() {
		file_proto_hasher_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hasher_proto_rawDescData)
	})
	return file_proto_hasher_proto_rawDescData
}

var file_proto_hasher_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_hasher_proto_goTypes = []interface{}{
	(*HashEmailRequest)(nil), // 0: hasher.HashEmailRequest
	(*HashPhoneRequest)(nil), // 1: hasher.HashPhoneRequest
	(*HashNameRequest)(nil),  // 2: hasher.HashNameRequest
	(*HashResponse)(nil),     // 3: hasher.HashResponse
}
var file_proto_hasher_proto_depIdxs = []int32{
	0, // 0: hasher.HasherService.HashEmail:input_type -> hasher.HashEmailRequest
	1, // 1: hasher.HasherService.HashPhone:input_type -> hasher.HashPhoneRequest
	2, // 2: hasher.HasherService.HashName:input_type -> hasher.HashNameRequest
	3, // 3: hasher.HasherService.HashEmail:output_type -> hasher.HashResponse
	3, // 4: hasher.HasherService.HashPhone:output_type -> hasher.HashResponse
	3, // 5: hasher.HasherService.HashName:output_type -> hasher.HashResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_hasher_proto_init() }
func file_proto_hasher_proto_init() {
	if File_proto_hasher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hasher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashEmailRequest); i {
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
		file_proto_hasher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashPhoneRequest); i {
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
		file_proto_hasher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashNameRequest); i {
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
		file_proto_hasher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashResponse); i {
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
			RawDescriptor: file_proto_hasher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hasher_proto_goTypes,
		DependencyIndexes: file_proto_hasher_proto_depIdxs,
		MessageInfos:      file_proto_hasher_proto_msgTypes,
	}.Build()
	File_proto_hasher_proto = out.File
	file_proto_hasher_proto_rawDesc = nil
	file_proto_hasher_proto_goTypes = nil
	file_proto_hasher_proto_depIdxs = nil
}