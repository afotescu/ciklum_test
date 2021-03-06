package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/afotescu/ciklum_test/PortDomainService/server"

	"github.com/afotescu/ciklum_test/PortDomainService/db"

	"github.com/pkg/errors"
)

var tcpPort = os.Getenv("PORT")

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-signalCh
		log.Println("shutting down the server")
		cancel()
	}()

	log.Println("connecting to the database")
	collection, client, err := db.NewMongoDB(ctx)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to initialize mongo database"))
	}

	log.Println("initializing grpc server")
	grpcServer, err := server.NewGRPCServer(ctx, collection)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to initialize grpc server"))
	}

	listener, err := net.Listen("tcp", tcpPort)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to initialize tcp listener"))
	}
	go func() {
		log.Println("starting the server")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal(errors.Wrap(err, "failed to start grpc server"))
		}
	}()

	<-ctx.Done()

	grpcServer.Stop()
	log.Println("grpc server was stopped")

	if err := client.Disconnect(ctx); err != nil {
		log.Println(errors.Wrap(err, "failed to disconnect from mongo database"))
	}
	log.Println("mongo client has disconnected")
	os.Exit(0)
}
