package main

import (
	"api-booking/database"
	"api-booking/handlers"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	addr := ":8081"
	listener, err := net.Listen("tcp", addr)
	db, err := database.Initialize()

	if err != nil {
		logrus.WithFields(logrus.Fields{}).Infof("Could not set up database: %v", err)
	}
	defer db.Close()

	httpHandler := handlers.NewHandler(db)

	server := &http.Server{
		Handler: httpHandler,
	}

	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)
	logrus.WithFields(logrus.Fields{}).Infof("Started server on %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	logrus.WithFields(logrus.Fields{}).Info("Stopping API server.")
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logrus.WithFields(logrus.Fields{}).Infof("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
