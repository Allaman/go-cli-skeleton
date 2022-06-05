# go-cli-skeleton

A skeleton for writing Go CLI applications.

## What's included

- global logging via [zerolog](https://github.com/rs/zerolog) (json and human friendly)
- simple HTTP client via `net/http` package returning dynamic JSON
- reading files
- writing files
- reading from stdin (e.g. Linux pipe)
- CLI commands, parameters and flags via [Kong](https://github.com/alecthomas/kong)
- flat file structure and one package

## What's not included

- sophisticated HTTP client
- modular file structure ([project-layout](https://github.com/golang-standards/project-layout))
- tests 😬

```sh
go run . -h
Usage: go-cli-skeleton <command>

A minimal skeleton for building CLI apps

Flags:
  -h, --help           Show context-sensitive help.
  -d, --debug          Enable debug output
  -j, --json-output    Log in json format

Commands:
  hello
    default command

  get
    GET query to postman echo

  post
    POST query to postman echo

  version
    Show version information

  file read --path=STRING
    Read a file

  file write --path=STRING
    Write a file

Run "go-cli-skeleton <command> --help" for more information on a command.
```
