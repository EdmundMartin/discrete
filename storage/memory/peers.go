package memory

import (
	"context"
	"github.com/EdmundMartin/discrete/protocol"
)

type PeerMemoryStore struct {
	peersByInfoHash map[string][]*protocol.Peer
}

func (m PeerMemoryStore) StorePeer(ctx context.Context, peer *protocol.Peer) error {
	peers, ok := m.peersByInfoHash[peer.InfoHash]
	if !ok {
		m.peersByInfoHash[peer.InfoHash] = []*protocol.Peer{peer}
		return nil
	}
	for idx, currentPeer := range peers {
		if currentPeer.ClientID == peer.ClientID {
			peers[idx] = peer
			return nil
		}
	}
	m.peersByInfoHash[peer.InfoHash] = append(m.peersByInfoHash[peer.InfoHash], peer)
	return nil
}

func (m PeerMemoryStore) LoadPeers(ctx context.Context, infoHash string) ([]*protocol.Peer, error) {
	peers, ok := m.peersByInfoHash[infoHash]
	if !ok {
		return []*protocol.Peer{}, nil
	}
	return peers, nil
}

func NewMemoryStore() *PeerMemoryStore {
	return &PeerMemoryStore{map[string][]*protocol.Peer{}}
}
