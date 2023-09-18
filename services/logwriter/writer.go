package logwriter

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

const FileName string = "driver.log"

func New() io.Writer {
	return &lumberjack.Logger{
		Filename:   FileName,
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 0,
		LocalTime:  true,
		Compress:   false,
	}
}
