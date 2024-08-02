package vps_monitoring_notifier

import (
	"fmt"
	"net/http"
)

func ping(url string) error {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("creating ping bot request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("ping bot request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ping bot request failed status code not ok: %s", resp.Status)
	}

	return nil
}
