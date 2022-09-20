// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: launch/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the launch module's genesis state.
type GenesisState struct {
	// this line is used by starport scaffolding # genesis/proto/state
	Chains            []Chain            `protobuf:"bytes,1,rep,name=chains,proto3" json:"chains"`
	ChainCounter      uint64             `protobuf:"varint,2,opt,name=chainCounter,proto3" json:"chainCounter,omitempty"`
	GenesisAccounts   []GenesisAccount   `protobuf:"bytes,3,rep,name=genesisAccounts,proto3" json:"genesisAccounts"`
	VestingAccounts   []VestingAccount   `protobuf:"bytes,4,rep,name=vestingAccounts,proto3" json:"vestingAccounts"`
	GenesisValidators []GenesisValidator `protobuf:"bytes,5,rep,name=genesisValidators,proto3" json:"genesisValidators"`
	Requests          []Request          `protobuf:"bytes,6,rep,name=requests,proto3" json:"requests"`
	RequestCounters   []RequestCounter   `protobuf:"bytes,7,rep,name=requestCounters,proto3" json:"requestCounters"`
	Params            Params             `protobuf:"bytes,8,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cd66d27edc51cd, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetChains() []Chain {
	if m != nil {
		return m.Chains
	}
	return nil
}

func (m *GenesisState) GetChainCounter() uint64 {
	if m != nil {
		return m.ChainCounter
	}
	return 0
}

func (m *GenesisState) GetGenesisAccounts() []GenesisAccount {
	if m != nil {
		return m.GenesisAccounts
	}
	return nil
}

func (m *GenesisState) GetVestingAccounts() []VestingAccount {
	if m != nil {
		return m.VestingAccounts
	}
	return nil
}

func (m *GenesisState) GetGenesisValidators() []GenesisValidator {
	if m != nil {
		return m.GenesisValidators
	}
	return nil
}

func (m *GenesisState) GetRequests() []Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *GenesisState) GetRequestCounters() []RequestCounter {
	if m != nil {
		return m.RequestCounters
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type RequestCounter struct {
	LaunchID uint64 `protobuf:"varint,1,opt,name=launchID,proto3" json:"launchID,omitempty"`
	Counter  uint64 `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (m *RequestCounter) Reset()         { *m = RequestCounter{} }
func (m *RequestCounter) String() string { return proto.CompactTextString(m) }
func (*RequestCounter) ProtoMessage()    {}
func (*RequestCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cd66d27edc51cd, []int{1}
}
func (m *RequestCounter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestCounter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestCounter.Merge(m, src)
}
func (m *RequestCounter) XXX_Size() int {
	return m.Size()
}
func (m *RequestCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestCounter.DiscardUnknown(m)
}

var xxx_messageInfo_RequestCounter proto.InternalMessageInfo

func (m *RequestCounter) GetLaunchID() uint64 {
	if m != nil {
		return m.LaunchID
	}
	return 0
}

func (m *RequestCounter) GetCounter() uint64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tendermint.spn.launch.GenesisState")
	proto.RegisterType((*RequestCounter)(nil), "tendermint.spn.launch.RequestCounter")
}

func init() { proto.RegisterFile("launch/genesis.proto", fileDescriptor_02cd66d27edc51cd) }

