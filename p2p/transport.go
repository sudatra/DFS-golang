package p2p

// Peer represents the remote node
type Peer interface {}

// Transport is used for communication between any 2 remote nodes
type Transport interface {
	ListenAndAccept() error
}