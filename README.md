# chh - ClickHouse ch-go Helper

Small wrapper for https://github.com/ClickHouse/ch-go package to simplify client code.

Install: `go get github.com/nikita-vanyasin/chh@latest`

## Usage example

Select whole result-set:
```go
type colsPerformanceResult struct {
	Date        []time.Time `json:"date"`
	Clicks      []uint64    `json:"clicks"`
	Impressions []uint64    `json:"impressions"`
	CTR         []float64   `json:"ctr"`
	Position    []float64   `json:"position"`
	Visibility  []float64   `json:"visibility"`
}

// ...

ret := &colsPerformanceResult{}
results := &chh.ColResults{}
chh.BindDate(results, "date", &ret.Date)
chh.BindUInt64(results, "clicks", &ret.Clicks)
chh.BindUInt64(results, "impressions", &ret.Impressions)
chh.BindFloat64(results, "ctr", &ret.CTR)
chh.BindFloat64(results, "avgPosition", &ret.Position)
chh.BindFloat64(results, "visibility", &ret.Visibility)

if err := chh.Select(ctx, conn, results, getPerformanceChartQuery(c)); err != nil {
	return nil, err
}
```

Select one row:
```go
totals := &Totals{}
results := &chh.ColResults{}
results.BindUInt64("count", &totals.Count)
results.BindUInt64("sumClicks", &totals.Clicks)
results.BindUInt64("sumImpressions", &totals.Impressions)
results.BindFloat64("avgCTR", &totals.CTR)
results.BindFloat64("avgPosition", &totals.Position)

if err := chh.SelectRow(ctx, conn, results, getQueriesTotalQuery(c)); err != nil {
	return nil, err
}
```


#### Links
- Inspired by https://clarktrimble.online/blog/funhouse/
- https://github.com/ClickHouse/ch-go