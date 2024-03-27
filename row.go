package chh

import (
	"context"
	"fmt"

	"github.com/ClickHouse/ch-go/proto"
)

// OnResultHandlerRow expects the columns results to be bind to slices representing columns.
func OnResultHandlerRow(cols *ColResults) func(ctx context.Context, block proto.Block) error {
	return func(ctx context.Context, block proto.Block) error {
		for _, resultCol := range cols.results {
			binding, ok := cols.bindings[resultCol.Name]
			if !ok {
				return fmt.Errorf("result column %s does not have binding", resultCol.Name)
			}

			binding.Write()
			resultCol.Data.Reset()
		}
		return nil
	}
}
