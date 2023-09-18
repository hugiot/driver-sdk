package providers

import (
	"github.com/hugiot/driver-sdk/common"
	"github.com/hugiot/driver-sdk/services/config"
	"github.com/hugiot/driver-sdk/services/http"
	"github.com/hugiot/driver-sdk/services/logger"
	"github.com/hugiot/driver-sdk/services/logwriter"
	"github.com/hugiot/gioc"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
)

type AppServiceProvider struct {
}

func (sp *AppServiceProvider) Register(c gioc.ServiceContainer) {
	// config
	c.Single(common.ConfigService, func(sc gioc.ServiceContainer) any {
		return config.New()
	})

	// log writer
	c.Single(common.LogWriterService, func(sc gioc.ServiceContainer) any {
		return logwriter.New()
	})

	// logger
	c.Single(common.LoggerService, func(sc gioc.ServiceContainer) any {
		w := sc.Make(common.LogWriterService).(*lumberjack.Logger)
		conf := sc.Make(common.ConfigService).(*viper.Viper)
		return logger.New(w, conf)
	})

	// http
	c.Single(common.HttpService, func(sc gioc.ServiceContainer) any {
		l := sc.Make(common.LoggerService).(*zap.Logger)
		return http.New(l)
	})
}

func (sp *AppServiceProvider) Boot(c gioc.ServiceContainer) {

}
