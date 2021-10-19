package handlers

import (
	"BookStore/internal/models"
	"BookStore/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHTTPHandler_Books(t *testing.T) {
	bStoreExample := models.BStore{
		ID:          0,
		Title:       "classic",
		Author:      "golang for por",
		PublicDate:  time.Date(2021, time.April, 21, 21, 21, 21, 21, time.UTC),
		PagesAmount: 2000,
		CreatedTime: time.Date(2022, time.May, 1, 1, 1, 1, 1, time.UTC),
		UpdatedTime: time.Time{},
	}
	var books []models.BStore
	books = append(books, bStoreExample)
	t.Run("get request for all book, all okey", func(t *testing.T) {
		controller := mocks.Controller{}
		controller.On("GetAllBooks").Return(books, nil)
		var handler = NewHTTPHandler(&controller)
		req, err := http.NewRequest("GET", "/book", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		endpoint := http.HandlerFunc(handler.GetAllBooks)
		endpoint.ServeHTTP(rr, req)
		require.Equal(t, http.StatusOK, rr.Code)
		var result []models.BStore
		err = json.NewDecoder(rr.Body).Decode(&result)
		if err != nil {
			t.Fatal("Wrong result, cannot decode to struct")
		}
		require.Equal(t, books, result)
	})
}
