package contextes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyStore struct {
	respose string
}

func (s *SpyStore) Fetch() string {
	return s.respose
}

func TestServer(t *testing.T) {
	data := "Hello, Arina"
	svr := Server(&SpyStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)
	if response.Body.String() != data {
		t.Errorf("got %s, want %s", response.Body.String(), data)
	}
}
