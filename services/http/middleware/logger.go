package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type defaultOutput struct {
	l *zap.Logger
}

func (d *defaultOutput) Write(p []byte) (n int, err error) {
	d.l.Info(string(p))
	return len(p), nil
}

var defaultFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}
	return fmt.Sprintf("GIN | %s %3d %s| %13v | %15s |%s %-7s %s %#v | %s",
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

func Logger(l *zap.Logger) gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: defaultFormatter,
		Output:    &defaultOutput{l: l},
		SkipPaths: nil,
	})
}
