package main

import "sync"

// Room ...
type Room struct {
	Peers *Peers
	// Hub
}

// Peers ...
type Peers struct {
	ListLock sync.RWMutex
	// Connections []PeerConnectionState
	// TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}
