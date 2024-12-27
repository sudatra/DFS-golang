package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over the established TCP connection
type TCPPeer struct {
	conn 		net.Conn
	// if we dial and retrieve a conn => outbound == true
	// if we accept and retrieve a conn => outbound == false
	outbound 	bool
}

type TCPTransportOps struct {
	ListenAddr		string
	HandshakeFunc	HandshakeFunc
	Decoder			Decoder

}

type TCPTransport struct {
	TCPTransportOps
	listener      net.Listener

	mu 			  sync.RWMutex
	peers 		  map[net.Addr]Peer
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn: conn,
		outbound: outbound,
	}
}

func NewTCPTransport(ops TCPTransportOps) *TCPTransport {
	return &TCPTransport{
		TCPTransportOps: ops,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr);
	if err != nil {
		return err
	}

	go t.startAcceptLoop();
	return nil;
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP Accept Error: %s\n: ", err)
		}

		fmt.Printf("New incoming connection %+v\n", conn)
		go t.handleConn(conn)
	}
}

type Temp struct {}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake Error: %s\n", err)

		return;
	}

	msg := &Temp{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP Error: %s\n", err)
			continue
		}
	}
}