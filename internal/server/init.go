/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package server

import (
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
)

func Start(port uint16) {
	var trok Trok
	if err := trok.Init(port); err != nil {
		log.Fatal().Msgf("failed to init trok %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	trok.Start()
	defer trok.Stop()

	<-signalChan
}
