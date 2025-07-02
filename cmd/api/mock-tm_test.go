package main

import (
	"context"
	"io"
	"time"
	"tm/pkg/contracts"
)

type mockTransactionManager struct {
	ImportCsvFunc      func(ctx context.Context, f io.Reader) (contracts.ImportCsvResponse, error)
	ReportTimelineFunc func(ctx context.Context, start, end time.Time) (contracts.TimelineResponse, error)
	ReportPeriodFunc   func(ctx context.Context, start, end time.Time, u100 bool) (contracts.ReportResponse, error)
}

func (m *mockTransactionManager) ImportCsv(ctx context.Context, f io.Reader) (contracts.ImportCsvResponse, error) {
	return m.ImportCsvFunc(ctx, f)
}

func (m *mockTransactionManager) ReportTimeline(ctx context.Context, start, end time.Time) (contracts.TimelineResponse, error) {
	return m.ReportTimelineFunc(ctx, start, end)
}

func (m *mockTransactionManager) ReportPeriod(ctx context.Context, start, end time.Time, u100 bool) (contracts.ReportResponse, error) {
	return m.ReportPeriodFunc(ctx, start, end, u100)
}
