package torrents

import (
	"github.com/EdmundMartin/discrete/protocol"
	"time"
)

type TorrentInfo struct {
	InfoHash   string
	Title      string
	Seeders    int
	Downloaded int
	Leechers   int
	Announces  int
	CreatedOn  time.Time
	UpdatedOn  time.Time
}

func UpdateFromPeer(torrentInfo *TorrentInfo, peer *protocol.Peer) {
	now := time.Now()
	torrentInfo.UpdatedOn = now
	if torrentInfo.CreatedOn.IsZero() {
		torrentInfo.CreatedOn = now
	}

	if peer.Event == protocol.ANNOUNCE {
		// Does announce count as active in the swarm?
		torrentInfo.Announces += 1
		return
	}

	if peer.Event == protocol.COMPLETE {
		torrentInfo.Downloaded += 1
	}
	if protocol.IsInactive(peer.Event) && peer.Seed {
		torrentInfo.Seeders -= 1
	}
	if protocol.IsInactive(peer.Event) && !peer.Seed {
		torrentInfo.Leechers -= 1
	}

	if protocol.IsActive(peer.Event) && peer.Seed {
		torrentInfo.Seeders += 1
	}
	if protocol.IsInactive(peer.Event) && !peer.Seed {
		torrentInfo.Leechers += 1
	}
}
