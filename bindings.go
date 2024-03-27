package chh

import (
	"time"

	"github.com/ClickHouse/ch-go/proto"
)

type bindingMap[K comparable, V any] struct {
	col    *proto.ColMap[K, V]
	valPtr interface{}
}

func (b *bindingMap[K, V]) Append() {
	s, _ := b.valPtr.(*[]map[K]V)
	appendToSlice(s, b.col)
}

func (b *bindingMap[K, V]) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*map[K]V)
		*s = b.col.Row(0)
	}
}

// ---

type bindingString struct {
	col    *proto.ColStr
	valPtr interface{}
}

func (b *bindingString) Append() {
	s, _ := b.valPtr.(*[]string)
	appendToSlice(s, b.col)
}

func (b *bindingString) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*string)
		*s = b.col.Row(0)
	}
}

// ---

type bindingDate struct {
	col    *proto.ColDate
	valPtr interface{}
}

func (b *bindingDate) Append() {
	s, _ := b.valPtr.(*[]time.Time)
	appendToSlice(s, b.col)
}

func (b *bindingDate) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*time.Time)
		*s = b.col.Row(0)
	}
}

// ----

type bindingUInt64 struct {
	col    *proto.ColUInt64
	valPtr interface{}
}

func (b *bindingUInt64) Append() {
	s, _ := b.valPtr.(*[]uint64)
	appendToSlice(s, b.col)
}

func (b *bindingUInt64) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*uint64)
		*s = b.col.Row(0)
	}
}

// ---

type bindingFloat64 struct {
	col    *proto.ColFloat64
	valPtr interface{}
}

func (b *bindingFloat64) Append() {
	s, _ := b.valPtr.(*[]float64)
	appendToSlice(s, b.col)
}

func (b *bindingFloat64) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*float64)
		*s = b.col.Row(0)
	}
}
