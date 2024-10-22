/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package server

import (
	"bufio"
	"net"

	"github.com/rs/zerolog/log"
)

type Trok struct {
	controlServer TCPServer
}

func (t *Trok) Init(port uint16) error {
	err := t.controlServer.Init(port, "Controller")
	return err
}

func (t *Trok) Start() {
	go t.controlServer.Start(t.ControlConnHandler)
	log.Info().Msgf("started Trok server on port %d", t.controlServer.Port())
}

func (t *Trok) Stop() {
	t.controlServer.Stop()
	log.Info().Msgf("stopped Trok server on port %d", t.controlServer.Port())
}

func (t *Trok) ControlConnHandler(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				log.Warn().Msgf("connection timed out: %s", conn.RemoteAddr())
			} else {
				log.Logger.Info().Msgf("connection closed: %s", err)
			}
			return
		}

		log.Info().Msgf(data)
	}
}
