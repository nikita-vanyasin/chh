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
