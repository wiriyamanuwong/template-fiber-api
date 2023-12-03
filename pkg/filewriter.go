package pkg

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog/diode"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/utahta/go-cronowriter"
)

// CronFileWriter create log file The file path is constructed based on current date and time, like cronolog.
func CronFileWriter(filename string) *cronowriter.CronoWriter {
	logDir := viper.GetString("log_dir")
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		log.Fatal().Str("dir", path.Dir(filename)).Msg(err.Error())
		return nil
	}
	filename = path.Join(logDir, filename)
	ext := path.Ext(filename)
	filename = filename[0:len(filename)-len(ext)] + "-%Y%m%d" + ext

	return cronowriter.MustNew(filename, cronowriter.WithInit(), cronowriter.WithMutex())
}

// NewDiodeWriter Thread-safe, lock-free, non-blocking writer
func NewDiodeWriter(w io.WriteCloser) io.WriteCloser {
	wr := diode.NewWriter(w, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages", missed)
	})
	return wr
}

// NewDiodeCronWriter  logfile rotation and Thread-safe, lock-free, non-blocking writer
func NewDiodeCronWriter(filename string) io.WriteCloser {
	w := CronFileWriter(filename)
	wr := NewDiodeWriter(w)
	return wr
}
