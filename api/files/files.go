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

type FileDetails struct {
	Href       string `json:"href"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	Size       int    `json:"size"`
	Project    string `json:"project"`
	CreatedOn  string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
	Storage    struct {
		Type     string `json:"type"`
		Volume   string `json:"volume"`
		Location string `json:"location"`
	} `json:"storage"`
	Origin struct {
		Dataset string `json:"dataset"`
	} `json:"origin"`
	Tags []string `json:"tags"`
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

func GetFileDetails(fileId string) (fDetails *FileDetails, err error) {

	respBody, err := api.CGCRequest("GET", api.UrlFiles+"/"+fileId, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &fDetails)
	if err != nil {
		return nil, err
	}

	return fDetails, nil
}
