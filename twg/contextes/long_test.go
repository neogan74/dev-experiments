package contextes

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	respose string
	t       *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.respose {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(time.Millisecond * 10)
				result += string(c)
			}
		}
		data <- result
	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
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
	})

}
