package main

import (
	"context"
	"log"
	"net/http"

	"ciklum_test/ClientAPI/router"

	"ciklum_test/ClientAPI/client"

	"github.com/pkg/errors"

	"ciklum_test/ClientAPI/utils"
)

const (
	port = ":5000"
)

func main() {
	ctx := context.Background()

	grpcClient, err := client.NewGRPCClient(ctx)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to create new grpc client"))
	}

	log.Println("parsing json")
	err = utils.ParseAndCallUpdateFunc(ctx, "./ports.json", grpcClient.CreateOrUpdatePorts)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "parse and call updateFunc"))
	}

	log.Println("initializing a new router")
	r := router.NewRouter(ctx, grpcClient)

	log.Println("starting the server")
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(errors.Wrap(err, "Could not start http server"))
	}
}
