package config

import (
	"os"

	"github.com/MigueelLago/goportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Checa se o banco de dados existe
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Banco de dados não existe, criando...")

		// Criar o diretório e o arquivo
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Criar e conectar ao banco
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.ErrorF("Erro ao executar SQLite: %v", err)
	}

	// Migração de Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("Erro na migração no SQLite: %v", err)
	}

	return db, nil
}
