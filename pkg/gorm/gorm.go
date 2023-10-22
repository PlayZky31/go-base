package gorm

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func NewGorm(host, port, user, pass, db string) (*gorm.DB, *sql.DB, error) {
	postgresConfig := postgres.Config{
		PreferSimpleProtocol: true,
		DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s search_path=%s",
			host,
			port,
			user,
			pass,
			db,
			"disable",      // ssl mode
			"ASIA/JAKARTA", // timezone
			"public",       // schema
		),
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	provider, err := gorm.Open(postgres.New(postgresConfig), &gorm.Config{
		PrepareStmt: true,
		Logger:      newLogger,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("model not connected : %s", err.Error())
	}

	sqlDB, err := provider.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("model not connected : %s", err.Error())
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, nil, fmt.Errorf("failed checking connection to database, err: %s", err.Error())
	}

	provider = provider.Debug()

	return provider, sqlDB, nil
}
