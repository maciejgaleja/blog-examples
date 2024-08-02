package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/maciejgaleja/blog-examples/serial-http/pkg/port"
)

func must(e error) {
	if e != nil {
		log.Print(e)
		os.Exit(1)
	}
}

func get(c http.Client, path string) string {
	resp, err := c.Get("http://serial/" + path)
	must(err)
	body, err := io.ReadAll(resp.Body)
	must(err)
	return string(body)
}

func main() {
	p := port.NewPort("/dev/ttyUSB0")
	dialer := func(context.Context, string, string) (net.Conn, error) {
		return p, nil
	}

	t := http.DefaultTransport
	t.(*http.Transport).DialContext = dialer
	c := http.Client{Transport: t}
	fmt.Println(get(c, "hostname"))
	fmt.Println(get(c, "platform"))
}
