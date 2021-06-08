// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/evm/v1alpha1/tx.proto

package types

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgEthereumTx encapsulates an Ethereum transaction as an SDK message.
type MsgEthereumTx struct {
	// inner transaction data
	Data *TxData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// encoded storage size of the transaction
	Size_ float64 `protobuf:"fixed64,2,opt,name=size,proto3" json:"-"`
	// transaction hash in hex format
	Hash string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty" rlp:"-"`
	// ethereum signer address in hex format. This address value is checked against
	// the address derived from the signature (V, R, S) using the secp256k1
	// elliptic curve
	From string `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
}

func (m *MsgEthereumTx) Reset()         { *m = MsgEthereumTx{} }
func (m *MsgEthereumTx) String() string { return proto.CompactTextString(m) }
func (*MsgEthereumTx) ProtoMessage()    {}
func (*MsgEthereumTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a305e80b084ab0e, []int{0}
}
func (m *MsgEthereumTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEthereumTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEthereumTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEthereumTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEthereumTx.Merge(m, src)
}
func (m *MsgEthereumTx) XXX_Size() int {
	return m.Size()
}
func (m *MsgEthereumTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEthereumTx.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEthereumTx proto.InternalMessageInfo

type ExtensionOptionsEthereumTx struct {
}

func (m *ExtensionOptionsEthereumTx) Reset()         { *m = ExtensionOptionsEthereumTx{} }
func (m *ExtensionOptionsEthereumTx) String() string { return proto.CompactTextString(m) }
func (*ExtensionOptionsEthereumTx) ProtoMessage()    {}
func (*ExtensionOptionsEthereumTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a305e80b084ab0e, []int{1}
}
func (m *ExtensionOptionsEthereumTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtensionOptionsEthereumTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtensionOptionsEthereumTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtensionOptionsEthereumTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionOptionsEthereumTx.Merge(m, src)
}
func (m *ExtensionOptionsEthereumTx) XXX_Size() int {
	return m.Size()
}
func (m *ExtensionOptionsEthereumTx) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionOptionsEthereumTx.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionOptionsEthereumTx proto.InternalMessageInfo

type ExtensionOptionsWeb3Tx struct {
}

func (m *ExtensionOptionsWeb3Tx) Reset()         { *m = ExtensionOptionsWeb3Tx{} }
func (m *ExtensionOptionsWeb3Tx) String() string { return proto.CompactTextString(m) }
func (*ExtensionOptionsWeb3Tx) ProtoMessage()    {}
func (*ExtensionOptionsWeb3Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a305e80b084ab0e, []int{2}
}
func (m *ExtensionOptionsWeb3Tx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtensionOptionsWeb3Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtensionOptionsWeb3Tx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtensionOptionsWeb3Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionOptionsWeb3Tx.Merge(m, src)
}
func (m *ExtensionOptionsWeb3Tx) XXX_Size() int {
	return m.Size()
}
func (m *ExtensionOptionsWeb3Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionOptionsWeb3Tx.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionOptionsWeb3Tx proto.InternalMessageInfo

// MsgEthereumTxResponse defines the Msg/EthereumTx response type.
type MsgEthereumTxResponse struct {
	// ethereum transaction hash. This hash differs from the Tendermint sha256 hash of the transaction
	// bytes. See https://github.com/tendermint/tendermint/issues/6539 for reference
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// logs contains the transaction hash and the proto-compatible ethereum
	// logs.
	Logs []*Log `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
	// returned data from evm function (result or data supplied with revert opcode)
	Ret []byte `protobuf:"bytes,3,opt,name=ret,proto3" json:"ret,omitempty"`
	// reverted flag is set to true when the call has been reverted
	Reverted bool `protobuf:"varint,4,opt,name=reverted,proto3" json:"reverted,omitempty"`
}

func (m *MsgEthereumTxResponse) Reset()         { *m = MsgEthereumTxResponse{} }
func (m *MsgEthereumTxResponse) String() string { return proto.CompactTextString(m) }
func (*MsgEthereumTxResponse) ProtoMessage()    {}
func (*MsgEthereumTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a305e80b084ab0e, []int{3}
}
func (m *MsgEthereumTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEthereumTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEthereumTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEthereumTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEthereumTxResponse.Merge(m, src)
}
func (m *MsgEthereumTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgEthereumTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEthereumTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEthereumTxResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgEthereumTx)(nil), "ethermint.evm.v1alpha1.MsgEthereumTx")
	proto.RegisterType((*ExtensionOptionsEthereumTx)(nil), "ethermint.evm.v1alpha1.ExtensionOptionsEthereumTx")
	proto.RegisterType((*ExtensionOptionsWeb3Tx)(nil), "ethermint.evm.v1alpha1.ExtensionOptionsWeb3Tx")
	proto.RegisterType((*MsgEthereumTxResponse)(nil), "ethermint.evm.v1alpha1.MsgEthereumTxResponse")
}

