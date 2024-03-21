package floodcontrol

import (
	"context"
	"task/internal/db"
	"time"
)

type FC struct {
	threshold int
	timeLimit time.Duration
	database db.DB
}

func New(threshold int, timeLimit time.Duration, db db.DB) *FC {
	return &FC{
		threshold: threshold,
		timeLimit: timeLimit,
		database: db,
	}
}

func (fc *FC) Check(ctx context.Context, userID int64) (bool, error) {
	now := time.Now()
	count, firstCall, err := fc.database.Get(ctx, userID)
	if err != nil {
		return false, err
	}
	durationBetween := now.Sub(firstCall)
	if fc.timeLimit > durationBetween {
		count++
		fc.database.Set(ctx, count, userID, firstCall)
		return count <= fc.threshold, nil
	}
	
	count = 0
	firstCall = now
	fc.database.Set(ctx, count, userID, firstCall)

	return true, nil
}