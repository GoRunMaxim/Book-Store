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

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
)

func TestHTTPHandler_AddBook(t *testing.T) {
	t.Run("good POST request, try to add book", func(t *testing.T) {
		t.Parallel()
		bStoreExample := models.BStore{
			ID:          5,
			Title:       "classic",
			Author:      "golang for por",
			PublicDate:  time.Date(2021, time.April, 21, 21, 21, 21, 21, time.UTC),
			PagesAmount: 2000,
			CreatedTime: time.Date(2022, time.May, 1, 1, 1, 1, 1, time.UTC),
			UpdatedTime: time.Time{},
		}
		controller := mocks.Controller{}
		controller.On("AddBook", mock.Anything).Return(nil)
		var handler = NewHTTPHandler(&controller)
		jsonData, _ := json.Marshal(bStoreExample)
		req, err := http.NewRequest(http.MethodPost, "/book/add", strings.NewReader(string(jsonData)))
		if err != nil {
			t.Fatal(err)
		}
		logrus.Info(strings.NewReader(string(jsonData)))
		rr := httptest.NewRecorder()
		endpoint := http.HandlerFunc(handler.AddBook)
		endpoint.ServeHTTP(rr, req)
		require.Equal(t, http.StatusOK, rr.Code)
	})
	t.Run("bad POST request, try to add book with nil body", func(t *testing.T) {
		t.Parallel()
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
		t.Parallel()
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
	t.Run("bad POST request, try to add book description with wrong id", func(t *testing.T) {
		t.Parallel()
		bStoreExample := models.BStore{
			ID:          5,
			Title:       "classic",
			Author:      "golang for por",
			PublicDate:  time.Date(2021, time.April, 21, 21, 21, 21, 21, time.UTC),
			PagesAmount: 2000,
			CreatedTime: time.Date(2022, time.May, 1, 1, 1, 1, 1, time.UTC),
			UpdatedTime: time.Time{},
		}
		controller := mocks.Controller{}
		controller.On("AddBook", mock.Anything).Return(nil)
		var handler = NewHTTPHandler(&controller)
		bStoreExample.ID = -1
		jsonData, _ := json.Marshal(bStoreExample)
		req, err := http.NewRequest(http.MethodPost, "/book/update", strings.NewReader(string(jsonData)))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		endpoint := http.HandlerFunc(handler.AddBook)
		endpoint.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.Equal(t, rr.Body.String(), "id cannot be less than 0\n")
	})
	t.Run("bad POST request, try to add book description with empty author", func(t *testing.T) {
		t.Parallel()
		bStoreExample := models.BStore{
			ID:          5,
			Title:       "classic",
			Author:      "golang for por",
			PublicDate:  time.Date(2021, time.April, 21, 21, 21, 21, 21, time.UTC),
			PagesAmount: 2000,
			CreatedTime: time.Date(2022, time.May, 1, 1, 1, 1, 1, time.UTC),
			UpdatedTime: time.Time{},
		}
		controller := mocks.Controller{}
		controller.On("AddBook", mock.Anything).Return(nil)
		var handler = NewHTTPHandler(&controller)
		bStoreExample.Author = ""
		jsonData, _ := json.Marshal(bStoreExample)
		req, err := http.NewRequest(http.MethodPost, "/book/add", strings.NewReader(string(jsonData)))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		endpoint := http.HandlerFunc(handler.AddBook)
		endpoint.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
		logrus.Info(rr.Body.String())
		require.Equal(t, rr.Body.String(), "'author' field cannot be empty\n")
	})
}

func TestHTTPHandler_GetAllBooks(t *testing.T) {
	t.Run("good GET request for all book", func(t *testing.T) {
		t.Parallel()
		bStoreExample := models.BStore{
			ID:          5,
			Title:       "classic",
			Author:      "golang for por",
			PublicDate:  time.Date(2021, time.April, 21, 21, 21, 21, 21, time.UTC),
			PagesAmount: 2000,
			CreatedTime: time.Date(2022, time.May, 1, 1, 1, 1, 1, time.UTC),
			UpdatedTime: time.Time{},
		}
		var books []models.BStore
		books = append(books, bStoreExample)
		controller := mocks.Controller{}
		controller.On("GetAllBooks").Return(books, nil)
		var handler = NewHTTPHandler(&controller)
		req, err := http.NewRequest(http.MethodGet, "/book/get", nil)
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

func TestHTTPHandler_DeleteBookByID(t *testing.T) {
	t.Run("bad DELETE request, try to delete book without ID", func(t *testing.T) {
		t.Parallel()
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
		t.Parallel()
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

func TestHTTPHandler_UpdateBookByID(t *testing.T) {
	t.Run("bad POST request, try to update books description with nil body", func(t *testing.T) {
		t.Parallel()
		controller := mocks.Controller{}
		controller.On("UpdateBookByID").Return(nil)
		var handler = NewHTTPHandler(&controller)
		req, err := http.NewRequest(http.MethodPost, "/book/update", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		endpoint := http.HandlerFunc(handler.UpdateBookByID)
		endpoint.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.Equal(t, rr.Body.String(), ErrPostForm+"\n")
	})
	t.Run("bad POST request, try to update books description with wrong body", func(t *testing.T) {
		t.Parallel()
		controller := mocks.Controller{}
		controller.On("UpdateBookByID", mock.Anything).Return(nil)
		var handler = NewHTTPHandler(&controller)
		req, err := http.NewRequest(http.MethodPost, "/book/update", strings.NewReader(""))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		endpoint := http.HandlerFunc(handler.UpdateBookByID)
		endpoint.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.Equal(t, rr.Body.String(), ErrRequestBody+"\n")
	})
	t.Run("bad POST request, try to update books description with wrong id", func(t *testing.T) {
		t.Parallel()
		bStoreExample := models.BStore{
			ID:          5,
			Title:       "classic",
			Author:      "golang for por",
			PublicDate:  time.Date(2021, time.April, 21, 21, 21, 21, 21, time.UTC),
			PagesAmount: 2000,
			CreatedTime: time.Date(2022, time.May, 1, 1, 1, 1, 1, time.UTC),
			UpdatedTime: time.Time{},
		}
		controller := mocks.Controller{}
		controller.On("UpdateBookByID", mock.Anything).Return(nil)
		var handler = NewHTTPHandler(&controller)
		bStoreExample.ID = -1
		jsonData, _ := json.Marshal(bStoreExample)
		req, err := http.NewRequest(http.MethodPost, "/book/update", strings.NewReader(string(jsonData)))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		endpoint := http.HandlerFunc(handler.UpdateBookByID)
		endpoint.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.Equal(t, rr.Body.String(), "id cannot be less than 0\n")
	})
	t.Run("bad POST request, try to update books description with empty author", func(t *testing.T) {
		t.Parallel()
		bStoreExample := models.BStore{
			ID:          5,
			Title:       "classic",
			Author:      "golang for por",
			PublicDate:  time.Date(2021, time.April, 21, 21, 21, 21, 21, time.UTC),
			PagesAmount: 2000,
			CreatedTime: time.Date(2022, time.May, 1, 1, 1, 1, 1, time.UTC),
			UpdatedTime: time.Time{},
		}
		controller := mocks.Controller{}
		controller.On("UpdateBookByID", mock.Anything).Return(nil)
		var handler = NewHTTPHandler(&controller)
		bStoreExample.Author = ""
		jsonData, _ := json.Marshal(bStoreExample)
		req, err := http.NewRequest(http.MethodPost, "/book/update", strings.NewReader(string(jsonData)))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		endpoint := http.HandlerFunc(handler.UpdateBookByID)
		endpoint.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.Equal(t, rr.Body.String(), "'author' field cannot be empty\n")
	})
	t.Run("bad POST request, try to update books description with empty title", func(t *testing.T) {
		t.Parallel()
		bStoreExample := models.BStore{
			ID:          5,
			Title:       "classic",
			Author:      "golang for por",
			PublicDate:  time.Date(2021, time.April, 21, 21, 21, 21, 21, time.UTC),
			PagesAmount: 2000,
			CreatedTime: time.Date(2022, time.May, 1, 1, 1, 1, 1, time.UTC),
			UpdatedTime: time.Time{},
		}
		controller := mocks.Controller{}
		controller.On("UpdateBookByID", mock.Anything).Return(nil)
		var handler = NewHTTPHandler(&controller)
		bStoreExample.Title = ""
		jsonData, _ := json.Marshal(bStoreExample)
		req, err := http.NewRequest(http.MethodPost, "/book/update", strings.NewReader(string(jsonData)))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		endpoint := http.HandlerFunc(handler.UpdateBookByID)
		endpoint.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.Equal(t, rr.Body.String(), "'title' field cannot be empty\n")
	})
}
