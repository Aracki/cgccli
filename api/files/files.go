package files

import (
	"encoding/json"
	"github.com/aracki/cgccli/api"
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

	respBody, err := api.CGCRequest("GET", api.UrlFiles+"?project="+project, nil)
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
