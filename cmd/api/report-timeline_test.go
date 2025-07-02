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

func TestReportTimeline(t *testing.T) {

	t.Run("Should return 400 if service returns with UserErr", func(t *testing.T) {
		mockTm := &mockTransactionManager{
			ReportTimelineFunc: func(ctx context.Context, start, end time.Time) (contracts.TimelineResponse, error) {
				return contracts.TimelineResponse{}, tm.UserErr("Invalid date range")
			},
		}
		server := &Server{tm: mockTm}
		reqBody := contracts.TimelineRequest{
			StartDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
		}
		requestBody, err := json.Marshal(reqBody)
		require.NoError(t, err, "Failed to marshal request body")

		// Simulate the HTTP request
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/report-timeline", bytes.NewBuffer(requestBody))
		r.Header.Set("Content-Type", "application/json")
		server.reportTimeline(w, r)

		require.Equal(t, 400, w.Code, "Expected status code 400 for UserErr")
	})

	t.Run("Should return 400 for invalid request body", func(t *testing.T) {
		mockTm := &mockTransactionManager{}
		server := &Server{tm: mockTm}

		// Simulate the HTTP request with an invalid body
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/report-timeline", bytes.NewBufferString("invalid body"))
		r.Header.Set("Content-Type", "application/json")

		server.reportTimeline(w, r)
		require.Equal(t, 400, w.Code, "Expected status code 400 for invalid request body")
	})

	t.Run("Should return report timeline data", func(t *testing.T) {
		startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)

		expectedResponse := contracts.TimelineResponse{
			Items: []contracts.TimelineResponseItem{
				{
					Month: "2023-01-01T00:00:00Z",
					Summary: contracts.Summary{
						SpendingCents: 1000,
						EarningCents:  2000,
						NetCents:      1000,
					},
				},
			},
		}

		mockTm := &mockTransactionManager{
			ReportTimelineFunc: func(ctx context.Context, start, end time.Time) (contracts.TimelineResponse, error) {
				if !start.Equal(startDate) && !end.Equal(endDate) {
					require.Fail(t, "Unexpected start or end date in ReportTimeline")
				}
				return expectedResponse, nil
			},
		}
		server := &Server{tm: mockTm}

		reqBody := contracts.TimelineRequest{
			StartDate: startDate,
			EndDate:   endDate,
		}

		requestBody, err := json.Marshal(reqBody)
		require.NoError(t, err, "Failed to marshal request body")

		// Simulate the HTTP request and response
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/report-timeline", bytes.NewBuffer(requestBody))
		r.Header.Set("Content-Type", "application/json")

		server.reportTimeline(w, r)

		require.Equal(t, 200, w.Code, "Expected status code 200")

		var response contracts.TimelineResponse
		err = json.NewDecoder(w.Body).Decode(&response)
		require.NoError(t, err, "Failed to decode response body")

    require.Equal(t, expectedResponse, response, "Expected response to match the expected data")
	})
}
