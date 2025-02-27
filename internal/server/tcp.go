/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package server

import (
	"net"
)

type TCPServer struct {
	title    string
	listener net.Listener
}

func (s *TCPServer) Init(addr, title string) error {
	ln, err := net.Listen("tcp", addr)
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

func (s *TCPServer) Addr() string {
	return s.listener.Addr().String()
}

func (c *TCPServer) Host() string {
	return c.listener.Addr().(*net.TCPAddr).IP.String()
}

func (s *TCPServer) Port() uint16 {
	return uint16(s.listener.Addr().(*net.TCPAddr).Port)
}
