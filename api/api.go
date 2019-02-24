// Package api provides different types of functions for making http request and reading responses.
package api

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	HeaderToken                = "X-SBG-Auth-Token"
	HeaderMaxOffset            = "X-Total-Matching-Query"
	HeaderJSONContentTypeKey   = "Content-Type"
	HeaderJSONContentTypeValue = "application/json"

	UrlBase     = "https://cgc-api.sbgenomics.com/v2"
	UrlProjects = UrlBase + "/projects"
	UrlFiles    = UrlBase + "/files"
)

// CGCRequest will make a http request based on method, url & body.
// It will return http response only if it's 200.
// If it is not 200, error with body message will be returned
func CGCRequest(method string, url string, body io.Reader) (resp *http.Response, err error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "cannot initialize *Request")
	}

	// add mandatory headers
	req.Header.Set(HeaderToken, viper.GetString("token"))
	req.Header.Set(HeaderJSONContentTypeKey, HeaderJSONContentTypeValue)

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failure to speak HTTP / network connectivity problem")
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read body")
		}
		return nil, errors.Wrap(errors.New(string(b)), resp.Status)
	}

	return resp, nil
}

// CGCRequestAndRead is wrapper around CGCRequest.
// Read the response body until EOF.
func CGCRequestAndRead(method string, url string, body io.Reader) (respBody []byte, err error) {

	resp, err := CGCRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// CGCRequestAndReadTotalOffset is wrapper around CGCRequest.
// Read the response body until EOF.
// Returns X-Total-Matching-Query header value.
func CGCRequestAndReadTotalOffset(method string, url string, body io.Reader) (bytesResp []byte, totalOffset int, err error) {

	resp, err := CGCRequest(method, url, body)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	totalOffset, err = strconv.Atoi(resp.Header.Get(HeaderMaxOffset))
	if err != nil {
		return nil, 0, errors.Wrap(err, fmt.Sprintf("%s header missing", HeaderMaxOffset))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return respBody, totalOffset, nil
}
