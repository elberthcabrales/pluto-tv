package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInitializeServer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	os.Setenv("TOKEN", "dummy-token")

	r := initializeServer()

	if gin.Mode() != gin.TestMode {
		req, _ := http.NewRequest("GET", "/swagger/index.html", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	}

	req, _ := http.NewRequest("GET", "/movies/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.True(t, w.Code == http.StatusNotFound || w.Code == http.StatusInternalServerError)
}
