package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"
	"tm/internal/tm"
	"tm/pkg/contracts"

	"github.com/stretchr/testify/require"
)

func TestReportPeriodHandler(
	t *testing.T,
) {

	t.Run("Should return 400 if service returns with UserErr", func(t *testing.T) {
		mockTm := &mockTransactionManager{
			ReportPeriodFunc: func(ctx context.Context, start, end time.Time, u100 bool) (contracts.ReportResponse, error) {
				return contracts.ReportResponse{}, tm.UserErr("Invalid date range")
			},
		}
		server := &Server{tm: mockTm}
		reqBody := contracts.ReportRequest{
			StartDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			U100:      true,
		}
		requestBody, err := json.Marshal(reqBody)
		require.NoError(t, err, "Failed to marshal request body")

		// Simulate the HTTP request
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/report-period", bytes.NewBuffer(requestBody))
		r.Header.Set("Content-Type", "application/json")
		server.reportPeriod(w, r)

		require.Equal(t, 400, w.Code, "Expected status code 400 for UserErr")
	})

	t.Run("Should return 400 for invalid request body", func(t *testing.T) {
		mockTm := &mockTransactionManager{}
		server := &Server{tm: mockTm}

		// Simulate the HTTP request with an invalid body
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/report-period", bytes.NewBufferString("invalid body"))
		r.Header.Set("Content-Type", "application/json")

		server.reportPeriod(w, r)
		require.Equal(t, 400, w.Code, "Expected status code 400 for invalid request body")
	})

	t.Run("Should return report period data", func(t *testing.T) {
		startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)
    u100Request := true

		expectedResponse := contracts.ReportResponse{
			Summary: contracts.Summary{
				SpendingCents: 1000,
				EarningCents:  2000,
				NetCents:      1000,
			},
			TopSpendings: []contracts.Transaction{
				{
					Id:          "1",
					Date:        "2023-01-01T00:00:00Z",
					Description: "Test Spending",
					AmountCents: -500,
				},
				{
					Id:          "2",
					Date:        "2023-01-01T00:00:00Z",
					Description: "Another Spending",
					AmountCents: -500,
				},
			},
		}

		mockTm := &mockTransactionManager{
			ReportPeriodFunc: func(ctx context.Context, start, end time.Time, u100 bool) (contracts.ReportResponse, error) {
				if !start.Equal(startDate) || !end.Equal(endDate) || u100 != u100Request {
					require.Fail(t, "Expected start and end dates to match")
				}
				return expectedResponse, nil
			},
		}
		server := &Server{tm: mockTm}

		reqBody := contracts.ReportRequest{
			StartDate: startDate,
			EndDate:   endDate,
			U100:      u100Request,
		}

		requestBody, err := json.Marshal(reqBody)
		require.NoError(t, err, "Failed to marshal request body")

		// Simulate the HTTP request and response
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/report-period", bytes.NewBuffer(requestBody))
		r.Header.Set("Content-Type", "application/json")

		server.reportPeriod(w, r)

		require.Equal(t, 200, w.Code, "Expected status code 200")

		var response contracts.ReportResponse
		err = json.NewDecoder(w.Body).Decode(&response)
		require.NoError(t, err, "Failed to decode response body")

		require.Equal(t, expectedResponse, response, "Expected response to match")
	})
}
