/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package client

import (
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
)

func Start(serverAddr, localAddr string) {
	trok, err := NewTrokClient(serverAddr, localAddr)
	if err != nil {
		log.Fatal().Msgf("failed init trok %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	trok.Start()
	defer trok.Stop()

	<-signalChan
}
