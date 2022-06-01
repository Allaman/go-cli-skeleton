## go-cli-skeleton

A skeleton for writing Go CLI applications.

### What's included

- global logging via [zerolog](https://github.com/rs/zerolog) (json and human friendly)
- simple HTTP client with `GET` and `POST` functions via `net/http` package returning dynamic JSON
- CLI commands, parameters and flags via [Kong](https://github.com/alecthomas/kong)
- flat file structure and one package

### What's not included

- sophisticated HTTP client (with tracing for instance)
- file structure for more complex apps ([project-layout](https://github.com/golang-standards/project-layout))
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

Run "go-cli-skeleton <command> --help" for more information on a command.
```
