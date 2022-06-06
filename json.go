package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/tidwall/gjson"
)

// generated with https://mholt.github.io/json-to-go/
type PostmanGETResponse struct {
	Args    interface{} `json:"args"`
	Headers struct {
		XForwardedProto string `json:"x-forwarded-proto"`
		XForwardedPort  string `json:"x-forwarded-port"`
		Host            string `json:"host"`
		XAmznTraceID    string `json:"x-amzn-trace-id"`
		UserAgent       string `json:"user-agent"`
		Accept          string `json:"accept"`
	} `json:"headers"`
	URL string `json:"url"`
}

func parsePostmanGETResponse() error {
	resp, err := get("/get?foo1=bar1&foo2=bar2")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var postmanGETResponse PostmanGETResponse
	if err := json.Unmarshal(buf, &postmanGETResponse); err != nil {
		return err
	}
	fmt.Printf("Args (via struct) are: %v\n", postmanGETResponse.Args)
	return nil
}

func parseDynamicJSON() error {
	resp, err := get("/get?foo1=bar1&foo2=bar2")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Args (via gjson) are: %s\n", getJSONKey(string(b), "args"))
	return nil
}

func getJSONKey(json string, key string) string {
	return gjson.Get(json, key).String()
}
