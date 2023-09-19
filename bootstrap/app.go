package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hugiot/driver-sdk/common"
	"github.com/hugiot/gioc"
	"github.com/spf13/viper"
)

// Start http service
func Start(driver any) error {
	initContainer()

	c := gioc.Make(common.ConfigService).(*viper.Viper)

	r := gioc.Make(common.HttpService).(*gin.Engine)
	addr := fmt.Sprintf("%s:%d", c.GetString(common.HttpBindConfig), c.GetInt(common.HttpPortConfig))
	return r.Run(addr)
}
