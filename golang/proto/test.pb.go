// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Test
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FOO int32

const (
	FOO_X FOO = 17
)

var FOO_name = map[int32]string{
	17: "X",
}
var FOO_value = map[string]int32{
	"X": 17,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}
func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}
func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}
func (FOO) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Test struct {
	Label            *string             `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type             *int32              `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps             []int64             `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	Optionalgroup    *Test_OptionalGroup `protobuf:"group,4,opt,name=OptionalGroup" json:"optionalgroup,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_Test_Type int32 = 77

func (m *Test) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Test_Type
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func (m *Test) GetOptionalgroup() *Test_OptionalGroup {
	if m != nil {
		return m.Optionalgroup
	}
	return nil
}

type Test_OptionalGroup struct {
	RequiredField    *string `protobuf:"bytes,5,req,name=RequiredField" json:"RequiredField,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Test_OptionalGroup) Reset()                    { *m = Test_OptionalGroup{} }
func (m *Test_OptionalGroup) String() string            { return proto.CompactTextString(m) }
func (*Test_OptionalGroup) ProtoMessage()               {}
func (*Test_OptionalGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Test_OptionalGroup) GetRequiredField() string {
	if m != nil && m.RequiredField != nil {
		return *m.RequiredField
	}
	return ""
}

func init() {
	proto.RegisterType((*Test)(nil), "main.Test")
	proto.RegisterType((*Test_OptionalGroup)(nil), "main.Test.OptionalGroup")
	proto.RegisterEnum("main.FOO", FOO_name, FOO_value)
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0xea, 0x63, 0xe4,
	0x62, 0x09, 0x01, 0x0a, 0x0a, 0xf1, 0x72, 0xb1, 0xe6, 0x24, 0x26, 0xa5, 0xe6, 0x48, 0x30, 0x2a,
	0x30, 0x69, 0x70, 0x0a, 0x09, 0x70, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x29, 0x30, 0x6a,
	0xb0, 0x5a, 0x31, 0x99, 0x9b, 0x0b, 0xf1, 0x70, 0xb1, 0x14, 0xa5, 0x16, 0x14, 0x4b, 0x30, 0x2b,
	0x30, 0x6b, 0x30, 0x0b, 0xe9, 0x73, 0xf1, 0xe6, 0x17, 0x94, 0x64, 0xe6, 0xe7, 0x25, 0xe6, 0xa4,
	0x17, 0xe5, 0x97, 0x16, 0x48, 0xb0, 0x00, 0x15, 0x72, 0x19, 0x49, 0xe8, 0x81, 0x4c, 0xd5, 0x03,
	0x99, 0xa8, 0xe7, 0x0f, 0x95, 0x77, 0x07, 0xc9, 0x4b, 0xa9, 0x71, 0xf1, 0xa2, 0x08, 0x08, 0x89,
	0x72, 0xf1, 0x06, 0xa5, 0x16, 0x96, 0x66, 0x16, 0xa5, 0xa6, 0xb8, 0x65, 0xa6, 0xe6, 0xa4, 0x48,
	0xb0, 0x82, 0x2c, 0xd6, 0xe2, 0xe1, 0x62, 0x76, 0xf3, 0xf7, 0x17, 0x62, 0xe5, 0x62, 0x8c, 0x10,
	0x10, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0x39, 0xd4, 0xb9, 0x2f, 0xb1, 0x00, 0x00, 0x00,
}
