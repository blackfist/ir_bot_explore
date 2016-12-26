package github

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

const url string = "https://api.github.com"

func New(name string, desc string) string {
	// POST /orgs/:org/repos
	// name string
	// description string
	// private boolean
	// curl -H "Authorization: token 0b6b166138a3c432043a5afce96d3c78c4374381" \
	// -H "Content-Type: application/json" \
	// -X POST -d '{"name":"twilight-shape-1279","description":"a simple incident to manage","private":true}' \
	// https://api.github.com/blackfist/repos

	// set up the post data and a struct to receive
	// the response
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Private     bool   `json:"private"`
	}

	type Response struct {
		Url string `json:"html_url"`
	}

	the_request, _ := json.Marshal(request{Name: name, Description: desc, Private: true})

	req, _ := http.NewRequest("POST", url+"/user/repos", bytes.NewBuffer(the_request))
	req.Header.Set("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var answer Response
	json.Unmarshal(body, &answer)

	return answer.Url
}
