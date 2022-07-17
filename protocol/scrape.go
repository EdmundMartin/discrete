package protocol

import (
	"github.com/EdmundMartin/discrete/torrent_errors"
	"net/http"
)

type ScrapeRequest struct {
	Hashes []string
}

func NewScrapeRequest(r *http.Request) (*ScrapeRequest, error) {
	scrape := &ScrapeRequest{}
	query := r.URL.Query()
	val, ok := query["info_hash"]
	if !ok {
		// Technically not malformed - but we do not support scrape all
		return nil, torrent_errors.MessageMalformedRequest
	}
	scrape.Hashes = val
	return scrape, nil
}
