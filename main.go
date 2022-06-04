package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/rs/zerolog"
)

// will be overwritten in release pipeline
var Version = "dev"

func main() {
	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("go-cli-skeleton"),
		kong.Description("A minimal skeleton for building CLI apps"),
		kong.UsageOnError())
	GL = NewLogger(zerolog.InfoLevel)
	if cli.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	if !cli.JsonOutput {
		*GL.logger = GL.logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	err := ctx.Run(&cli)
	if err != nil {
		GL.logger.Fatal().Err(err).Msg("Shit is on fire")
	}
}

func sayHello(pw, addr string) error {
	GL.logger.Info().Msg("Hello")
	GL.logger.Debug().Str("PW", pw).Str("Addr", addr).Msg("You are curious!")
	return nil
}
