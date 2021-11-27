// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: person/person.proto

package person

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

type Person_Gender int32

const (
	Person_MAN    Person_Gender = 0
	Person_MALE   Person_Gender = 0
	Person_WOMAN  Person_Gender = 1
	Person_FEMALE Person_Gender = 1
	Person_OTHER  Person_Gender = 2
)

// Enum value maps for Person_Gender.
var (
	Person_Gender_name = map[int32]string{
		0: "MAN",
		// Duplicate value: 0: "MALE",
		1: "WOMAN",
		// Duplicate value: 1: "FEMALE",
		2: "OTHER",
	}
	Person_Gender_value = map[string]int32{
		"MAN":    0,
		"MALE":   0,
		"WOMAN":  1,
		"FEMALE": 1,
		"OTHER":  2,
	}
)

func (x Person_Gender) Enum() *Person_Gender {
	p := new(Person_Gender)
	*p = x
	return p
}

func (x Person_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Person_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_person_person_proto_enumTypes[0].Descriptor()
}

func (Person_Gender) Type() protoreflect.EnumType {
	return &file_person_person_proto_enumTypes[0]
}

func (x Person_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Person_Gender.Descriptor instead.
func (Person_Gender) EnumDescriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{0, 0}
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age       int32             `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Gender    Person_Gender     `protobuf:"varint,3,opt,name=gender,proto3,enum=person.Person_Gender" json:"gender,omitempty"`
	TestSlice []string          `protobuf:"bytes,4,rep,name=testSlice,proto3" json:"testSlice,omitempty"`                                                                                     // 切片
	TestMap   map[string]string `protobuf:"bytes,5,rep,name=testMap,proto3" json:"testMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map <key:int/string, value: any type>
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_person_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetGender() Person_Gender {
	if x != nil {
		return x.Gender
	}
	return Person_MAN
}

func (x *Person) GetTestSlice() []string {
	if x != nil {
		return x.TestSlice
	}
	return nil
}

func (x *Person) GetTestMap() map[string]string {
	if x != nil {
		return x.TestMap
	}
	return nil
}

type Home struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persons []*Person     `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	Visitor *Home_Visitor `protobuf:"bytes,2,opt,name=visitor,proto3" json:"visitor,omitempty"`
}

func (x *Home) Reset() {
	*x = Home{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Home) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Home) ProtoMessage() {}

func (x *Home) ProtoReflect() protoreflect.Message {
	mi := &file_person_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Home.ProtoReflect.Descriptor instead.
func (*Home) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{1}
}

func (x *Home) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

func (x *Home) GetVisitor() *Home_Visitor {
	if x != nil {
		return x.Visitor
	}
	return nil
}

type Home_Visitor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Home_Visitor) Reset() {
	*x = Home_Visitor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_person_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Home_Visitor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Home_Visitor) ProtoMessage() {}

func (x *Home_Visitor) ProtoReflect() protoreflect.Message {
	mi := &file_person_person_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Home_Visitor.ProtoReflect.Descriptor instead.
func (*Home_Visitor) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Home_Visitor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_person_person_proto protoreflect.FileDescriptor

var file_person_person_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0xb1, 0x02,
	0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x2d,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x65, 0x73, 0x74, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x74,
	0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x74, 0x65, 0x73, 0x74, 0x4d,
	0x61, 0x70, 0x1a, 0x3a, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x41,
	0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x4e, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x57,
	0x4f, 0x4d, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x02, 0x1a, 0x02, 0x10,
	0x01, 0x22, 0x7f, 0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x48, 0x6f,
	0x6d, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x1a, 0x1d, 0x0a, 0x07, 0x56, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x6c, 0x65, 0x61,
	0x72, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_person_person_proto_rawDescOnce sync.Once
	file_person_person_proto_rawDescData = file_person_person_proto_rawDesc
)

func file_person_person_proto_rawDescGZIP() []byte {
	file_person_person_proto_rawDescOnce.Do(func() {
		file_person_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_person_proto_rawDescData)
	})
	return file_person_person_proto_rawDescData
}

var file_person_person_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_person_person_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_person_person_proto_goTypes = []interface{}{
	(Person_Gender)(0),   // 0: person.Person.Gender
	(*Person)(nil),       // 1: person.Person
	(*Home)(nil),         // 2: person.Home
	nil,                  // 3: person.Person.TestMapEntry
	(*Home_Visitor)(nil), // 4: person.Home.Visitor
}
var file_person_person_proto_depIdxs = []int32{
	0, // 0: person.Person.gender:type_name -> person.Person.Gender
	3, // 1: person.Person.testMap:type_name -> person.Person.TestMapEntry
	1, // 2: person.Home.persons:type_name -> person.Person
	4, // 3: person.Home.visitor:type_name -> person.Home.Visitor
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_person_person_proto_init() }
func file_person_person_proto_init() {
	if File_person_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_person_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_person_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Home); i {
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
		file_person_person_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Home_Visitor); i {
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
			RawDescriptor: file_person_person_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_person_person_proto_goTypes,
		DependencyIndexes: file_person_person_proto_depIdxs,
		EnumInfos:         file_person_person_proto_enumTypes,
		MessageInfos:      file_person_person_proto_msgTypes,
	}.Build()
	File_person_person_proto = out.File
	file_person_person_proto_rawDesc = nil
	file_person_person_proto_goTypes = nil
	file_person_person_proto_depIdxs = nil
}