package dynamo

import (
	"context"
	"github.com/EdmundMartin/discrete/protocol"
)

func (d DataStore) StorePeer(ctx context.Context, peer *protocol.Peer) error {
	//TODO implement me
	panic("implement me")
}

func (d DataStore) LoadPeers(ctx context.Context, infoHash string) ([]*protocol.Peer, error) {
	//TODO implement me
	panic("implement me")
}
