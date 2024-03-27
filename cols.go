package chh

import (
	"context"
	"fmt"

	"github.com/ClickHouse/ch-go/proto"
)

// OnResultHandler expects the columns results to be bind to slices representing columns.
func OnResultHandler(cols *ColResults) func(ctx context.Context, block proto.Block) error {
	return func(ctx context.Context, block proto.Block) error {
		for _, resultCol := range cols.results {
			binding, ok := cols.bindings[resultCol.Name]
			if !ok {
				return fmt.Errorf("result column %s does not have binding", resultCol.Name)
			}

			binding.Append()
			resultCol.Data.Reset()
		}
		return nil
	}
}

func appendToSlice[T any](slice *[]T, col proto.ColumnOf[T]) {
	for i := 0; i < col.Rows(); i++ {
		*slice = append(*slice, col.Row(i))
	}
}
