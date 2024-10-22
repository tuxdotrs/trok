/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package config

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		NoColor:    false,
	}

	log.Logger = zerolog.New(output).
		With().
		Timestamp().
		Caller().
		Stack().
		Logger()

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}
