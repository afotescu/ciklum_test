package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"

	pb "ciklum_test/protobuf"
)

const (
	address = "localhost:50051"
)

func notMain(ctx context.Context, client pb.TransporterClient) {
	jsonFile, err := os.Open("./ports.json")
	if err != nil {
		log.Fatal(err)
	}

	//
	decoder := json.NewDecoder(jsonFile)
	t, err := decoder.Token()
	if delim, ok := t.(json.Delim); !ok || delim != '{' {
		log.Fatal("Expected object")
	}
	for decoder.More() {
		token, err := decoder.Token()
		tokenStr, ok := token.(string)
		if !ok && tokenStr != strings.ToUpper(tokenStr) {
			log.Fatal("token must be an uppercase string")
		}
		if err != nil {
			log.Fatalf("1an error occured: %v", err)
		}

		for decoder.More() {
			port := &pb.Port{}
			err = decoder.Decode(port)
			if err != nil {
				if err.Error() == "not at beginning of value" {
					break
				}
				log.Fatalf("2an error occured: %v", err)
			}
			_, err := client.CreateOrUpdatePorts(ctx, port)
			if err != nil {
				log.Fatalf("1an error occured: %v", err)
			}
		}
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		if err = conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	c := pb.NewTransporterClient(conn)

	ctx := context.Background()
	notMain(ctx, c)
}
