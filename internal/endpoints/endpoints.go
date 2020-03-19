// Package endpoints provides variables exposing access to various services.
package endpoints

import (
	"github.com/spf13/viper"
	"strings"
)

// Endpoint used to create a new probe.
var CreateProbe = Endpoint{
	url: "/api/v1/probe/create",
}

// Endpoint used to retrieve one probe.
var GetProbe = Endpoint{
	url: "/api/v1/probe/{name}",
}

var FindAllProbe = Endpoint{
	url: "/api/v1/probe",
}

var DeleteProbe = Endpoint{
	url: "/api/v1/probe/{name}",
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

// GetURLWithParam is a function that replaces any urlParameter
// existing within the endpoint by it's given value.
// Basically it constructs the final API URL.
func (e *Endpoint) GetURLWithParam(urlParameters map[string]string) string {
	baseURL := viper.GetString("server")
	apiURL := e.url
	for k, v := range urlParameters {
		apiURL = strings.ReplaceAll(apiURL, "{"+k+"}", v)
	}
	return baseURL + apiURL
}
