package server

import (
	"context"
	"fmt"
	"github.com/EdmundMartin/discrete/bencode"
	"github.com/EdmundMartin/discrete/config"
	"github.com/EdmundMartin/discrete/protocol"
	"github.com/EdmundMartin/discrete/storage"
	"github.com/EdmundMartin/discrete/torrent_errors"
	"net/http"
)

type TorrentServer struct {
	PeerDB    storage.PeerStore
	TorrentDB storage.TorrentStore
	Config    config.ConfigStore
}

func (t *TorrentServer) handleAutoRegister(ctx context.Context, peer *protocol.Peer) error {
	if err := t.PeerDB.StorePeer(ctx, peer); err != nil {
		// As this will be an internal storage error we want to return a Generic error to client
		return torrent_errors.GenericError
	}

	if err := t.TorrentDB.UpdateTorrentStatus(ctx, peer.InfoHash); err != nil {
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

	err = t.TorrentDB.UpdateTorrentStatus(ctx, peer.InfoHash)
	if err != nil {
		return torrent_errors.GenericError
	}

	return nil
}

func (t *TorrentServer) Announce(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	peer, err := protocol.NewPeerFromRequest(r)
	if err != nil {
		ErrorResponse(w, torrent_errors.InvalidRequestType)
		return
	}

	// When auto register is allowed we blindly store the peer information
	if t.Config.AutoRegister() {
		err = t.handleAutoRegister(ctx, peer)
		if err != nil {
			ErrorResponse(w, err)
			return
		}
	} else {
		err = t.handleRequiredRegister(ctx, peer)
		if err != nil {
			ErrorResponse(w, err)
			return
		}
	}

	peers, err := t.PeerDB.LoadPeers(ctx, peer.InfoHash)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	activePeers := protocol.ActivePeers(peers)

	contents := bencode.EncodePeerResponse(activePeers)
	fmt.Println(contents)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(contents))
}
