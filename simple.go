package chh

import (
	"context"

	"github.com/ClickHouse/ch-go"
)

//go:generate go run ./internal/gen

type Doer interface {
	Do(ctx context.Context, q ch.Query) (err error)
}

func Select(ctx context.Context, conn Doer, results *ColResults, q ch.Query) error {
	q.Result = results.Proto()
	q.OnResult = OnResultHandler(results)

	return conn.Do(ctx, q)
}

func SelectRow(ctx context.Context, conn Doer, results *ColResults, q ch.Query) error {
	q.Result = results.Proto()
	q.OnResult = OnResultHandlerRow(results)

	return conn.Do(ctx, q)
}

func Do(ctx context.Context, conn Doer, q ch.Query) error {
	return conn.Do(ctx, q)
}
