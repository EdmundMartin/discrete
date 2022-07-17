package memory

import (
	"context"
	"github.com/EdmundMartin/discrete/torrents"
)

type TorrentInfoMemoryStore struct {
	internalStorage map[string]*torrents.TorrentInfo
}

func (t TorrentInfoMemoryStore) ListInfo(ctx context.Context, infoHash []string) ([]*torrents.TorrentInfo, error) {
	var results []*torrents.TorrentInfo
	for _, hash := range infoHash {
		val, ok := t.internalStorage[hash]
		if ok {
			results = append(results, val)
		}
	}
	return results, nil
}

func (t TorrentInfoMemoryStore) LoadTorrentInfo(ctx context.Context, infoHash string) (*torrents.TorrentInfo, error) {
	val, ok := t.internalStorage[infoHash]
	if !ok {
		return nil, nil
	}
	return val, nil
}

func (t TorrentInfoMemoryStore) UpdateTorrentStatus(ctx context.Context, info *torrents.TorrentInfo) error {
	t.internalStorage[info.InfoHash] = info
	return nil
}

func NewTorrentInfoMemoryStore() *TorrentInfoMemoryStore {
	return &TorrentInfoMemoryStore{internalStorage: map[string]*torrents.TorrentInfo{}}
}
