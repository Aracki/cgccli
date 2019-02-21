package projects

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
	Href  string    `json:"href"`
	Items []Project `json:"items"`
}

type Project struct {
	Href string `json:"href"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetProjects() (projects []Project, err error) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", api.UrlProjects, nil)
	req.Header.Set(api.TokenHeader, viper.GetString("token"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("GET request on %s returns %x instead of 200", api.UrlProjects, resp.StatusCode))
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
