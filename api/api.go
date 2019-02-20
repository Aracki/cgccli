package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

type ProjectsResponse struct {
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
	req, _ := http.NewRequest("GET", "https://cgc-api.sbgenomics.com/v2/projects", nil)
	req.Header.Set("X-SBG-Auth-Token", viper.GetString("token"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}

	var jsonResp ProjectsResponse
	err = json.Unmarshal(respBody, &jsonResp)
	if err != nil {
		return nil, err
	}

	return  jsonResp.Items, nil
}
