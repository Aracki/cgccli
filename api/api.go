package api

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

func GetProjects() (projects string, err error) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://cgc-api.sbgenomics.com/v2/projects", nil)
	req.Header.Set("X-SBG-Auth-Token", viper.GetString("token"))
	resp, err := client.Do(req)
	if err != nil {
		return"", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("1")

	return fmt.Sprintf("%s", respBody), nil
}
