package gins

import (
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/sorsby/gin-rating-api/claims"
	"github.com/sorsby/gin-rating-api/data"
)

func TestList(t *testing.T) {
	testCases := []struct {
		desc      string
		req       *http.Request
		auth      claims.Authorizer
		ginLister data.GinLister
		expStatus int
		expBody   string
	}{
		{
			desc: "gin lister fails",
			req:  httptest.NewRequest(http.MethodGet, "/gins", strings.NewReader("{ whatever: 'trevor' }")),
			auth: func(r *http.Request) (claims.Claims, bool, error) {
				return claims.Claims{}, true, nil
			},
			ginLister: func() (data.ListGinOutput, error) {
				return data.ListGinOutput{}, errors.New("failure")
			},
			expStatus: http.StatusInternalServerError,
			expBody:   `"failure"` + "\n",
		},
		{
			desc: "success",
			req:  httptest.NewRequest(http.MethodGet, "/gins", strings.NewReader("{ whatever: 'trevor' }")),
			auth: func(r *http.Request) (claims.Claims, bool, error) {
				return claims.Claims{}, true, nil
			},
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

			h := NewHandler(tC.auth, tC.ginLister, nil)
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

func TestPost(t *testing.T) {

	invalidBody := `{"name":"","quantity":"300ml","abv":"44"}`
	validBody := `{"name":"valid-gin","quantity":"300ml","abv":"44"}`

	testCases := []struct {
		desc       string
		req        *http.Request
		auth       claims.Authorizer
		ginCreator data.GinCreater
		expStatus  int
		expBody    string
	}{
		{
			desc: "invalid json request body",
			req:  httptest.NewRequest(http.MethodPost, "/gins", strings.NewReader(`{{...}`)),
			auth: func(r *http.Request) (claims.Claims, bool, error) {
				return claims.Claims{}, true, nil
			},
			expStatus: http.StatusInternalServerError,
			expBody:   `"invalid character '{' looking for beginning of object key string"` + "\n",
		},
		{
			desc: "invalid post request body content",
			req:  httptest.NewRequest(http.MethodPost, "/gins", strings.NewReader(invalidBody)),
			auth: func(r *http.Request) (claims.Claims, bool, error) {
				return claims.Claims{}, true, nil
			},
			expStatus: http.StatusInternalServerError,
			expBody:   `"name must not be an empty string"` + "\n",
		},
		{
			desc: "gin creator fails",
			req:  httptest.NewRequest(http.MethodPost, "/gins", strings.NewReader(validBody)),
			auth: func(r *http.Request) (claims.Claims, bool, error) {
				return claims.Claims{}, true, nil
			},
			ginCreator: func(in data.CreateGinInput) error {
				return errors.New("failure")
			},
			expStatus: http.StatusInternalServerError,
			expBody:   `"failure"` + "\n",
		},
		{
			desc: "success",
			req:  httptest.NewRequest(http.MethodPost, "/gins", strings.NewReader(validBody)),
			auth: func(r *http.Request) (claims.Claims, bool, error) {
				return claims.Claims{}, true, nil
			},
			ginCreator: func(in data.CreateGinInput) error {
				return nil
			},
			expStatus: http.StatusOK,
			expBody:   `"{\"ok\":true}"` + "\n",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			router := mux.NewRouter()

			h := NewHandler(tC.auth, nil, tC.ginCreator)
			router.Path("/gins").Methods(http.MethodPost).HandlerFunc(h.Post)

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
