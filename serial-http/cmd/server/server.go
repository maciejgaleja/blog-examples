package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/maciejgaleja/blog-examples/serial-http/pkg/port"
)

func main() {
	p := port.NewPort("/dev/ttyUSB1")

	mux := http.NewServeMux()
	mux.HandleFunc("/hostname", func(w http.ResponseWriter, r *http.Request) {
		h, _ := os.Hostname()
		fmt.Fprint(w, h)
	})
	mux.HandleFunc("/issue", func(w http.ResponseWriter, r *http.Request) {
		h, _ := os.ReadFile("/etc/issue")
		fmt.Fprint(w, string(h))
	})
	mux.HandleFunc("/platform", func(w http.ResponseWriter, r *http.Request) {
		h, _ := exec.Command("uname", "-p").CombinedOutput()
		fmt.Fprint(w, string(h))
	})

	http.Serve(p, mux)
}
