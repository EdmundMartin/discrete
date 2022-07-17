package storage

import (
	"context"
	"github.com/EdmundMartin/discrete/protocol"
)

type PeerStore interface {
	StorePeer(ctx context.Context, peer *protocol.Peer) error
	LoadPeers(ctx context.Context, infoHash string) ([]*protocol.Peer, error)
}
