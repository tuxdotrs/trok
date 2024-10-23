/*
Copyright © 2024 tux <0xtux@pm.me>
*/

package lib

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

type Message struct {
	CMD string
	ARG string
}

type ProtocolHandler struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

func InitProtocolHandler(conn net.Conn) *ProtocolHandler {
	return &ProtocolHandler{
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}
}

func (p *ProtocolHandler) ReadMessage() (*Message, error) {
	data, err := p.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	cmd, arg, err := p.parseMessage(data)
	if err != nil {
		return nil, errors.New("can't parse data")
	}

	return &Message{
		CMD: cmd,
		ARG: arg,
	}, nil
}

func (p *ProtocolHandler) WriteMessage(m *Message) error {
	_, err := p.writer.WriteString(fmt.Sprintf("%s %s\n", m.CMD, m.ARG))
	if err != nil {
		return err
	}
	return p.writer.Flush()
}

func (p *ProtocolHandler) parseMessage(data string) (string, string, error) {
	data = strings.TrimSpace(data)
	d := strings.Fields(data)

	if len(d) != 2 {
		return "", "", errors.New("invalid command")
	}

	return d[0], d[1], nil
}
