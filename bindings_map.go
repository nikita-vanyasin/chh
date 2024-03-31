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

func (r *ColResults) BindMapStringToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[string, float32](&proto.ColStr{}, &proto.ColFloat32{})
	r.Bind(name, col, &BindingMap[string, float32]{col: col, valPtr: valPtr})
}
