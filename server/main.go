package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/prateek041/real-time-chat-app/handlers"
)

func createMux(logger *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	wsHandler := handlers.NewSocket(logger)
	homeHandler := handlers.NewHomePage()
	mux.Handle("/", homeHandler)
	mux.Handle("/ws", wsHandler)
	return mux
}

func main() {
	logger := log.New(os.Stdout, "CHAT-APP:", log.LstdFlags)
	serveMux := createMux(logger)

	server := &http.Server{
		Addr:    ":9090",
		Handler: serveMux,
	}

	go func() {
		logger.Println("The server is starting")

		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Error in starting the server")
		}
	}()

	interruptSig := make(chan os.Signal)
	signal.Notify(interruptSig, os.Interrupt)
	signal.Notify(interruptSig, os.Kill)
	sig := <-interruptSig
	logger.Println("Server shutting down", sig)
	tc, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	_ = cancelFunc
	server.Shutdown(tc)
}
