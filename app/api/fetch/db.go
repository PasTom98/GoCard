package fetch

import (
	"context"
	"database/sql"
	"time"
)

func New(address string, maxOpenConnections int, maxIdleConnections int, maxIdleTime string) *sql.DB {
	db, err := sql.Open("postgres", address)
	if err != nil {
		return nil
	}

	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxIdleConnections)

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil
	}

	return db
}
