package fwriter

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

var defaultFileName = "driver.log"

func New() io.Writer {
	return &lumberjack.Logger{
		Filename:   defaultFileName,
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 0,
		LocalTime:  true,
		Compress:   false,
	}
}
