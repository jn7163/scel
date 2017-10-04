// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scel_data.proto

/*
Package sceldata is a generated protocol buffer package.

It is generated from these files:
	scel_data.proto

It has these top-level messages:
	ScelData
	Word
	Info
*/
package sceldata

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

type ScelData struct {
	Info    *Info    `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Pinyins []string `protobuf:"bytes,2,rep,name=pinyins" json:"pinyins,omitempty"`
	Words   []*Word  `protobuf:"bytes,3,rep,name=words" json:"words,omitempty"`
}

func (m *ScelData) Reset()                    { *m = ScelData{} }
func (m *ScelData) String() string            { return proto.CompactTextString(m) }
func (*ScelData) ProtoMessage()               {}
func (*ScelData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ScelData) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ScelData) GetPinyins() []string {
	if m != nil {
		return m.Pinyins
	}
	return nil
}

func (m *ScelData) GetWords() []*Word {
	if m != nil {
		return m.Words
	}
	return nil
}

type Word struct {
	Pinyins []int32  `protobuf:"varint,1,rep,packed,name=pinyins" json:"pinyins,omitempty"`
	Words   []string `protobuf:"bytes,2,rep,name=words" json:"words,omitempty"`
	Exts    [][]byte `protobuf:"bytes,3,rep,name=exts,proto3" json:"exts,omitempty"`
}

func (m *Word) Reset()                    { *m = Word{} }
func (m *Word) String() string            { return proto.CompactTextString(m) }
func (*Word) ProtoMessage()               {}
func (*Word) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Word) GetPinyins() []int32 {
	if m != nil {
		return m.Pinyins
	}
	return nil
}

func (m *Word) GetWords() []string {
	if m != nil {
		return m.Words
	}
	return nil
}

func (m *Word) GetExts() [][]byte {
	if m != nil {
		return m.Exts
	}
	return nil
}

type Info struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Example     string `protobuf:"bytes,4,opt,name=example" json:"example,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Info) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Info) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Info) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Info) GetExample() string {
	if m != nil {
		return m.Example
	}
	return ""
}

func init() {
	proto.RegisterType((*ScelData)(nil), "wener.teleattr.v1.ScelData")
	proto.RegisterType((*Word)(nil), "wener.teleattr.v1.Word")
	proto.RegisterType((*Info)(nil), "wener.teleattr.v1.Info")
}

func init() { proto.RegisterFile("scel_data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x4f, 0x4b, 0x3b, 0x31,
	0x10, 0x65, 0xbb, 0xdb, 0xdf, 0xcf, 0xa6, 0x8a, 0x18, 0x04, 0x73, 0x5c, 0xea, 0xa5, 0x20, 0x66,
	0xa9, 0xe2, 0xc9, 0x93, 0xc5, 0x4b, 0x3d, 0x2d, 0xeb, 0xa1, 0xe0, 0x45, 0xd2, 0xdd, 0x69, 0x8d,
	0xec, 0x26, 0x4b, 0x36, 0xf6, 0xcf, 0xcd, 0xcf, 0xe2, 0x27, 0x95, 0x49, 0x1a, 0x2c, 0x88, 0xa7,
	0xbc, 0x37, 0xf3, 0xe6, 0x65, 0xfe, 0x90, 0xd3, 0xae, 0x84, 0xfa, 0xb5, 0x12, 0x56, 0xf0, 0xd6,
	0x68, 0xab, 0xe9, 0xd9, 0x06, 0x14, 0x18, 0x6e, 0xa1, 0x06, 0x61, 0xad, 0xe1, 0xeb, 0xc9, 0xe8,
	0x33, 0x22, 0x47, 0xcf, 0x25, 0xd4, 0x8f, 0xc2, 0x0a, 0x7a, 0x45, 0x12, 0xa9, 0x96, 0x9a, 0x45,
	0x69, 0x34, 0x1e, 0xde, 0x5c, 0xf0, 0x5f, 0x72, 0x3e, 0x53, 0x4b, 0x5d, 0x38, 0x11, 0x65, 0xe4,
	0x7f, 0x2b, 0xd5, 0x4e, 0xaa, 0x8e, 0xf5, 0xd2, 0x78, 0x3c, 0x28, 0x02, 0xa5, 0xd7, 0xa4, 0xbf,
	0xd1, 0xa6, 0xea, 0x58, 0x9c, 0xc6, 0x7f, 0xf8, 0xcc, 0xb5, 0xa9, 0x0a, 0xaf, 0x1a, 0x3d, 0x91,
	0x04, 0xe9, 0xa1, 0x61, 0x94, 0xc6, 0xe3, 0xfe, 0x8f, 0xe1, 0x79, 0x30, 0xf4, 0x1f, 0x79, 0x42,
	0x29, 0x49, 0x60, 0x6b, 0xfd, 0x2f, 0xc7, 0x85, 0xc3, 0xa3, 0x77, 0x92, 0x60, 0x8b, 0x98, 0x53,
	0xa2, 0x01, 0x37, 0xc9, 0xa0, 0x70, 0x18, 0x63, 0x76, 0xd7, 0x02, 0xeb, 0xf9, 0x18, 0x62, 0x9a,
	0x92, 0x61, 0x05, 0x5d, 0x69, 0x64, 0x6b, 0xa5, 0x56, 0x2c, 0x76, 0xa9, 0xc3, 0x10, 0x76, 0x05,
	0x5b, 0xd1, 0xb4, 0x35, 0xb0, 0xc4, 0x65, 0x03, 0x9d, 0x5a, 0x72, 0x59, 0xea, 0x86, 0xaf, 0xa4,
	0x7d, 0xfb, 0x58, 0xf8, 0x19, 0x1b, 0xe0, 0xb8, 0x73, 0x9c, 0x10, 0x5f, 0x5c, 0xfd, 0xf4, 0x24,
	0xac, 0x37, 0xc7, 0x1b, 0xe4, 0xd1, 0xcb, 0xdd, 0xbe, 0xa2, 0xd4, 0x4d, 0xb6, 0xaf, 0xca, 0x50,
	0x9d, 0xad, 0x40, 0xb9, 0x3b, 0x65, 0xeb, 0x49, 0x16, 0xca, 0xef, 0x03, 0xf8, 0xea, 0x25, 0xf3,
	0x87, 0x7c, 0xb6, 0xf8, 0xe7, 0x24, 0xb7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x62, 0x76, 0x8d,
	0x54, 0xdd, 0x01, 0x00, 0x00,
}
