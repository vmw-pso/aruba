package service

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/vmw-pso/aruba/resource-service/data"
)

type Configuration struct {
	Port int64
	Env  string
	DB   struct {
		DSN          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
}

type service struct {
	cfg  Configuration
	repo data.Repository
	wg   sync.WaitGroup
}

func NewService(cfg Configuration, db *sql.DB) *service {
	return &service{
		cfg:  cfg,
		repo: data.NewInMemoryRepository(),
		wg:   sync.WaitGroup{},
	}
}

func (s *service) Run() error {
	fmt.Println("Hello World")
	return nil
}
