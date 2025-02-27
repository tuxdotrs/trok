/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package client

import (
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/tuxdotrs/trok/internal/lib"
)

type Trok struct {
	controlClient *TCPClient
	serverAddr    string
	localAddr     string
}

func NewTrokClient(serverAddr, localAddr string) (*Trok, error) {
	controlClient, err := NewTCPClient(serverAddr, "Controller")

	return &Trok{
		controlClient: controlClient,
		serverAddr:    serverAddr,
		localAddr:     localAddr,
	}, err
}

func (t *Trok) Start() {
	go t.controlClient.Start(t.ControlConnHandler)
	parts := strings.Split(t.serverAddr, ":")
	log.Info().Msgf("started Trok client on %s", parts[0])
}

func (t *Trok) Stop() {
	t.controlClient.Stop()
	parts := strings.Split(t.serverAddr, ":")
	log.Info().Msgf("stopped Trok client on %s", parts[0])
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
	parts := strings.Split(t.serverAddr, ":")
	log.Info().Msgf("[CMD] %s [ARG] %s:%s", m.CMD, parts[0], m.ARG)
}

func (t *Trok) handleCMDCNCT(m *lib.Message) {
	log.Info().Msgf("[CMD] %s [ARG] %s", m.CMD, m.ARG)

	upStream, err := NewTCPClient(t.localAddr, "UpStream")
	if err != nil {
		log.Error().Msgf("can't connect to upstream socket: %v", err)
		return
	}

	downStream, err := NewTCPClient(t.serverAddr, "DownStream")
	if err != nil {
		log.Error().Msgf("can't connect to downstream socket: %v", err)
		return
	}

	downStream.conn.Write([]byte(fmt.Sprintf("ACPT %s\n", m.ARG)))
	go io.Copy(upStream.conn, downStream.conn)
	io.Copy(downStream.conn, upStream.conn)
}
