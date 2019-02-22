package files

import (
	"encoding/json"
	"fmt"
	"github.com/aracki/cgccli/api"
	"github.com/pkg/errors"
	"strconv"
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
	Tags     []string          `json:"tags"`
	Metadata map[string]string `json:"metadata"`
}

// GetFiles will get first 50 (default limit) files for the given project. The totalOffset is obtained from X-Total-Matching-Query header.
// Based on that value we know how many times to make a new GET request.
// On each GET request we get []byte which we unmarshal to array of files.
// All file arrays gets merged into one which this function returns.
func GetFiles(project string) (files []File, err error) {

	// if omitted; default is 50
	// The maximum value for the query parameter limit is 100.
	limit := 100
	limitS := strconv.Itoa(limit)

	url := api.UrlFiles + "?limit=" + limitS + "&project=" + project
	respBody, totalOffset, err := api.CGCRequestFiles("GET", url, nil)
	if err != nil {
		return nil, err
	}
	jsonResp := JsonResponse{}
	err = json.Unmarshal(respBody, &jsonResp)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s", "unmarshaling failed"))
	}

	files = append(files, jsonResp.Items...)

	blocksNum := totalOffset / limit
	for i := 1; i <= blocksNum; i++ {
		offset := strconv.Itoa(limit * i)

		url := api.UrlFiles + "?offset=" + offset + "&limit=" + limitS + "&project=" + project
		respBody, _, err := api.CGCRequestFiles("GET", url, nil)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprint("CGCRequestFiles failed"))
		}

		jsonResp = JsonResponse{}
		err = json.Unmarshal(respBody, &jsonResp)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("unmarshaling failed for block num %d", i))
		}

		files = append(files, jsonResp.Items...)
	}
	return files, nil
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
