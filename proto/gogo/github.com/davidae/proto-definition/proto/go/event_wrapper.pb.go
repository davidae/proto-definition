// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: event_wrapper.proto

package _go

import (
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GreetingWrapp struct {
	From                 *types.StringValue `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	LastUpdated          *types.Timestamp   `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	To                   *types.StringValue `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	PublisherCategories  []int32            `protobuf:"zigzag32,4,rep,packed,name=publisher_categories,json=publisherCategories,proto3" json:"publisher_categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GreetingWrapp) Reset()         { *m = GreetingWrapp{} }
func (m *GreetingWrapp) String() string { return proto.CompactTextString(m) }
func (*GreetingWrapp) ProtoMessage()    {}
func (*GreetingWrapp) Descriptor() ([]byte, []int) {
	return fileDescriptor_598b41a4843f2c76, []int{0}
}
func (m *GreetingWrapp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GreetingWrapp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GreetingWrapp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GreetingWrapp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingWrapp.Merge(m, src)
}
func (m *GreetingWrapp) XXX_Size() int {
	return m.Size()
}
func (m *GreetingWrapp) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingWrapp.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingWrapp proto.InternalMessageInfo

func (m *GreetingWrapp) GetFrom() *types.StringValue {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *GreetingWrapp) GetLastUpdated() *types.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

func (m *GreetingWrapp) GetTo() *types.StringValue {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *GreetingWrapp) GetPublisherCategories() []int32 {
	if m != nil {
		return m.PublisherCategories
	}
	return nil
}

func init() {
	proto.RegisterType((*GreetingWrapp)(nil), "event.GreetingWrapp")
}

func init() { proto.RegisterFile("event_wrapper.proto", fileDescriptor_598b41a4843f2c76) }

var fileDescriptor_598b41a4843f2c76 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x71, 0x5a, 0x18, 0xdc, 0x32, 0x90, 0x32, 0x44, 0x11, 0x0a, 0x15, 0x53, 0x87, 0x92,
	0xf0, 0x33, 0xc3, 0x00, 0x03, 0x7b, 0xf8, 0x93, 0x58, 0x22, 0xa7, 0xb9, 0x71, 0x2d, 0x25, 0xb9,
	0x96, 0x73, 0x53, 0x5e, 0x85, 0x47, 0x62, 0xe4, 0x11, 0x20, 0xbc, 0x08, 0x6a, 0xdc, 0x74, 0x80,
	0x85, 0xcd, 0x3a, 0xdf, 0x77, 0x8e, 0x2d, 0xf3, 0x09, 0xac, 0xa0, 0xa2, 0xe4, 0xd5, 0x08, 0xad,
	0xc1, 0x84, 0xda, 0x20, 0xa1, 0xbb, 0xdb, 0x85, 0xfe, 0xb1, 0x44, 0x94, 0x05, 0x44, 0x5d, 0x98,
	0x36, 0x79, 0x44, 0xaa, 0x84, 0x9a, 0x44, 0xa9, 0xad, 0xe7, 0x8f, 0x49, 0x48, 0xd9, 0xb7, 0xfc,
	0xe0, 0xb7, 0xbe, 0x19, 0xad, 0x2d, 0x3f, 0xf9, 0x62, 0x7c, 0xff, 0xce, 0x00, 0x90, 0xaa, 0xe4,
	0xf3, 0x1a, 0xb9, 0x67, 0x7c, 0x98, 0x1b, 0x2c, 0x3d, 0x36, 0x65, 0xb3, 0xd1, 0xc5, 0x51, 0x68,
	0x07, 0xc2, 0x7e, 0x20, 0xbc, 0x27, 0xa3, 0x2a, 0xf9, 0x24, 0x8a, 0x06, 0xe2, 0xce, 0x74, 0xaf,
	0xf8, 0xb8, 0x10, 0x35, 0x25, 0x8d, 0xce, 0x04, 0x41, 0xe6, 0x39, 0x5d, 0xd3, 0xff, 0xd3, 0x7c,
	0xe8, 0x5f, 0x1a, 0x8f, 0xd6, 0xfe, 0xa3, 0xd5, 0xdd, 0x39, 0x77, 0x08, 0xbd, 0xc1, 0x3f, 0xae,
	0x73, 0x08, 0xdd, 0x73, 0x7e, 0xa8, 0x9b, 0xb4, 0x50, 0xf5, 0x12, 0x4c, 0xb2, 0x10, 0x04, 0x12,
	0x8d, 0x82, 0xda, 0x1b, 0x4e, 0x07, 0xb3, 0x83, 0x78, 0xb2, 0x65, 0xb7, 0x5b, 0x74, 0x73, 0xfd,
	0xde, 0x06, 0xec, 0xa3, 0x0d, 0xd8, 0x67, 0x1b, 0xb0, 0xb7, 0xef, 0x60, 0xe7, 0x65, 0x2e, 0x15,
	0x2d, 0x9b, 0x34, 0x5c, 0x60, 0x19, 0x65, 0x62, 0xa5, 0x32, 0xb1, 0xf9, 0xa1, 0xd3, 0x0c, 0x72,
	0x55, 0x29, 0x52, 0x58, 0xd9, 0x20, 0x92, 0x98, 0xee, 0x75, 0xa7, 0xcb, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x78, 0x07, 0xcb, 0x3c, 0x97, 0x01, 0x00, 0x00,
}

func (m *GreetingWrapp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GreetingWrapp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GreetingWrapp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PublisherCategories) > 0 {
		dAtA1 := make([]byte, len(m.PublisherCategories)*5)
		var j2 int
		for _, num := range m.PublisherCategories {
			x3 := (uint32(num) << 1) ^ uint32((num >> 31))
			for x3 >= 1<<7 {
				dAtA1[j2] = uint8(uint64(x3)&0x7f | 0x80)
				j2++
				x3 >>= 7
			}
			dAtA1[j2] = uint8(x3)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA1[:j2])
		i = encodeVarintEventWrapper(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0x22
	}
	if m.To != nil {
		{
			size, err := m.To.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEventWrapper(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.LastUpdated != nil {
		{
			size, err := m.LastUpdated.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEventWrapper(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.From != nil {
		{
			size, err := m.From.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEventWrapper(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEventWrapper(dAtA []byte, offset int, v uint64) int {
	offset -= sovEventWrapper(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GreetingWrapp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovEventWrapper(uint64(l))
	}
	if m.LastUpdated != nil {
		l = m.LastUpdated.Size()
		n += 1 + l + sovEventWrapper(uint64(l))
	}
	if m.To != nil {
		l = m.To.Size()
		n += 1 + l + sovEventWrapper(uint64(l))
	}
	if len(m.PublisherCategories) > 0 {
		l = 0
		for _, e := range m.PublisherCategories {
			l += sozEventWrapper(uint64(e))
		}
		n += 1 + sovEventWrapper(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEventWrapper(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEventWrapper(x uint64) (n int) {
	return sovEventWrapper(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GreetingWrapp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventWrapper
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
			return fmt.Errorf("proto: GreetingWrapp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GreetingWrapp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventWrapper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEventWrapper
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEventWrapper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &types.StringValue{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventWrapper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEventWrapper
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEventWrapper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastUpdated == nil {
				m.LastUpdated = &types.Timestamp{}
			}
			if err := m.LastUpdated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventWrapper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEventWrapper
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEventWrapper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.To == nil {
				m.To = &types.StringValue{}
			}
			if err := m.To.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEventWrapper
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
				m.PublisherCategories = append(m.PublisherCategories, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEventWrapper
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthEventWrapper
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthEventWrapper
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.PublisherCategories) == 0 {
					m.PublisherCategories = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEventWrapper
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
					m.PublisherCategories = append(m.PublisherCategories, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field PublisherCategories", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEventWrapper(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEventWrapper
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEventWrapper(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEventWrapper
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
					return 0, ErrIntOverflowEventWrapper
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEventWrapper
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
				return 0, ErrInvalidLengthEventWrapper
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEventWrapper
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEventWrapper
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEventWrapper        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEventWrapper          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEventWrapper = fmt.Errorf("proto: unexpected end of group")
)