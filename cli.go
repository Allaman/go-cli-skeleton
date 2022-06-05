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
	File       struct {
		Read  readFileCmd  `cmd:"" help:"Read a file"`
		Write writeFileCmd `cmd:"" help:"Write a file"`
	} `cmd:"" help:"File operations"`
}

type versionCmd struct {
	Version string
}

func (c *versionCmd) Run() error {
	fmt.Println(Version)
	return nil
}

type readFileCmd struct {
	Path string `short:"p" required:"" help:"Path to the file to be read"`
}

func (c *readFileCmd) Run() error {
	return func() error {
		f, err := readFile(c.Path)
		if err != nil {
			return err
		}
		fmt.Println(string(f))
		return nil
	}()
}

type writeFileCmd struct {
	Path    string `short:"p" required:"" help:"Path to the file to be written"`
	Content string `short:"c" help:"Content to be written to the file (stdin also works)"`
}

func (c *writeFileCmd) Run() error {
	return writeBytesToFile([]byte(c.Content), c.Path)
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
