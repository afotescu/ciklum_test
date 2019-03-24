package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"ciklum_test/PortDomainService/server"

	"ciklum_test/PortDomainService/db"

	"github.com/pkg/errors"
)

const (
	tcpPort = ":5001"
)

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

	mdb, client, err := db.NewMongoDB(ctx)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to initialize mongo database"))
	}

	grpcServer, err := server.NewGRPCServer(ctx, mdb)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to initialize grpc server"))
	}

	listener, err := net.Listen("tcp", tcpPort)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to initialize tcp listener"))
	}
	go func() {
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
