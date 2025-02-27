/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package client

import (
	"fmt"
	"io"
	"net"

	"github.com/0xtux/trok/internal/lib"
	"github.com/rs/zerolog/log"
)

type Trok struct {
	controlClient TCPClient
}

func (t *Trok) Init(port uint16) error {
	err := t.controlClient.Init(port, "Controller")
	return err
}

func (t *Trok) Start() {
	go t.controlClient.Start(t.ControlConnHandler)
	log.Info().Msgf("started Trok client on port %d", t.controlClient.Port())
}

func (t *Trok) Stop() {
	t.controlClient.Stop()
	log.Info().Msgf("stopped Trok client on port %d", t.controlClient.Port())
}

func (t *Trok) ControlConnHandler(conn net.Conn) {
	p := lib.InitProtocolHandler(conn)

	p.WriteMessage(&lib.Message{CMD: "HELO", ARG: "Trok"})

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

		case "EHLO":
			go t.hanldeCMDEHLO(m)

		case "CNCT":
			go t.handleCMDCNCT(m)

		default:
			log.Info().Msgf("invalid command")
		}
	}
}

func (t *Trok) hanldeCMDEHLO(m *lib.Message) {
	log.Info().Msgf("[CMD] %s [ARG] %s", m.CMD, m.ARG)
}

func (t *Trok) handleCMDCNCT(m *lib.Message) {
	log.Info().Msgf("[CMD] %s [ARG] %s", m.CMD, m.ARG)

	var upstream TCPClient
	var downstream TCPClient

	err := upstream.Init(3000, "UpStream")
	if err != nil {
		log.Error().Msgf("can't connect to upstream socket: %v", err)
		return
	}

	err = downstream.Init(1421, "DownStream")
	if err != nil {
		log.Error().Msgf("can't connect to downstream socket: %v", err)
		return
	}

	downstream.conn.Write([]byte(fmt.Sprintf("ACPT %s\n", m.ARG)))
	go io.Copy(upstream.conn, downstream.conn)
	io.Copy(downstream.conn, upstream.conn)
}
