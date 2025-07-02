package main

import (
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"testing"
	"tm/internal/tm"
	"tm/pkg/contracts"

	"github.com/stretchr/testify/require"
)

func TestImportCsv(t *testing.T) {

	t.Run("Should return 400 if the manager returns a user error", func(t *testing.T) {
		mockTm := &mockTransactionManager{
			ImportCsvFunc: func(ctx context.Context, f io.Reader) (contracts.ImportCsvResponse, error) {
				return contracts.ImportCsvResponse{}, tm.UserErr("user error")
			},
		}

		server := &Server{tm: mockTm}
		pr, pw := io.Pipe()
		mw := multipart.NewWriter(pw)
		go func() {
			defer mw.Close()
			part, err := mw.CreateFormFile("file", "test.csv")
			require.NoError(t, err, "Failed to create form file")
			part.Write([]byte("CSV_DATA"))
		}()
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/import-csv", pr)
		r.Header.Set("Content-Type", mw.FormDataContentType())

		server.importCsv(w, r)
		require.Equal(t, 400, w.Code, "Expected status code 400 for user error")
	})

	t.Run("Should return 400 if no file is provided", func(t *testing.T) {
		server := &Server{tm: &mockTransactionManager{}}
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/import-csv", nil)
		server.importCsv(w, r)
		require.Equal(t, 400, w.Code, "Expected status code 400 for missing file")
	})

	t.Run("Should return the response from the manager", func(t *testing.T) {
		expectedResponse := contracts.ImportCsvResponse{
			Duplicates: 5,
			Total:      10,
		}

		csvData := `CSV_DATA`
		mockTm := &mockTransactionManager{
			ImportCsvFunc: func(ctx context.Context, f io.Reader) (contracts.ImportCsvResponse, error) {
				str, err := io.ReadAll(f)
				require.NoError(t, err, "Failed to read CSV data")
				if string(str) != csvData {
					require.Fail(t, "CSV data does not match expected data")
				}
				return expectedResponse, nil
			},
		}

		server := &Server{tm: mockTm}

		pr, pw := io.Pipe()
		mw := multipart.NewWriter(pw)

		go func() {
			defer mw.Close()
			part, err := mw.CreateFormFile("file", "test.csv")

			if err != nil {
				t.Errorf("Failed to create form file: %v", err)
			}
			part.Write([]byte(csvData))
		}()

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/import-csv", pr)
		r.Header.Set("Content-Type", mw.FormDataContentType())

		server.importCsv(w, r)
		require.Equal(t, 200, w.Code, "Expected status code 200")

		var response contracts.ImportCsvResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		require.NoError(t, err, "Failed to decode response")

		require.Equal(t, expectedResponse, response, "Response does not match expected response")
	})
}
