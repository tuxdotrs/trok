/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package server

import (
	"github.com/rs/zerolog/log"
)

func Start(port uint16) {
	log.Info().Msgf("Hello from server: %d", port)
}
