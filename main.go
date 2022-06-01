package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// will be overwritten in release pipeline
var Version = "dev"

func main() {
	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("go-cli-skeleton"),
		kong.Description("A minimal skeleton for building CLI apps"))
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cli.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	if !cli.JsonOutput {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	err := ctx.Run(&cli)
	if err != nil {
		log.Fatal().Err(err).Msg("Shit is on fire")
	}
}

func sayHello(pw, addr string) error {
	log.Info().Msg("Hello")
	log.Debug().Str("PW", pw).Str("Addr", addr).Msg("You are curious!")
	return nil
}
