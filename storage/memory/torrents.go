package memory

import (
	"context"
	"github.com/EdmundMartin/discrete/torrents"
	"time"
)

type TorrentInfoMemoryStore struct {
	internalStorage map[string]*torrents.TorrentInfo
}

func (t TorrentInfoMemoryStore) LoadTorrentInfo(ctx context.Context, infoHash string) (*torrents.TorrentInfo, error) {
	val, ok := t.internalStorage[infoHash]
	if !ok {
		return nil, nil
	}
	return val, nil
}

func (t TorrentInfoMemoryStore) UpdateTorrentStatus(ctx context.Context, infoHash string) error {
	val, ok := t.internalStorage[infoHash]
	if !ok {
		now := time.Now()
		t.internalStorage[infoHash] = &torrents.TorrentInfo{
			InfoHash:  infoHash,
			CreatedOn: now,
			UpdatedOn: now,
		}
		return nil
	}
	val.UpdatedOn = time.Now()
	return nil
}

func NewTorrentInfoMemoryStore() *TorrentInfoMemoryStore {
	return &TorrentInfoMemoryStore{internalStorage: map[string]*torrents.TorrentInfo{}}
}
