/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package server

import (
	"net"

	"github.com/0xtux/trok/internal/lib"
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
	p := lib.InitProtocolHandler(conn)

	for {
		m, err := p.ReadMessage()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				log.Warn().Msgf("connection timed out: %s", conn.RemoteAddr())
			} else {
				log.Logger.Info().Msgf("connection closed: %s", err)
			}
			return
		}

		switch m.CMD {

		case "HELO":
			t.handleCMDHELO(p, m)

		case "ACPT":
			t.handleCMDACPT(p, m)

		default:
			log.Info().Msgf("invalid command")
		}
	}
}

func (t *Trok) handleCMDHELO(p *lib.ProtocolHandler, m *lib.Message) {
	log.Info().Msgf("[CMD] %s [ARG] %s", m.CMD, m.ARG)
	p.WriteMessage(m)
}

func (t *Trok) handleCMDACPT(p *lib.ProtocolHandler, m *lib.Message) {
	log.Info().Msgf("[CMD] %s [ARG] %s", m.CMD, m.ARG)
	p.WriteMessage(m)
}
