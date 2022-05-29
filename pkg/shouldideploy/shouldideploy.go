package shouldideploy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var API_ADDRESS = "https://shouldideploy.today/api?tz=Brazil%2FAcre"

type ShouldIDeployResponse struct {
	Message       string `json:"message"`
	ShouldIdeploy bool   `json:"shouldideploy"`
	Timezone      string `json:"timezone"`
}

func Request() (*ShouldIDeployResponse, error) {
	response := &ShouldIDeployResponse{}

	resp, err := http.Get(API_ADDRESS)

	fmt.Println(resp)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	fmt.Println(body)
	err = json.Unmarshal(body, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
