package api

import (
	"errors"
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

func CGCRequest(method string, url string, body io.Reader) (bytesResp []byte, err error) {

	req, err := http.NewRequest(method, url, body)
	req.Header.Set(TokenHeader, viper.GetString("token"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// If not OK
	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(string(b))
	}

	// If OK
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func CGCRequestFiles(method string, url string, body io.Reader) (bytesResp []byte, totalOffset int, err error) {

	req, err := http.NewRequest(method, url, body)
	req.Header.Set(TokenHeader, viper.GetString("token"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	// If not OK
	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		return nil, 0, errors.New(string(b))
	}

	// If OK
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	totalOffset, err = strconv.Atoi(resp.Header.Get("X-Total-Matching-Query"))
	if err != nil {
		return nil, 0, err
	}
	return respBody, totalOffset, nil
}
