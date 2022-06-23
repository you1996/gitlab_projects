package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// helper for making tests
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func TestShouldReturn500WhenParamNameIsWrong(t *testing.T) {
	r := setupRouter()
	t.Run("", func(t *testing.T) {
		w := performRequest(r, "GET", "/projects-names-and-stars?proj-number=20")
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "{\"code\":500,\"message\":\"Invalid parameter !\"}", w.Body.String())
	})
}

func TestShouldReturn500IfParamIsNotInteger(t *testing.T) {
	r := setupRouter()
	t.Run("", func(t *testing.T) {
		w := performRequest(r, "GET", "/projects-names-and-stars?number-of-projects=gitlap")
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "{\"code\":500,\"message\":\"Can't parse query string !\"}", w.Body.String())
	})
}

func TestShouldReturn500IfParamIsZero(t *testing.T) {
	r := setupRouter()
	t.Run("", func(t *testing.T) {
		w := performRequest(r, "GET", "/projects-names-and-stars?number-of-projects=0")
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "{\"code\":500,\"message\":\"Number of projects to fetch cannot be zero !\"}", w.Body.String())
	})
}
