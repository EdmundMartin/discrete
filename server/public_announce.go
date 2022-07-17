package server

import (
	"context"
	"github.com/EdmundMartin/discrete/bencoding"
	"github.com/EdmundMartin/discrete/config"
	"github.com/EdmundMartin/discrete/protocol"
	"github.com/EdmundMartin/discrete/storage"
	"github.com/EdmundMartin/discrete/torrent_errors"
	"github.com/EdmundMartin/discrete/torrents"
	"log"
	"net/http"
	"time"
)

type TorrentServer struct {
	Logger    *log.Logger // TODO - Replace with interface
	PeerDB    storage.PeerStore
	TorrentDB storage.TorrentStore
	Config    config.ConfigStore
}

func (t *TorrentServer) handleAutoRegister(ctx context.Context, peer *protocol.Peer) error {
	if err := t.PeerDB.StorePeer(ctx, peer); err != nil {
		// As this will be an internal storage error we want to return a Generic error to client
		return torrent_errors.GenericError
	}

	tor, err := t.TorrentDB.LoadTorrentInfo(ctx, peer.InfoHash)
	if err != nil {
		return torrent_errors.GenericError
	}

	if tor == nil {
		tor = &torrents.TorrentInfo{
			InfoHash:   peer.InfoHash,
			Seeders:    0,
			Downloaded: 0,
			Leechers:   0,
			Announces:  0,
			CreatedOn:  time.Now(),
			UpdatedOn:  time.Now(),
		}
	}

	torrents.UpdateFromPeer(tor, peer)

	if err := t.TorrentDB.UpdateTorrentStatus(ctx, tor); err != nil {
		return torrent_errors.GenericError
	}

	return nil
}

func (t *TorrentServer) handleRequiredRegister(ctx context.Context, peer *protocol.Peer) error {
	info, err := t.TorrentDB.LoadTorrentInfo(ctx, peer.InfoHash)
	// As this will be an internal storage error we want to return a Generic error to client
	if err != nil {
		return torrent_errors.GenericError
	}
	if info == nil {
		return torrent_errors.InfoHashNotFound
	}

	err = t.PeerDB.StorePeer(ctx, peer)
	if err != nil {
		return torrent_errors.GenericError
	}

	torrents.UpdateFromPeer(info, peer)
	err = t.TorrentDB.UpdateTorrentStatus(ctx, info)
	if err != nil {
		return torrent_errors.GenericError
	}

	return nil
}

func (t *TorrentServer) PublicAnnounce(w http.ResponseWriter, r *http.Request) {

	if t.Config.IsPrivateOnly() {
		// If we don't support public announcements simple respond with Invalid request
		errorResponse(w, torrent_errors.InvalidRequestType)
	}

	ctx := r.Context()
	peer, err := protocol.NewPeerFromRequest(r)
	if err != nil {
		errorResponse(w, torrent_errors.InvalidRequestType)
		return
	}
	t.Logger.Printf("got event: %s, from clientID: %s", peer.Event, peer.ClientID)

	// When auto register is allowed we blindly store the peer information
	if t.Config.AutoRegister() {
		err = t.handleAutoRegister(ctx, peer)
		if err != nil {
			errorResponse(w, err)
			return
		}
	} else {
		err = t.handleRequiredRegister(ctx, peer)
		if err != nil {
			errorResponse(w, err)
			return
		}
	}

	peers, err := t.PeerDB.LoadPeers(ctx, peer.InfoHash)
	if err != nil {
		errorResponse(w, err)
		return
	}
	activePeers := protocol.FilterPeers(peers, peer.ClientID)
	contents, err := bencoding.AnnounceResponse(activePeers, t.Config.TrackerInterval(), peer.IpV6)

	if err != nil {
		errorResponse(w, err)
		return
	}
	scrapeResponse(w, contents)
}
