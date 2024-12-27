package main

import (
	"log"

	"github.com/sudatra/DFS-golang/p2p"
)

func main() {
	tcpOps := p2p.TCPTransportOps{
		ListenAddr: ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.GOBDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOps);

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}