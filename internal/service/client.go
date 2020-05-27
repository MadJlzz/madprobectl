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

// ProbeDetails represents the data structure
// of a probe after it has been unmarshal from JSON.
type ProbeDetails struct {
	Name   string
	URL    string
	Status string
	Delay  uint
}

// GetWithParam is a convenient way to apply a GET request to a specified endpoint with the given urlParameters.
func GetWithParam(endpoint endpoints.Endpoint, urlParameters map[string]string) (ProbeDetails, error) {
	var pd ProbeDetails

	url := endpoint.GetURLWithParam(urlParameters)
	resp, err := http.Get(url)
	if err != nil {
		return pd, fmt.Errorf("error occured when doing the request. %v\n", err)
	}
	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return pd, fmt.Errorf("could not read response body. %v\n", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return pd, fmt.Errorf("server returned an error: %v. status code %d\n", strings.TrimSpace(string(bodyData)), resp.StatusCode)
	}

	err = unmarshalToStruct(bodyData, &pd)
	if err != nil {
		return pd, err
	}

	return pd, nil
}

// GetAll is a convenient way to apply a GET request to a specified endpoint and retrieve a collection of probes.
func GetAll(endpoint endpoints.Endpoint) ([]ProbeDetails, error) {
	var pds []ProbeDetails

	url := endpoint.GetURL()
	resp, err := http.Get(url)
	if err != nil {
		return pds, fmt.Errorf("error occured when doing the request. %v\n", err)
	}
	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return pds, fmt.Errorf("could not read response body. %v\n", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return pds, fmt.Errorf("server returned an error: %v. status code %d\n", strings.TrimSpace(string(bodyData)), resp.StatusCode)
	}

	err = unmarshalToStruct(bodyData, &pds)
	if err != nil {
		return pds, err
	}

	return pds, nil
}

// DeleteWithParam is a convenient way to apply a DELETE request to a specified endpoint with the given urlParameters.
func DeleteWithParam(endpoint endpoints.Endpoint, urlParameters map[string]string) error {
	client := http.DefaultClient
	url := endpoint.GetURLWithParam(urlParameters)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("error occured when creating the request. %v\n", err)
	}

	resp, err := client.Do(req)
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

	fmt.Println(string(bodyData))
	return nil
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

	fmt.Println(string(bodyData))
	return nil
}
