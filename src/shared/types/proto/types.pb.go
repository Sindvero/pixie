// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/shared/types/proto/types.proto

package typespb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type DataType int32

const (
	DATA_TYPE_UNKNOWN DataType = 0
	BOOLEAN           DataType = 1
	INT64             DataType = 2
	FLOAT64           DataType = 3
	STRING            DataType = 4
	TIME64NS          DataType = 5
)

var DataType_name = map[int32]string{
	0: "DATA_TYPE_UNKNOWN",
	1: "BOOLEAN",
	2: "INT64",
	3: "FLOAT64",
	4: "STRING",
	5: "TIME64NS",
}
var DataType_value = map[string]int32{
	"DATA_TYPE_UNKNOWN": 0,
	"BOOLEAN":           1,
	"INT64":             2,
	"FLOAT64":           3,
	"STRING":            4,
	"TIME64NS":          5,
}

func (DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_types_923ebdc07becc380, []int{0}
}

func init() {
	proto.RegisterEnum("pl.types.DataType", DataType_name, DataType_value)
}
func (x DataType) String() string {
	s, ok := DataType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() {
	proto.RegisterFile("src/shared/types/proto/types.proto", fileDescriptor_types_923ebdc07becc380)
}

var fileDescriptor_types_923ebdc07becc380 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x2e, 0x4a, 0xd6,
	0x2f, 0xce, 0x48, 0x2c, 0x4a, 0x4d, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x28, 0xca,
	0x2f, 0xc9, 0x87, 0xb0, 0xf5, 0xc0, 0x6c, 0x21, 0x8e, 0x82, 0x1c, 0x3d, 0x30, 0x5f, 0x2b, 0x81,
	0x8b, 0xc3, 0x25, 0xb1, 0x24, 0x31, 0xa4, 0xb2, 0x20, 0x55, 0x48, 0x94, 0x4b, 0xd0, 0xc5, 0x31,
	0xc4, 0x31, 0x3e, 0x24, 0x32, 0xc0, 0x35, 0x3e, 0xd4, 0xcf, 0xdb, 0xcf, 0x3f, 0xdc, 0x4f, 0x80,
	0x41, 0x88, 0x9b, 0x8b, 0xdd, 0xc9, 0xdf, 0xdf, 0xc7, 0xd5, 0xd1, 0x4f, 0x80, 0x51, 0x88, 0x93,
	0x8b, 0xd5, 0xd3, 0x2f, 0xc4, 0xcc, 0x44, 0x80, 0x09, 0x24, 0xee, 0xe6, 0xe3, 0xef, 0x08, 0xe2,
	0x30, 0x0b, 0x71, 0x71, 0xb1, 0x05, 0x87, 0x04, 0x79, 0xfa, 0xb9, 0x0b, 0xb0, 0x08, 0xf1, 0x70,
	0x71, 0x84, 0x78, 0xfa, 0xba, 0x9a, 0x99, 0xf8, 0x05, 0x0b, 0xb0, 0x3a, 0xd9, 0x5e, 0x78, 0x28,
	0xc7, 0x70, 0xe3, 0xa1, 0x1c, 0xc3, 0x87, 0x87, 0x72, 0x8c, 0x0d, 0x8f, 0xe4, 0x18, 0x57, 0x3c,
	0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x5f,
	0x3c, 0x92, 0x63, 0xf8, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0x62, 0x07, 0x3b, 0xaf, 0x20, 0x29, 0x89, 0x0d, 0xec, 0x62, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x25, 0x87, 0xa1, 0xd7, 0x00, 0x00, 0x00,
}
