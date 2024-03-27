package chh

import (
	"github.com/ClickHouse/ch-go/proto"
)

type ColumnBinding interface {
	// Append adds chunk column values to valPtr
	Append()
	// Write overwrites pointer value with value from current chunk
	Write()
}

type ColResults struct {
	bindings map[string]ColumnBinding
	results  proto.Results
}

func (r *ColResults) init() {
	if r.bindings == nil {
		r.bindings = make(map[string]ColumnBinding)
	}
}

func (r *ColResults) Proto() proto.Result {
	return r.results
}

func (r *ColResults) Bind(name string, col proto.ColResult, binding ColumnBinding) {
	r.init()

	r.results = append(r.results, proto.ResultColumn{Name: name, Data: col})
	r.bindings[name] = binding
}

func (r *ColResults) BindMapStringToFloat32(name string, valPtr interface{}) {
	col := proto.NewMap[string, float32](&proto.ColStr{}, &proto.ColFloat32{})
	r.Bind(name, col, &bindingMap[string, float32]{col: col, valPtr: valPtr})
}

func (r *ColResults) BindString(name string, valPtr interface{}) {
	col := &proto.ColStr{}
	r.Bind(name, col, &bindingString{col: col, valPtr: valPtr})
}

func (r *ColResults) BindDate(name string, valPtr interface{}) {
	col := &proto.ColDate{}
	r.Bind(name, col, &bindingDate{col: col, valPtr: valPtr})
}

func (r *ColResults) BindUInt64(name string, valPtr interface{}) {
	col := &proto.ColUInt64{}
	r.Bind(name, col, &bindingUInt64{col: col, valPtr: valPtr})
}

func (r *ColResults) BindFloat64(name string, valPtr interface{}) {
	col := &proto.ColFloat64{}
	r.Bind(name, col, &bindingFloat64{col: col, valPtr: valPtr})
}
