package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/quintsys/ga-exporter/middleware"
)

func TestPostOnly(t *testing.T) {
	tests := []struct {
		name string
		r    *http.Request
		want int
	}{
		{
			name: "GET",
			r: &http.Request{
				Method: "GET",
			},
			want: http.StatusMethodNotAllowed,
		},
		{
			name: "POST",
			r: &http.Request{
				Method: "POST",
			},
			want: http.StatusOK,
		},
	}

	for _, tt := range tests {
		tt := tt // https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := httptest.NewRecorder()
			m := middleware.PostOnly(mockHandler())
			m.ServeHTTP(w, tt.r)
			if w.Code != tt.want {
				t.Errorf("got code %d, want %d", w.Code, tt.want)
			}
		})
	}
}
