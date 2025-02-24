package models

import (
	"Go_Day06/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

func New(cfg *config.DataBaseConfig) (*sqlx.DB, error) {
	location, err := time.LoadLocation(cfg.TimeZone)
	if err != nil {
		return nil, fmt.Errorf("load time location fail: %v", err)
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable timeZone=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Name,
		cfg.Password,
		location.String(),
	)
	dbConn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("connect postgres fail: %v", err)
	}
	dbConn.SetMaxIdleConns(cfg.MaxIdleConnections)
	dbConn.SetMaxOpenConns(cfg.MaxOpenConnections)
	if err := dbConn.Ping(); err != nil {
		return nil, fmt.Errorf("ping postgres fail: %v", err)
	}
	return dbConn, nil
}
