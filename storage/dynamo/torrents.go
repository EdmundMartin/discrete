package dynamo

import (
	"context"
	"github.com/EdmundMartin/discrete/torrents"
)

func (d DataStore) LoadTorrentInfo(ctx context.Context, infoHash string) (*torrents.TorrentInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (d DataStore) UpdateTorrentStatus(ctx context.Context, info *torrents.TorrentInfo) error {
	//TODO implement me
	panic("implement me")
}
