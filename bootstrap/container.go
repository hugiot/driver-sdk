package bootstrap

import (
	"github.com/hugiot/driver-sdk/providers"
	"github.com/hugiot/gioc"
)

// initContainer init ioc container
func initContainer() {
	// add providers
	gioc.AddServerProvider(&providers.AppServiceProvider{})

	// boot
	gioc.Boot()
}
