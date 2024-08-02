package vps_monitoring_notifier

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func pingPG(pgURI string) error {

	db, err := sql.Open("postgres", pgURI)
	if err != nil {
		return fmt.Errorf("open pg: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("ping pg: %w", err)
	}

	err = db.Close()
	if err != nil {
		return fmt.Errorf("close pg: %w", err)
	}

	return nil
}
