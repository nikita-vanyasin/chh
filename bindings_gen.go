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

// -- type Map[K]V

func (r *ColResults) BindMapStrToStr(name string, valPtr interface{}) {
	col := proto.NewMap[string, string](&proto.ColStr{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[string, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[string, float32](&proto.ColStr{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[string, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[string, float64](&proto.ColStr{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[string, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[string, proto.IPv4](&proto.ColStr{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[string, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[string, proto.IPv6](&proto.ColStr{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[string, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[string, time.Time](&proto.ColStr{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[string, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[string, time.Time](&proto.ColStr{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[string, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToDate(name string, valPtr interface{}) {
	col := proto.NewMap[string, time.Time](&proto.ColStr{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[string, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[string, time.Time](&proto.ColStr{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[string, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[string, proto.Enum8](&proto.ColStr{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[string, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[string, proto.Enum16](&proto.ColStr{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[string, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[string, proto.Decimal32](&proto.ColStr{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[string, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[string, proto.Decimal64](&proto.ColStr{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[string, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[string, proto.Decimal128](&proto.ColStr{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[string, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[string, proto.Decimal256](&proto.ColStr{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[string, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[string, int8](&proto.ColStr{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[string, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[string, uint8](&proto.ColStr{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[string, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[string, int16](&proto.ColStr{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[string, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[string, uint16](&proto.ColStr{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[string, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[string, int32](&proto.ColStr{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[string, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[string, uint32](&proto.ColStr{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[string, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[string, int64](&proto.ColStr{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[string, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[string, uint64](&proto.ColStr{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[string, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[string, proto.Int128](&proto.ColStr{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[string, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[string, proto.UInt128](&proto.ColStr{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[string, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[string, proto.Int256](&proto.ColStr{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[string, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[string, proto.UInt256](&proto.ColStr{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[string, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[string, [8]byte](&proto.ColStr{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[string, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[string, [16]byte](&proto.ColStr{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[string, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[string, [32]byte](&proto.ColStr{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[string, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[string, [64]byte](&proto.ColStr{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[string, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[string, [128]byte](&proto.ColStr{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[string, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[string, [256]byte](&proto.ColStr{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[string, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapStrToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[string, [512]byte](&proto.ColStr{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[string, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[float32, string](&proto.ColFloat32{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[float32, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[float32, float32](&proto.ColFloat32{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[float32, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[float32, float64](&proto.ColFloat32{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[float32, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[float32, proto.IPv4](&proto.ColFloat32{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[float32, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[float32, proto.IPv6](&proto.ColFloat32{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[float32, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[float32, time.Time](&proto.ColFloat32{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[float32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[float32, time.Time](&proto.ColFloat32{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[float32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[float32, time.Time](&proto.ColFloat32{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[float32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[float32, time.Time](&proto.ColFloat32{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[float32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[float32, proto.Enum8](&proto.ColFloat32{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[float32, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[float32, proto.Enum16](&proto.ColFloat32{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[float32, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[float32, proto.Decimal32](&proto.ColFloat32{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[float32, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[float32, proto.Decimal64](&proto.ColFloat32{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[float32, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[float32, proto.Decimal128](&proto.ColFloat32{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[float32, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[float32, proto.Decimal256](&proto.ColFloat32{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[float32, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[float32, int8](&proto.ColFloat32{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[float32, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[float32, uint8](&proto.ColFloat32{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[float32, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[float32, int16](&proto.ColFloat32{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[float32, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[float32, uint16](&proto.ColFloat32{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[float32, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[float32, int32](&proto.ColFloat32{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[float32, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[float32, uint32](&proto.ColFloat32{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[float32, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[float32, int64](&proto.ColFloat32{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[float32, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[float32, uint64](&proto.ColFloat32{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[float32, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[float32, proto.Int128](&proto.ColFloat32{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[float32, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[float32, proto.UInt128](&proto.ColFloat32{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[float32, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[float32, proto.Int256](&proto.ColFloat32{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[float32, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[float32, proto.UInt256](&proto.ColFloat32{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[float32, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[float32, [8]byte](&proto.ColFloat32{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[float32, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[float32, [16]byte](&proto.ColFloat32{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[float32, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[float32, [32]byte](&proto.ColFloat32{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[float32, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[float32, [64]byte](&proto.ColFloat32{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[float32, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[float32, [128]byte](&proto.ColFloat32{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[float32, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[float32, [256]byte](&proto.ColFloat32{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[float32, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat32ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[float32, [512]byte](&proto.ColFloat32{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[float32, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[float64, string](&proto.ColFloat64{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[float64, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[float64, float32](&proto.ColFloat64{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[float64, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[float64, float64](&proto.ColFloat64{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[float64, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[float64, proto.IPv4](&proto.ColFloat64{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[float64, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[float64, proto.IPv6](&proto.ColFloat64{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[float64, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[float64, time.Time](&proto.ColFloat64{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[float64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[float64, time.Time](&proto.ColFloat64{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[float64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[float64, time.Time](&proto.ColFloat64{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[float64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[float64, time.Time](&proto.ColFloat64{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[float64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[float64, proto.Enum8](&proto.ColFloat64{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[float64, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[float64, proto.Enum16](&proto.ColFloat64{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[float64, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[float64, proto.Decimal32](&proto.ColFloat64{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[float64, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[float64, proto.Decimal64](&proto.ColFloat64{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[float64, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[float64, proto.Decimal128](&proto.ColFloat64{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[float64, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[float64, proto.Decimal256](&proto.ColFloat64{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[float64, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[float64, int8](&proto.ColFloat64{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[float64, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[float64, uint8](&proto.ColFloat64{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[float64, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[float64, int16](&proto.ColFloat64{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[float64, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[float64, uint16](&proto.ColFloat64{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[float64, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[float64, int32](&proto.ColFloat64{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[float64, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[float64, uint32](&proto.ColFloat64{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[float64, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[float64, int64](&proto.ColFloat64{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[float64, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[float64, uint64](&proto.ColFloat64{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[float64, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[float64, proto.Int128](&proto.ColFloat64{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[float64, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[float64, proto.UInt128](&proto.ColFloat64{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[float64, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[float64, proto.Int256](&proto.ColFloat64{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[float64, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[float64, proto.UInt256](&proto.ColFloat64{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[float64, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[float64, [8]byte](&proto.ColFloat64{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[float64, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[float64, [16]byte](&proto.ColFloat64{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[float64, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[float64, [32]byte](&proto.ColFloat64{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[float64, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[float64, [64]byte](&proto.ColFloat64{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[float64, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[float64, [128]byte](&proto.ColFloat64{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[float64, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[float64, [256]byte](&proto.ColFloat64{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[float64, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFloat64ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[float64, [512]byte](&proto.ColFloat64{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[float64, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, string](&proto.ColIPv4{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[proto.IPv4, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, float32](&proto.ColIPv4{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[proto.IPv4, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, float64](&proto.ColIPv4{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[proto.IPv4, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, proto.IPv4](&proto.ColIPv4{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[proto.IPv4, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, proto.IPv6](&proto.ColIPv4{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[proto.IPv4, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, time.Time](&proto.ColIPv4{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[proto.IPv4, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, time.Time](&proto.ColIPv4{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[proto.IPv4, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, time.Time](&proto.ColIPv4{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[proto.IPv4, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, time.Time](&proto.ColIPv4{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[proto.IPv4, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, proto.Enum8](&proto.ColIPv4{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[proto.IPv4, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, proto.Enum16](&proto.ColIPv4{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[proto.IPv4, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, proto.Decimal32](&proto.ColIPv4{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[proto.IPv4, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, proto.Decimal64](&proto.ColIPv4{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[proto.IPv4, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, proto.Decimal128](&proto.ColIPv4{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[proto.IPv4, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, proto.Decimal256](&proto.ColIPv4{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[proto.IPv4, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, int8](&proto.ColIPv4{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[proto.IPv4, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, uint8](&proto.ColIPv4{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[proto.IPv4, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, int16](&proto.ColIPv4{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[proto.IPv4, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, uint16](&proto.ColIPv4{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[proto.IPv4, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, int32](&proto.ColIPv4{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[proto.IPv4, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, uint32](&proto.ColIPv4{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[proto.IPv4, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, int64](&proto.ColIPv4{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[proto.IPv4, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, uint64](&proto.ColIPv4{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[proto.IPv4, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, proto.Int128](&proto.ColIPv4{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[proto.IPv4, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, proto.UInt128](&proto.ColIPv4{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[proto.IPv4, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, proto.Int256](&proto.ColIPv4{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[proto.IPv4, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, proto.UInt256](&proto.ColIPv4{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[proto.IPv4, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, [8]byte](&proto.ColIPv4{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[proto.IPv4, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, [16]byte](&proto.ColIPv4{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[proto.IPv4, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, [32]byte](&proto.ColIPv4{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[proto.IPv4, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, [64]byte](&proto.ColIPv4{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[proto.IPv4, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, [128]byte](&proto.ColIPv4{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[proto.IPv4, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, [256]byte](&proto.ColIPv4{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[proto.IPv4, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv4ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv4, [512]byte](&proto.ColIPv4{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[proto.IPv4, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, string](&proto.ColIPv6{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[proto.IPv6, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, float32](&proto.ColIPv6{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[proto.IPv6, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, float64](&proto.ColIPv6{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[proto.IPv6, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, proto.IPv4](&proto.ColIPv6{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[proto.IPv6, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, proto.IPv6](&proto.ColIPv6{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[proto.IPv6, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, time.Time](&proto.ColIPv6{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[proto.IPv6, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, time.Time](&proto.ColIPv6{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[proto.IPv6, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, time.Time](&proto.ColIPv6{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[proto.IPv6, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, time.Time](&proto.ColIPv6{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[proto.IPv6, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, proto.Enum8](&proto.ColIPv6{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[proto.IPv6, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, proto.Enum16](&proto.ColIPv6{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[proto.IPv6, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, proto.Decimal32](&proto.ColIPv6{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[proto.IPv6, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, proto.Decimal64](&proto.ColIPv6{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[proto.IPv6, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, proto.Decimal128](&proto.ColIPv6{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[proto.IPv6, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, proto.Decimal256](&proto.ColIPv6{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[proto.IPv6, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, int8](&proto.ColIPv6{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[proto.IPv6, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, uint8](&proto.ColIPv6{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[proto.IPv6, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, int16](&proto.ColIPv6{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[proto.IPv6, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, uint16](&proto.ColIPv6{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[proto.IPv6, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, int32](&proto.ColIPv6{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[proto.IPv6, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, uint32](&proto.ColIPv6{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[proto.IPv6, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, int64](&proto.ColIPv6{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[proto.IPv6, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, uint64](&proto.ColIPv6{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[proto.IPv6, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, proto.Int128](&proto.ColIPv6{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[proto.IPv6, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, proto.UInt128](&proto.ColIPv6{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[proto.IPv6, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, proto.Int256](&proto.ColIPv6{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[proto.IPv6, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, proto.UInt256](&proto.ColIPv6{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[proto.IPv6, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, [8]byte](&proto.ColIPv6{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[proto.IPv6, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, [16]byte](&proto.ColIPv6{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[proto.IPv6, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, [32]byte](&proto.ColIPv6{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[proto.IPv6, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, [64]byte](&proto.ColIPv6{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[proto.IPv6, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, [128]byte](&proto.ColIPv6{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[proto.IPv6, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, [256]byte](&proto.ColIPv6{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[proto.IPv6, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapIPv6ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[proto.IPv6, [512]byte](&proto.ColIPv6{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[proto.IPv6, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToStr(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, string](&proto.ColDateTime{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[time.Time, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, float32](&proto.ColDateTime{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[time.Time, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, float64](&proto.ColDateTime{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[time.Time, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.IPv4](&proto.ColDateTime{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[time.Time, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.IPv6](&proto.ColDateTime{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[time.Time, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDateTime{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDateTime{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToDate(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDateTime{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDateTime{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Enum8](&proto.ColDateTime{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Enum16](&proto.ColDateTime{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal32](&proto.ColDateTime{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal64](&proto.ColDateTime{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal128](&proto.ColDateTime{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal256](&proto.ColDateTime{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int8](&proto.ColDateTime{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[time.Time, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint8](&proto.ColDateTime{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[time.Time, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int16](&proto.ColDateTime{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[time.Time, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint16](&proto.ColDateTime{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[time.Time, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int32](&proto.ColDateTime{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[time.Time, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint32](&proto.ColDateTime{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[time.Time, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int64](&proto.ColDateTime{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[time.Time, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint64](&proto.ColDateTime{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[time.Time, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Int128](&proto.ColDateTime{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.UInt128](&proto.ColDateTime{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[time.Time, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Int256](&proto.ColDateTime{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.UInt256](&proto.ColDateTime{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[time.Time, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [8]byte](&proto.ColDateTime{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[time.Time, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [16]byte](&proto.ColDateTime{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[time.Time, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [32]byte](&proto.ColDateTime{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[time.Time, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [64]byte](&proto.ColDateTime{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[time.Time, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [128]byte](&proto.ColDateTime{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[time.Time, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [256]byte](&proto.ColDateTime{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[time.Time, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTimeToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [512]byte](&proto.ColDateTime{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[time.Time, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, string](&proto.ColDateTime64{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[time.Time, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, float32](&proto.ColDateTime64{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[time.Time, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, float64](&proto.ColDateTime64{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[time.Time, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.IPv4](&proto.ColDateTime64{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[time.Time, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.IPv6](&proto.ColDateTime64{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[time.Time, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDateTime64{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDateTime64{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDateTime64{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDateTime64{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Enum8](&proto.ColDateTime64{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Enum16](&proto.ColDateTime64{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal32](&proto.ColDateTime64{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal64](&proto.ColDateTime64{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal128](&proto.ColDateTime64{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal256](&proto.ColDateTime64{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int8](&proto.ColDateTime64{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[time.Time, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint8](&proto.ColDateTime64{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[time.Time, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int16](&proto.ColDateTime64{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[time.Time, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint16](&proto.ColDateTime64{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[time.Time, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int32](&proto.ColDateTime64{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[time.Time, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint32](&proto.ColDateTime64{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[time.Time, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int64](&proto.ColDateTime64{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[time.Time, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint64](&proto.ColDateTime64{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[time.Time, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Int128](&proto.ColDateTime64{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.UInt128](&proto.ColDateTime64{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[time.Time, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Int256](&proto.ColDateTime64{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.UInt256](&proto.ColDateTime64{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[time.Time, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [8]byte](&proto.ColDateTime64{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[time.Time, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [16]byte](&proto.ColDateTime64{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[time.Time, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [32]byte](&proto.ColDateTime64{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[time.Time, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [64]byte](&proto.ColDateTime64{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[time.Time, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [128]byte](&proto.ColDateTime64{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[time.Time, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [256]byte](&proto.ColDateTime64{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[time.Time, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateTime64ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [512]byte](&proto.ColDateTime64{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[time.Time, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToStr(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, string](&proto.ColDate{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[time.Time, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, float32](&proto.ColDate{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[time.Time, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, float64](&proto.ColDate{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[time.Time, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.IPv4](&proto.ColDate{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[time.Time, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.IPv6](&proto.ColDate{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[time.Time, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDate{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDate{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToDate(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDate{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDate{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Enum8](&proto.ColDate{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Enum16](&proto.ColDate{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal32](&proto.ColDate{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal64](&proto.ColDate{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal128](&proto.ColDate{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal256](&proto.ColDate{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int8](&proto.ColDate{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[time.Time, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint8](&proto.ColDate{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[time.Time, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int16](&proto.ColDate{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[time.Time, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint16](&proto.ColDate{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[time.Time, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int32](&proto.ColDate{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[time.Time, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint32](&proto.ColDate{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[time.Time, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int64](&proto.ColDate{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[time.Time, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint64](&proto.ColDate{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[time.Time, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Int128](&proto.ColDate{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.UInt128](&proto.ColDate{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[time.Time, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Int256](&proto.ColDate{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.UInt256](&proto.ColDate{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[time.Time, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [8]byte](&proto.ColDate{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[time.Time, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [16]byte](&proto.ColDate{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[time.Time, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [32]byte](&proto.ColDate{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[time.Time, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [64]byte](&proto.ColDate{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[time.Time, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [128]byte](&proto.ColDate{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[time.Time, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [256]byte](&proto.ColDate{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[time.Time, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDateToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [512]byte](&proto.ColDate{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[time.Time, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, string](&proto.ColDate32{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[time.Time, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, float32](&proto.ColDate32{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[time.Time, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, float64](&proto.ColDate32{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[time.Time, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.IPv4](&proto.ColDate32{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[time.Time, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.IPv6](&proto.ColDate32{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[time.Time, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDate32{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDate32{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDate32{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, time.Time](&proto.ColDate32{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[time.Time, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Enum8](&proto.ColDate32{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Enum16](&proto.ColDate32{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal32](&proto.ColDate32{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal64](&proto.ColDate32{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal128](&proto.ColDate32{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Decimal256](&proto.ColDate32{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int8](&proto.ColDate32{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[time.Time, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint8](&proto.ColDate32{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[time.Time, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int16](&proto.ColDate32{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[time.Time, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint16](&proto.ColDate32{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[time.Time, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int32](&proto.ColDate32{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[time.Time, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint32](&proto.ColDate32{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[time.Time, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, int64](&proto.ColDate32{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[time.Time, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, uint64](&proto.ColDate32{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[time.Time, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Int128](&proto.ColDate32{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.UInt128](&proto.ColDate32{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[time.Time, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.Int256](&proto.ColDate32{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[time.Time, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, proto.UInt256](&proto.ColDate32{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[time.Time, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [8]byte](&proto.ColDate32{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[time.Time, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [16]byte](&proto.ColDate32{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[time.Time, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [32]byte](&proto.ColDate32{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[time.Time, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [64]byte](&proto.ColDate32{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[time.Time, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [128]byte](&proto.ColDate32{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[time.Time, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [256]byte](&proto.ColDate32{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[time.Time, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDate32ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[time.Time, [512]byte](&proto.ColDate32{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[time.Time, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, string](&proto.ColEnum8{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[proto.Enum8, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, float32](&proto.ColEnum8{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[proto.Enum8, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, float64](&proto.ColEnum8{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[proto.Enum8, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, proto.IPv4](&proto.ColEnum8{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[proto.Enum8, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, proto.IPv6](&proto.ColEnum8{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[proto.Enum8, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, time.Time](&proto.ColEnum8{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[proto.Enum8, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, time.Time](&proto.ColEnum8{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[proto.Enum8, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, time.Time](&proto.ColEnum8{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[proto.Enum8, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, time.Time](&proto.ColEnum8{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[proto.Enum8, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, proto.Enum8](&proto.ColEnum8{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[proto.Enum8, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, proto.Enum16](&proto.ColEnum8{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[proto.Enum8, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, proto.Decimal32](&proto.ColEnum8{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[proto.Enum8, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, proto.Decimal64](&proto.ColEnum8{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[proto.Enum8, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, proto.Decimal128](&proto.ColEnum8{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[proto.Enum8, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, proto.Decimal256](&proto.ColEnum8{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[proto.Enum8, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, int8](&proto.ColEnum8{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[proto.Enum8, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, uint8](&proto.ColEnum8{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[proto.Enum8, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, int16](&proto.ColEnum8{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[proto.Enum8, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, uint16](&proto.ColEnum8{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[proto.Enum8, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, int32](&proto.ColEnum8{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[proto.Enum8, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, uint32](&proto.ColEnum8{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[proto.Enum8, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, int64](&proto.ColEnum8{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[proto.Enum8, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, uint64](&proto.ColEnum8{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[proto.Enum8, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, proto.Int128](&proto.ColEnum8{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[proto.Enum8, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, proto.UInt128](&proto.ColEnum8{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[proto.Enum8, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, proto.Int256](&proto.ColEnum8{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[proto.Enum8, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, proto.UInt256](&proto.ColEnum8{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[proto.Enum8, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, [8]byte](&proto.ColEnum8{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[proto.Enum8, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, [16]byte](&proto.ColEnum8{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[proto.Enum8, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, [32]byte](&proto.ColEnum8{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[proto.Enum8, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, [64]byte](&proto.ColEnum8{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[proto.Enum8, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, [128]byte](&proto.ColEnum8{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[proto.Enum8, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, [256]byte](&proto.ColEnum8{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[proto.Enum8, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum8ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum8, [512]byte](&proto.ColEnum8{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[proto.Enum8, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, string](&proto.ColEnum16{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[proto.Enum16, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, float32](&proto.ColEnum16{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[proto.Enum16, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, float64](&proto.ColEnum16{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[proto.Enum16, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, proto.IPv4](&proto.ColEnum16{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[proto.Enum16, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, proto.IPv6](&proto.ColEnum16{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[proto.Enum16, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, time.Time](&proto.ColEnum16{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[proto.Enum16, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, time.Time](&proto.ColEnum16{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[proto.Enum16, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, time.Time](&proto.ColEnum16{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[proto.Enum16, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, time.Time](&proto.ColEnum16{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[proto.Enum16, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, proto.Enum8](&proto.ColEnum16{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[proto.Enum16, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, proto.Enum16](&proto.ColEnum16{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[proto.Enum16, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, proto.Decimal32](&proto.ColEnum16{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[proto.Enum16, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, proto.Decimal64](&proto.ColEnum16{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[proto.Enum16, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, proto.Decimal128](&proto.ColEnum16{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[proto.Enum16, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, proto.Decimal256](&proto.ColEnum16{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[proto.Enum16, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, int8](&proto.ColEnum16{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[proto.Enum16, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, uint8](&proto.ColEnum16{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[proto.Enum16, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, int16](&proto.ColEnum16{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[proto.Enum16, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, uint16](&proto.ColEnum16{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[proto.Enum16, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, int32](&proto.ColEnum16{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[proto.Enum16, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, uint32](&proto.ColEnum16{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[proto.Enum16, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, int64](&proto.ColEnum16{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[proto.Enum16, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, uint64](&proto.ColEnum16{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[proto.Enum16, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, proto.Int128](&proto.ColEnum16{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[proto.Enum16, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, proto.UInt128](&proto.ColEnum16{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[proto.Enum16, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, proto.Int256](&proto.ColEnum16{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[proto.Enum16, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, proto.UInt256](&proto.ColEnum16{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[proto.Enum16, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, [8]byte](&proto.ColEnum16{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[proto.Enum16, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, [16]byte](&proto.ColEnum16{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[proto.Enum16, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, [32]byte](&proto.ColEnum16{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[proto.Enum16, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, [64]byte](&proto.ColEnum16{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[proto.Enum16, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, [128]byte](&proto.ColEnum16{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[proto.Enum16, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, [256]byte](&proto.ColEnum16{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[proto.Enum16, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapEnum16ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Enum16, [512]byte](&proto.ColEnum16{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[proto.Enum16, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, string](&proto.ColDecimal32{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, float32](&proto.ColDecimal32{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, float64](&proto.ColDecimal32{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, proto.IPv4](&proto.ColDecimal32{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, proto.IPv6](&proto.ColDecimal32{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, time.Time](&proto.ColDecimal32{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, time.Time](&proto.ColDecimal32{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, time.Time](&proto.ColDecimal32{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, time.Time](&proto.ColDecimal32{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, proto.Enum8](&proto.ColDecimal32{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, proto.Enum16](&proto.ColDecimal32{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, proto.Decimal32](&proto.ColDecimal32{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, proto.Decimal64](&proto.ColDecimal32{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, proto.Decimal128](&proto.ColDecimal32{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, proto.Decimal256](&proto.ColDecimal32{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, int8](&proto.ColDecimal32{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, uint8](&proto.ColDecimal32{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, int16](&proto.ColDecimal32{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, uint16](&proto.ColDecimal32{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, int32](&proto.ColDecimal32{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, uint32](&proto.ColDecimal32{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, int64](&proto.ColDecimal32{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, uint64](&proto.ColDecimal32{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, proto.Int128](&proto.ColDecimal32{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, proto.UInt128](&proto.ColDecimal32{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, proto.Int256](&proto.ColDecimal32{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, proto.UInt256](&proto.ColDecimal32{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, [8]byte](&proto.ColDecimal32{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, [16]byte](&proto.ColDecimal32{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, [32]byte](&proto.ColDecimal32{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, [64]byte](&proto.ColDecimal32{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, [128]byte](&proto.ColDecimal32{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, [256]byte](&proto.ColDecimal32{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal32ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal32, [512]byte](&proto.ColDecimal32{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[proto.Decimal32, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, string](&proto.ColDecimal64{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, float32](&proto.ColDecimal64{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, float64](&proto.ColDecimal64{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, proto.IPv4](&proto.ColDecimal64{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, proto.IPv6](&proto.ColDecimal64{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, time.Time](&proto.ColDecimal64{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, time.Time](&proto.ColDecimal64{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, time.Time](&proto.ColDecimal64{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, time.Time](&proto.ColDecimal64{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, proto.Enum8](&proto.ColDecimal64{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, proto.Enum16](&proto.ColDecimal64{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, proto.Decimal32](&proto.ColDecimal64{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, proto.Decimal64](&proto.ColDecimal64{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, proto.Decimal128](&proto.ColDecimal64{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, proto.Decimal256](&proto.ColDecimal64{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, int8](&proto.ColDecimal64{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, uint8](&proto.ColDecimal64{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, int16](&proto.ColDecimal64{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, uint16](&proto.ColDecimal64{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, int32](&proto.ColDecimal64{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, uint32](&proto.ColDecimal64{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, int64](&proto.ColDecimal64{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, uint64](&proto.ColDecimal64{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, proto.Int128](&proto.ColDecimal64{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, proto.UInt128](&proto.ColDecimal64{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, proto.Int256](&proto.ColDecimal64{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, proto.UInt256](&proto.ColDecimal64{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, [8]byte](&proto.ColDecimal64{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, [16]byte](&proto.ColDecimal64{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, [32]byte](&proto.ColDecimal64{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, [64]byte](&proto.ColDecimal64{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, [128]byte](&proto.ColDecimal64{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, [256]byte](&proto.ColDecimal64{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal64ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal64, [512]byte](&proto.ColDecimal64{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[proto.Decimal64, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, string](&proto.ColDecimal128{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, float32](&proto.ColDecimal128{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, float64](&proto.ColDecimal128{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, proto.IPv4](&proto.ColDecimal128{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, proto.IPv6](&proto.ColDecimal128{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, time.Time](&proto.ColDecimal128{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, time.Time](&proto.ColDecimal128{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, time.Time](&proto.ColDecimal128{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, time.Time](&proto.ColDecimal128{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, proto.Enum8](&proto.ColDecimal128{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, proto.Enum16](&proto.ColDecimal128{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, proto.Decimal32](&proto.ColDecimal128{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, proto.Decimal64](&proto.ColDecimal128{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, proto.Decimal128](&proto.ColDecimal128{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, proto.Decimal256](&proto.ColDecimal128{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, int8](&proto.ColDecimal128{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, uint8](&proto.ColDecimal128{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, int16](&proto.ColDecimal128{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, uint16](&proto.ColDecimal128{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, int32](&proto.ColDecimal128{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, uint32](&proto.ColDecimal128{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, int64](&proto.ColDecimal128{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, uint64](&proto.ColDecimal128{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, proto.Int128](&proto.ColDecimal128{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, proto.UInt128](&proto.ColDecimal128{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, proto.Int256](&proto.ColDecimal128{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, proto.UInt256](&proto.ColDecimal128{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, [8]byte](&proto.ColDecimal128{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, [16]byte](&proto.ColDecimal128{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, [32]byte](&proto.ColDecimal128{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, [64]byte](&proto.ColDecimal128{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, [128]byte](&proto.ColDecimal128{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, [256]byte](&proto.ColDecimal128{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal128ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal128, [512]byte](&proto.ColDecimal128{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[proto.Decimal128, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, string](&proto.ColDecimal256{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, float32](&proto.ColDecimal256{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, float64](&proto.ColDecimal256{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, proto.IPv4](&proto.ColDecimal256{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, proto.IPv6](&proto.ColDecimal256{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, time.Time](&proto.ColDecimal256{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, time.Time](&proto.ColDecimal256{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, time.Time](&proto.ColDecimal256{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, time.Time](&proto.ColDecimal256{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, proto.Enum8](&proto.ColDecimal256{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, proto.Enum16](&proto.ColDecimal256{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, proto.Decimal32](&proto.ColDecimal256{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, proto.Decimal64](&proto.ColDecimal256{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, proto.Decimal128](&proto.ColDecimal256{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, proto.Decimal256](&proto.ColDecimal256{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, int8](&proto.ColDecimal256{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, uint8](&proto.ColDecimal256{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, int16](&proto.ColDecimal256{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, uint16](&proto.ColDecimal256{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, int32](&proto.ColDecimal256{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, uint32](&proto.ColDecimal256{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, int64](&proto.ColDecimal256{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, uint64](&proto.ColDecimal256{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, proto.Int128](&proto.ColDecimal256{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, proto.UInt128](&proto.ColDecimal256{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, proto.Int256](&proto.ColDecimal256{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, proto.UInt256](&proto.ColDecimal256{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, [8]byte](&proto.ColDecimal256{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, [16]byte](&proto.ColDecimal256{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, [32]byte](&proto.ColDecimal256{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, [64]byte](&proto.ColDecimal256{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, [128]byte](&proto.ColDecimal256{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, [256]byte](&proto.ColDecimal256{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapDecimal256ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Decimal256, [512]byte](&proto.ColDecimal256{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[proto.Decimal256, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[int8, string](&proto.ColInt8{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[int8, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[int8, float32](&proto.ColInt8{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[int8, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[int8, float64](&proto.ColInt8{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[int8, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[int8, proto.IPv4](&proto.ColInt8{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[int8, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[int8, proto.IPv6](&proto.ColInt8{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[int8, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[int8, time.Time](&proto.ColInt8{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[int8, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[int8, time.Time](&proto.ColInt8{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[int8, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[int8, time.Time](&proto.ColInt8{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[int8, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[int8, time.Time](&proto.ColInt8{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[int8, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[int8, proto.Enum8](&proto.ColInt8{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[int8, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[int8, proto.Enum16](&proto.ColInt8{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[int8, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[int8, proto.Decimal32](&proto.ColInt8{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[int8, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[int8, proto.Decimal64](&proto.ColInt8{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[int8, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[int8, proto.Decimal128](&proto.ColInt8{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[int8, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[int8, proto.Decimal256](&proto.ColInt8{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[int8, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[int8, int8](&proto.ColInt8{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[int8, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[int8, uint8](&proto.ColInt8{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[int8, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[int8, int16](&proto.ColInt8{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[int8, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[int8, uint16](&proto.ColInt8{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[int8, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[int8, int32](&proto.ColInt8{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[int8, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[int8, uint32](&proto.ColInt8{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[int8, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[int8, int64](&proto.ColInt8{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[int8, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[int8, uint64](&proto.ColInt8{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[int8, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[int8, proto.Int128](&proto.ColInt8{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[int8, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[int8, proto.UInt128](&proto.ColInt8{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[int8, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[int8, proto.Int256](&proto.ColInt8{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[int8, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[int8, proto.UInt256](&proto.ColInt8{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[int8, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[int8, [8]byte](&proto.ColInt8{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[int8, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[int8, [16]byte](&proto.ColInt8{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[int8, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[int8, [32]byte](&proto.ColInt8{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[int8, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[int8, [64]byte](&proto.ColInt8{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[int8, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[int8, [128]byte](&proto.ColInt8{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[int8, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[int8, [256]byte](&proto.ColInt8{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[int8, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt8ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[int8, [512]byte](&proto.ColInt8{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[int8, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, string](&proto.ColUInt8{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[uint8, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, float32](&proto.ColUInt8{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[uint8, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, float64](&proto.ColUInt8{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[uint8, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, proto.IPv4](&proto.ColUInt8{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[uint8, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, proto.IPv6](&proto.ColUInt8{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[uint8, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, time.Time](&proto.ColUInt8{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[uint8, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, time.Time](&proto.ColUInt8{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[uint8, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, time.Time](&proto.ColUInt8{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[uint8, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, time.Time](&proto.ColUInt8{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[uint8, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, proto.Enum8](&proto.ColUInt8{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[uint8, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, proto.Enum16](&proto.ColUInt8{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[uint8, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, proto.Decimal32](&proto.ColUInt8{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[uint8, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, proto.Decimal64](&proto.ColUInt8{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[uint8, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, proto.Decimal128](&proto.ColUInt8{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[uint8, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, proto.Decimal256](&proto.ColUInt8{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[uint8, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, int8](&proto.ColUInt8{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[uint8, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, uint8](&proto.ColUInt8{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[uint8, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, int16](&proto.ColUInt8{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[uint8, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, uint16](&proto.ColUInt8{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[uint8, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, int32](&proto.ColUInt8{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[uint8, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, uint32](&proto.ColUInt8{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[uint8, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, int64](&proto.ColUInt8{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[uint8, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, uint64](&proto.ColUInt8{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[uint8, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, proto.Int128](&proto.ColUInt8{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[uint8, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, proto.UInt128](&proto.ColUInt8{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[uint8, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, proto.Int256](&proto.ColUInt8{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[uint8, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, proto.UInt256](&proto.ColUInt8{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[uint8, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, [8]byte](&proto.ColUInt8{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[uint8, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, [16]byte](&proto.ColUInt8{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[uint8, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, [32]byte](&proto.ColUInt8{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[uint8, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, [64]byte](&proto.ColUInt8{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[uint8, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, [128]byte](&proto.ColUInt8{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[uint8, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, [256]byte](&proto.ColUInt8{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[uint8, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt8ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[uint8, [512]byte](&proto.ColUInt8{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[uint8, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[int16, string](&proto.ColInt16{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[int16, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[int16, float32](&proto.ColInt16{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[int16, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[int16, float64](&proto.ColInt16{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[int16, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[int16, proto.IPv4](&proto.ColInt16{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[int16, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[int16, proto.IPv6](&proto.ColInt16{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[int16, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[int16, time.Time](&proto.ColInt16{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[int16, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[int16, time.Time](&proto.ColInt16{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[int16, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[int16, time.Time](&proto.ColInt16{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[int16, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[int16, time.Time](&proto.ColInt16{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[int16, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[int16, proto.Enum8](&proto.ColInt16{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[int16, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[int16, proto.Enum16](&proto.ColInt16{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[int16, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[int16, proto.Decimal32](&proto.ColInt16{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[int16, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[int16, proto.Decimal64](&proto.ColInt16{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[int16, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[int16, proto.Decimal128](&proto.ColInt16{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[int16, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[int16, proto.Decimal256](&proto.ColInt16{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[int16, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[int16, int8](&proto.ColInt16{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[int16, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[int16, uint8](&proto.ColInt16{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[int16, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[int16, int16](&proto.ColInt16{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[int16, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[int16, uint16](&proto.ColInt16{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[int16, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[int16, int32](&proto.ColInt16{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[int16, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[int16, uint32](&proto.ColInt16{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[int16, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[int16, int64](&proto.ColInt16{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[int16, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[int16, uint64](&proto.ColInt16{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[int16, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[int16, proto.Int128](&proto.ColInt16{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[int16, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[int16, proto.UInt128](&proto.ColInt16{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[int16, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[int16, proto.Int256](&proto.ColInt16{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[int16, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[int16, proto.UInt256](&proto.ColInt16{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[int16, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[int16, [8]byte](&proto.ColInt16{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[int16, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[int16, [16]byte](&proto.ColInt16{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[int16, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[int16, [32]byte](&proto.ColInt16{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[int16, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[int16, [64]byte](&proto.ColInt16{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[int16, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[int16, [128]byte](&proto.ColInt16{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[int16, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[int16, [256]byte](&proto.ColInt16{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[int16, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt16ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[int16, [512]byte](&proto.ColInt16{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[int16, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, string](&proto.ColUInt16{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[uint16, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, float32](&proto.ColUInt16{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[uint16, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, float64](&proto.ColUInt16{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[uint16, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, proto.IPv4](&proto.ColUInt16{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[uint16, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, proto.IPv6](&proto.ColUInt16{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[uint16, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, time.Time](&proto.ColUInt16{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[uint16, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, time.Time](&proto.ColUInt16{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[uint16, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, time.Time](&proto.ColUInt16{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[uint16, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, time.Time](&proto.ColUInt16{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[uint16, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, proto.Enum8](&proto.ColUInt16{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[uint16, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, proto.Enum16](&proto.ColUInt16{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[uint16, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, proto.Decimal32](&proto.ColUInt16{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[uint16, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, proto.Decimal64](&proto.ColUInt16{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[uint16, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, proto.Decimal128](&proto.ColUInt16{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[uint16, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, proto.Decimal256](&proto.ColUInt16{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[uint16, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, int8](&proto.ColUInt16{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[uint16, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, uint8](&proto.ColUInt16{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[uint16, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, int16](&proto.ColUInt16{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[uint16, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, uint16](&proto.ColUInt16{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[uint16, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, int32](&proto.ColUInt16{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[uint16, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, uint32](&proto.ColUInt16{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[uint16, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, int64](&proto.ColUInt16{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[uint16, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, uint64](&proto.ColUInt16{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[uint16, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, proto.Int128](&proto.ColUInt16{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[uint16, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, proto.UInt128](&proto.ColUInt16{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[uint16, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, proto.Int256](&proto.ColUInt16{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[uint16, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, proto.UInt256](&proto.ColUInt16{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[uint16, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, [8]byte](&proto.ColUInt16{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[uint16, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, [16]byte](&proto.ColUInt16{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[uint16, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, [32]byte](&proto.ColUInt16{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[uint16, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, [64]byte](&proto.ColUInt16{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[uint16, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, [128]byte](&proto.ColUInt16{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[uint16, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, [256]byte](&proto.ColUInt16{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[uint16, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt16ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[uint16, [512]byte](&proto.ColUInt16{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[uint16, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[int32, string](&proto.ColInt32{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[int32, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[int32, float32](&proto.ColInt32{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[int32, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[int32, float64](&proto.ColInt32{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[int32, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[int32, proto.IPv4](&proto.ColInt32{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[int32, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[int32, proto.IPv6](&proto.ColInt32{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[int32, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[int32, time.Time](&proto.ColInt32{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[int32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[int32, time.Time](&proto.ColInt32{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[int32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[int32, time.Time](&proto.ColInt32{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[int32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[int32, time.Time](&proto.ColInt32{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[int32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[int32, proto.Enum8](&proto.ColInt32{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[int32, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[int32, proto.Enum16](&proto.ColInt32{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[int32, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[int32, proto.Decimal32](&proto.ColInt32{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[int32, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[int32, proto.Decimal64](&proto.ColInt32{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[int32, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[int32, proto.Decimal128](&proto.ColInt32{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[int32, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[int32, proto.Decimal256](&proto.ColInt32{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[int32, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[int32, int8](&proto.ColInt32{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[int32, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[int32, uint8](&proto.ColInt32{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[int32, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[int32, int16](&proto.ColInt32{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[int32, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[int32, uint16](&proto.ColInt32{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[int32, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[int32, int32](&proto.ColInt32{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[int32, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[int32, uint32](&proto.ColInt32{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[int32, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[int32, int64](&proto.ColInt32{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[int32, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[int32, uint64](&proto.ColInt32{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[int32, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[int32, proto.Int128](&proto.ColInt32{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[int32, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[int32, proto.UInt128](&proto.ColInt32{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[int32, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[int32, proto.Int256](&proto.ColInt32{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[int32, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[int32, proto.UInt256](&proto.ColInt32{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[int32, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[int32, [8]byte](&proto.ColInt32{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[int32, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[int32, [16]byte](&proto.ColInt32{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[int32, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[int32, [32]byte](&proto.ColInt32{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[int32, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[int32, [64]byte](&proto.ColInt32{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[int32, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[int32, [128]byte](&proto.ColInt32{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[int32, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[int32, [256]byte](&proto.ColInt32{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[int32, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt32ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[int32, [512]byte](&proto.ColInt32{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[int32, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, string](&proto.ColUInt32{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[uint32, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, float32](&proto.ColUInt32{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[uint32, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, float64](&proto.ColUInt32{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[uint32, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, proto.IPv4](&proto.ColUInt32{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[uint32, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, proto.IPv6](&proto.ColUInt32{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[uint32, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, time.Time](&proto.ColUInt32{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[uint32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, time.Time](&proto.ColUInt32{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[uint32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, time.Time](&proto.ColUInt32{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[uint32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, time.Time](&proto.ColUInt32{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[uint32, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, proto.Enum8](&proto.ColUInt32{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[uint32, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, proto.Enum16](&proto.ColUInt32{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[uint32, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, proto.Decimal32](&proto.ColUInt32{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[uint32, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, proto.Decimal64](&proto.ColUInt32{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[uint32, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, proto.Decimal128](&proto.ColUInt32{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[uint32, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, proto.Decimal256](&proto.ColUInt32{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[uint32, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, int8](&proto.ColUInt32{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[uint32, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, uint8](&proto.ColUInt32{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[uint32, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, int16](&proto.ColUInt32{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[uint32, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, uint16](&proto.ColUInt32{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[uint32, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, int32](&proto.ColUInt32{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[uint32, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, uint32](&proto.ColUInt32{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[uint32, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, int64](&proto.ColUInt32{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[uint32, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, uint64](&proto.ColUInt32{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[uint32, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, proto.Int128](&proto.ColUInt32{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[uint32, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, proto.UInt128](&proto.ColUInt32{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[uint32, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, proto.Int256](&proto.ColUInt32{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[uint32, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, proto.UInt256](&proto.ColUInt32{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[uint32, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, [8]byte](&proto.ColUInt32{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[uint32, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, [16]byte](&proto.ColUInt32{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[uint32, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, [32]byte](&proto.ColUInt32{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[uint32, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, [64]byte](&proto.ColUInt32{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[uint32, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, [128]byte](&proto.ColUInt32{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[uint32, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, [256]byte](&proto.ColUInt32{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[uint32, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt32ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[uint32, [512]byte](&proto.ColUInt32{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[uint32, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[int64, string](&proto.ColInt64{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[int64, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[int64, float32](&proto.ColInt64{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[int64, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[int64, float64](&proto.ColInt64{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[int64, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[int64, proto.IPv4](&proto.ColInt64{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[int64, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[int64, proto.IPv6](&proto.ColInt64{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[int64, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[int64, time.Time](&proto.ColInt64{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[int64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[int64, time.Time](&proto.ColInt64{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[int64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[int64, time.Time](&proto.ColInt64{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[int64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[int64, time.Time](&proto.ColInt64{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[int64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[int64, proto.Enum8](&proto.ColInt64{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[int64, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[int64, proto.Enum16](&proto.ColInt64{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[int64, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[int64, proto.Decimal32](&proto.ColInt64{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[int64, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[int64, proto.Decimal64](&proto.ColInt64{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[int64, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[int64, proto.Decimal128](&proto.ColInt64{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[int64, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[int64, proto.Decimal256](&proto.ColInt64{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[int64, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[int64, int8](&proto.ColInt64{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[int64, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[int64, uint8](&proto.ColInt64{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[int64, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[int64, int16](&proto.ColInt64{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[int64, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[int64, uint16](&proto.ColInt64{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[int64, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[int64, int32](&proto.ColInt64{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[int64, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[int64, uint32](&proto.ColInt64{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[int64, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[int64, int64](&proto.ColInt64{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[int64, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[int64, uint64](&proto.ColInt64{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[int64, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[int64, proto.Int128](&proto.ColInt64{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[int64, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[int64, proto.UInt128](&proto.ColInt64{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[int64, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[int64, proto.Int256](&proto.ColInt64{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[int64, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[int64, proto.UInt256](&proto.ColInt64{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[int64, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[int64, [8]byte](&proto.ColInt64{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[int64, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[int64, [16]byte](&proto.ColInt64{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[int64, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[int64, [32]byte](&proto.ColInt64{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[int64, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[int64, [64]byte](&proto.ColInt64{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[int64, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[int64, [128]byte](&proto.ColInt64{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[int64, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[int64, [256]byte](&proto.ColInt64{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[int64, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt64ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[int64, [512]byte](&proto.ColInt64{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[int64, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, string](&proto.ColUInt64{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[uint64, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, float32](&proto.ColUInt64{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[uint64, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, float64](&proto.ColUInt64{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[uint64, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, proto.IPv4](&proto.ColUInt64{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[uint64, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, proto.IPv6](&proto.ColUInt64{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[uint64, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, time.Time](&proto.ColUInt64{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[uint64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, time.Time](&proto.ColUInt64{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[uint64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, time.Time](&proto.ColUInt64{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[uint64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, time.Time](&proto.ColUInt64{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[uint64, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, proto.Enum8](&proto.ColUInt64{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[uint64, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, proto.Enum16](&proto.ColUInt64{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[uint64, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, proto.Decimal32](&proto.ColUInt64{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[uint64, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, proto.Decimal64](&proto.ColUInt64{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[uint64, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, proto.Decimal128](&proto.ColUInt64{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[uint64, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, proto.Decimal256](&proto.ColUInt64{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[uint64, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, int8](&proto.ColUInt64{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[uint64, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, uint8](&proto.ColUInt64{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[uint64, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, int16](&proto.ColUInt64{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[uint64, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, uint16](&proto.ColUInt64{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[uint64, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, int32](&proto.ColUInt64{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[uint64, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, uint32](&proto.ColUInt64{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[uint64, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, int64](&proto.ColUInt64{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[uint64, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, uint64](&proto.ColUInt64{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[uint64, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, proto.Int128](&proto.ColUInt64{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[uint64, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, proto.UInt128](&proto.ColUInt64{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[uint64, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, proto.Int256](&proto.ColUInt64{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[uint64, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, proto.UInt256](&proto.ColUInt64{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[uint64, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, [8]byte](&proto.ColUInt64{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[uint64, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, [16]byte](&proto.ColUInt64{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[uint64, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, [32]byte](&proto.ColUInt64{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[uint64, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, [64]byte](&proto.ColUInt64{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[uint64, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, [128]byte](&proto.ColUInt64{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[uint64, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, [256]byte](&proto.ColUInt64{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[uint64, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt64ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[uint64, [512]byte](&proto.ColUInt64{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[uint64, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, string](&proto.ColInt128{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[proto.Int128, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, float32](&proto.ColInt128{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[proto.Int128, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, float64](&proto.ColInt128{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[proto.Int128, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, proto.IPv4](&proto.ColInt128{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[proto.Int128, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, proto.IPv6](&proto.ColInt128{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[proto.Int128, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, time.Time](&proto.ColInt128{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[proto.Int128, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, time.Time](&proto.ColInt128{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[proto.Int128, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, time.Time](&proto.ColInt128{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[proto.Int128, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, time.Time](&proto.ColInt128{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[proto.Int128, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, proto.Enum8](&proto.ColInt128{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[proto.Int128, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, proto.Enum16](&proto.ColInt128{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[proto.Int128, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, proto.Decimal32](&proto.ColInt128{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[proto.Int128, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, proto.Decimal64](&proto.ColInt128{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[proto.Int128, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, proto.Decimal128](&proto.ColInt128{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[proto.Int128, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, proto.Decimal256](&proto.ColInt128{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[proto.Int128, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, int8](&proto.ColInt128{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[proto.Int128, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, uint8](&proto.ColInt128{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[proto.Int128, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, int16](&proto.ColInt128{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[proto.Int128, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, uint16](&proto.ColInt128{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[proto.Int128, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, int32](&proto.ColInt128{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[proto.Int128, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, uint32](&proto.ColInt128{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[proto.Int128, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, int64](&proto.ColInt128{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[proto.Int128, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, uint64](&proto.ColInt128{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[proto.Int128, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, proto.Int128](&proto.ColInt128{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[proto.Int128, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, proto.UInt128](&proto.ColInt128{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[proto.Int128, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, proto.Int256](&proto.ColInt128{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[proto.Int128, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, proto.UInt256](&proto.ColInt128{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[proto.Int128, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, [8]byte](&proto.ColInt128{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[proto.Int128, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, [16]byte](&proto.ColInt128{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[proto.Int128, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, [32]byte](&proto.ColInt128{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[proto.Int128, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, [64]byte](&proto.ColInt128{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[proto.Int128, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, [128]byte](&proto.ColInt128{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[proto.Int128, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, [256]byte](&proto.ColInt128{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[proto.Int128, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt128ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int128, [512]byte](&proto.ColInt128{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[proto.Int128, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, string](&proto.ColUInt128{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[proto.UInt128, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, float32](&proto.ColUInt128{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[proto.UInt128, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, float64](&proto.ColUInt128{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[proto.UInt128, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, proto.IPv4](&proto.ColUInt128{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[proto.UInt128, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, proto.IPv6](&proto.ColUInt128{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[proto.UInt128, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, time.Time](&proto.ColUInt128{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[proto.UInt128, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, time.Time](&proto.ColUInt128{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[proto.UInt128, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, time.Time](&proto.ColUInt128{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[proto.UInt128, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, time.Time](&proto.ColUInt128{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[proto.UInt128, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, proto.Enum8](&proto.ColUInt128{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[proto.UInt128, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, proto.Enum16](&proto.ColUInt128{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[proto.UInt128, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, proto.Decimal32](&proto.ColUInt128{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[proto.UInt128, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, proto.Decimal64](&proto.ColUInt128{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[proto.UInt128, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, proto.Decimal128](&proto.ColUInt128{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[proto.UInt128, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, proto.Decimal256](&proto.ColUInt128{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[proto.UInt128, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, int8](&proto.ColUInt128{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[proto.UInt128, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, uint8](&proto.ColUInt128{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[proto.UInt128, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, int16](&proto.ColUInt128{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[proto.UInt128, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, uint16](&proto.ColUInt128{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[proto.UInt128, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, int32](&proto.ColUInt128{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[proto.UInt128, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, uint32](&proto.ColUInt128{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[proto.UInt128, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, int64](&proto.ColUInt128{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[proto.UInt128, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, uint64](&proto.ColUInt128{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[proto.UInt128, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, proto.Int128](&proto.ColUInt128{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[proto.UInt128, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, proto.UInt128](&proto.ColUInt128{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[proto.UInt128, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, proto.Int256](&proto.ColUInt128{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[proto.UInt128, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, proto.UInt256](&proto.ColUInt128{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[proto.UInt128, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, [8]byte](&proto.ColUInt128{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[proto.UInt128, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, [16]byte](&proto.ColUInt128{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[proto.UInt128, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, [32]byte](&proto.ColUInt128{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[proto.UInt128, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, [64]byte](&proto.ColUInt128{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[proto.UInt128, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, [128]byte](&proto.ColUInt128{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[proto.UInt128, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, [256]byte](&proto.ColUInt128{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[proto.UInt128, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt128ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt128, [512]byte](&proto.ColUInt128{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[proto.UInt128, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, string](&proto.ColInt256{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[proto.Int256, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, float32](&proto.ColInt256{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[proto.Int256, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, float64](&proto.ColInt256{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[proto.Int256, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, proto.IPv4](&proto.ColInt256{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[proto.Int256, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, proto.IPv6](&proto.ColInt256{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[proto.Int256, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, time.Time](&proto.ColInt256{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[proto.Int256, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, time.Time](&proto.ColInt256{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[proto.Int256, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, time.Time](&proto.ColInt256{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[proto.Int256, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, time.Time](&proto.ColInt256{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[proto.Int256, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, proto.Enum8](&proto.ColInt256{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[proto.Int256, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, proto.Enum16](&proto.ColInt256{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[proto.Int256, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, proto.Decimal32](&proto.ColInt256{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[proto.Int256, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, proto.Decimal64](&proto.ColInt256{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[proto.Int256, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, proto.Decimal128](&proto.ColInt256{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[proto.Int256, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, proto.Decimal256](&proto.ColInt256{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[proto.Int256, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, int8](&proto.ColInt256{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[proto.Int256, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, uint8](&proto.ColInt256{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[proto.Int256, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, int16](&proto.ColInt256{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[proto.Int256, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, uint16](&proto.ColInt256{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[proto.Int256, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, int32](&proto.ColInt256{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[proto.Int256, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, uint32](&proto.ColInt256{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[proto.Int256, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, int64](&proto.ColInt256{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[proto.Int256, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, uint64](&proto.ColInt256{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[proto.Int256, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, proto.Int128](&proto.ColInt256{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[proto.Int256, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, proto.UInt128](&proto.ColInt256{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[proto.Int256, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, proto.Int256](&proto.ColInt256{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[proto.Int256, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, proto.UInt256](&proto.ColInt256{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[proto.Int256, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, [8]byte](&proto.ColInt256{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[proto.Int256, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, [16]byte](&proto.ColInt256{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[proto.Int256, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, [32]byte](&proto.ColInt256{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[proto.Int256, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, [64]byte](&proto.ColInt256{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[proto.Int256, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, [128]byte](&proto.ColInt256{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[proto.Int256, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, [256]byte](&proto.ColInt256{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[proto.Int256, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapInt256ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[proto.Int256, [512]byte](&proto.ColInt256{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[proto.Int256, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, string](&proto.ColUInt256{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[proto.UInt256, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, float32](&proto.ColUInt256{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[proto.UInt256, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, float64](&proto.ColUInt256{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[proto.UInt256, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, proto.IPv4](&proto.ColUInt256{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[proto.UInt256, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, proto.IPv6](&proto.ColUInt256{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[proto.UInt256, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, time.Time](&proto.ColUInt256{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[proto.UInt256, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, time.Time](&proto.ColUInt256{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[proto.UInt256, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, time.Time](&proto.ColUInt256{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[proto.UInt256, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, time.Time](&proto.ColUInt256{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[proto.UInt256, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, proto.Enum8](&proto.ColUInt256{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[proto.UInt256, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, proto.Enum16](&proto.ColUInt256{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[proto.UInt256, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, proto.Decimal32](&proto.ColUInt256{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[proto.UInt256, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, proto.Decimal64](&proto.ColUInt256{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[proto.UInt256, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, proto.Decimal128](&proto.ColUInt256{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[proto.UInt256, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, proto.Decimal256](&proto.ColUInt256{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[proto.UInt256, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, int8](&proto.ColUInt256{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[proto.UInt256, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, uint8](&proto.ColUInt256{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[proto.UInt256, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, int16](&proto.ColUInt256{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[proto.UInt256, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, uint16](&proto.ColUInt256{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[proto.UInt256, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, int32](&proto.ColUInt256{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[proto.UInt256, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, uint32](&proto.ColUInt256{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[proto.UInt256, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, int64](&proto.ColUInt256{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[proto.UInt256, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, uint64](&proto.ColUInt256{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[proto.UInt256, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, proto.Int128](&proto.ColUInt256{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[proto.UInt256, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, proto.UInt128](&proto.ColUInt256{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[proto.UInt256, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, proto.Int256](&proto.ColUInt256{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[proto.UInt256, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, proto.UInt256](&proto.ColUInt256{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[proto.UInt256, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, [8]byte](&proto.ColUInt256{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[proto.UInt256, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, [16]byte](&proto.ColUInt256{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[proto.UInt256, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, [32]byte](&proto.ColUInt256{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[proto.UInt256, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, [64]byte](&proto.ColUInt256{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[proto.UInt256, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, [128]byte](&proto.ColUInt256{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[proto.UInt256, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, [256]byte](&proto.ColUInt256{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[proto.UInt256, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapUInt256ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[proto.UInt256, [512]byte](&proto.ColUInt256{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[proto.UInt256, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, string](&proto.ColFixedStr8{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[[8]byte, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, float32](&proto.ColFixedStr8{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[[8]byte, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, float64](&proto.ColFixedStr8{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[[8]byte, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, proto.IPv4](&proto.ColFixedStr8{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[[8]byte, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, proto.IPv6](&proto.ColFixedStr8{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[[8]byte, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, time.Time](&proto.ColFixedStr8{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[[8]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, time.Time](&proto.ColFixedStr8{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[[8]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, time.Time](&proto.ColFixedStr8{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[[8]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, time.Time](&proto.ColFixedStr8{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[[8]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, proto.Enum8](&proto.ColFixedStr8{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[[8]byte, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, proto.Enum16](&proto.ColFixedStr8{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[[8]byte, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, proto.Decimal32](&proto.ColFixedStr8{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[[8]byte, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, proto.Decimal64](&proto.ColFixedStr8{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[[8]byte, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, proto.Decimal128](&proto.ColFixedStr8{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[[8]byte, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, proto.Decimal256](&proto.ColFixedStr8{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[[8]byte, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, int8](&proto.ColFixedStr8{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[[8]byte, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, uint8](&proto.ColFixedStr8{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[[8]byte, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, int16](&proto.ColFixedStr8{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[[8]byte, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, uint16](&proto.ColFixedStr8{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[[8]byte, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, int32](&proto.ColFixedStr8{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[[8]byte, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, uint32](&proto.ColFixedStr8{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[[8]byte, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, int64](&proto.ColFixedStr8{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[[8]byte, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, uint64](&proto.ColFixedStr8{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[[8]byte, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, proto.Int128](&proto.ColFixedStr8{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[[8]byte, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, proto.UInt128](&proto.ColFixedStr8{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[[8]byte, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, proto.Int256](&proto.ColFixedStr8{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[[8]byte, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, proto.UInt256](&proto.ColFixedStr8{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[[8]byte, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, [8]byte](&proto.ColFixedStr8{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[[8]byte, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, [16]byte](&proto.ColFixedStr8{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[[8]byte, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, [32]byte](&proto.ColFixedStr8{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[[8]byte, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, [64]byte](&proto.ColFixedStr8{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[[8]byte, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, [128]byte](&proto.ColFixedStr8{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[[8]byte, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, [256]byte](&proto.ColFixedStr8{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[[8]byte, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr8ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[[8]byte, [512]byte](&proto.ColFixedStr8{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[[8]byte, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, string](&proto.ColFixedStr16{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[[16]byte, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, float32](&proto.ColFixedStr16{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[[16]byte, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, float64](&proto.ColFixedStr16{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[[16]byte, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, proto.IPv4](&proto.ColFixedStr16{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[[16]byte, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, proto.IPv6](&proto.ColFixedStr16{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[[16]byte, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, time.Time](&proto.ColFixedStr16{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[[16]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, time.Time](&proto.ColFixedStr16{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[[16]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, time.Time](&proto.ColFixedStr16{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[[16]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, time.Time](&proto.ColFixedStr16{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[[16]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, proto.Enum8](&proto.ColFixedStr16{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[[16]byte, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, proto.Enum16](&proto.ColFixedStr16{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[[16]byte, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, proto.Decimal32](&proto.ColFixedStr16{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[[16]byte, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, proto.Decimal64](&proto.ColFixedStr16{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[[16]byte, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, proto.Decimal128](&proto.ColFixedStr16{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[[16]byte, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, proto.Decimal256](&proto.ColFixedStr16{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[[16]byte, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, int8](&proto.ColFixedStr16{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[[16]byte, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, uint8](&proto.ColFixedStr16{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[[16]byte, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, int16](&proto.ColFixedStr16{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[[16]byte, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, uint16](&proto.ColFixedStr16{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[[16]byte, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, int32](&proto.ColFixedStr16{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[[16]byte, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, uint32](&proto.ColFixedStr16{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[[16]byte, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, int64](&proto.ColFixedStr16{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[[16]byte, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, uint64](&proto.ColFixedStr16{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[[16]byte, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, proto.Int128](&proto.ColFixedStr16{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[[16]byte, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, proto.UInt128](&proto.ColFixedStr16{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[[16]byte, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, proto.Int256](&proto.ColFixedStr16{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[[16]byte, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, proto.UInt256](&proto.ColFixedStr16{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[[16]byte, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, [8]byte](&proto.ColFixedStr16{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[[16]byte, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, [16]byte](&proto.ColFixedStr16{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[[16]byte, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, [32]byte](&proto.ColFixedStr16{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[[16]byte, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, [64]byte](&proto.ColFixedStr16{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[[16]byte, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, [128]byte](&proto.ColFixedStr16{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[[16]byte, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, [256]byte](&proto.ColFixedStr16{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[[16]byte, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr16ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[[16]byte, [512]byte](&proto.ColFixedStr16{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[[16]byte, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, string](&proto.ColFixedStr32{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[[32]byte, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, float32](&proto.ColFixedStr32{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[[32]byte, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, float64](&proto.ColFixedStr32{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[[32]byte, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, proto.IPv4](&proto.ColFixedStr32{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[[32]byte, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, proto.IPv6](&proto.ColFixedStr32{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[[32]byte, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, time.Time](&proto.ColFixedStr32{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[[32]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, time.Time](&proto.ColFixedStr32{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[[32]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, time.Time](&proto.ColFixedStr32{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[[32]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, time.Time](&proto.ColFixedStr32{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[[32]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, proto.Enum8](&proto.ColFixedStr32{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[[32]byte, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, proto.Enum16](&proto.ColFixedStr32{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[[32]byte, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, proto.Decimal32](&proto.ColFixedStr32{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[[32]byte, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, proto.Decimal64](&proto.ColFixedStr32{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[[32]byte, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, proto.Decimal128](&proto.ColFixedStr32{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[[32]byte, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, proto.Decimal256](&proto.ColFixedStr32{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[[32]byte, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, int8](&proto.ColFixedStr32{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[[32]byte, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, uint8](&proto.ColFixedStr32{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[[32]byte, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, int16](&proto.ColFixedStr32{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[[32]byte, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, uint16](&proto.ColFixedStr32{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[[32]byte, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, int32](&proto.ColFixedStr32{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[[32]byte, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, uint32](&proto.ColFixedStr32{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[[32]byte, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, int64](&proto.ColFixedStr32{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[[32]byte, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, uint64](&proto.ColFixedStr32{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[[32]byte, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, proto.Int128](&proto.ColFixedStr32{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[[32]byte, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, proto.UInt128](&proto.ColFixedStr32{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[[32]byte, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, proto.Int256](&proto.ColFixedStr32{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[[32]byte, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, proto.UInt256](&proto.ColFixedStr32{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[[32]byte, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, [8]byte](&proto.ColFixedStr32{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[[32]byte, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, [16]byte](&proto.ColFixedStr32{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[[32]byte, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, [32]byte](&proto.ColFixedStr32{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[[32]byte, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, [64]byte](&proto.ColFixedStr32{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[[32]byte, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, [128]byte](&proto.ColFixedStr32{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[[32]byte, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, [256]byte](&proto.ColFixedStr32{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[[32]byte, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr32ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[[32]byte, [512]byte](&proto.ColFixedStr32{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[[32]byte, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, string](&proto.ColFixedStr64{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[[64]byte, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, float32](&proto.ColFixedStr64{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[[64]byte, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, float64](&proto.ColFixedStr64{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[[64]byte, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, proto.IPv4](&proto.ColFixedStr64{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[[64]byte, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, proto.IPv6](&proto.ColFixedStr64{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[[64]byte, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, time.Time](&proto.ColFixedStr64{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[[64]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, time.Time](&proto.ColFixedStr64{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[[64]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, time.Time](&proto.ColFixedStr64{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[[64]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, time.Time](&proto.ColFixedStr64{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[[64]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, proto.Enum8](&proto.ColFixedStr64{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[[64]byte, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, proto.Enum16](&proto.ColFixedStr64{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[[64]byte, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, proto.Decimal32](&proto.ColFixedStr64{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[[64]byte, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, proto.Decimal64](&proto.ColFixedStr64{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[[64]byte, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, proto.Decimal128](&proto.ColFixedStr64{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[[64]byte, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, proto.Decimal256](&proto.ColFixedStr64{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[[64]byte, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, int8](&proto.ColFixedStr64{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[[64]byte, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, uint8](&proto.ColFixedStr64{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[[64]byte, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, int16](&proto.ColFixedStr64{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[[64]byte, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, uint16](&proto.ColFixedStr64{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[[64]byte, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, int32](&proto.ColFixedStr64{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[[64]byte, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, uint32](&proto.ColFixedStr64{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[[64]byte, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, int64](&proto.ColFixedStr64{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[[64]byte, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, uint64](&proto.ColFixedStr64{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[[64]byte, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, proto.Int128](&proto.ColFixedStr64{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[[64]byte, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, proto.UInt128](&proto.ColFixedStr64{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[[64]byte, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, proto.Int256](&proto.ColFixedStr64{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[[64]byte, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, proto.UInt256](&proto.ColFixedStr64{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[[64]byte, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, [8]byte](&proto.ColFixedStr64{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[[64]byte, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, [16]byte](&proto.ColFixedStr64{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[[64]byte, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, [32]byte](&proto.ColFixedStr64{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[[64]byte, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, [64]byte](&proto.ColFixedStr64{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[[64]byte, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, [128]byte](&proto.ColFixedStr64{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[[64]byte, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, [256]byte](&proto.ColFixedStr64{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[[64]byte, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr64ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[[64]byte, [512]byte](&proto.ColFixedStr64{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[[64]byte, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, string](&proto.ColFixedStr128{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[[128]byte, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, float32](&proto.ColFixedStr128{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[[128]byte, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, float64](&proto.ColFixedStr128{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[[128]byte, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, proto.IPv4](&proto.ColFixedStr128{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[[128]byte, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, proto.IPv6](&proto.ColFixedStr128{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[[128]byte, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, time.Time](&proto.ColFixedStr128{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[[128]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, time.Time](&proto.ColFixedStr128{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[[128]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, time.Time](&proto.ColFixedStr128{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[[128]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, time.Time](&proto.ColFixedStr128{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[[128]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, proto.Enum8](&proto.ColFixedStr128{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[[128]byte, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, proto.Enum16](&proto.ColFixedStr128{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[[128]byte, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, proto.Decimal32](&proto.ColFixedStr128{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[[128]byte, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, proto.Decimal64](&proto.ColFixedStr128{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[[128]byte, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, proto.Decimal128](&proto.ColFixedStr128{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[[128]byte, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, proto.Decimal256](&proto.ColFixedStr128{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[[128]byte, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, int8](&proto.ColFixedStr128{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[[128]byte, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, uint8](&proto.ColFixedStr128{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[[128]byte, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, int16](&proto.ColFixedStr128{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[[128]byte, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, uint16](&proto.ColFixedStr128{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[[128]byte, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, int32](&proto.ColFixedStr128{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[[128]byte, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, uint32](&proto.ColFixedStr128{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[[128]byte, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, int64](&proto.ColFixedStr128{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[[128]byte, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, uint64](&proto.ColFixedStr128{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[[128]byte, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, proto.Int128](&proto.ColFixedStr128{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[[128]byte, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, proto.UInt128](&proto.ColFixedStr128{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[[128]byte, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, proto.Int256](&proto.ColFixedStr128{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[[128]byte, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, proto.UInt256](&proto.ColFixedStr128{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[[128]byte, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, [8]byte](&proto.ColFixedStr128{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[[128]byte, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, [16]byte](&proto.ColFixedStr128{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[[128]byte, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, [32]byte](&proto.ColFixedStr128{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[[128]byte, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, [64]byte](&proto.ColFixedStr128{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[[128]byte, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, [128]byte](&proto.ColFixedStr128{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[[128]byte, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, [256]byte](&proto.ColFixedStr128{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[[128]byte, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr128ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[[128]byte, [512]byte](&proto.ColFixedStr128{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[[128]byte, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, string](&proto.ColFixedStr256{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[[256]byte, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, float32](&proto.ColFixedStr256{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[[256]byte, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, float64](&proto.ColFixedStr256{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[[256]byte, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, proto.IPv4](&proto.ColFixedStr256{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[[256]byte, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, proto.IPv6](&proto.ColFixedStr256{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[[256]byte, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, time.Time](&proto.ColFixedStr256{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[[256]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, time.Time](&proto.ColFixedStr256{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[[256]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, time.Time](&proto.ColFixedStr256{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[[256]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, time.Time](&proto.ColFixedStr256{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[[256]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, proto.Enum8](&proto.ColFixedStr256{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[[256]byte, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, proto.Enum16](&proto.ColFixedStr256{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[[256]byte, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, proto.Decimal32](&proto.ColFixedStr256{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[[256]byte, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, proto.Decimal64](&proto.ColFixedStr256{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[[256]byte, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, proto.Decimal128](&proto.ColFixedStr256{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[[256]byte, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, proto.Decimal256](&proto.ColFixedStr256{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[[256]byte, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, int8](&proto.ColFixedStr256{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[[256]byte, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, uint8](&proto.ColFixedStr256{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[[256]byte, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, int16](&proto.ColFixedStr256{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[[256]byte, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, uint16](&proto.ColFixedStr256{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[[256]byte, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, int32](&proto.ColFixedStr256{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[[256]byte, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, uint32](&proto.ColFixedStr256{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[[256]byte, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, int64](&proto.ColFixedStr256{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[[256]byte, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, uint64](&proto.ColFixedStr256{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[[256]byte, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, proto.Int128](&proto.ColFixedStr256{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[[256]byte, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, proto.UInt128](&proto.ColFixedStr256{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[[256]byte, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, proto.Int256](&proto.ColFixedStr256{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[[256]byte, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, proto.UInt256](&proto.ColFixedStr256{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[[256]byte, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, [8]byte](&proto.ColFixedStr256{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[[256]byte, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, [16]byte](&proto.ColFixedStr256{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[[256]byte, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, [32]byte](&proto.ColFixedStr256{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[[256]byte, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, [64]byte](&proto.ColFixedStr256{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[[256]byte, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, [128]byte](&proto.ColFixedStr256{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[[256]byte, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, [256]byte](&proto.ColFixedStr256{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[[256]byte, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr256ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[[256]byte, [512]byte](&proto.ColFixedStr256{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[[256]byte, [512]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToStr(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, string](&proto.ColFixedStr512{}, &proto.ColStr{})
	r.Bind(name, col, &BindingMap[[512]byte, string]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, float32](&proto.ColFixedStr512{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[[512]byte, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToFloat64(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, float64](&proto.ColFixedStr512{}, &proto.ColFloat64{})
	r.Bind(name, col, &BindingMap[[512]byte, float64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToIPv4(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, proto.IPv4](&proto.ColFixedStr512{}, &proto.ColIPv4{})
	r.Bind(name, col, &BindingMap[[512]byte, proto.IPv4]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToIPv6(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, proto.IPv6](&proto.ColFixedStr512{}, &proto.ColIPv6{})
	r.Bind(name, col, &BindingMap[[512]byte, proto.IPv6]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToDateTime(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, time.Time](&proto.ColFixedStr512{}, &proto.ColDateTime{})
	r.Bind(name, col, &BindingMap[[512]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToDateTime64(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, time.Time](&proto.ColFixedStr512{}, &proto.ColDateTime64{})
	r.Bind(name, col, &BindingMap[[512]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToDate(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, time.Time](&proto.ColFixedStr512{}, &proto.ColDate{})
	r.Bind(name, col, &BindingMap[[512]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToDate32(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, time.Time](&proto.ColFixedStr512{}, &proto.ColDate32{})
	r.Bind(name, col, &BindingMap[[512]byte, time.Time]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToEnum8(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, proto.Enum8](&proto.ColFixedStr512{}, &proto.ColEnum8{})
	r.Bind(name, col, &BindingMap[[512]byte, proto.Enum8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToEnum16(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, proto.Enum16](&proto.ColFixedStr512{}, &proto.ColEnum16{})
	r.Bind(name, col, &BindingMap[[512]byte, proto.Enum16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToDecimal32(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, proto.Decimal32](&proto.ColFixedStr512{}, &proto.ColDecimal32{})
	r.Bind(name, col, &BindingMap[[512]byte, proto.Decimal32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToDecimal64(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, proto.Decimal64](&proto.ColFixedStr512{}, &proto.ColDecimal64{})
	r.Bind(name, col, &BindingMap[[512]byte, proto.Decimal64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToDecimal128(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, proto.Decimal128](&proto.ColFixedStr512{}, &proto.ColDecimal128{})
	r.Bind(name, col, &BindingMap[[512]byte, proto.Decimal128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToDecimal256(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, proto.Decimal256](&proto.ColFixedStr512{}, &proto.ColDecimal256{})
	r.Bind(name, col, &BindingMap[[512]byte, proto.Decimal256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, int8](&proto.ColFixedStr512{}, &proto.ColInt8{})
	r.Bind(name, col, &BindingMap[[512]byte, int8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToUInt8(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, uint8](&proto.ColFixedStr512{}, &proto.ColUInt8{})
	r.Bind(name, col, &BindingMap[[512]byte, uint8]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, int16](&proto.ColFixedStr512{}, &proto.ColInt16{})
	r.Bind(name, col, &BindingMap[[512]byte, int16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToUInt16(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, uint16](&proto.ColFixedStr512{}, &proto.ColUInt16{})
	r.Bind(name, col, &BindingMap[[512]byte, uint16]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, int32](&proto.ColFixedStr512{}, &proto.ColInt32{})
	r.Bind(name, col, &BindingMap[[512]byte, int32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToUInt32(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, uint32](&proto.ColFixedStr512{}, &proto.ColUInt32{})
	r.Bind(name, col, &BindingMap[[512]byte, uint32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, int64](&proto.ColFixedStr512{}, &proto.ColInt64{})
	r.Bind(name, col, &BindingMap[[512]byte, int64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToUInt64(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, uint64](&proto.ColFixedStr512{}, &proto.ColUInt64{})
	r.Bind(name, col, &BindingMap[[512]byte, uint64]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, proto.Int128](&proto.ColFixedStr512{}, &proto.ColInt128{})
	r.Bind(name, col, &BindingMap[[512]byte, proto.Int128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToUInt128(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, proto.UInt128](&proto.ColFixedStr512{}, &proto.ColUInt128{})
	r.Bind(name, col, &BindingMap[[512]byte, proto.UInt128]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, proto.Int256](&proto.ColFixedStr512{}, &proto.ColInt256{})
	r.Bind(name, col, &BindingMap[[512]byte, proto.Int256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToUInt256(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, proto.UInt256](&proto.ColFixedStr512{}, &proto.ColUInt256{})
	r.Bind(name, col, &BindingMap[[512]byte, proto.UInt256]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToFixedStr8(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, [8]byte](&proto.ColFixedStr512{}, &proto.ColFixedStr8{})
	r.Bind(name, col, &BindingMap[[512]byte, [8]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToFixedStr16(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, [16]byte](&proto.ColFixedStr512{}, &proto.ColFixedStr16{})
	r.Bind(name, col, &BindingMap[[512]byte, [16]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToFixedStr32(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, [32]byte](&proto.ColFixedStr512{}, &proto.ColFixedStr32{})
	r.Bind(name, col, &BindingMap[[512]byte, [32]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToFixedStr64(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, [64]byte](&proto.ColFixedStr512{}, &proto.ColFixedStr64{})
	r.Bind(name, col, &BindingMap[[512]byte, [64]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToFixedStr128(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, [128]byte](&proto.ColFixedStr512{}, &proto.ColFixedStr128{})
	r.Bind(name, col, &BindingMap[[512]byte, [128]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToFixedStr256(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, [256]byte](&proto.ColFixedStr512{}, &proto.ColFixedStr256{})
	r.Bind(name, col, &BindingMap[[512]byte, [256]byte]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindMapFixedStr512ToFixedStr512(name string, valPtr interface{}) {
	col := proto.NewMap[[512]byte, [512]byte](&proto.ColFixedStr512{}, &proto.ColFixedStr512{})
	r.Bind(name, col, &BindingMap[[512]byte, [512]byte]{col: col, valPtr: valPtr})
}
