package server

import (
	"github.com/EdmundMartin/discrete/bencode"
	"github.com/EdmundMartin/discrete/torrent_errors"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, err error) {
	switch t := err.(type) {
	default:
		w.WriteHeader(torrent_errors.GenericError.Code)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(bencode.KeyValue("failure reason", torrent_errors.GenericError.Message)))
	case torrent_errors.TrackerError:
		w.WriteHeader(t.Code)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(bencode.KeyValue("failure reason", t.Message)))
	}
}
