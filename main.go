package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/m-neves/goclock/database"
	"github.com/m-neves/goclock/handler"
)

func main() {
	sm := http.NewServeMux()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	rh := handler.NewRoutesHandler(sm)
	rh.SetupRoutes()

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	database.GetDb().Close()
	s.Shutdown(context.Background())
}
