package gins

import (
	"io/ioutil"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		desc      string
		req       *http.Request
		expStatus int
		expBody   string
	}{
		{
			desc:      "handler should always return hello world",
			req:       httptest.NewRequest(http.MethodGet, "/gins", strings.NewReader("{ whatever: 'trevor' }")),
			expStatus: http.StatusOK,
			expBody:   "Hello World!",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			router := mux.NewRouter()

			h := NewHandler()
			router.Path("/gins").Methods(http.MethodGet).HandlerFunc(h.Get)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, tC.req)
			if tC.expStatus != w.Code {
				t.Errorf("expected status %v, got %v", tC.expStatus, w.Code)
			}
			actualBody, err := ioutil.ReadAll(w.Body)
			if err != nil {
				t.Errorf("unexpected error reading body: '%v'", err)
			}
			if tC.expBody != string(actualBody) {
				t.Errorf("expected body '%v', got '%v'", tC.expBody, string(actualBody))
			}
		})
	}
}
