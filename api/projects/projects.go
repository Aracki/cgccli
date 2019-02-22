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

// GetProjects will list the projects owned by and accessible to a particular user.
// Each project's ID and URL will be returned.
func GetProjects() (projects []Project, err error) {

	respBody, err := api.CGCRequestBody("GET", api.UrlProjects, nil)
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
