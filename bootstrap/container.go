package bootstrap

import (
	"github.com/hugiot/driver-sdk/providers"
	"github.com/hugiot/gioc"
)

func InitContainer() {
	gioc.AddServerProvider(&providers.AppServiceProvider{})
	gioc.Boot()
}
