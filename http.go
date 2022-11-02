package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	baseURL = "https://postman-echo.com"
	c       = http.Client{Timeout: time.Duration(1) * time.Second}
)

func get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "foo")
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if !isHTTPResponseCode200(resp) {
		GL.logger.Warn().Msgf("statuscode was '%d'", resp.StatusCode)
	}
	return resp, nil
}

func post(path string, body interface{}) (*http.Response, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", baseURL, path), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if !isHTTPResponseCode200(resp) {
		GL.logger.Warn().Msgf("statuscode was '%d'", resp.StatusCode)
	}
	return resp, nil
}

func getJSONFromHTTPResponse(resp *http.Response) ([]byte, error) {
	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func printHTTPResponse(resp *http.Response) error {
	data, err := getJSONFromHTTPResponse(resp)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func isHTTPResponseCode200(resp *http.Response) bool {
	return resp.StatusCode == 200
}
