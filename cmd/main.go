package main

import (
	"context"
	"github.com/EdmundMartin/discrete/config/defaultconfig"
	"github.com/EdmundMartin/discrete/server"
	"github.com/EdmundMartin/discrete/storage/memory"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	t := server.TorrentServer{
		TorrentDB: memory.NewTorrentInfoMemoryStore(),
		PeerDB:    memory.NewMemoryStore(),
		Config:    defaultconfig.NewDefaultConfig(),
		Logger:    log.New(os.Stdout, "", log.Flags()),
	}

	r := mux.NewRouter()
	r.HandleFunc("/announce", t.PublicAnnounce)

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Logger.Fatalf("Listen: %s\n", err)
		}
	}()
	t.Logger.Println("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		t.Logger.Fatalf("Server shutdown failed: %v", err)
	}

	t.Logger.Println("Server exited successfully")
}
