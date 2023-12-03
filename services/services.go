package services

import (
	"github.com/attapon-th/template-fiber-api/models"
	"github.com/attapon-th/template-fiber-api/pkg"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// NewServices initialize global database and configuration
func NewServices() {
	// test connect database
	db := pkg.ConnectPostgreSQL(viper.GetString("DB_DSN"))

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Failed to Ping to database")
	}
}

func MigrateDatabase() error {
	db := pkg.ConnectPostgreSQL(viper.GetString("DB_DSN"))
	db = db.Debug()
	return db.AutoMigrate(
		&models.Todo{},
	)
}
