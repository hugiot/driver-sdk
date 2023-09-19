package providers

import (
	"github.com/hugiot/driver-sdk/common"
	"github.com/hugiot/driver-sdk/services/config"
	"github.com/hugiot/driver-sdk/services/fwriter"
	"github.com/hugiot/driver-sdk/services/http"
	"github.com/hugiot/driver-sdk/services/logger"
	"github.com/hugiot/gioc"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io"
)

type AppServiceProvider struct {
}

func (sp *AppServiceProvider) Register(c gioc.ServiceContainer) {
	// config
	c.Single(common.ConfigService, func(sc gioc.ServiceContainer) any {
		return config.New()
	})

	// file writer
	c.Single(common.FileWriterService, func(sc gioc.ServiceContainer) any {
		return fwriter.New()
	})

	// logger
	c.Single(common.LoggerService, func(sc gioc.ServiceContainer) any {
		w := sc.Make(common.FileWriterService).(io.Writer)
		conf := sc.Make(common.ConfigService).(*viper.Viper)
		return logger.New(w, conf.GetString(common.LogLevelConfig))
	})

	// http
	c.Single(common.HttpService, func(sc gioc.ServiceContainer) any {
		l := sc.Make(common.LoggerService).(*zap.Logger)
		return http.New(l)
	})
}

func (sp *AppServiceProvider) Boot(c gioc.ServiceContainer) {

}
