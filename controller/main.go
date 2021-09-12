package main

import (
	"api-booking/controller/server/api"
	"api-booking/database"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

var client *redis.Client

func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr:       dsn,
		Password:   "admin123",
		DB:         0,
		MaxRetries: 3,
	})
	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Printf("Cannot Ping: %v\n", err.Error())
	} else {
		fmt.Printf("Pong: %v\n", pong)
	}
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
		ForceColors:   true,
	})
	logrus.SetReportCaller(true)
}

func main() {
	addr := ":8081"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Infof("Could not set up tcp: %v", err)
	}
	db, err := database.Initialize()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Infof("Could not set up database: %v", err)
	}

	defer db.Close()

	var httpHandler = api.NewHandler(db, client)
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
