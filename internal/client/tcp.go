/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package client

import (
	"net"
)

type TCPClient struct {
	title string
	conn  net.Conn
}

func NewTCPClient(addr, title string) (*TCPClient, error) {
	conn, err := net.Dial("tcp", addr)

	return &TCPClient{
		title: title,
		conn:  conn,
	}, err
}

func (c *TCPClient) Start(handler func(conn net.Conn)) {
	handler(c.conn)
}

func (c *TCPClient) Stop() error {
	return c.conn.Close()
}

func (s *TCPClient) Addr() string {
	return s.conn.RemoteAddr().String()
}

func (c *TCPClient) Host() string {
	return c.conn.RemoteAddr().(*net.TCPAddr).IP.String()
}

func (c *TCPClient) Port() uint16 {
	return uint16(c.conn.RemoteAddr().(*net.TCPAddr).Port)
}
