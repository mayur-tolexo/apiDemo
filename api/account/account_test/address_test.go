package account_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAddress(t *testing.T) {
	tc := []struct {
		name     string
		expected string
		status   int
		err      string
	}{
		{
			name:     "basic case",
			status:   http.StatusOK,
			expected: `{"data":{"addresss":[],"total_count":0}}`,
		},
	}

	for _, ctc := range tc {
		t.Run(ctc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/account/v1/address/", nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, ctc.status, w.Code)
			assert.Equal(t, ctc.expected, strings.TrimSuffix(w.Body.String(), "\n"))
		})
	}
}
