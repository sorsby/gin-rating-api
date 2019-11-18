package gins

import (
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/sorsby/gin-rating-api/data"
)

func TestList(t *testing.T) {
	testCases := []struct {
		desc      string
		req       *http.Request
		ginLister data.GinLister
		expStatus int
		expBody   string
	}{
		{
			desc: "gin lister fails",
			req:  httptest.NewRequest(http.MethodGet, "/gins", strings.NewReader("{ whatever: 'trevor' }")),
			ginLister: func() (data.ListGinOutput, error) {
				return data.ListGinOutput{}, errors.New("failure")
			},
			expStatus: http.StatusInternalServerError,
			expBody:   `"failure"` + "\n",
		},
		{
			desc: "success",
			req:  httptest.NewRequest(http.MethodGet, "/gins", strings.NewReader("{ whatever: 'trevor' }")),
			ginLister: func() (data.ListGinOutput, error) {
				return data.ListGinOutput{
					GinItems: []data.GinItem{
						data.GinItem{
							ID:           "123",
							Name:         "gin-1",
							Quantity:     "300ml",
							ABV:          "40",
							LastModified: "123",
						},
					},
				}, nil
			},
			expStatus: http.StatusOK,
			expBody:   `"{\"gins\":[{\"ID\":\"123\",\"name\":\"gin-1\",\"quantity\":\"300ml\",\"abv\":\"40\",\"imageUrl\":\"\",\"lastModified\":\"123\"}]}"` + "\n",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			router := mux.NewRouter()

			h := NewHandler(tC.ginLister)
			router.Path("/gins").Methods(http.MethodGet).HandlerFunc(h.List)

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
