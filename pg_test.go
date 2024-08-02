package vps_monitoring_notifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingPG(t *testing.T) {
	cases := []struct {
		name string

		url string

		wantErr bool
	}{
		{
			name: "ok",
			url:  "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable",
		},
		{
			name:    "ping error",
			url:     "postgres://postgres:mysecretpassword@localhost:5433/postgres?sslmode=disable",
			wantErr: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			err := pingPG(tc.url)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}
