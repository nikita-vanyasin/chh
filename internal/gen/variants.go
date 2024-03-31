package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// File is based on https://github.com/ClickHouse/ch-go/blob/main/proto/cmd/ch-gen-col/main.go

type Kind byte

const (
	KindInt Kind = iota
	KindFloat
	KindIP
	KindDateTime
	KindDate
	KindEnum
	KindDecimal
	KindFixedStr
	KindStr
)

type Variant struct {
	GoType         string
	Kind           Kind
	Signed         bool
	Bits           int
	GenerateUnsafe bool
}

func (v Variant) Bytes() int {
	return v.Bits / 8
}

type Variants []Variant

func (v Variant) IsFloat() bool {
	return v.Kind == KindFloat
}

func (v Variant) IsInt() bool {
	return v.Kind == KindInt
}

func (v Variant) SingleByte() bool {
	return v.Bits == 8
}

func (v Variant) Byte() bool {
	return v.Bits/8 == 1 && !v.Signed && v.IsInt()
}

func (v Variant) Type() string {
	return "Col" + v.Name()
}

func (v Variant) ColumnType() string {
	if v.Kind == KindFixedStr {
		return fmt.Sprintf(`ColumnTypeFixedString.With("%d")`, v.Bytes())
	}
	return "ColumnType" + v.Name()
}

func (v Variant) New() string {
	if v.Kind == KindFixedStr {
		return "newByte" + strconv.Itoa(v.Bytes())
	}
	if v.Big() {
		return v.ElemType() + "FromInt"
	}
	return v.ElemType()
}

func (v Variant) GoName() string {
	if v.GoType != "" {
		return v.GoType
	}

	return fmt.Sprintf("proto.%s", v.Name())
}

func (v Variant) Name() string {
	if v.Kind == KindStr {
		return "Str"
	}
	if v.Kind == KindFixedStr {
		return "FixedStr" + strconv.Itoa(v.Bytes())
	}
	if v.Kind != KindInt && v.Kind != KindFloat {
		return v.ElemType()
	}
	var b strings.Builder
	if !v.Signed {
		b.WriteString("U")
	}
	switch v.Kind {
	case KindFloat:
		b.WriteString("Float")
	case KindInt:
		b.WriteString("Int")
	}
	b.WriteString(strconv.Itoa(v.Bits))
	return b.String()
}

func (v Variant) BinFunc() string {
	return fmt.Sprintf("Uint%d", v.Bits)
}

func (v Variant) BinGet() string {
	if v.IPv6() {
		return "binIPv6"
	}
	if v.Big() {
		return fmt.Sprintf("binUInt%d", v.Bits)
	}
	return "binary.LittleEndian." + v.BinFunc()
}

func (v Variant) IsIP() bool {
	return v.Kind == KindIP
}

func (v Variant) IPv6() bool {
	return v.IsIP() && v.Bits == 128
}

func (v Variant) IPv4() bool {
	return v.IsIP() && v.Bits == 32
}

func (v Variant) BinPut() string {
	if v.FixedStr() {
		return "copy"
	}
	if v.IPv6() {
		return "binPutIPv6"
	}
	if v.Big() {
		return fmt.Sprintf("binPutUInt%d", v.Bits)
	}
	return "binary.LittleEndian.Put" + v.BinFunc()
}

func (v Variant) Big() bool {
	return v.Bits > 64
}

func (v Variant) Cast() bool {
	return v.Signed || v.IPv4()
}

func (v Variant) UnsignedType() string {
	var b strings.Builder
	if v.Big() {
		b.WriteString("UInt")
	} else {
		b.WriteString("uint")
	}
	b.WriteString(strconv.Itoa(v.Bits))
	return b.String()
}

func (v Variant) ElemLower() string {
	if v.Kind == KindFixedStr {
		return "byte" + strconv.Itoa(v.Bytes())
	}
	return strings.ToLower(v.ElemType())
}

func (v Variant) Complex() bool {
	return v.Time()
}

func (v Variant) Time() bool {
	switch v.Kind {
	case KindDate, KindDateTime:
		return true
	default:
		return false
	}
}

func (v Variant) Date() bool {
	return v.Kind == KindDate
}

func (v Variant) DateTime() bool {
	return v.Kind == KindDateTime
}

func (v Variant) FixedStr() bool {
	return v.Kind == KindFixedStr
}

