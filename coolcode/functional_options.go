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

//使用一个builder类来做包装
type ServerBuilder struct {
	Server
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
	sb.Server.Addr = addr
	sb.Server.Port = port
	//其它代码设置其它成员的默认值
	return sb
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	sb.Server.Protocol = protocol
	return sb
}

func (sb *ServerBuilder) WithMaxConnect(maxConnect int) *ServerBuilder{
	sb.Server.MaxConnects = maxConnect
	return sb
}

func (sb *ServerBuilder) WithTimeOut( timeout time.Duration) *ServerBuilder {
	sb.Server.Timeout = timeout
	return sb
}
func (sb *ServerBuilder) WithTLS( tls *tls.Config) *ServerBuilder {
	sb.Server.TLS = tls
	return sb
}
func (sb *ServerBuilder) Build() (Server) {
	return  sb.Server
}

type Option func(*Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func MaxConnects(maxConnect int) Option {
	return func(s *Server) {
		s.MaxConnects = maxConnect
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConnects: 1000,
		TLS:      nil,
	}
	for _, option := range options {
		option(&srv)
	}
	//...
	return &srv, nil
}

func main() {
	//sb := ServerBuilder{}
	//server, err := sb.Create("127.0.0.1", 8080).
	//	WithProtocol("udp").
	//	WithMaxConnect(1024).
	//	WithTimeOut(30*time.Second).
	//	Build()

	//s1, _ := NewServer("localhost", 1024)
	//s2, _ := NewServer("localhost", 2048, Protocol("udp"))
	//s3, _ := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConnects(1000))
}