package chh

import (
	"time"

	"github.com/ClickHouse/ch-go/proto"
)

// -- type Str

func (r *ColResults) BindStr(name string, valPtr interface{}) {
	col := &proto.ColStr{}
	r.Bind(name, col, &bindingColStr{col: col, valPtr: valPtr})
}

type bindingColStr struct {
	col    *proto.ColStr
	valPtr interface{}
}

func (b *bindingColStr) Append() {
	s, _ := b.valPtr.(*[]string)
	appendToSlice(s, b.col)
}

func (b *bindingColStr) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*string)
		*s = b.col.Row(0)
	}
}

// -- type Float32

func (r *ColResults) BindFloat32(name string, valPtr interface{}) {
	col := &proto.ColFloat32{}
	r.Bind(name, col, &bindingColFloat32{col: col, valPtr: valPtr})
}

type bindingColFloat32 struct {
	col    *proto.ColFloat32
	valPtr interface{}
}

func (b *bindingColFloat32) Append() {
	s, _ := b.valPtr.(*[]float32)
	appendToSlice(s, b.col)
}

func (b *bindingColFloat32) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*float32)
		*s = b.col.Row(0)
	}
}

// -- type Float64

func (r *ColResults) BindFloat64(name string, valPtr interface{}) {
	col := &proto.ColFloat64{}
	r.Bind(name, col, &bindingColFloat64{col: col, valPtr: valPtr})
}

type bindingColFloat64 struct {
	col    *proto.ColFloat64
	valPtr interface{}
}

func (b *bindingColFloat64) Append() {
	s, _ := b.valPtr.(*[]float64)
	appendToSlice(s, b.col)
}

func (b *bindingColFloat64) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*float64)
		*s = b.col.Row(0)
	}
}

// -- type IPv4

func (r *ColResults) BindIPv4(name string, valPtr interface{}) {
	col := &proto.ColIPv4{}
	r.Bind(name, col, &bindingColIPv4{col: col, valPtr: valPtr})
}

type bindingColIPv4 struct {
	col    *proto.ColIPv4
	valPtr interface{}
}

func (b *bindingColIPv4) Append() {
	s, _ := b.valPtr.(*[]proto.IPv4)
	appendToSlice(s, b.col)
}

func (b *bindingColIPv4) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*proto.IPv4)
		*s = b.col.Row(0)
	}
}

// -- type IPv6

func (r *ColResults) BindIPv6(name string, valPtr interface{}) {
	col := &proto.ColIPv6{}
	r.Bind(name, col, &bindingColIPv6{col: col, valPtr: valPtr})
}

type bindingColIPv6 struct {
	col    *proto.ColIPv6
	valPtr interface{}
}

func (b *bindingColIPv6) Append() {
	s, _ := b.valPtr.(*[]proto.IPv6)
	appendToSlice(s, b.col)
}

func (b *bindingColIPv6) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*proto.IPv6)
		*s = b.col.Row(0)
	}
}

// -- type DateTime

func (r *ColResults) BindDateTime(name string, valPtr interface{}) {
	col := &proto.ColDateTime{}
	r.Bind(name, col, &bindingColDateTime{col: col, valPtr: valPtr})
}

type bindingColDateTime struct {
	col    *proto.ColDateTime
	valPtr interface{}
}

func (b *bindingColDateTime) Append() {
	s, _ := b.valPtr.(*[]time.Time)
	appendToSlice(s, b.col)
}

func (b *bindingColDateTime) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*time.Time)
		*s = b.col.Row(0)
	}
}

// -- type DateTime64

func (r *ColResults) BindDateTime64(name string, valPtr interface{}) {
	col := &proto.ColDateTime64{}
	r.Bind(name, col, &bindingColDateTime64{col: col, valPtr: valPtr})
}

type bindingColDateTime64 struct {
	col    *proto.ColDateTime64
	valPtr interface{}
}

func (b *bindingColDateTime64) Append() {
	s, _ := b.valPtr.(*[]time.Time)
	appendToSlice(s, b.col)
}

