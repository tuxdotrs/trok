/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package server

import (
	"fmt"
	"net"
)

type TCPServer struct {
	title    string
	listener net.Listener
}

func (s *TCPServer) Init(port uint16, title string) error {
	address := ":"
	if port > 0 {
		address = fmt.Sprintf(":%d", port)
	}

	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	s.title = title
	s.listener = ln
	return nil
}

func (s *TCPServer) Start(handler func(conn net.Conn)) error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}
		go handler(conn)
	}
}

func (s *TCPServer) Stop() error {
	return s.listener.Close()
}

func (s *TCPServer) Port() uint16 {
	return uint16(s.listener.Addr().(*net.TCPAddr).Port)
}
