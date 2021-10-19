package handlers

import (
	"BookStore/internal/models"
	"BookStore/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

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

	t.Run("good GET request for all book", func(t *testing.T) {
		controller := mocks.Controller{}
		controller.On("GetAllBooks").Return(books, nil)
		var handler = NewHTTPHandler(&controller)
		req, err := http.NewRequest(http.MethodGet, "/book/add", nil)
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

	t.Run("bad POST request, try to add book with nil body", func(t *testing.T) {
		controller := mocks.Controller{}
		controller.On("AddBook").Return(nil)
		var handler = NewHTTPHandler(&controller)
		req, err := http.NewRequest(http.MethodPost, "/book/update", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		endpoint := http.HandlerFunc(handler.AddBook)
		endpoint.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.Equal(t, rr.Body.String(), "invalid post form\n")
	})
	t.Run("bad POST request, try to add book with wrong post body", func(t *testing.T) {
		controller := mocks.Controller{}
		controller.On("AddBook", mock.Anything).Return(nil)
		var handler = NewHTTPHandler(&controller)
		req, err := http.NewRequest(http.MethodPost, "/book/update", strings.NewReader(""))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		endpoint := http.HandlerFunc(handler.AddBook)
		endpoint.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.Equal(t, rr.Body.String(), ErrRequestBody+"\n")
	})

	t.Run("bad DELETE request, try to delete book without ID", func(t *testing.T) {
		controller := mocks.Controller{}
		controller.On("DeleteBookByID", mock.Anything).Return(nil)
		var handler = NewHTTPHandler(&controller)
		req, err := http.NewRequest(http.MethodDelete, "/book/delete", strings.NewReader(""))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		endpoint := http.HandlerFunc(handler.DeleteBookByID)
		endpoint.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.Equal(t, rr.Body.String(), ErrRequestBody+"\n")
	})
	t.Run("bad DELETE request, try to delete book with wrong ID", func(t *testing.T) {
		controller := mocks.Controller{}
		controller.On("DeleteBookByID", mock.Anything).Return(nil)
		var handler = NewHTTPHandler(&controller)
		req, err := http.NewRequest(http.MethodDelete, "/book/delete", strings.NewReader("-1"))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		endpoint := http.HandlerFunc(handler.DeleteBookByID)
		endpoint.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.Equal(t, rr.Body.String(), "Invalid ID\n")
	})
}