func (b *bindingColDateTime64) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*time.Time)
		*s = b.col.Row(0)
	}
}

// -- type Date

func (r *ColResults) BindDate(name string, valPtr interface{}) {
	col := &proto.ColDate{}
	r.Bind(name, col, &bindingColDate{col: col, valPtr: valPtr})
}

type bindingColDate struct {
	col    *proto.ColDate
	valPtr interface{}
}

func (b *bindingColDate) Append() {
	s, _ := b.valPtr.(*[]time.Time)
	appendToSlice(s, b.col)
}

func (b *bindingColDate) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*time.Time)
		*s = b.col.Row(0)
	}
}

// -- type Date32

func (r *ColResults) BindDate32(name string, valPtr interface{}) {
	col := &proto.ColDate32{}
	r.Bind(name, col, &bindingColDate32{col: col, valPtr: valPtr})
}

type bindingColDate32 struct {
	col    *proto.ColDate32
	valPtr interface{}
}

func (b *bindingColDate32) Append() {
	s, _ := b.valPtr.(*[]time.Time)
	appendToSlice(s, b.col)
}

func (b *bindingColDate32) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*time.Time)
		*s = b.col.Row(0)
	}
}

// -- type Enum8

func (r *ColResults) BindEnum8(name string, valPtr interface{}) {
	col := &proto.ColEnum8{}
	r.Bind(name, col, &bindingColEnum8{col: col, valPtr: valPtr})
}

type bindingColEnum8 struct {
	col    *proto.ColEnum8
	valPtr interface{}
}

func (b *bindingColEnum8) Append() {
	s, _ := b.valPtr.(*[]proto.Enum8)
	appendToSlice(s, b.col)
}

func (b *bindingColEnum8) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*proto.Enum8)
		*s = b.col.Row(0)
	}
}

// -- type Enum16

func (r *ColResults) BindEnum16(name string, valPtr interface{}) {
	col := &proto.ColEnum16{}
	r.Bind(name, col, &bindingColEnum16{col: col, valPtr: valPtr})
}

type bindingColEnum16 struct {
	col    *proto.ColEnum16
	valPtr interface{}
}

func (b *bindingColEnum16) Append() {
	s, _ := b.valPtr.(*[]proto.Enum16)
	appendToSlice(s, b.col)
}

func (b *bindingColEnum16) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*proto.Enum16)
		*s = b.col.Row(0)
	}
}

// -- type Decimal32

func (r *ColResults) BindDecimal32(name string, valPtr interface{}) {
	col := &proto.ColDecimal32{}
	r.Bind(name, col, &bindingColDecimal32{col: col, valPtr: valPtr})
}

type bindingColDecimal32 struct {
	col    *proto.ColDecimal32
	valPtr interface{}
}

func (b *bindingColDecimal32) Append() {
	s, _ := b.valPtr.(*[]proto.Decimal32)
	appendToSlice(s, b.col)
}

func (b *bindingColDecimal32) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*proto.Decimal32)
		*s = b.col.Row(0)
	}
}

// -- type Decimal64

func (r *ColResults) BindDecimal64(name string, valPtr interface{}) {
	col := &proto.ColDecimal64{}
	r.Bind(name, col, &bindingColDecimal64{col: col, valPtr: valPtr})
}

type bindingColDecimal64 struct {
	col    *proto.ColDecimal64
	valPtr interface{}
}

func (b *bindingColDecimal64) Append() {
	s, _ := b.valPtr.(*[]proto.Decimal64)
	appendToSlice(s, b.col)
}

func (b *bindingColDecimal64) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*proto.Decimal64)
		*s = b.col.Row(0)
	}
}

// -- type Decimal128

func (r *ColResults) BindDecimal128(name string, valPtr interface{}) {
	col := &proto.ColDecimal128{}
	r.Bind(name, col, &bindingColDecimal128{col: col, valPtr: valPtr})
}

type bindingColDecimal128 struct {
	col    *proto.ColDecimal128
	valPtr interface{}
}

