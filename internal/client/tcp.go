/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package client

import (
	"fmt"
	"net"
)

type TCPClient struct {
	title string
	conn  net.Conn
}

func (c *TCPClient) Init(port uint16, title string) error {
	address := fmt.Sprintf(":%d", port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	c.title = title
	c.conn = conn
	return nil
}

func (c *TCPClient) Start(handler func(conn net.Conn)) {
	handler(c.conn)
}

func (c *TCPClient) Stop() error {
	return c.conn.Close()
}

func (c *TCPClient) Port() uint16 {
	return uint16(c.conn.RemoteAddr().(*net.TCPAddr).Port)
}
