// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: internal/protos/partners.proto

package partners

import (
	orders "github.com/TienMinh25/delivery-system/internal/orders"
	products "github.com/TienMinh25/delivery-system/internal/products"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PartnerID     int64                  `protobuf:"varint,1,opt,name=partnerID,proto3" json:"partnerID,omitempty"`
	TotalAmount   int64                  `protobuf:"varint,2,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
	Products      []*orders.Product      `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	mi := &file_internal_protos_partners_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_partners_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_internal_protos_partners_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetPartnerID() int64 {
	if x != nil {
		return x.PartnerID
	}
	return 0
}

func (x *CheckRequest) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *CheckRequest) GetProducts() []*orders.Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type CheckResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PartnerTitle  string                 `protobuf:"bytes,1,opt,name=partnerTitle,proto3" json:"partnerTitle,omitempty"`
	PartnerBrand  string                 `protobuf:"bytes,2,opt,name=partnerBrand,proto3" json:"partnerBrand,omitempty"`
	Products      []*orders.Product      `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	mi := &file_internal_protos_partners_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_partners_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_internal_protos_partners_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResponse) GetPartnerTitle() string {
	if x != nil {
		return x.PartnerTitle
	}
	return ""
}

func (x *CheckResponse) GetPartnerBrand() string {
	if x != nil {
		return x.PartnerBrand
	}
	return ""
}

func (x *CheckResponse) GetProducts() []*orders.Product {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_internal_protos_partners_proto protoreflect.FileDescriptor

var file_internal_protos_partners_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74,
	0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x22, 0x7d, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x24, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x32, 0x78, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x12,
	0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x0d,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a,
	0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x69, 0x65, 0x6e,
	0x4d, 0x69, 0x6e, 0x68, 0x32, 0x35, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2d,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_internal_protos_partners_proto_rawDescOnce sync.Once
	file_internal_protos_partners_proto_rawDescData []byte
)

func file_internal_protos_partners_proto_rawDescGZIP() []byte {
	file_internal_protos_partners_proto_rawDescOnce.Do(func() {
		file_internal_protos_partners_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_protos_partners_proto_rawDesc), len(file_internal_protos_partners_proto_rawDesc)))
	})
	return file_internal_protos_partners_proto_rawDescData
}

var file_internal_protos_partners_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_protos_partners_proto_goTypes = []any{
	(*CheckRequest)(nil),            // 0: CheckRequest
	(*CheckResponse)(nil),           // 1: CheckResponse
	(*orders.Product)(nil),          // 2: Product
	(*products.GetAllRequest)(nil),  // 3: GetAllRequest
	(*products.GetAllResponse)(nil), // 4: GetAllResponse
}
var file_internal_protos_partners_proto_depIdxs = []int32{
	2, // 0: CheckRequest.products:type_name -> Product
	2, // 1: CheckResponse.products:type_name -> Product
	3, // 2: Partners.GetPartnerProducts:input_type -> GetAllRequest
	0, // 3: Partners.CheckPartnerProducts:input_type -> CheckRequest
	4, // 4: Partners.GetPartnerProducts:output_type -> GetAllResponse
	1, // 5: Partners.CheckPartnerProducts:output_type -> CheckResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_protos_partners_proto_init() }
func file_internal_protos_partners_proto_init() {
	if File_internal_protos_partners_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_protos_partners_proto_rawDesc), len(file_internal_protos_partners_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_protos_partners_proto_goTypes,
		DependencyIndexes: file_internal_protos_partners_proto_depIdxs,
		MessageInfos:      file_internal_protos_partners_proto_msgTypes,
	}.Build()
	File_internal_protos_partners_proto = out.File
	file_internal_protos_partners_proto_goTypes = nil
	file_internal_protos_partners_proto_depIdxs = nil
}
