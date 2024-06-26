package chh

import (
	"github.com/ClickHouse/ch-go/proto"
)

// -- type Map[K]V

type BindingMap[K comparable, V any] struct {
	col    *proto.ColMap[K, V]
	valPtr interface{}
}

func (b *BindingMap[K, V]) Append() {
	s, _ := b.valPtr.(*[]map[K]V)
	appendToSlice(s, b.col)
}

func (b *BindingMap[K, V]) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*map[K]V)
		*s = b.col.Row(0)
	}
}
