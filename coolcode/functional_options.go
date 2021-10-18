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

func NewDefaultServer(addr string, port int) (*Server,error) {
	return &Server{addr,port,"tcp",30 * time.Second,100,nil},nil
}

func NewTLSServer(addr string,port int ,tls *tls.Config) (*Server,error) {
	return &Server{addr,port,"tcp",30*time.Second,100,tls},nil
}

func NewServerWithTimeout(addr string,port int,timeout time.Duration) (*Server, error) {
	return &Server{addr,port,"tcp",timeout,100,nil},nil
}

func NewServerTLSServerWithMaxConnAndTimeout(addr string,port int,maxConnects int,timeout time.Duration,tls *tls.Config)(*Server, error) {
	return &Server{addr,port,"tcp",30*time.Second,maxConnects,tls},nil
}

