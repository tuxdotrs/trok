/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package server

import (
	"fmt"
	"io"
	"net"

	"github.com/0xtux/trok/internal/lib"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/rs/zerolog/log"
)

type Trok struct {
	controlServer TCPServer
	publicConns   map[string]net.Conn
	tunnels       map[uint16]*lib.ProtocolHandler
}

func (t *Trok) Init(port uint16) error {
	t.publicConns = make(map[string]net.Conn)
	t.tunnels = make(map[uint16]*lib.ProtocolHandler)
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
			t.handleCMDACPT(conn, m)

		default:
			log.Info().Msgf("invalid command")
		}
	}
}

func (t *Trok) handleCMDHELO(p *lib.ProtocolHandler, m *lib.Message) {
	log.Info().Msgf("[CMD] %s [ARG] %s", m.CMD, m.ARG)
	var s TCPServer

	err := s.Init(0, "Handler")
	if err != nil {
		log.Error().Msgf("error handling HELO cmd: %v", err)
		return
	}

	port := s.Port()
	go s.Start(t.PublicConnHandler)
	t.tunnels[port] = p

	p.WriteMessage(&lib.Message{CMD: "EHLO", ARG: fmt.Sprintf("%d", port)})
}

func (t *Trok) handleCMDACPT(conn net.Conn, m *lib.Message) {
	log.Info().Msgf("[CMD] %s [ARG] %s", m.CMD, m.ARG)

	pConn, ok := t.publicConns[m.ARG]
	if !ok {
		log.Error().Msgf("error finding public connection")
	}

	go io.Copy(pConn, conn)
	io.Copy(conn, pConn)
}

func (t *Trok) PublicConnHandler(conn net.Conn) {
	id, err := gonanoid.New(12)
	if err != nil {
		log.Error().Msgf("error generating uid: %v", err)
		return
	}

	port := uint16(conn.LocalAddr().(*net.TCPAddr).Port)
	tnl, ok := t.tunnels[port]
	if !ok {
		log.Error().Msgf("error finding tunnel connection")
		return
	}
	tnl.WriteMessage(&lib.Message{CMD: "CNCT", ARG: id})

	t.publicConns[id] = conn
}
