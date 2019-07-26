// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/shared/types/proto/types.proto

package typespb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

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
	UINT128           DataType = 3
	FLOAT64           DataType = 4
	STRING            DataType = 5
	TIME64NS          DataType = 6
)

var DataType_name = map[int32]string{
	0: "DATA_TYPE_UNKNOWN",
	1: "BOOLEAN",
	2: "INT64",
	3: "UINT128",
	4: "FLOAT64",
	5: "STRING",
	6: "TIME64NS",
}

var DataType_value = map[string]int32{
	"DATA_TYPE_UNKNOWN": 0,
	"BOOLEAN":           1,
	"INT64":             2,
	"UINT128":           3,
	"FLOAT64":           4,
	"STRING":            5,
	"TIME64NS":          6,
}

func (DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a0f3bae72116e90, []int{0}
}

type PatternType int32

const (
	UNSPECIFIED    PatternType = 0
	GENERAL        PatternType = 100
	GENERAL_ENUM   PatternType = 101
	STRUCTURED     PatternType = 200
	METRIC_COUNTER PatternType = 300
	METRIC_GAUGE   PatternType = 301
)

var PatternType_name = map[int32]string{
	0:   "UNSPECIFIED",
	100: "GENERAL",
	101: "GENERAL_ENUM",
	200: "STRUCTURED",
	300: "METRIC_COUNTER",
	301: "METRIC_GAUGE",
}

var PatternType_value = map[string]int32{
	"UNSPECIFIED":    0,
	"GENERAL":        100,
	"GENERAL_ENUM":   101,
	"STRUCTURED":     200,
	"METRIC_COUNTER": 300,
	"METRIC_GAUGE":   301,
}

func (PatternType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a0f3bae72116e90, []int{1}
}

type UInt128 struct {
	Low  uint64 `protobuf:"varint,1,opt,name=low,proto3" json:"low,omitempty"`
	High uint64 `protobuf:"varint,2,opt,name=high,proto3" json:"high,omitempty"`
}

func (m *UInt128) Reset()      { *m = UInt128{} }
func (*UInt128) ProtoMessage() {}
func (*UInt128) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0f3bae72116e90, []int{0}
}
func (m *UInt128) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UInt128) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UInt128.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UInt128) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UInt128.Merge(m, src)
}
func (m *UInt128) XXX_Size() int {
	return m.Size()
}
func (m *UInt128) XXX_DiscardUnknown() {
	xxx_messageInfo_UInt128.DiscardUnknown(m)
}

var xxx_messageInfo_UInt128 proto.InternalMessageInfo

func (m *UInt128) GetLow() uint64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *UInt128) GetHigh() uint64 {
	if m != nil {
		return m.High
	}
	return 0
}

func init() {
	proto.RegisterEnum("pl.types.DataType", DataType_name, DataType_value)
	proto.RegisterEnum("pl.types.PatternType", PatternType_name, PatternType_value)
	proto.RegisterType((*UInt128)(nil), "pl.types.UInt128")
}

func init() { proto.RegisterFile("src/shared/types/proto/types.proto", fileDescriptor_6a0f3bae72116e90) }

