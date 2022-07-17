package server

import (
	"bytes"
	"github.com/EdmundMartin/discrete/bencoding"
	"github.com/EdmundMartin/discrete/protocol"
	"github.com/EdmundMartin/discrete/torrent_errors"
	"net/http"
)

func (t *TorrentServer) Scrape(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	if !t.Config.SupportScraping() {
		errorResponse(w, torrent_errors.MessageMalformedRequest)
		return
	}

	scrape, err := protocol.NewScrapeRequest(r)
	if err != nil {
		errorResponse(w, err)
		return
	}

	info, err := t.TorrentDB.ListInfo(ctx, scrape.Hashes)
	if err != nil {
		errorResponse(w, err)
		return
	}

	resp := make(bencoding.Dict, len(info))

	for _, tor := range info {
		resp[tor.InfoHash] = bencoding.Dict{
			"complete":   tor.Seeders,
			"downloaded": tor.Downloaded,
			"incomplete": tor.Leechers,
		}
	}

	var buffer bytes.Buffer

	if err := bencoding.NewEncoder(&buffer).Encode(resp); err != nil {
		errorResponse(w, err)
		return
	}
	scrapeResponse(w, buffer.Bytes())
}
