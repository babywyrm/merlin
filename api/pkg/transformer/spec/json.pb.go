// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: transformer/spec/json.proto

package spec

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

type JsonOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsonTemplate *JsonTemplate `protobuf:"bytes,1,opt,name=jsonTemplate,proto3" json:"jsonTemplate,omitempty"`
}

func (x *JsonOutput) Reset() {
	*x = JsonOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_json_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonOutput) ProtoMessage() {}

func (x *JsonOutput) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_json_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonOutput.ProtoReflect.Descriptor instead.
func (*JsonOutput) Descriptor() ([]byte, []int) {
	return file_transformer_spec_json_proto_rawDescGZIP(), []int{0}
}

func (x *JsonOutput) GetJsonTemplate() *JsonTemplate {
	if x != nil {
		return x.JsonTemplate
	}
	return nil
}

type JsonTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseJson *BaseJson `protobuf:"bytes,1,opt,name=baseJson,proto3" json:"baseJson,omitempty"`
	Fields   []*Field  `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *JsonTemplate) Reset() {
	*x = JsonTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_json_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonTemplate) ProtoMessage() {}

func (x *JsonTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_json_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonTemplate.ProtoReflect.Descriptor instead.
func (*JsonTemplate) Descriptor() ([]byte, []int) {
	return file_transformer_spec_json_proto_rawDescGZIP(), []int{1}
}

func (x *JsonTemplate) GetBaseJson() *BaseJson {
	if x != nil {
		return x.BaseJson
	}
	return nil
}

func (x *JsonTemplate) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type BaseJson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to BaseJson:
	//	*BaseJson_FromJson
	//	*BaseJson_FromTable
	BaseJson isBaseJson_BaseJson `protobuf_oneof:"baseJson"`
}

func (x *BaseJson) Reset() {
	*x = BaseJson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_json_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseJson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseJson) ProtoMessage() {}

func (x *BaseJson) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_json_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseJson.ProtoReflect.Descriptor instead.
func (*BaseJson) Descriptor() ([]byte, []int) {
	return file_transformer_spec_json_proto_rawDescGZIP(), []int{2}
}

func (m *BaseJson) GetBaseJson() isBaseJson_BaseJson {
	if m != nil {
		return m.BaseJson
	}
	return nil
}

func (x *BaseJson) GetFromJson() *FromJson {
	if x, ok := x.GetBaseJson().(*BaseJson_FromJson); ok {
		return x.FromJson
	}
	return nil
}

func (x *BaseJson) GetFromTable() *FromTable {
	if x, ok := x.GetBaseJson().(*BaseJson_FromTable); ok {
		return x.FromTable
	}
	return nil
}

type isBaseJson_BaseJson interface {
	isBaseJson_BaseJson()
}

type BaseJson_FromJson struct {
	FromJson *FromJson `protobuf:"bytes,1,opt,name=fromJson,proto3,oneof"`
}

type BaseJson_FromTable struct {
	FromTable *FromTable `protobuf:"bytes,2,opt,name=fromTable,proto3,oneof"`
}

func (*BaseJson_FromJson) isBaseJson_BaseJson() {}

func (*BaseJson_FromTable) isBaseJson_BaseJson() {}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []*Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	// Types that are assignable to Value:
	//	*Field_FromJson
	//	*Field_FromTable
	//	*Field_Expression
	Value isField_Value `protobuf_oneof:"value"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_json_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_json_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_transformer_spec_json_proto_rawDescGZIP(), []int{3}
}

func (x *Field) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (m *Field) GetValue() isField_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Field) GetFromJson() *FromJson {
	if x, ok := x.GetValue().(*Field_FromJson); ok {
		return x.FromJson
	}
	return nil
}

func (x *Field) GetFromTable() *FromTable {
	if x, ok := x.GetValue().(*Field_FromTable); ok {
		return x.FromTable
	}
	return nil
}

func (x *Field) GetExpression() string {
	if x, ok := x.GetValue().(*Field_Expression); ok {
		return x.Expression
	}
	return ""
}

type isField_Value interface {
	isField_Value()
}

type Field_FromJson struct {
	FromJson *FromJson `protobuf:"bytes,2,opt,name=fromJson,proto3,oneof"`
}

type Field_FromTable struct {
	FromTable *FromTable `protobuf:"bytes,3,opt,name=fromTable,proto3,oneof"`
}

