package http

import (
	"crypto/tls"
	"net"
	gohttp "net/http"
)

type Server struct {
	l net.Listener
}

func New(addr, target string, tlsConfig *tls.Config) (*Server, error) {
	s := gohttp.Server{
		Addr: addr,
	}
}
