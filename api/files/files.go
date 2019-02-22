package files

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aracki/cgccli/api"
	"github.com/pkg/errors"
	"net/url"
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

type FileDetailsMap map[string]interface{}
type FileDetailsMetadataMap map[string]string

// GetFiles will get first 1-100 (limit) files for the given project.
// The totalOffset is obtained from X-Total-Matching-Query header.
// Based on that value we know how many times to make a new GET request.
// On each GET request we get []byte which we unmarshal to array of files.
// All file arrays gets merged into one which this function returns.
func GetFiles(project string) (files []File, err error) {

	// If omitted, default is 50
	// The minimum value for the query parameter limit is 1.
	// The maximum value for the query parameter limit is 100.
	limit := 100

	url, _ := url.Parse(api.UrlFiles)
	q := url.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("project", project)
	url.RawQuery = q.Encode()

	respBody, totalOffset, err := api.CGCRequestBodyTotalOffset("GET", url.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "CGCRequestBodyTotalOffset request failed")
	}
	jsonResp := JsonResponse{}
	err = json.Unmarshal(respBody, &jsonResp)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshaling CGCRequestBodyTotalOffset response failed")
	}

	files = append(files, jsonResp.Items...)

	blocksNum := totalOffset / limit
	for i := 1; i <= blocksNum; i++ {
		offset := strconv.Itoa(limit * i)

		q.Set("offset", offset)
		url.RawQuery = q.Encode()
		respBody, _, err := api.CGCRequestBodyTotalOffset("GET", url.String(), nil)
		if err != nil {
			return nil, errors.Wrap(err, "CGCRequestFiles request failed")
		}

		jsonResp = JsonResponse{}
		err = json.Unmarshal(respBody, &jsonResp)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("unmarshaling CGCRequestFiles response failed for block num %d", i))
		}

		files = append(files, jsonResp.Items...)
	}
	return files, nil
}

// GetFileDetails will get all the details for that fileId.
func GetFileDetails(fileId string) (fDetails *FileDetails, err error) {

	respBody, err := api.CGCRequestBody("GET", api.UrlFiles+"/"+fileId, nil)
	if err != nil {
		return nil, errors.Wrap(err, "CGCRequestBody request failed")
	}

	err = json.Unmarshal(respBody, &fDetails)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshaling CGCRequestBody response failed")
	}

	return fDetails, nil
}

// UpdateFileDetails will update file details for that fileId.
func UpdateFileDetails(fileId string, fdMap FileDetailsMap) (respBody []byte, err error) {

	urlFile := api.UrlFiles + "/" + fileId
	urlFileMetadata := api.UrlFiles + "/" + fileId + "/metadata"

	if _, ok := fdMap["metadata"]; ok {

		jsonMetadata, err := json.Marshal(fdMap["metadata"])
		if err != nil {
			return nil, errors.Wrap(err, "marshaling metadata map failed")
		}
		_, err = api.CGCRequestBody("PATCH", urlFileMetadata, bytes.NewBuffer(jsonMetadata))
		if err != nil {
			return nil, errors.Wrap(err, "CGCRequestBody PATCH metadata failed")
		}
	}

	delete(fdMap, "metadata")

	jsonBody, err := json.Marshal(fdMap)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling file details map failed")
	}

	respBody, err = api.CGCRequestBody("PATCH", urlFile, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, errors.Wrap(err, "CGCRequestBody PATCH file details failed")
	}

	return respBody, nil
}
