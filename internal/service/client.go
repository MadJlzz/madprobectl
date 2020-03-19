// Package service provides a simple way to do HTTP(s) calls to the remote server.
package service

import (
	"fmt"
	"github.com/madjlzz/madprobectl/internal/endpoints"
	"github.com/madjlzz/madprobectl/internal/parser"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	jsonMimeType = "application/json"
)

// CreateProbeRequest represents the data structure
// encoded in JSON just before doing the create HTTP request.
type CreateProbeRequest struct {
	Name  string
	URL   string
	Delay uint
}

// Post is a convenient way to apply a POST request to a specified endpoint.
func Post(endpoint endpoints.Endpoint, yamlProbe parser.HttpProbe) error {
	// TODO: manage the creation of all probes set in the yaml. Maybe new endpoint ?
	// TODO: pass CreateProbeRequest as a parameter so we can pass whatever payload we want.
	cpr := CreateProbeRequest{
		Name:  yamlProbe.Http[0].Name,
		URL:   yamlProbe.Http[0].URL,
		Delay: yamlProbe.Http[0].Delay,
	}
	jsonCpr, err := marshalToJSON(cpr)
	if err != nil {
		return err
	}

	url := endpoint.GetURL()
	resp, err := http.Post(url, jsonMimeType, strings.NewReader(jsonCpr))
	if err != nil {
		return fmt.Errorf("error occured when doing the request. %v\n", err)
	}
	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body. %v\n", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("server returned an error: %v. status code %d\n", strings.TrimSpace(string(bodyData)), resp.StatusCode)
	}

	return nil
}
