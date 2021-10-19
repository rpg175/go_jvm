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

func main() {
	//sb := ServerBuilder{}
	//server, err := sb.Create("127.0.0.1", 8080).
	//	WithProtocol("udp").
	//	WithMaxConnect(1024).
	//	WithTimeOut(30*time.Second).
	//	Build()
}