var fileDescriptor_02cd66d27edc51cd = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0xcb, 0xd3, 0x40,
	0x10, 0xc6, 0x13, 0xdf, 0x98, 0x96, 0x6d, 0x51, 0x5c, 0x2b, 0x2c, 0xa5, 0xae, 0xa5, 0x20, 0xe6,
	0x94, 0x40, 0xbd, 0xe9, 0x45, 0x5b, 0xb1, 0x78, 0x93, 0x88, 0x3d, 0xe8, 0x41, 0xb6, 0xe9, 0x92,
	0x06, 0xda, 0x4d, 0xcc, 0x6e, 0x8a, 0x7e, 0x0b, 0x3f, 0x56, 0x6f, 0xf6, 0xe8, 0x49, 0xa4, 0xfd,
	0x22, 0x92, 0xdd, 0x49, 0x24, 0xf1, 0x6d, 0x6f, 0xbb, 0x33, 0xcf, 0xfc, 0xe6, 0x0f, 0x0f, 0x1a,
	0x6c, 0x59, 0x21, 0xa2, 0x4d, 0x10, 0x73, 0xc1, 0x65, 0x22, 0xfd, 0x2c, 0x4f, 0x55, 0x8a, 0x1f,
	0x29, 0x2e, 0xd6, 0x3c, 0xdf, 0x25, 0x42, 0xf9, 0x32, 0x13, 0xbe, 0x11, 0x0d, 0x07, 0x71, 0x1a,
	0xa7, 0x5a, 0x11, 0x94, 0x2f, 0x23, 0x1e, 0x56, 0x88, 0x9c, 0x7f, 0x2d, 0xb8, 0x54, 0x10, 0x1d,
	0x41, 0x74, 0xcf, 0xa5, 0x4a, 0x44, 0xfc, 0x85, 0x45, 0x51, 0x5a, 0x88, 0x76, 0x16, 0xda, 0xb6,
	0xb2, 0xb4, 0x95, 0xdd, 0xb3, 0x6d, 0xb2, 0x66, 0x2a, 0xcd, 0x21, 0x8f, 0x21, 0x1f, 0x6d, 0x58,
	0x22, 0x20, 0xf6, 0x10, 0x62, 0x19, 0xcb, 0xd9, 0x0e, 0xf6, 0x98, 0xfc, 0x74, 0x50, 0x7f, 0x61,
	0x20, 0x1f, 0x14, 0x53, 0x1c, 0xbf, 0x40, 0xae, 0x2e, 0x92, 0xc4, 0x1e, 0xdf, 0x78, 0xbd, 0xe9,
	0xc8, 0xbf, 0x75, 0x53, 0x7f, 0x5e, 0x8a, 0x66, 0xce, 0xe1, 0xf7, 0x13, 0x2b, 0x84, 0x0a, 0x3c,
	0x41, 0x7d, 0xfd, 0x9a, 0x97, 0x93, 0xf2, 0x9c, 0xdc, 0x19, 0xdb, 0x9e, 0x13, 0x36, 0x62, 0xf8,
	0x23, 0xba, 0x0f, 0x43, 0xbf, 0x36, 0x1b, 0x49, 0x72, 0xa3, 0x1b, 0x3d, 0xbd, 0xd0, 0x68, 0xd1,
	0x50, 0x43, 0xc7, 0x36, 0xa3, 0xc4, 0xc2, 0x1d, 0x6b, 0xac, 0x73, 0x15, 0xbb, 0x6c, 0xa8, 0x2b,
	0x6c, 0x8b, 0x81, 0x3f, 0xa3, 0x07, 0xd0, 0x69, 0x59, 0x5d, 0x58, 0x92, 0xbb, 0x1a, 0xfc, 0xec,
	0xfa, 0xbc, 0xb5, 0x1e, 0xd0, 0xff, 0x73, 0xf0, 0x2b, 0xd4, 0x05, 0x47, 0x48, 0xe2, 0x6a, 0x26,
	0xbd, 0xc0, 0x0c, 0x8d, 0x0c, 0x50, 0x75, 0x55, 0xb9, 0x35, 0xbc, 0xe1, 0xbc, 0x92, 0x74, 0xae,
	0x6e, 0x1d, 0x36, 0xd4, 0xd5, 0xd6, 0x2d, 0x06, 0x7e, 0x89, 0x5c, 0x63, 0x12, 0xd2, 0x1d, 0xdb,
	0x5e, 0x6f, 0xfa, 0xf8, 0x02, 0xed, 0xbd, 0x16, 0x55, 0x26, 0x30, 0x25, 0x93, 0xb7, 0xe8, 0x5e,
	0xb3, 0x0b, 0x1e, 0xa2, 0xae, 0x29, 0x78, 0xf7, 0x86, 0xd8, 0xda, 0x12, 0xf5, 0x1f, 0x13, 0xd4,
	0x89, 0x1a, 0x6e, 0xa9, 0xbe, 0xb3, 0xd9, 0xe1, 0x44, 0xed, 0xe3, 0x89, 0xda, 0x7f, 0x4e, 0xd4,
	0xfe, 0x71, 0xa6, 0xd6, 0xf1, 0x4c, 0xad, 0x5f, 0x67, 0x6a, 0x7d, 0xf2, 0xe2, 0x44, 0x6d, 0x8a,
	0x95, 0x1f, 0xa5, 0xbb, 0xe0, 0xdf, 0x60, 0x81, 0xcc, 0x44, 0xf0, 0x2d, 0x00, 0x93, 0xab, 0xef,
	0x19, 0x97, 0x2b, 0x57, 0x9b, 0xfc, 0xf9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0x5f, 0xfc,
	0xc9, 0xc4, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.RequestCounters) > 0 {
		for iNdEx := len(m.RequestCounters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RequestCounters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Requests) > 0 {
		for iNdEx := len(m.Requests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.GenesisValidators) > 0 {
		for iNdEx := len(m.GenesisValidators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenesisValidators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.VestingAccounts) > 0 {
		for iNdEx := len(m.VestingAccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingAccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.GenesisAccounts) > 0 {
		for iNdEx := len(m.GenesisAccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenesisAccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ChainCounter != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ChainCounter))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Chains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RequestCounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestCounter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestCounter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Counter != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Counter))
		i--
		dAtA[i] = 0x10
	}
	if m.LaunchID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LaunchID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Chains) > 0 {
		for _, e := range m.Chains {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ChainCounter != 0 {
		n += 1 + sovGenesis(uint64(m.ChainCounter))
	}
	if len(m.GenesisAccounts) > 0 {
		for _, e := range m.GenesisAccounts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.VestingAccounts) > 0 {
		for _, e := range m.VestingAccounts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GenesisValidators) > 0 {
		for _, e := range m.GenesisValidators {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Requests) > 0 {
		for _, e := range m.Requests {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RequestCounters) > 0 {
		for _, e := range m.RequestCounters {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *RequestCounter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LaunchID != 0 {
		n += 1 + sovGenesis(uint64(m.LaunchID))
	}
	if m.Counter != 0 {
		n += 1 + sovGenesis(uint64(m.Counter))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, Chain{})
			if err := m.Chains[len(m.Chains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainCounter", wireType)
			}
			m.ChainCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainCounter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisAccounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisAccounts = append(m.GenesisAccounts, GenesisAccount{})
			if err := m.GenesisAccounts[len(m.GenesisAccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingAccounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingAccounts = append(m.VestingAccounts, VestingAccount{})
			if err := m.VestingAccounts[len(m.VestingAccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisValidators = append(m.GenesisValidators, GenesisValidator{})
			if err := m.GenesisValidators[len(m.GenesisValidators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requests = append(m.Requests, Request{})
			if err := m.Requests[len(m.Requests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestCounters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestCounters = append(m.RequestCounters, RequestCounter{})
			if err := m.RequestCounters[len(m.RequestCounters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *RequestCounter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: RequestCounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestCounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchID", wireType)
			}
			m.LaunchID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LaunchID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
