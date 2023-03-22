package logger

import (
	"fmt"
	"log"
	"os"
)

type AppLogger struct {
	logger *log.Logger
  
}

func NewAppLogger(appName string) *AppLogger {
	return &AppLogger{
		logger: log.New(
			os.Stdout,
			fmt.Sprintf("[%s]", appName),
			log.Lmsgprefix|log.LstdFlags,
		),
	}
}
