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
	resp, err := c.Get(fmt.Sprintf("%s%s", baseURL, path))
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
	resp, err := c.Post(fmt.Sprintf("%s%s", baseURL, path), "application/json", bytes.NewBuffer(jsonData))
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
