// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: supply/v1beta1/supply.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CoinDetails represents bonded and unbonded coin details
type CoinDetails struct {
	Bonded   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=bonded,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"bonded" yaml:"bonded"`
	Unbonded github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=unbonded,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"unbonded" yaml:"unbonded"`
}

func (m *CoinDetails) Reset()         { *m = CoinDetails{} }
func (m *CoinDetails) String() string { return proto.CompactTextString(m) }
func (*CoinDetails) ProtoMessage()    {}
func (*CoinDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5cccb41288e84e, []int{0}
}
func (m *CoinDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoinDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoinDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoinDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinDetails.Merge(m, src)
}
func (m *CoinDetails) XXX_Size() int {
	return m.Size()
}
func (m *CoinDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CoinDetails proto.InternalMessageInfo

// Supply represents total coins vested, available and circulating supply
type Supply struct {
	Vesting     CoinDetails                              `protobuf:"bytes,1,opt,name=vesting,proto3" json:"vesting" yaml:"vesting"`
	Available   CoinDetails                              `protobuf:"bytes,2,opt,name=available,proto3" json:"available" yaml:"available"`
	Circulating github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=circulating,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"circulating" yaml:"circulating"`
	Total       github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=total,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total" yaml:"total"`
}

func (m *Supply) Reset()         { *m = Supply{} }
func (m *Supply) String() string { return proto.CompactTextString(m) }
func (*Supply) ProtoMessage()    {}
func (*Supply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5cccb41288e84e, []int{1}
}
func (m *Supply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Supply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Supply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Supply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Supply.Merge(m, src)
}
func (m *Supply) XXX_Size() int {
	return m.Size()
}
func (m *Supply) XXX_DiscardUnknown() {
	xxx_messageInfo_Supply.DiscardUnknown(m)
}

var xxx_messageInfo_Supply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CoinDetails)(nil), "supply.v1beta1.CoinDetails")
	proto.RegisterType((*Supply)(nil), "supply.v1beta1.Supply")
}

func init() { proto.RegisterFile("supply/v1beta1/supply.proto", fileDescriptor_0a5cccb41288e84e) }

var fileDescriptor_0a5cccb41288e84e = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0x80, 0xe3, 0x16, 0xca, 0xe1, 0xc0, 0x81, 0x2c, 0x86, 0x70, 0x27, 0xc5, 0x47, 0x24, 0xa4,
	0x2e, 0x97, 0xe8, 0x0e, 0x89, 0xe1, 0xc6, 0xc2, 0x82, 0x18, 0x40, 0x45, 0x48, 0x88, 0x09, 0x27,
	0x8d, 0x42, 0x74, 0x4e, 0x1c, 0xc5, 0x4e, 0x44, 0x18, 0x18, 0x11, 0x23, 0xe2, 0x17, 0x74, 0xe6,
	0x97, 0xdc, 0x78, 0x23, 0x93, 0x8b, 0xda, 0x05, 0x31, 0xe6, 0x17, 0xa0, 0xc6, 0x4e, 0x1a, 0xa6,
	0xa8, 0x53, 0xfb, 0xde, 0xf3, 0xfb, 0xde, 0x97, 0x67, 0x19, 0x1e, 0xf3, 0x22, 0xcb, 0x68, 0xe5,
	0x95, 0x67, 0x7e, 0x28, 0xc8, 0x99, 0xa7, 0x42, 0x37, 0xcb, 0x99, 0x60, 0xe8, 0x50, 0x47, 0xba,
	0x78, 0xf4, 0x20, 0x62, 0x11, 0x6b, 0x4a, 0xde, 0xf6, 0x9f, 0x3a, 0x75, 0x64, 0x07, 0x8c, 0x27,
	0x8c, 0x7b, 0x3e, 0xe1, 0x61, 0xc7, 0x09, 0x58, 0x9c, 0xaa, 0xba, 0xb3, 0x1c, 0x41, 0xf3, 0x19,
	0x8b, 0xd3, 0xe7, 0xa1, 0x20, 0x31, 0xe5, 0xe8, 0x0b, 0x9c, 0xf8, 0x2c, 0x5d, 0x84, 0x0b, 0x0b,
	0x9c, 0x8c, 0xa7, 0xe6, 0xf9, 0x43, 0x57, 0x01, 0xdc, 0x2d, 0xa0, 0x9d, 0xe5, 0x6e, 0x3b, 0x66,
	0x2f, 0xaf, 0x24, 0x36, 0xfe, 0x4a, 0xac, 0x1b, 0x6a, 0x89, 0xef, 0x56, 0x24, 0xa1, 0x17, 0x8e,
	0x8a, 0x9d, 0x9f, 0x2b, 0x3c, 0x8d, 0x62, 0xf1, 0xb1, 0xf0, 0xdd, 0x80, 0x25, 0x9e, 0x16, 0x51,
	0x3f, 0xa7, 0x7c, 0x71, 0xe9, 0x89, 0x2a, 0x0b, 0x79, 0xc3, 0xe2, 0x73, 0x0d, 0x41, 0x5f, 0x01,
	0x3c, 0x28, 0x52, 0xad, 0x30, 0x1a, 0x52, 0x78, 0xa5, 0x15, 0xba, 0x96, 0x5a, 0xe2, 0x7b, 0x4a,
	0xa2, 0xcd, 0xec, 0xa7, 0xd1, 0x81, 0x2e, 0x0e, 0xbe, 0x2d, 0xb1, 0xf1, 0x67, 0x89, 0x0d, 0x67,
	0x35, 0x86, 0x93, 0x37, 0xcd, 0xae, 0xd1, 0x3b, 0x78, 0xab, 0x0c, 0xb9, 0x88, 0xd3, 0xc8, 0x02,
	0x27, 0x60, 0x6a, 0x9e, 0x1f, 0xbb, 0xff, 0xdf, 0x82, 0xdb, 0xdb, 0xe5, 0xec, 0x91, 0xb6, 0x6b,
	0x7b, 0x6a, 0x89, 0x0f, 0x95, 0x9c, 0x4e, 0x38, 0xf3, 0xb6, 0x84, 0x3e, 0xc0, 0xdb, 0xa4, 0x24,
	0x31, 0x25, 0x3e, 0x0d, 0xad, 0xd1, 0x30, 0xfb, 0xb1, 0x66, 0xef, 0xba, 0x6a, 0x89, 0xef, 0x2b,
	0x7a, 0x97, 0x72, 0xe6, 0xbb, 0x32, 0xfa, 0x01, 0xa0, 0x19, 0xc4, 0x79, 0x50, 0x50, 0xd2, 0x7c,
	0xc0, 0x78, 0x68, 0xb9, 0x6f, 0xf5, 0x88, 0x7e, 0x57, 0x2d, 0x31, 0x52, 0x43, 0x7a, 0xc9, 0xfd,
	0x56, 0xdc, 0xc7, 0xa1, 0xcf, 0xf0, 0xa6, 0x60, 0x82, 0x50, 0xeb, 0xc6, 0x90, 0xcd, 0x0b, 0x6d,
	0xa3, 0xce, 0xd7, 0x12, 0xdf, 0x51, 0x1e, 0x4d, 0xb8, 0x9f, 0x81, 0x42, 0xec, 0x6e, 0x78, 0xf6,
	0xfa, 0x6a, 0x6d, 0x83, 0xeb, 0xb5, 0x0d, 0x7e, 0xaf, 0x6d, 0xf0, 0x7d, 0x63, 0x1b, 0xd7, 0x1b,
	0xdb, 0xf8, 0xb5, 0xb1, 0x8d, 0xf7, 0x4f, 0x7b, 0x50, 0x56, 0xe6, 0x01, 0xbd, 0xec, 0xa0, 0xcd,
	0xdd, 0x9c, 0xf2, 0x22, 0x49, 0x48, 0x5e, 0x79, 0x9f, 0xf4, 0xe3, 0x54, 0x83, 0xfc, 0x49, 0xf3,
	0xba, 0x9e, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x75, 0xf9, 0xd7, 0xc2, 0x03, 0x00, 0x00,
}

