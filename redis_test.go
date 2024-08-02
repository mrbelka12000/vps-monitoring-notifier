package vps_monitoring_notifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRedis(t *testing.T) {
	cases := []struct {
		name string

		redisAddr string

		wantErr bool
	}{
		{
			name:      "ok",
			redisAddr: "localhost:6379",
		},
		{
			name:      "ping error",
			redisAddr: "localhost:6378",
			wantErr:   true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			err := pingRedis(tc.redisAddr)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}
