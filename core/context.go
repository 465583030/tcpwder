package core

import "net"

type Context interface {
	String() string
	Ip() net.IP
	Port() int
	Sni() string
}

/**
 * Proxy tcp context
 */
type TcpContext struct {
	Hostname string
	/**
	 * Current client connection
	 */
	Conn net.Conn
}

func (t TcpContext) String() string {
	return t.Conn.RemoteAddr().String()
}

func (t TcpContext) Ip() net.IP {
	return t.Conn.RemoteAddr().(*net.TCPAddr).IP
}

func (t TcpContext) Port() int {
	return t.Conn.RemoteAddr().(*net.TCPAddr).Port
}

func (t TcpContext) Sni() string {
	return t.Hostname
}

/*
 * Proxy udp context
 */
type UdpContext struct {

	/**
	 * Current client remote address
	 */
	RemoteAddr net.UDPAddr
}

func (u UdpContext) String() string {
	return u.RemoteAddr.String()
}

func (u UdpContext) Ip() net.IP {
	return u.RemoteAddr.IP
}

func (u UdpContext) Port() int {
	return u.RemoteAddr.Port
}

func (u UdpContext) Sni() string {
	return ""
}