package main

import (
	"context"
	"log"
	"task/internal/config"
	"task/internal/db"
	"task/internal/floodcontrol"
	"time"
)



func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := db.NewPostgres(cfg.DBConnect)
	if err != nil {
		log.Fatal(err)
	}

	threshold := cfg.Threshold
	timeLimit := time.Second * time.Duration(cfg.TimeLimit)

	floodcontrol.New(threshold, timeLimit, db)

}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
