package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aracki/cgccli/api"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type JsonResponse struct {
	Href  string `json:"href"`
	Items []File `json:"items"`
}

type File struct {
	Href    string `json:"href"`
	Id      string `json:"id"`
	Name    string `json:"name"`
	Parent  string `json:"parent"`
	Project string `json:"project"`
	Type    string `json:"type"`
}

func GetFiles(project string) (files []File, err error) {

	client := &http.Client{}

	method := "GET"
	url := api.UrlFiles + "?project=" + project

	req, _ := http.NewRequest(method, url, nil)
	req.Header.Set(api.TokenHeader, viper.GetString("token"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("%s method on %s returns %x instead of 200", method, url, resp.StatusCode))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsonResp JsonResponse
	err = json.Unmarshal(respBody, &jsonResp)
	if err != nil {
		return nil, err
	}

	return jsonResp.Items, nil
}