type Field_Expression struct {
	Expression string `protobuf:"bytes,4,opt,name=expression,proto3,oneof"`
}

func (*Field_FromJson) isField_Value() {}

func (*Field_FromTable) isField_Value() {}

func (*Field_Expression) isField_Value() {}

var File_transformer_spec_json_proto protoreflect.FileDescriptor

var file_transformer_spec_json_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2f, 0x73, 0x70,
	0x65, 0x63, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6d,
	0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x72, 0x1a, 0x1d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2f, 0x73,
	0x70, 0x65, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x52, 0x0a, 0x0a, 0x4a, 0x73, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x44,
	0x0a, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x22, 0x7b, 0x0a, 0x0c, 0x4a, 0x73, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x4a, 0x73, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x31,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x22, 0x91, 0x01, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x3a,
	0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x09, 0x66, 0x72,
	0x6f, 0x6d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x65, 0x72, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x09,
	0x66, 0x72, 0x6f, 0x6d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x62, 0x61, 0x73,
	0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x22, 0xe0, 0x01, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x31, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x3a, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x4a, 0x73,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x3d,
	0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x48, 0x00, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6a, 0x65, 0x6b, 0x2f, 0x6d, 0x65, 0x72,
	0x6c, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x65, 0x72, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transformer_spec_json_proto_rawDescOnce sync.Once
	file_transformer_spec_json_proto_rawDescData = file_transformer_spec_json_proto_rawDesc
)

func file_transformer_spec_json_proto_rawDescGZIP() []byte {
	file_transformer_spec_json_proto_rawDescOnce.Do(func() {
		file_transformer_spec_json_proto_rawDescData = protoimpl.X.CompressGZIP(file_transformer_spec_json_proto_rawDescData)
	})
	return file_transformer_spec_json_proto_rawDescData
}

var file_transformer_spec_json_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_transformer_spec_json_proto_goTypes = []interface{}{
	(*JsonOutput)(nil),   // 0: merlin.spec.JsonOutput
	(*JsonTemplate)(nil), // 1: merlin.spec.JsonTemplate
	(*BaseJson)(nil),     // 2: merlin.spec.BaseJson
	(*Field)(nil),        // 3: merlin.spec.Field
	(*FromJson)(nil),     // 4: merlin.spec.FromJson
	(*FromTable)(nil),    // 5: merlin.spec.FromTable
}
var file_transformer_spec_json_proto_depIdxs = []int32{
	1, // 0: merlin.spec.JsonOutput.jsonTemplate:type_name -> merlin.spec.JsonTemplate
	2, // 1: merlin.spec.JsonTemplate.baseJson:type_name -> merlin.spec.BaseJson
	3, // 2: merlin.spec.JsonTemplate.fields:type_name -> merlin.spec.Field
	4, // 3: merlin.spec.BaseJson.fromJson:type_name -> merlin.spec.FromJson
	5, // 4: merlin.spec.BaseJson.fromTable:type_name -> merlin.spec.FromTable
	3, // 5: merlin.spec.Field.fields:type_name -> merlin.spec.Field
	4, // 6: merlin.spec.Field.fromJson:type_name -> merlin.spec.FromJson
	5, // 7: merlin.spec.Field.fromTable:type_name -> merlin.spec.FromTable
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_transformer_spec_json_proto_init() }
func file_transformer_spec_json_proto_init() {
	if File_transformer_spec_json_proto != nil {
		return
	}
	file_transformer_spec_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transformer_spec_json_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonOutput); i {
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
		file_transformer_spec_json_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonTemplate); i {
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
		file_transformer_spec_json_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseJson); i {
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
		file_transformer_spec_json_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
	file_transformer_spec_json_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*BaseJson_FromJson)(nil),
		(*BaseJson_FromTable)(nil),
	}
	file_transformer_spec_json_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Field_FromJson)(nil),
		(*Field_FromTable)(nil),
		(*Field_Expression)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transformer_spec_json_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transformer_spec_json_proto_goTypes,
		DependencyIndexes: file_transformer_spec_json_proto_depIdxs,
		MessageInfos:      file_transformer_spec_json_proto_msgTypes,
	}.Build()
	File_transformer_spec_json_proto = out.File
	file_transformer_spec_json_proto_rawDesc = nil
	file_transformer_spec_json_proto_goTypes = nil
	file_transformer_spec_json_proto_depIdxs = nil
}