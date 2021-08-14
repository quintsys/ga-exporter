package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/quintsys/ga-exporter/middleware"
)

func TestContentTypeJSON(t *testing.T) {
	tests := []struct {
		name string
		r    *http.Request
		m    http.HandlerFunc
		want string
	}{
		{
			name: "JSON",
			r: &http.Request{
				Method: "GET",
			},
			m:    middleware.ContentTypeJSON(mockHandler()),
			want: "application/json",
		},
		{
			name: "EMPTY",
			r: &http.Request{
				Method: "GET",
			},
			m:    mockHandler(),
			want: "",
		},
	}

	for _, tt := range tests {
		tt := tt // https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := httptest.NewRecorder()
			tt.m.ServeHTTP(w, tt.r)
			if w.Header().Get("Content-Type") != tt.want {
				t.Errorf("Content-Type = %q, want %q",
					w.Header().Get("Content-Type"),
					tt.want)
			}
		})
	}
}
