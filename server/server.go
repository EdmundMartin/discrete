package server

import (
	"fmt"
	"github.com/EdmundMartin/discrete/bencode"
	"github.com/EdmundMartin/discrete/protocol"
	"github.com/EdmundMartin/discrete/storage"
	"github.com/EdmundMartin/discrete/torrent_errors"
	"net/http"
)

type TorrentServer struct {
	PeerDB storage.PeerStore
}

func (t *TorrentServer) Announce(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	peer, err := protocol.NewPeerFromRequest(r)
	if err != nil {
		ErrorResponse(w, torrent_errors.InvalidRequestType)
		return
	}
	fmt.Println(peer)
	err = t.PeerDB.StorePeer(ctx, peer)
	if err != nil {
		ErrorResponse(w, torrent_errors.GenericError)
		return
	}
	peers, err := t.PeerDB.LoadPeers(ctx, peer.InfoHash)
	contents := bencode.EncodePeerResponse(peers)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(contents))
}