func (b *bindingColDecimal128) Append() {
	s, _ := b.valPtr.(*[]proto.Decimal128)
	appendToSlice(s, b.col)
}

func (b *bindingColDecimal128) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*proto.Decimal128)
		*s = b.col.Row(0)
	}
}

// -- type Decimal256

func (r *ColResults) BindDecimal256(name string, valPtr interface{}) {
	col := &proto.ColDecimal256{}
	r.Bind(name, col, &bindingColDecimal256{col: col, valPtr: valPtr})
}

type bindingColDecimal256 struct {
	col    *proto.ColDecimal256
	valPtr interface{}
}

func (b *bindingColDecimal256) Append() {
	s, _ := b.valPtr.(*[]proto.Decimal256)
	appendToSlice(s, b.col)
}

func (b *bindingColDecimal256) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*proto.Decimal256)
		*s = b.col.Row(0)
	}
}

// -- type Int8

func (r *ColResults) BindInt8(name string, valPtr interface{}) {
	col := &proto.ColInt8{}
	r.Bind(name, col, &bindingColInt8{col: col, valPtr: valPtr})
}

type bindingColInt8 struct {
	col    *proto.ColInt8
	valPtr interface{}
}

func (b *bindingColInt8) Append() {
	s, _ := b.valPtr.(*[]int8)
	appendToSlice(s, b.col)
}

func (b *bindingColInt8) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*int8)
		*s = b.col.Row(0)
	}
}

// -- type UInt8

func (r *ColResults) BindUInt8(name string, valPtr interface{}) {
	col := &proto.ColUInt8{}
	r.Bind(name, col, &bindingColUInt8{col: col, valPtr: valPtr})
}

type bindingColUInt8 struct {
	col    *proto.ColUInt8
	valPtr interface{}
}

func (b *bindingColUInt8) Append() {
	s, _ := b.valPtr.(*[]uint8)
	appendToSlice(s, b.col)
}

func (b *bindingColUInt8) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*uint8)
		*s = b.col.Row(0)
	}
}

// -- type Int16

func (r *ColResults) BindInt16(name string, valPtr interface{}) {
	col := &proto.ColInt16{}
	r.Bind(name, col, &bindingColInt16{col: col, valPtr: valPtr})
}

type bindingColInt16 struct {
	col    *proto.ColInt16
	valPtr interface{}
}

func (b *bindingColInt16) Append() {
	s, _ := b.valPtr.(*[]int16)
	appendToSlice(s, b.col)
}

func (b *bindingColInt16) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*int16)
		*s = b.col.Row(0)
	}
}

// -- type UInt16

func (r *ColResults) BindUInt16(name string, valPtr interface{}) {
	col := &proto.ColUInt16{}
	r.Bind(name, col, &bindingColUInt16{col: col, valPtr: valPtr})
}

type bindingColUInt16 struct {
	col    *proto.ColUInt16
	valPtr interface{}
}

func (b *bindingColUInt16) Append() {
	s, _ := b.valPtr.(*[]uint16)
	appendToSlice(s, b.col)
}

func (b *bindingColUInt16) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*uint16)
		*s = b.col.Row(0)
	}
}

// -- type Int32

func (r *ColResults) BindInt32(name string, valPtr interface{}) {
	col := &proto.ColInt32{}
	r.Bind(name, col, &bindingColInt32{col: col, valPtr: valPtr})
}

type bindingColInt32 struct {
	col    *proto.ColInt32
	valPtr interface{}
}

func (b *bindingColInt32) Append() {
	s, _ := b.valPtr.(*[]int32)
	appendToSlice(s, b.col)
}

func (b *bindingColInt32) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*int32)
		*s = b.col.Row(0)
	}
}

// -- type UInt32

func (r *ColResults) BindUInt32(name string, valPtr interface{}) {
	col := &proto.ColUInt32{}
	r.Bind(name, col, &bindingColUInt32{col: col, valPtr: valPtr})
}

type bindingColUInt32 struct {
	col    *proto.ColUInt32
	valPtr interface{}
}

