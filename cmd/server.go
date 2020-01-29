package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SuddenGunter/pandaren/browser"
)

type server struct {
	browser *browser.Browser

	handler *http.Server
}

func RunServer(shutdown chan os.Signal) {
	s := &server{
		browser: browser.New(),
	}
	s.configureRoutes()
	s.runUntilShutdown(shutdown)
}

func (s *server) configureRoutes() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", PdfHandlerFunc(s.browser))
	s.handler = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

func (s *server) runUntilShutdown(shutdown chan os.Signal) {
	go func() {
		if err := s.handler.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// graceful shutdown
	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
		s.browser.Close()
	}()

	if err := s.handler.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed:%+v", err)
	}
	log.Print("Server shutdown properly")
}
