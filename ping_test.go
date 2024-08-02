package vps_monitoring_notifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingBot(t *testing.T) {
	cases := []struct {
		name string

		url string

		wantErr bool
	}{
		{
			name: "ok",
			url:  "https://google.com",
		},
		{
			name:    "error request create",
			url:     "https://go  ogle.com",
			wantErr: true,
		},
		{
			name:    "error request get",
			url:     "https://example.com",
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			err := ping(tc.url)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}