func init() { proto.RegisterFile("ethermint/evm/v1alpha1/tx.proto", fileDescriptor_6a305e80b084ab0e) }

var fileDescriptor_6a305e80b084ab0e = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x8b, 0xdb, 0x30,
	0x18, 0x87, 0xad, 0xd8, 0x6d, 0x53, 0x25, 0x85, 0x22, 0xda, 0xe0, 0xba, 0x20, 0x1b, 0x43, 0xa9,
	0x97, 0xd8, 0xc4, 0xd9, 0xb2, 0x35, 0x34, 0x5b, 0x43, 0x41, 0x04, 0x0a, 0xdd, 0xec, 0x44, 0xb5,
	0x0d, 0xb1, 0x65, 0x2c, 0xc5, 0xb8, 0xfd, 0x04, 0x1d, 0xb3, 0x76, 0xeb, 0xc7, 0xb9, 0x31, 0xe3,
	0x4d, 0xe1, 0x48, 0xb6, 0x1b, 0xef, 0x13, 0x1c, 0x56, 0xfe, 0x5d, 0x8e, 0x0b, 0xdc, 0xf6, 0xda,
	0xef, 0x23, 0xbd, 0xbf, 0x47, 0xbc, 0xd0, 0xa4, 0x22, 0xa6, 0x45, 0x9a, 0x64, 0xc2, 0xa3, 0x65,
	0xea, 0x95, 0xbd, 0x60, 0x9e, 0xc7, 0x41, 0xcf, 0x13, 0x95, 0x9b, 0x17, 0x4c, 0x30, 0xd4, 0x39,
	0x02, 0x2e, 0x2d, 0x53, 0xf7, 0x00, 0x18, 0xef, 0x22, 0x16, 0x31, 0x89, 0x78, 0x75, 0xb5, 0xa3,
	0x0d, 0xeb, 0xc2, 0x75, 0xf5, 0x51, 0x49, 0xd8, 0xff, 0x00, 0x7c, 0x33, 0xe6, 0xd1, 0xa8, 0xe6,
	0xe8, 0x22, 0x9d, 0x54, 0xc8, 0x87, 0xda, 0x2c, 0x10, 0x81, 0x0e, 0x2c, 0xe0, 0xb4, 0x7c, 0xec,
	0x3e, 0x3d, 0xd0, 0x9d, 0x54, 0x5f, 0x03, 0x11, 0x10, 0xc9, 0xa2, 0x0f, 0x50, 0xe3, 0xc9, 0x1f,
	0xaa, 0x37, 0x2c, 0xe0, 0x80, 0xe1, 0x8b, 0xdb, 0xb5, 0x09, 0xba, 0x44, 0xfe, 0x42, 0x26, 0xd4,
	0xe2, 0x80, 0xc7, 0xba, 0x6a, 0x01, 0xe7, 0xf5, 0xb0, 0x75, 0xb7, 0x36, 0x5f, 0x15, 0xf3, 0x7c,
	0x60, 0x77, 0x6d, 0x22, 0x1b, 0x08, 0x41, 0xed, 0x57, 0xc1, 0x52, 0x5d, 0xab, 0x01, 0x22, 0xeb,
	0x81, 0xf6, 0xf7, 0xbf, 0xa9, 0xd8, 0x36, 0x34, 0x46, 0x95, 0xa0, 0x19, 0x4f, 0x58, 0xf6, 0x3d,
	0x17, 0x09, 0xcb, 0xf8, 0x29, 0xe7, 0x9e, 0xc1, 0xb0, 0xf3, 0x98, 0xf9, 0x41, 0xc3, 0xfe, 0xb1,
	0xbf, 0x04, 0xf0, 0xfd, 0x99, 0x1f, 0xa1, 0x3c, 0x67, 0x19, 0xa7, 0xf5, 0x5c, 0x19, 0xac, 0xf6,
	0x6c, 0xef, 0xb3, 0x78, 0x50, 0x9b, 0xb3, 0x88, 0xeb, 0x0d, 0x4b, 0x75, 0x5a, 0xfe, 0xc7, 0x4b,
	0xee, 0xdf, 0x58, 0x44, 0x24, 0x88, 0xde, 0x42, 0xb5, 0xa0, 0x42, 0xca, 0xb5, 0x49, 0x5d, 0x22,
	0x03, 0x36, 0x0b, 0x5a, 0xd2, 0x42, 0xd0, 0x99, 0x54, 0x6a, 0x92, 0xe3, 0xf7, 0x2e, 0x92, 0x9f,
	0x40, 0x75, 0xcc, 0x23, 0x14, 0x42, 0xf8, 0xe0, 0xd5, 0x3f, 0x5d, 0x9a, 0x75, 0x16, 0xde, 0xe8,
	0x3e, 0x0b, 0x3b, 0x38, 0x0e, 0xbf, 0x5c, 0x6d, 0x30, 0x58, 0x6d, 0x30, 0xb8, 0xd9, 0x60, 0xb0,
	0xdc, 0x62, 0x65, 0xb5, 0xc5, 0xca, 0xf5, 0x16, 0x2b, 0x3f, 0x3f, 0x47, 0x89, 0x88, 0x17, 0xa1,
	0x3b, 0x65, 0xa9, 0x37, 0x65, 0x3c, 0x65, 0xdc, 0x3b, 0xed, 0x4a, 0x25, 0xb7, 0x45, 0xfc, 0xce,
	0x29, 0x0f, 0x5f, 0xca, 0x3d, 0xe9, 0xdf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x90, 0xf9, 0x0d, 0xd1,
	0x9a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// EthereumTx defines a method submitting Ethereum transactions.
	EthereumTx(ctx context.Context, in *MsgEthereumTx, opts ...grpc.CallOption) (*MsgEthereumTxResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) EthereumTx(ctx context.Context, in *MsgEthereumTx, opts ...grpc.CallOption) (*MsgEthereumTxResponse, error) {
	out := new(MsgEthereumTxResponse)
	err := c.cc.Invoke(ctx, "/ethermint.evm.v1alpha1.Msg/EthereumTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// EthereumTx defines a method submitting Ethereum transactions.
	EthereumTx(context.Context, *MsgEthereumTx) (*MsgEthereumTxResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) EthereumTx(ctx context.Context, req *MsgEthereumTx) (*MsgEthereumTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthereumTx not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_EthereumTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEthereumTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EthereumTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethermint.evm.v1alpha1.Msg/EthereumTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EthereumTx(ctx, req.(*MsgEthereumTx))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethermint.evm.v1alpha1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EthereumTx",
			Handler:    _Msg_EthereumTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ethermint/evm/v1alpha1/tx.proto",
}

func (m *MsgEthereumTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEthereumTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEthereumTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Size_ != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Size_))))
		i--
		dAtA[i] = 0x11
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ExtensionOptionsEthereumTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtensionOptionsEthereumTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtensionOptionsEthereumTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ExtensionOptionsWeb3Tx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtensionOptionsWeb3Tx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtensionOptionsWeb3Tx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgEthereumTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEthereumTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEthereumTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Reverted {
		i--
		if m.Reverted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Ret) > 0 {
		i -= len(m.Ret)
		copy(dAtA[i:], m.Ret)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Ret)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Logs) > 0 {
		for iNdEx := len(m.Logs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Logs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgEthereumTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Size_ != 0 {
		n += 9
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *ExtensionOptionsEthereumTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ExtensionOptionsWeb3Tx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgEthereumTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Logs) > 0 {
		for _, e := range m.Logs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Ret)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Reverted {
		n += 2
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgEthereumTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgEthereumTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEthereumTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &TxData{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Size_ = float64(math.Float64frombits(v))
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *ExtensionOptionsEthereumTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: ExtensionOptionsEthereumTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtensionOptionsEthereumTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *ExtensionOptionsWeb3Tx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: ExtensionOptionsWeb3Tx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtensionOptionsWeb3Tx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgEthereumTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgEthereumTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEthereumTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, &Log{})
			if err := m.Logs[len(m.Logs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ret = append(m.Ret[:0], dAtA[iNdEx:postIndex]...)
			if m.Ret == nil {
				m.Ret = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reverted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Reverted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
