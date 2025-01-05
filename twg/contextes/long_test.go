package contextes

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	respose  string
	canceled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.respose
}

func (s *SpyStore) Cancel() {
	s.canceled = true
}

func TestServer(t *testing.T) {
	t.Run("Simple test", func(t *testing.T) {
		data := "Hello, Arina"
		svr := Server(&SpyStore{data, false})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		if response.Body.String() != data {
			t.Errorf("got %s, want %s", response.Body.String(), data)
		}
	})

	t.Run("returns data from store", func(t *testing.T) {
		data := "Hello, Amaliya"
		store := &SpyStore{respose: data}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		responce := httptest.NewRecorder()

		svr.ServeHTTP(responce, req)

		if responce.Body.String() != data {
			t.Errorf("got %s, want %s", responce.Body.String(), data)
		}
		if store.canceled {
			t.Errorf("it should not have canceled store")
		}
	})

	t.Run("tells store to cancel work if request is canceled", func(t *testing.T) {
		data := "Hello, Geo"
		strore := &SpyStore{respose: data}
		svr := Server(strore)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		canceledCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Microsecond, cancel)
		req = req.WithContext(canceledCtx)

		response := httptest.NewRecorder()
		svr.ServeHTTP(response, req)

		if !strore.canceled {
			t.Errorf("store was not told to cancel")
		}
	})

}
