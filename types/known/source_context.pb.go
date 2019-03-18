// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/source_context.proto

package known_proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

// `SourceContext` represents information about the source of a
// protobuf element, like the file in which it is defined.
type SourceContext struct {
	// The path-qualified name of the .proto file that contained the associated
	// protobuf element.  For example: `"google/protobuf/source_context.proto"`.
	FileName             string   `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceContext) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_source_context_proto_messageTypes[0].MessageOf(m)
}
func (m *SourceContext) Reset()         { *m = SourceContext{} }
func (m *SourceContext) String() string { return proto.CompactTextString(m) }
func (*SourceContext) ProtoMessage()    {}

// Deprecated: Use SourceContext.ProtoReflect.Type instead.
func (*SourceContext) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_source_context_proto_rawdesc_gzipped, []int{0}
}

func (m *SourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceContext.Unmarshal(m, b)
}
func (m *SourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceContext.Marshal(b, m, deterministic)
}
func (m *SourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceContext.Merge(m, src)
}
func (m *SourceContext) XXX_Size() int {
	return xxx_messageInfo_SourceContext.Size(m)
}
func (m *SourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_SourceContext proto.InternalMessageInfo

func (m *SourceContext) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func init() {
	proto.RegisterFile("google/protobuf/source_context.proto", xxx_File_google_protobuf_source_context_proto_rawdesc_gzipped)
	proto.RegisterType((*SourceContext)(nil), "google.protobuf.SourceContext")
}

var xxx_File_google_protobuf_source_context_proto_rawdesc = []byte{
	// 249 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x2c, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x89, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x12, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x3b, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x47, 0x50, 0x42,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var xxx_File_google_protobuf_source_context_proto_rawdesc_gzipped = protoimpl.X.CompressGZIP(xxx_File_google_protobuf_source_context_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_google_protobuf_source_context_proto protoreflect.FileDescriptor

var xxx_File_google_protobuf_source_context_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_google_protobuf_source_context_proto_goTypes = []interface{}{
	(*SourceContext)(nil), // 0: google.protobuf.SourceContext
}
var xxx_File_google_protobuf_source_context_proto_depIdxs = []int32{}

func init() { xxx_File_google_protobuf_source_context_proto_init() }
func xxx_File_google_protobuf_source_context_proto_init() {
	if File_google_protobuf_source_context_proto != nil {
		return
	}
	messageTypes := make([]protoreflect.MessageType, 1)
	File_google_protobuf_source_context_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_google_protobuf_source_context_proto_rawdesc,
		GoTypes:            xxx_File_google_protobuf_source_context_proto_goTypes,
		DependencyIndexes:  xxx_File_google_protobuf_source_context_proto_depIdxs,
		MessageOutputTypes: messageTypes,
	}.Init()
	messageGoTypes := xxx_File_google_protobuf_source_context_proto_goTypes[0:][:1]
	for i, mt := range messageTypes {
		xxx_File_google_protobuf_source_context_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_google_protobuf_source_context_proto_messageTypes[i].PBType = mt
	}
	xxx_File_google_protobuf_source_context_proto_goTypes = nil
	xxx_File_google_protobuf_source_context_proto_depIdxs = nil
}