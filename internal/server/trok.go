/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package server

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/0xtux/trok/internal/lib"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/rs/zerolog/log"
)

type Conn struct {
	conn      net.Conn
	timestamp time.Time
}

type Trok struct {
	controlServer TCPServer
	publicConns   map[string]Conn
	mutex         sync.Mutex
}

func (t *Trok) Init(addr string) error {
	t.publicConns = make(map[string]Conn)
	err := t.controlServer.Init(addr, "Controller")
	return err
}

func (t *Trok) Start() {
	go t.controlServer.Start(t.ControlConnHandler)
	log.Info().Msgf("started Trok server on %s", t.controlServer.Addr())
}

func (t *Trok) Stop() {
	t.controlServer.Stop()
	log.Info().Msgf("stopped Trok server on %s", t.controlServer.Addr())
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
	err := s.Init(":", "Handler")
	if err != nil {
		log.Error().Msgf("error handling HELO cmd: %v", err)
		return
	}

	port := s.Port()
	p.WriteMessage(&lib.Message{CMD: "EHLO", ARG: fmt.Sprintf("%d", port)})

	uidChan := make(chan string)
	defer close(uidChan)

	go t.PublicConnHandler(s.listener, uidChan)

	for id := range uidChan {
		err := p.WriteMessage(&lib.Message{CMD: "CNCT", ARG: id})
		if err != nil {
			break
		}
	}
}

func (t *Trok) handleCMDACPT(conn net.Conn, m *lib.Message) {
	log.Info().Msgf("[CMD] %s [ARG] %s", m.CMD, m.ARG)

	t.mutex.Lock()
	pc, exists := t.publicConns[m.ARG]
	delete(t.publicConns, m.ARG)
	t.mutex.Unlock()

	if !exists || time.Since(pc.timestamp) > 10*time.Second {
		conn.Close()
		if exists {
			pc.conn.Close()
		}
		return
	}

	go io.Copy(pc.conn, conn)
	io.Copy(conn, pc.conn)
}

func (t *Trok) PublicConnHandler(ln net.Listener, uidChan chan<- string) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}

		id, err := gonanoid.New(12)
		if err != nil {
			log.Error().Msgf("error generating uid for public conn: %v", err)
			return
		}

		t.mutex.Lock()
		t.publicConns[id] = Conn{
			conn:      conn,
			timestamp: time.Now(),
		}
		t.mutex.Unlock()

		uidChan <- id
	}
}
