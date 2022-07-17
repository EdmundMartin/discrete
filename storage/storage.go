package storage

import (
	"context"
	"github.com/EdmundMartin/discrete/protocol"
	"github.com/EdmundMartin/discrete/torrents"
)

type PeerStore interface {
	StorePeer(ctx context.Context, peer *protocol.Peer) error
	LoadPeers(ctx context.Context, infoHash string) ([]*protocol.Peer, error)
}

type TorrentStore interface {
	LoadTorrentInfo(ctx context.Context, infoHash string) (*torrents.TorrentInfo, error)
	UpdateTorrentStatus(ctx context.Context, infoHash string) error
	ListInfo(ctx context.Context, infoHashes []string) ([]*torrents.TorrentInfo, error)
}
