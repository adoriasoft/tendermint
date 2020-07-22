// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/google/api/annotations.proto

package annotations

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*HttpRule)(nil),
	Field:         72295728,
	Name:          "google.api.http",
	Tag:           "bytes,72295728,opt,name=http",
	Filename:      "tendermint/google/api/annotations.proto",
}

func init() {
	proto.RegisterExtension(E_Http)
}

func init() {
	proto.RegisterFile("tendermint/google/api/annotations.proto", fileDescriptor_813e733033efb621)
}

var fileDescriptor_813e733033efb621 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c,
	0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0xc8, 0xea, 0x25, 0x16, 0x64, 0x4a, 0x29, 0x60, 0xd7, 0x94,
	0x51, 0x52, 0x52, 0x00, 0x51, 0x2d, 0xa5, 0x00, 0x15, 0x06, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53,
	0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0xf2, 0x8b, 0x20, 0x2a, 0xac, 0xbc, 0xb9, 0x58, 0x40,
	0xea, 0x85, 0xe4, 0xf4, 0xa0, 0x06, 0xc3, 0x94, 0xea, 0xf9, 0xa6, 0x96, 0x64, 0xe4, 0xa7, 0xf8,
	0x17, 0x80, 0x6d, 0x97, 0xd8, 0x70, 0x6a, 0x8f, 0x92, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x88, 0x1e,
	0xc2, 0x05, 0x7a, 0x1e, 0x25, 0x25, 0x05, 0x41, 0xa5, 0x39, 0xa9, 0x41, 0x60, 0x43, 0x9c, 0xda,
	0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x81, 0x8b, 0x2f, 0x39, 0x3f, 0x17,
	0x49, 0xa7, 0x93, 0x80, 0x23, 0xc2, 0x6b, 0x01, 0x20, 0x2b, 0x03, 0x18, 0xa3, 0x1c, 0xa1, 0xf2,
	0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0xfa, 0xe9, 0xa9, 0x79, 0x60, 0x07,
	0x41, 0xfd, 0x97, 0x58, 0x90, 0x59, 0x8c, 0x1e, 0x30, 0xd6, 0x48, 0xec, 0x45, 0x4c, 0x2c, 0xee,
	0x8e, 0x01, 0x9e, 0x49, 0x6c, 0x60, 0x4d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x38,
	0x23, 0x0e, 0x57, 0x01, 0x00, 0x00,
}
