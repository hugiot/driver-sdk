package interfaces

import (
	"github.com/hugiot/driver-sdk/models"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Driver interface {
	// Initialize init driver
	Initialize(config *viper.Viper, logger *zap.Logger, asyncCh chan<- string, deviceCh chan<- []models.Device) error
	// HandleRead handle read resource
	HandleRead(sn string, protocol map[string]string)
	// HandleWrite handle write resource
	HandleWrite(sn string, protocol map[string]string)
	// HandleCommand handle command
	HandleCommand(sn string, protocol map[string]string, command any)
}
