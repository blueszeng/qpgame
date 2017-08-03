// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	Person
	AddressBook
*/
package msg

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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// [START messages]
type Person struct {
	Name   string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id     int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email  string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=msg.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "msg.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "msg.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "msg.AddressBook")
	proto.RegisterEnum("msg.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)

}

func init() { proto.RegisterFile("msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0xb5, 0x5d, 0x17, 0xdc, 0x2d, 0x8c, 0x19, 0xa6, 0x16, 0x9f, 0xc6, 0xf4, 0x61, 0x22, 0x44,
	0x98, 0x5f, 0x60, 0xa1, 0xa8, 0xe8, 0x6c, 0x09, 0x82, 0xcf, 0xad, 0xbd, 0xd6, 0x62, 0xd3, 0x84,
	0xa6, 0x05, 0xf7, 0x4b, 0x7e, 0x9e, 0x5f, 0x20, 0x49, 0x8b, 0x14, 0x7c, 0x3b, 0xe7, 0x9e, 0x93,
	0x73, 0x92, 0x5c, 0x98, 0x09, 0x5d, 0x30, 0xd5, 0xc8, 0x56, 0xd2, 0x89, 0xd0, 0xc5, 0xfa, 0xc7,
	0x01, 0x92, 0x60, 0xa3, 0x65, 0x4d, 0x29, 0x78, 0x75, 0x2a, 0x30, 0x70, 0x56, 0xce, 0x66, 0xc6,
	0x2d, 0xa6, 0x73, 0x70, 0xcb, 0x3c, 0x70, 0x57, 0xce, 0x66, 0xca, 0xdd, 0x32, 0xa7, 0x4b, 0x98,
	0xa2, 0x48, 0xcb, 0x2a, 0x98, 0x58, 0x53, 0x4f, 0xe8, 0x35, 0x10, 0xf5, 0x21, 0x6b, 0xd4, 0x81,
	0xb7, 0x9a, 0x6c, 0xfc, 0xed, 0x29, 0x33, 0x2d, 0x7d, 0x2c, 0x4b, 0x8c, 0xf2, 0xdc, 0x89, 0x0c,
	0x1b, 0x3e, 0xd8, 0xce, 0x12, 0xf0, 0x47, 0x63, 0x7a, 0x02, 0xa4, 0xb6, 0x68, 0xe8, 0x1e, 0x18,
	0xbd, 0x04, 0xaf, 0xdd, 0x2b, 0xb4, 0xfd, 0xf3, 0xed, 0xf1, 0xbf, 0xd4, 0x97, 0xbd, 0x42, 0x6e,
	0x2d, 0xeb, 0x2b, 0x98, 0xfd, 0x8d, 0x28, 0x00, 0xd9, 0xc5, 0xe1, 0xc3, 0x53, 0xb4, 0x38, 0xa0,
	0x87, 0xe0, 0xdd, 0xc7, 0xbb, 0x68, 0xe1, 0x18, 0xf4, 0x1a, 0xf3, 0xc7, 0x85, 0xbb, 0xde, 0x82,
	0x7f, 0x9b, 0xe7, 0x0d, 0x6a, 0x1d, 0x4a, 0xf9, 0x49, 0xcf, 0x81, 0x28, 0x94, 0xaa, 0x32, 0x4f,
	0x37, 0xd7, 0xf7, 0x47, 0x45, 0x7c, 0x90, 0xc2, 0x04, 0x96, 0x6f, 0x52, 0x30, 0xfc, 0x4a, 0x85,
	0xaa, 0x90, 0xb5, 0x5d, 0x2b, 0x9b, 0x32, 0xad, 0xc2, 0xa3, 0x51, 0x52, 0x62, 0xfe, 0x55, 0x7f,
	0xbb, 0x17, 0x77, 0x52, 0x16, 0x15, 0x32, 0xcb, 0xb3, 0xee, 0x9d, 0x45, 0xfd, 0x29, 0xcd, 0x46,
	0xe6, 0x8c, 0xd8, 0x35, 0xdc, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xed, 0xa3, 0x63, 0x93,
	0x01, 0x00, 0x00,
}
