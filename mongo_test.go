package vps_monitoring_notifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingMongo(t *testing.T) {
	cases := []struct {
		name string

		url string

		wantErr bool
	}{
		{
			name: "ok",
			url:  "mongodb://localhost:27017",
		},
		{
			name:    "ping error",
			url:     "mongodb://localhost:27018",
			wantErr: true,
		},
		{
			name:    "connect error",
			url:     "mongo://localhost:27017",
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			err := pingMongo(tc.url)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}