func (v Variant) ElemType() string {
	if v.Kind == KindFixedStr {
		return fmt.Sprintf("[%d]byte", v.Bytes())
	}
	if v.Kind == KindEnum {
		return fmt.Sprintf("Enum%d", v.Bits)
	}
	if v.IPv4() {
		return "IPv4"
	}
	if v.IPv6() {
		return "IPv6"
	}
	if v.Kind == KindDecimal {
		return fmt.Sprintf("Decimal%d", v.Bits)
	}
	if v.Kind == KindDateTime {
		if v.Bits == 64 {
			return "DateTime64"
		}
		return "DateTime"
	}
	if v.Kind == KindDate {
		if v.Bits == 32 {
			return "Date32"
		}
		return "Date"
	}
	var b strings.Builder
	var (
		unsigned = "u"
		integer  = "int"
		float    = "float"
	)
	if v.Big() {
		unsigned = "U"
		integer = "Int"
	}
	if !v.Signed {
		b.WriteString(unsigned)
	}
	if v.IsFloat() {
		b.WriteString(float)
	} else {
		b.WriteString(integer)
	}
	b.WriteString(strconv.Itoa(v.Bits))
	return b.String()
}

func getVariants() Variants {
	variants := Variants{
		{
			// String
			GoType: "string",
			Kind:   KindStr,
		},
		{ // Float32
			GoType: "float32",
			Bits:   32,
			Kind:   KindFloat,
			Signed: true,
		},
		{ // Float64
			GoType: "float64",
			Bits:   64,
			Kind:   KindFloat,
			Signed: true,
		},
		{ // IPv4
			Bits: 32,
			Kind: KindIP,
		},
		{ // IPv6
			Bits: 128,
			Kind: KindIP,
		},
		{ // DateTime
			GoType: "time.Time",
			Bits:   32,
			Signed: true,
			Kind:   KindDateTime,
		},
		{ // DateTime64
			GoType: "time.Time",
			Bits:   64,
			Signed: true,
			Kind:   KindDateTime,
		},
		{ // Date
			GoType: "time.Time",
			Bits:   16,
			Signed: true,
			Kind:   KindDate,
		},
		{ // Date32
			GoType: "time.Time",
			Bits:   32,
			Signed: true,
			Kind:   KindDate,
		},
		{ // Enum8
			Bits:   8,
			Signed: true,
			Kind:   KindEnum,
		},
		{ // Enum16
			Bits:   16,
			Signed: true,
			Kind:   KindEnum,
		},
		{ // Decimal32
			Bits:   32,
			Signed: true,
			Kind:   KindDecimal,
		},
		{ // Decimal64
			Bits:   64,
			Signed: true,
			Kind:   KindDecimal,
		},
		{ // Decimal128
			Bits:   128,
			Signed: true,
			Kind:   KindDecimal,
		},
		{ // Decimal256
			Bits:   256,
			Signed: true,
			Kind:   KindDecimal,
		},
	}
	for _, bits := range []int{
		8,
		16,
		32,
		64,
		128,
		256,
	} {
		for _, signed := range []bool{true, false} {
			variant := Variant{
				Kind:   KindInt,
				Bits:   bits,
				Signed: signed,
			}
			if bits < 128 {
				goType := fmt.Sprintf("int%d", bits)
				if !signed {
					goType = "u" + goType
				}

				variant.GoType = goType
			}

			variants = append(variants, variant)
		}
	}
	for _, b := range []int{
		8,
		16,
		32,
		64,
		128,
		256,
		512,
	} {
		variant := Variant{
			Kind: KindFixedStr,
			Bits: b * 8,
		}
		variant.GoType = variant.ElemType()
		variants = append(variants, variant)
	}
	return variants
}

type MapVariant struct {
	key Variant
	val Variant
}

func (m MapVariant) Key() string {
	return m.key.Name()
}

func (m MapVariant) Value() string {
	return m.val.Name()
}

func (m MapVariant) GoKey() string {
	return m.key.GoName()
}

func (m MapVariant) GoValue() string {
	return m.val.GoName()
}

func getMapVariants() []MapVariant {
	// all combinations:
	base := getVariants()
	variants := make([]MapVariant, 0, len(base)*len(base))
	for _, key := range base {
		for _, val := range base {
			variants = append(variants, MapVariant{
				key: key,
				val: val,
			})
		}
	}
	return variants
}
