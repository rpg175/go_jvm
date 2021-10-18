package main

import (
	"crypto/tls"
	"time"
)

type Server struct {
	Addr string
	Port int
	Protocol string
	Timeout time.Duration
	MaxConnects int
	TLS *tls.Config
}


