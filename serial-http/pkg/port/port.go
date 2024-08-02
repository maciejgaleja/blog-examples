package port

import (
	"io"
	"log"
	"net"
	"time"

	"go.bug.st/serial"
)

type Port struct {
	port     io.ReadWriter
	accepted bool
}

func NewPort(path string) *Port {
	mode := &serial.Mode{BaudRate: 115200}

	port, err := serial.Open(path, mode)
	if err != nil {
		panic(err)
	}
	return &Port{port: port}
}

func (p *Port) Accept() (net.Conn, error) {
	if !p.accepted {
		p.accepted = true
		return p, nil
	} else {
		for {
			time.Sleep(time.Hour)
		}
	}
}
func (p *Port) Read(b []byte) (n int, err error) {
	n, err = p.port.Read(b)
	log.Printf("Read(%d) -> %d", len(b), n)
	return
}
func (p *Port) Write(b []byte) (n int, err error) {
	log.Printf("Write(%d)", len(b))
	return p.port.Write(b)
}
func (p *Port) Close() error                       { return nil }
func (p *Port) Addr() net.Addr                     { return &net.UnixAddr{Name: "serial"} }
func (p *Port) LocalAddr() net.Addr                { return &net.UnixAddr{Name: "serial"} }
func (p *Port) RemoteAddr() net.Addr               { return &net.UnixAddr{Name: "serial"} }
func (p *Port) SetDeadline(t time.Time) error      { return nil }
func (p *Port) SetReadDeadline(t time.Time) error  { return nil }
func (p *Port) SetWriteDeadline(t time.Time) error { return nil }
