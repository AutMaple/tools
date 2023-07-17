package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"sync"
	"tools/internal/log"
)

var (
	db    *sql.DB
	cache = make(map[string]*sql.Stmt)
)

type Configs struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func Open(configs Configs) {
	var err error
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.User, configs.Password, configs.Host, configs.Port, configs.Database)
	db, err = sql.Open("mysql", uri)
	if err != nil {
		log.PanicMsg(fmt.Sprintf("Failed to connect to database: %s", uri), errors.WithStack(err))
		return
	}
	if err := db.Ping(); err != nil {
		log.PanicMsg(fmt.Sprintf("Failed to ping database: %s", uri), errors.WithStack(err))
		return
	}
	log.Info("Connected to mysql database successfully!")
}

func PrepareStatement(query string) (*sql.Stmt, error) {
	var mu sync.Mutex
	if _, ok := cache[query]; !ok {
		mu.Lock()
		if _, ok := cache[query]; !ok {
			stmt, err := db.Prepare(query)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			cache[query] = stmt
		}
		mu.Unlock()
	}
	return cache[query], nil
}