func (b *bindingColUInt32) Append() {
	s, _ := b.valPtr.(*[]uint32)
	appendToSlice(s, b.col)
}

func (b *bindingColUInt32) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*uint32)
		*s = b.col.Row(0)
	}
}

// -- type Int64

func (r *ColResults) BindInt64(name string, valPtr interface{}) {
	col := &proto.ColInt64{}
	r.Bind(name, col, &bindingColInt64{col: col, valPtr: valPtr})
}

type bindingColInt64 struct {
	col    *proto.ColInt64
	valPtr interface{}
}

func (b *bindingColInt64) Append() {
	s, _ := b.valPtr.(*[]int64)
	appendToSlice(s, b.col)
}

func (b *bindingColInt64) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*int64)
		*s = b.col.Row(0)
	}
}

// -- type UInt64

func (r *ColResults) BindUInt64(name string, valPtr interface{}) {
	col := &proto.ColUInt64{}
	r.Bind(name, col, &bindingColUInt64{col: col, valPtr: valPtr})
}

type bindingColUInt64 struct {
	col    *proto.ColUInt64
	valPtr interface{}
}

func (b *bindingColUInt64) Append() {
	s, _ := b.valPtr.(*[]uint64)
	appendToSlice(s, b.col)
}

func (b *bindingColUInt64) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*uint64)
		*s = b.col.Row(0)
	}
}

// -- type Int128

func (r *ColResults) BindInt128(name string, valPtr interface{}) {
	col := &proto.ColInt128{}
	r.Bind(name, col, &bindingColInt128{col: col, valPtr: valPtr})
}

type bindingColInt128 struct {
	col    *proto.ColInt128
	valPtr interface{}
}

func (b *bindingColInt128) Append() {
	s, _ := b.valPtr.(*[]proto.Int128)
	appendToSlice(s, b.col)
}

func (b *bindingColInt128) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*proto.Int128)
		*s = b.col.Row(0)
	}
}

// -- type UInt128

func (r *ColResults) BindUInt128(name string, valPtr interface{}) {
	col := &proto.ColUInt128{}
	r.Bind(name, col, &bindingColUInt128{col: col, valPtr: valPtr})
}

type bindingColUInt128 struct {
	col    *proto.ColUInt128
	valPtr interface{}
}

func (b *bindingColUInt128) Append() {
	s, _ := b.valPtr.(*[]proto.UInt128)
	appendToSlice(s, b.col)
}

func (b *bindingColUInt128) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*proto.UInt128)
		*s = b.col.Row(0)
	}
}

// -- type Int256

func (r *ColResults) BindInt256(name string, valPtr interface{}) {
	col := &proto.ColInt256{}
	r.Bind(name, col, &bindingColInt256{col: col, valPtr: valPtr})
}

type bindingColInt256 struct {
	col    *proto.ColInt256
	valPtr interface{}
}

func (b *bindingColInt256) Append() {
	s, _ := b.valPtr.(*[]proto.Int256)
	appendToSlice(s, b.col)
}

func (b *bindingColInt256) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*proto.Int256)
		*s = b.col.Row(0)
	}
}

// -- type UInt256

func (r *ColResults) BindUInt256(name string, valPtr interface{}) {
	col := &proto.ColUInt256{}
	r.Bind(name, col, &bindingColUInt256{col: col, valPtr: valPtr})
}

type bindingColUInt256 struct {
	col    *proto.ColUInt256
	valPtr interface{}
}

func (b *bindingColUInt256) Append() {
	s, _ := b.valPtr.(*[]proto.UInt256)
	appendToSlice(s, b.col)
}

func (b *bindingColUInt256) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*proto.UInt256)
		*s = b.col.Row(0)
	}
}

// -- type FixedStr8

func (r *ColResults) BindFixedStr8(name string, valPtr interface{}) {
	col := &proto.ColFixedStr8{}
	r.Bind(name, col, &bindingColFixedStr8{col: col, valPtr: valPtr})
}

type bindingColFixedStr8 struct {
	col    *proto.ColFixedStr8
	valPtr interface{}
}

