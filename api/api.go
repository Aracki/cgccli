package api

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	TokenHeader = "X-SBG-Auth-Token"

	UrlBase = "https://cgc-api.sbgenomics.com/v2"

	UrlProjects = UrlBase + "/projects"
	UrlFiles    = UrlBase + "/files"
)

// CGCRequest will make a http request based on method, url & body.
// It will return http response only if it's 200
func CGCRequest(method string, url string, body io.Reader) (resp *http.Response, err error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "cannot initialize *Request")
	}

	req.Header.Set(TokenHeader, viper.GetString("token"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failure to speak HTTP / network connectivity problem")
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read body")
		}
		return nil, errors.Wrap(errors.New(string(b)), resp.Status)
	}

	return resp, nil
}

func CGCRequestBody(method string, url string, body io.Reader) (respBody []byte, err error) {

	resp, err := CGCRequest(method, url, body)
	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func CGCRequestBodyTotalOffset(method string, url string, body io.Reader) (bytesResp []byte, totalOffset int, err error) {

	resp, err := CGCRequest(method, url, body)
	defer resp.Body.Close()

	totalOffset, err = strconv.Atoi(resp.Header.Get("X-Total-Matching-Query"))
	if err != nil {
		return nil, 0, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return respBody, totalOffset, nil
}