func (m *CoinDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoinDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoinDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Unbonded) > 0 {
		for iNdEx := len(m.Unbonded) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Unbonded[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSupply(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Bonded) > 0 {
		for iNdEx := len(m.Bonded) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Bonded[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSupply(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Supply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Supply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Supply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Total) > 0 {
		for iNdEx := len(m.Total) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Total[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSupply(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Circulating) > 0 {
		for iNdEx := len(m.Circulating) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Circulating[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSupply(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Available.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSupply(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Vesting.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSupply(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintSupply(dAtA []byte, offset int, v uint64) int {
	offset -= sovSupply(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CoinDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Bonded) > 0 {
		for _, e := range m.Bonded {
			l = e.Size()
			n += 1 + l + sovSupply(uint64(l))
		}
	}
	if len(m.Unbonded) > 0 {
		for _, e := range m.Unbonded {
			l = e.Size()
			n += 1 + l + sovSupply(uint64(l))
		}
	}
	return n
}

func (m *Supply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Vesting.Size()
	n += 1 + l + sovSupply(uint64(l))
	l = m.Available.Size()
	n += 1 + l + sovSupply(uint64(l))
	if len(m.Circulating) > 0 {
		for _, e := range m.Circulating {
			l = e.Size()
			n += 1 + l + sovSupply(uint64(l))
		}
	}
	if len(m.Total) > 0 {
		for _, e := range m.Total {
			l = e.Size()
			n += 1 + l + sovSupply(uint64(l))
		}
	}
	return n
}

func sovSupply(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSupply(x uint64) (n int) {
	return sovSupply(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CoinDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSupply
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
			return fmt.Errorf("proto: CoinDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoinDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bonded", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupply
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
				return ErrInvalidLengthSupply
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupply
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bonded = append(m.Bonded, types.Coin{})
			if err := m.Bonded[len(m.Bonded)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unbonded", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupply
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
				return ErrInvalidLengthSupply
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupply
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unbonded = append(m.Unbonded, types.Coin{})
			if err := m.Unbonded[len(m.Unbonded)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSupply(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSupply
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSupply
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
func (m *Supply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSupply
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
			return fmt.Errorf("proto: Supply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Supply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vesting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupply
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
				return ErrInvalidLengthSupply
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupply
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Vesting.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Available", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupply
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
				return ErrInvalidLengthSupply
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupply
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Available.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Circulating", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupply
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
				return ErrInvalidLengthSupply
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupply
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Circulating = append(m.Circulating, types.Coin{})
			if err := m.Circulating[len(m.Circulating)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupply
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
				return ErrInvalidLengthSupply
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupply
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Total = append(m.Total, types.Coin{})
			if err := m.Total[len(m.Total)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSupply(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSupply
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSupply
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
func skipSupply(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSupply
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
					return 0, ErrIntOverflowSupply
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
					return 0, ErrIntOverflowSupply
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
				return 0, ErrInvalidLengthSupply
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSupply
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSupply
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSupply        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSupply          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSupply = fmt.Errorf("proto: unexpected end of group")
)
