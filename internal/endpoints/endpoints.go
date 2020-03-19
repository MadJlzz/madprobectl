// Package endpoints provides variables exposing access to various services.
package endpoints

import (
	"github.com/spf13/viper"
)

// Endpoint used to create a new probe.
var CreateProbe = Endpoint{
	url: "/api/v1/probe/create",
}

// Endpoint represent a madprobe API endpoint.
type Endpoint struct {
	url string
}

// GetURL is a simple function joining the URL of the server and the API endpoint.
// baseURL is coming from the configuration file in $HOME/.madprobectl
func (e *Endpoint) GetURL() string {
	baseURL := viper.GetString("server")
	return baseURL + e.url
}
