package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlers(t *testing.T) {
	type test struct {
		method               string
		acceptLanguageHeader string
		handler              func(w http.ResponseWriter, r *http.Request)
		expectedCode         int
		expectedContent      string
	}
	tests := []test{
		{"GET", "cy", indexHandler, 302, "http://localhost:8080/croeso"},
		{"GET", "de", indexHandler, 302, "http://localhost:8080/welcome"},
		{"GET", "en", indexHandler, 302, "http://localhost:8080/welcome"},
		{"GET", "fr", indexHandler, 302, "http://localhost:8080/welcome"},
		{"GET", "", indexHandler, 302, "http://localhost:8080/welcome"},
		{"HEAD", "", indexHandler, 405, "Unsupported method"},
		{"POST", "", indexHandler, 405, "Unsupported method"},
		{"PUT", "", indexHandler, 405, "Unsupported method"},
		{"DELETE", "", indexHandler, 405, "Unsupported method"},
		{"CONNECT", "", indexHandler, 405, "Unsupported method"},
		{"OPTIONS", "", indexHandler, 405, "Unsupported method"},
		{"TRACE", "", indexHandler, 405, "Unsupported method"},
		{"PATCH", "", indexHandler, 405, "Unsupported method"},
		{"GET", "cy", checkHandler, 200, "Prefer cy: true"},
		{"GET", "de", checkHandler, 200, "Prefer cy: false"},
		{"GET", "en", checkHandler, 200, "Prefer cy: false"},
		{"GET", "fr", checkHandler, 200, "Prefer cy: false"},
		{"GET", "", checkHandler, 200, "Prefer cy: false"},
		{"HEAD", "", checkHandler, 405, "Unsupported method"},
		{"POST", "", checkHandler, 405, "Unsupported method"},
		{"PUT", "", checkHandler, 405, "Unsupported method"},
		{"DELETE", "", checkHandler, 405, "Unsupported method"},
		{"CONNECT", "", checkHandler, 405, "Unsupported method"},
		{"OPTIONS", "", checkHandler, 405, "Unsupported method"},
		{"TRACE", "", checkHandler, 405, "Unsupported method"},
		{"PATCH", "", checkHandler, 405, "Unsupported method"},
		{"GET", "", croesoHandler, 200, "Croeso i'r tudalen Cymraeg"},
		{"HEAD", "", croesoHandler, 405, "Unsupported method"},
		{"POST", "", croesoHandler, 405, "Unsupported method"},
		{"PUT", "", croesoHandler, 405, "Unsupported method"},
		{"DELETE", "", croesoHandler, 405, "Unsupported method"},
		{"CONNECT", "", croesoHandler, 405, "Unsupported method"},
		{"OPTIONS", "", croesoHandler, 405, "Unsupported method"},
		{"TRACE", "", croesoHandler, 405, "Unsupported method"},
		{"PATCH", "", croesoHandler, 405, "Unsupported method"},
		{"GET", "", welcomeHandler, 200, "Welcome to the English page"},
		{"HEAD", "", welcomeHandler, 405, "Unsupported method"},
		{"POST", "", welcomeHandler, 405, "Unsupported method"},
		{"PUT", "", welcomeHandler, 405, "Unsupported method"},
		{"DELETE", "", welcomeHandler, 405, "Unsupported method"},
		{"CONNECT", "", welcomeHandler, 405, "Unsupported method"},
		{"OPTIONS", "", welcomeHandler, 405, "Unsupported method"},
		{"TRACE", "", welcomeHandler, 405, "Unsupported method"},
		{"PATCH", "", welcomeHandler, 405, "Unsupported method"},
	}
	for _, tc := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(tc.method, "http://irrelevant", nil)
		if tc.acceptLanguageHeader != "" {
			r.Header.Set("Accept-Language", tc.acceptLanguageHeader)
		}
		h := http.HandlerFunc(tc.handler)
		h.ServeHTTP(w, r)
		assert.Equal(t, tc.expectedCode, w.Code)
		assert.Contains(t, w.Body.String(), tc.expectedContent)
	}
}
