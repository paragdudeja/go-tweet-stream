package rules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Data []struct {
		Id string `json:"id"`
	} `json:"data"`
}

func Bearer() string {
	bearer := fmt.Sprintf("Bearer %s", os.Getenv("API_BEARER"))
	return bearer
}

func GetRules() []string {
	url := "https://api.twitter.com/2/tweets/search/stream/rules"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	var ids []string

	if err != nil {
		fmt.Println(err)
		return ids
	}
	req.Header.Add("Authorization", Bearer())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ids
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ids
	}
	defer res.Body.Close()

	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	for _, rec := range result.Data {
		ids = append(ids, rec.Id)
	}

	return ids
}

func DeleteRules(ids []string) {
	url := "https://api.twitter.com/2/tweets/search/stream/rules"
	method := "POST"

	var payload = fmt.Sprintf(`{
		"delete": {
			"ids": ["%s"]
		}
	}`, strings.Join(ids, `", "`))

	var jsonData = []byte(payload)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", Bearer())
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = res
}

func AddRules(query string) {
	url := "https://api.twitter.com/2/tweets/search/stream/rules"
	method := "POST"

	payload := fmt.Sprintf(`{
		"add": [
			{
				"value": "%s"
			}
		]
	}`, query)

	var jsonData = []byte(payload)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", Bearer())
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = res
}
