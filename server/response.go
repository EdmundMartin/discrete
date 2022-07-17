package server

import (
	"github.com/EdmundMartin/discrete/simple_bencode"
	"github.com/EdmundMartin/discrete/torrent_errors"
	"net/http"
)

func errorResponse(w http.ResponseWriter, err error) {
	switch t := err.(type) {
	default:
		w.WriteHeader(torrent_errors.GenericError.Code)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(simple_bencode.KeyValue("failure reason", torrent_errors.GenericError.Message)))
	case torrent_errors.TrackerError:
		w.WriteHeader(t.Code)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(simple_bencode.KeyValue("failure reason", t.Message)))
	}
}

func announceResponse(w http.ResponseWriter, contents string) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(contents))
}

func scrapeResponse(w http.ResponseWriter, contents []byte) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write(contents)
}