func (b *bindingColFixedStr8) Append() {
	s, _ := b.valPtr.(*[][8]byte)
	appendToSlice(s, b.col)
}

func (b *bindingColFixedStr8) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*[8]byte)
		*s = b.col.Row(0)
	}
}

// -- type FixedStr16

func (r *ColResults) BindFixedStr16(name string, valPtr interface{}) {
	col := &proto.ColFixedStr16{}
	r.Bind(name, col, &bindingColFixedStr16{col: col, valPtr: valPtr})
}

type bindingColFixedStr16 struct {
	col    *proto.ColFixedStr16
	valPtr interface{}
}

func (b *bindingColFixedStr16) Append() {
	s, _ := b.valPtr.(*[][16]byte)
	appendToSlice(s, b.col)
}

func (b *bindingColFixedStr16) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*[16]byte)
		*s = b.col.Row(0)
	}
}

// -- type FixedStr32

func (r *ColResults) BindFixedStr32(name string, valPtr interface{}) {
	col := &proto.ColFixedStr32{}
	r.Bind(name, col, &bindingColFixedStr32{col: col, valPtr: valPtr})
}

type bindingColFixedStr32 struct {
	col    *proto.ColFixedStr32
	valPtr interface{}
}

func (b *bindingColFixedStr32) Append() {
	s, _ := b.valPtr.(*[][32]byte)
	appendToSlice(s, b.col)
}

func (b *bindingColFixedStr32) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*[32]byte)
		*s = b.col.Row(0)
	}
}

// -- type FixedStr64

func (r *ColResults) BindFixedStr64(name string, valPtr interface{}) {
	col := &proto.ColFixedStr64{}
	r.Bind(name, col, &bindingColFixedStr64{col: col, valPtr: valPtr})
}

type bindingColFixedStr64 struct {
	col    *proto.ColFixedStr64
	valPtr interface{}
}

func (b *bindingColFixedStr64) Append() {
	s, _ := b.valPtr.(*[][64]byte)
	appendToSlice(s, b.col)
}

func (b *bindingColFixedStr64) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*[64]byte)
		*s = b.col.Row(0)
	}
}

// -- type FixedStr128

func (r *ColResults) BindFixedStr128(name string, valPtr interface{}) {
	col := &proto.ColFixedStr128{}
	r.Bind(name, col, &bindingColFixedStr128{col: col, valPtr: valPtr})
}

type bindingColFixedStr128 struct {
	col    *proto.ColFixedStr128
	valPtr interface{}
}

func (b *bindingColFixedStr128) Append() {
	s, _ := b.valPtr.(*[][128]byte)
	appendToSlice(s, b.col)
}

func (b *bindingColFixedStr128) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*[128]byte)
		*s = b.col.Row(0)
	}
}

// -- type FixedStr256

func (r *ColResults) BindFixedStr256(name string, valPtr interface{}) {
	col := &proto.ColFixedStr256{}
	r.Bind(name, col, &bindingColFixedStr256{col: col, valPtr: valPtr})
}

type bindingColFixedStr256 struct {
	col    *proto.ColFixedStr256
	valPtr interface{}
}

func (b *bindingColFixedStr256) Append() {
	s, _ := b.valPtr.(*[][256]byte)
	appendToSlice(s, b.col)
}

func (b *bindingColFixedStr256) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*[256]byte)
		*s = b.col.Row(0)
	}
}

// -- type FixedStr512

func (r *ColResults) BindFixedStr512(name string, valPtr interface{}) {
	col := &proto.ColFixedStr512{}
	r.Bind(name, col, &bindingColFixedStr512{col: col, valPtr: valPtr})
}

type bindingColFixedStr512 struct {
	col    *proto.ColFixedStr512
	valPtr interface{}
}

func (b *bindingColFixedStr512) Append() {
	s, _ := b.valPtr.(*[][512]byte)
	appendToSlice(s, b.col)
}

func (b *bindingColFixedStr512) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*[512]byte)
		*s = b.col.Row(0)
	}
}
