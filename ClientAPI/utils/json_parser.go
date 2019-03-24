package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	pb "ciklum_test/protobuf"
)

type createOrUpdateFunc func(context.Context, *pb.Port, ...grpc.CallOption) (*pb.Response, error)

// ParseAndCallUpdateService will parse json file and then call updateFunc on each record
func ParseAndCallUpdateFunc(ctx context.Context, filePath string, updateFunc createOrUpdateFunc) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(jsonFile)
	t, err := decoder.Token()
	if delim, ok := t.(json.Delim); !ok || delim != '{' {
		return errors.New(fmt.Sprintf("expected first token to be \"{\", got %s ", delim))
	}
	for decoder.More() {
		token, err := decoder.Token()
		tokenStr, ok := token.(string)
		if !ok && tokenStr != strings.ToUpper(tokenStr) {
			return errors.New("token must be an uppercase string")
		}
		if err != nil {
			return errors.Wrap(err, "failed to get json token")
		}

		for decoder.More() {
			port := &pb.Port{}
			err = decoder.Decode(port)
			if err != nil {
				if err.Error() == "not at beginning of value" {
					break
				}
				return errors.Wrap(err, "could not decode data")
			}
			_, err := updateFunc(ctx, port)
			if err != nil {
				return errors.Wrap(err, "failed to call updateFunc")
			}
		}
	}
	return nil
}
