package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// ConnectPostgreSQL connects to PostgreSQL using the provided URI and returns a SQL DB connection
func ConnectPostgreSQL(uri string) (*sql.DB, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Create a context with timeout for the ping
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Verify the connection
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
