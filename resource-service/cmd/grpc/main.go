package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vmw-pso/aruba/resource-service/internal/service"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg := service.Configuration{}
	db, err := connectToDB()
	if err != nil {
		return err
	}
	svc := service.NewService(cfg, db)
	if err := svc.Run(); err != nil {
		return err
	}
	return nil
}

func connectToDB() (*sql.DB, error) {
	dsn := os.Getenv("DSN")
	counts := 0
	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Printf("Postgres not yet ready: %v\n", err)
			counts++
		} else {
			return conn, nil
		}
		if counts > 10 {
			return nil, errors.New("could not connect to database")
		}
		log.Println("Backing off and waiting for 2 seconds...")
		time.Sleep(2 * time.Second)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
