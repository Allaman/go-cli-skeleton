package main

import (
	"fmt"
)

type CLI struct {
	Hello      helloCmd   `cmd:"" default:"1" help:"default command"`
	Get        httpGET    `cmd:"" help:"GET query to postman echo"`
	Post       httpPOST   `cmd:"" help:"POST query to postman echo"`
	Version    versionCmd `cmd:"" help:"Show version information"`
	Debug      bool       `short:"d" help:"Enable debug output"`
	JsonOutput bool       `short:"j" default:"false" help:"Log in json format"`
}

type versionCmd struct {
	Version string
}

func (c *versionCmd) Run() error {
	fmt.Println(Version)
	return nil
}

type helloCmd struct {
	Addr     string `default:"localhost:8080" short:"a" help:"An address string with no purpose"`
	Password string `default:"" short:"p" help:"Secret password with no purpose"`
}

func (c *helloCmd) Run() error {
	return sayHello(c.Password, c.Addr)
}

type httpGET struct{}

func (c *httpGET) Run() error {
	return func() error {
		resp, err := get("/get?foo1=bar1&foo2=bar2")
		if err != nil {
			return err
		}
		return printHTTPResponse(resp)
	}()
}

type httpPOST struct{}

func (c *httpPOST) Run() error {
	return func() error {
		resp, err := post("/post", map[string]string{"foo": "bar"})
		if err != nil {
			return err
		}
		return printHTTPResponse(resp)
	}()
}
