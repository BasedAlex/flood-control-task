package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type DB interface {
	Set(context.Context, int, int64, time.Time) error
	Get(context.Context, int64) (int, time.Time, error)
}

type Postgres struct {
	db *sql.DB
}

func NewPostgres(dbConnect string) (*Postgres, error){
	db, err := sql.Open("postgres", dbConnect)
	if err != nil {
		return nil, err
	}
	return &Postgres{db:db}, nil
}

func (db *Postgres) Get(ctx context.Context, userID int64) (int, time.Time, error) {
	stmt := `SELECT call_count, first_call from floodcontrol WHERE user_id = ($1)`

	var callCount int
	var firstCall time.Time

	err := db.db.QueryRowContext(ctx, stmt, userID).Scan(&callCount, &firstCall)
	if err != nil {
		return 0, time.Time{}, err
	}

	return callCount, firstCall, nil

}


func (db *Postgres) Set(ctx context.Context, count int, userID int64, firstCall time.Time) error {

	stmt := `INSERT INTO floodcontrol (user_id, first_call, call_count) VALUES ($1, $2, $3)`

	_, err := db.db.ExecContext(ctx, stmt, userID, firstCall, count)

	if err != nil {
		return err
	}
	return nil
}