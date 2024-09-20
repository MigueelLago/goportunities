package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Inicializa o DB SQLite
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("erro ao inicializar sqlite: %v", err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	// Inicializa o logger
	logger = NewLogger(prefix)
	return logger
}
