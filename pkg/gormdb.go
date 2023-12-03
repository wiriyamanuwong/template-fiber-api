package pkg

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/attapon-th/template-fiber-api/helper"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbs map[string]*gorm.DB = make(map[string]*gorm.DB)

	defaultGormLoggerConfig = logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	}
)

// ConnectPostgreSQL connect postgresql database
func ConnectPostgreSQL(dsn string, cfgs ...*gorm.Config) *gorm.DB {

	h := sha256.New224()
	h.Write(helper.S2B(dsn))
	hStr := fmt.Sprintf("%x", h.Sum(nil))
	log.Debug().Str("dsn", dsn).Str("hash", hStr).Send()

	if db, ok := dbs[hStr]; ok && db != nil {
		return db
	}

	var cfg *gorm.Config
	if len(cfgs) == 0 {
		cfg = &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			DisableNestedTransaction:                 true,
			Logger:                                   getGormLogger(),
		}
	} else {
		cfg = cfgs[0]
	}

	db, err := gorm.Open(postgres.Open(dsn), cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		return nil
	}
	dbs[hStr] = db
	return db

}

func getGormLogger() logger.Interface {
	logSQL := viper.GetString("log_sql")
	var l logger.Interface
	if logSQL == "console" {
		l = logger.New(newGormLoggerConsoleWriter(), defaultGormLoggerConfig)
	} else if strings.HasSuffix(logSQL, ".log") {
		cfg := defaultGormLoggerConfig
		cfg.Colorful = false
		cfg.LogLevel = logger.Info
		l = logger.New(newGormLoggerFileWriter(logSQL), cfg)
	} else {
		l = logger.Discard
	}
	return l
}

type gormLoggerWriter struct {
	log  zerolog.Logger
	mode string
}

func (g *gormLoggerWriter) Printf(s string, a ...any) {
	g.log.Log().Str("logger", "gorm").Str("mode", g.mode).Msgf(s, a...)
}

func newGormLoggerConsoleWriter() *gormLoggerWriter {
	cs := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: zerolog.TimeFieldFormat,
		NoColor:    false,
	}
	return &gormLoggerWriter{
		log:  zerolog.New(cs).With().Timestamp().Logger(),
		mode: "console",
	}
}

func newGormLoggerFileWriter(filename string) *gormLoggerWriter {
	cs := NewDiodeCronWriter(filename)
	return &gormLoggerWriter{
		log:  zerolog.New(cs).With().Timestamp().Logger(),
		mode: "filemode",
	}
}
