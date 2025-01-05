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
	t        *testing.T
}

func (s *SpyStore) assertWasCanceled() {
	s.t.Helper()
	if !s.canceled {
		s.t.Errorf("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCanceled() {
	s.t.Helper()
	if s.canceled {
		s.t.Errorf("store was told to canceled")
	}
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.respose
}

func (s *SpyStore) Cancel() {
	s.canceled = true
}

func TestServer(t *testing.T) {
	data := "Hello, Arina"
	t.Run("Simple test", func(t *testing.T) {
		svr := Server(&SpyStore{respose: data, t: t})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		if response.Body.String() != data {
			t.Errorf("got %s, want %s", response.Body.String(), data)
		}
	})

	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{respose: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		responce := httptest.NewRecorder()

		svr.ServeHTTP(responce, req)

		if responce.Body.String() != data {
			t.Errorf("got %s, want %s", responce.Body.String(), data)
		}
		store.assertWasNotCanceled()
	})

	t.Run("tells store to cancel work if request is canceled", func(t *testing.T) {
		store := &SpyStore{respose: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		canceledCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Microsecond, cancel)
		req = req.WithContext(canceledCtx)

		response := httptest.NewRecorder()
		svr.ServeHTTP(response, req)

		store.assertWasCanceled()
	})

}
