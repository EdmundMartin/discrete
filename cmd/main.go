package main

import (
	"github.com/EdmundMartin/discrete/server"
	"github.com/EdmundMartin/discrete/storage/memory"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	t := server.TorrentServer{
		PeerDB: memory.NewMemoryStore(),
	}

	r := mux.NewRouter()
	r.HandleFunc("/announce", t.Announce)
	http.ListenAndServe("0.0.0.0:8080", r)
}
