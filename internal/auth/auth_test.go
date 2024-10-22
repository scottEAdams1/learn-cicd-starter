package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name: "valid api key",
			headers: func() http.Header {
				h := make(http.Header)
				h.Add("Authorization", "ApiKey your-test-key")
				return h
			}(),
			want:    "your-test-key",
			wantErr: nil,
		},
		{
			name:    "missing authorization header",
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "malformed authorization header",
			headers: make(http.Header),
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("expected error: %v, got: %v", tc.wantErr, err)
			}
			if got != tc.want {
				t.Errorf("expected: %v, got %v", tc.want, got)
			}
		})
	}
}
