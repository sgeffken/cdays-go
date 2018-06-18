package routing_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sgeffken/cdays-go/internal/routing"
)

func TestNewBLRouter(t *testing.T) {
	r := routing.NewBLRouter()
	srv := httptest.NewServer(r)
	srv.Start()

	// tell go that I am going to close the server when this function terminates
	defer srv.Close()

	testCases := []struct {
		route              string
		expectedStatusCode int
	}{
		{"/home", http.StatusOK},
		{"/", http.StatusNotFound},
	}

	for _, c := range testCases {
		resp, err := http.Get(srv.URL + c.route)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != c.expectedStatusCode {
			t.Errorf("Status code is %v, but %v expected", resp.StatusCode, c.expectedStatusCode)
		}

	}
}
