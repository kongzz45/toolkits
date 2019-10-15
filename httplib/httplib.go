package httplib

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// GetContent get http content
func GetContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return json.Marshal(resp.Body)
}

// PostContent http post
func PostContent(url string, content interface{}) ([]byte, error) {
	client := &http.Client{}
	cBytes, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(cBytes)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return json.Marshal(resp.Body)
}
