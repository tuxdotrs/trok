/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package client

import (
	"github.com/rs/zerolog/log"
)

func Start(port uint16) {
	log.Info().Msgf("Hello from client: %d", port)
}
