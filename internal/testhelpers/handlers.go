package testhelpers

import (
	"io"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type Header struct {
	Key   string
	Value string
}

func PerformAnonRequest(f gin.HandlerFunc, method, path string, body io.Reader, headers ...Header) *httptest.ResponseRecorder {
	r := gin.Default()
	r.Handle(method, path, f)

	return performRequest(r, method, path, body, headers...)
}

func performRequest(r *gin.Engine, method, path string, body io.Reader, headers ...Header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)

	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}
