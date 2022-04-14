package main

import (
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		label string
		want  string
	}{
		{
			label: "default",
			want:  "Hello World!\n",
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()
		handler(rr, req)

		if got := rr.Body.String(); got != test.want {
			t.Errorf("%s: got %q, want %q", test.label, got, test.want)
		}
	}
}
