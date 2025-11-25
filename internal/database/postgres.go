package database

import (
	"fmt"
	"logiTrack/config"
	"logiTrack/internal/logger"
	"time"

	"github.com/jmoiron/sqlx"
)

type PostgresDB struct {
	DB *sqlx.DB
}

func NewPostgresDB(cfg *config.DatabaseConfig) (*PostgresDB, error) {
	startTime := time.Now()

	//connStrForLog := fmt.Sprintf(
	//	"host=%s port=%s user=%s dbname=%s sslmode=%s",
	//	cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.SSLMode)
	//
	//logger.Log.WithFields(logger.Fields{
	//	"host":    cfg.Host,
	//	"port":    cfg.Port,
	//	"user":    cfg.User,
	//	"dbname":  cfg.DBName,
	//	"sslmode": cfg.SSLMode,
	//}).Info("Попытка подключения к PostgreSQL")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к PostgreSQL: %w", err)
	}

	duration := time.Since(startTime)

	logger.Log.WithFields(logger.Fields{
		"duration_ms": duration.Milliseconds(),
	})

	return &PostgresDB{DB: db}, nil
}

func (p *PostgresDB) Close() error {
	return p.DB.Close()
}