var fileDescriptor_6a0f3bae72116e90 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x90, 0xcf, 0xea, 0xda, 0x40,
	0x14, 0x85, 0x33, 0xf1, 0x6f, 0xaf, 0xd2, 0x8e, 0x53, 0x0a, 0xae, 0x86, 0xe2, 0xaa, 0xb8, 0x30,
	0x68, 0x45, 0xdc, 0x74, 0x11, 0x93, 0x31, 0x0c, 0xd5, 0x89, 0x24, 0x33, 0x94, 0x76, 0x13, 0x62,
	0x0d, 0xb5, 0x60, 0x35, 0xc4, 0xd0, 0xe2, 0xae, 0x8f, 0xd0, 0xc7, 0xe8, 0xa2, 0x7d, 0x0f, 0x97,
	0x2e, 0x5d, 0xd6, 0xb8, 0xe9, 0xd2, 0x47, 0xf8, 0x91, 0xe8, 0xee, 0x3b, 0xdf, 0xbd, 0x70, 0xe0,
	0x40, 0x67, 0x9f, 0x7c, 0x36, 0xf6, 0xeb, 0x30, 0x89, 0x56, 0x46, 0x7a, 0x88, 0xa3, 0xbd, 0x11,
	0x27, 0xbb, 0x74, 0x77, 0xe7, 0x5e, 0xc1, 0xa4, 0x1e, 0x6f, 0x7a, 0x45, 0xee, 0x18, 0x50, 0x53,
	0x7c, 0x9b, 0xf6, 0x07, 0x63, 0x82, 0xa1, 0xb4, 0xd9, 0xfd, 0x68, 0xa3, 0xd7, 0xe8, 0x4d, 0xd9,
	0xcb, 0x91, 0x10, 0x28, 0xaf, 0xbf, 0x7e, 0x59, 0xb7, 0xf5, 0x42, 0x15, 0xdc, 0xfd, 0x06, 0x75,
	0x3b, 0x4c, 0x43, 0x79, 0x88, 0x23, 0xf2, 0x0a, 0x5a, 0xb6, 0x29, 0xcd, 0x40, 0x7e, 0x5c, 0xb0,
	0x40, 0x89, 0xf7, 0xc2, 0xfd, 0x20, 0xb0, 0x46, 0x1a, 0x50, 0x9b, 0xb8, 0xee, 0x8c, 0x99, 0x02,
	0x23, 0xf2, 0x0c, 0x2a, 0x5c, 0xc8, 0xd1, 0x10, 0xeb, 0xb9, 0x57, 0x5c, 0xc8, 0xfe, 0x60, 0x8c,
	0x4b, 0x79, 0x98, 0xce, 0x5c, 0x33, 0xbf, 0x94, 0x09, 0x40, 0xd5, 0x97, 0x1e, 0x17, 0x0e, 0xae,
	0x90, 0x26, 0xd4, 0x25, 0x9f, 0xb3, 0xd1, 0x50, 0xf8, 0xb8, 0xda, 0xfd, 0x0e, 0x8d, 0x45, 0x98,
	0xa6, 0x51, 0xb2, 0x2d, 0x1a, 0x5f, 0x40, 0x43, 0x09, 0x7f, 0xc1, 0x2c, 0x3e, 0xe5, 0xcc, 0xbe,
	0x77, 0x39, 0x4c, 0x30, 0xcf, 0x9c, 0xe1, 0x15, 0xc1, 0xd0, 0x7c, 0x84, 0x80, 0x09, 0x35, 0xc7,
	0xf9, 0x3f, 0xf8, 0xd2, 0x53, 0x96, 0x54, 0x1e, 0xb3, 0xf1, 0x11, 0x91, 0x97, 0xf0, 0x7c, 0xce,
	0xa4, 0xc7, 0xad, 0xc0, 0x72, 0x95, 0x90, 0xcc, 0xc3, 0x7f, 0x74, 0xd2, 0x82, 0xe6, 0x43, 0x3a,
	0xa6, 0x72, 0x18, 0xfe, 0xab, 0x4f, 0xde, 0x9d, 0x2e, 0x54, 0x3b, 0x5f, 0xa8, 0x76, 0xbb, 0x50,
	0xf4, 0x33, 0xa3, 0xe8, 0x77, 0x46, 0xd1, 0x31, 0xa3, 0xe8, 0x94, 0x51, 0xf4, 0x2f, 0xa3, 0xe8,
	0x7f, 0x46, 0xb5, 0x5b, 0x46, 0xd1, 0xaf, 0x2b, 0xd5, 0x4e, 0x57, 0xaa, 0x9d, 0xaf, 0x54, 0xfb,
	0x54, 0x2b, 0x46, 0x8d, 0x97, 0xcb, 0x6a, 0xb1, 0xf3, 0xdb, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x1a, 0xdd, 0x28, 0x3b, 0x8d, 0x01, 0x00, 0x00,
}

func (x DataType) String() string {
	s, ok := DataType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x PatternType) String() string {
	s, ok := PatternType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *UInt128) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UInt128)
	if !ok {
		that2, ok := that.(UInt128)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Low != that1.Low {
		return false
	}
	if this.High != that1.High {
		return false
	}
	return true
}
func (this *UInt128) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&typespb.UInt128{")
	s = append(s, "Low: "+fmt.Sprintf("%#v", this.Low)+",\n")
	s = append(s, "High: "+fmt.Sprintf("%#v", this.High)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTypes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *UInt128) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UInt128) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Low != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTypes(dAtA, i, uint64(m.Low))
	}
	if m.High != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTypes(dAtA, i, uint64(m.High))
	}
	return i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UInt128) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Low != 0 {
		n += 1 + sovTypes(uint64(m.Low))
	}
	if m.High != 0 {
		n += 1 + sovTypes(uint64(m.High))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UInt128) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UInt128{`,
		`Low:` + fmt.Sprintf("%v", this.Low) + `,`,
		`High:` + fmt.Sprintf("%v", this.High) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTypes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UInt128) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UInt128: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UInt128: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Low", wireType)
			}
			m.Low = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Low |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field High", wireType)
			}
			m.High = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.High |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTypes(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTypes
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes   = fmt.Errorf("proto: integer overflow")
)
