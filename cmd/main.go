package main

import (
	"github.com/EdmundMartin/discrete/config/defaultconfig"
	"github.com/EdmundMartin/discrete/server"
	"github.com/EdmundMartin/discrete/storage/memory"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	t := server.TorrentServer{
		TorrentDB: memory.NewTorrentInfoMemoryStore(),
		PeerDB:    memory.NewMemoryStore(),
		Config:    defaultconfig.NewDefaultConfig(),
	}

	r := mux.NewRouter()
	r.HandleFunc("/announce", t.Announce)
	http.ListenAndServe("0.0.0.0:8080", r)
}
