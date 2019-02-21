package projects

import (
	"encoding/json"
	"github.com/aracki/cgccli/api"
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

	respBody, err := api.CGCRequest("GET", api.UrlProjects, nil)
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
