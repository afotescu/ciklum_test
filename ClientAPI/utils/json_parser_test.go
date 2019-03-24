package utils

import (
	"context"
	"log"
	"os"
	"testing"

	"google.golang.org/grpc"

	pb "ciklum_test/protobuf"
)

func TestParseAndCallUpdateFunc(t *testing.T) {
	ctx := context.TODO()
	fileName := "test.json"

	t.Run("should fail when file doesn't exist", func(t *testing.T) {
		err := ParseAndCallUpdateFunc(ctx, fileName, nil)
		if err == nil {
			t.Fatal("error should not be nil")
		}
	})

	t.Run("should open file and fail because of wrong or no content", func(t *testing.T) {
		createTestFile(fileName, false)
		defer removeTestFile(fileName)
		err := ParseAndCallUpdateFunc(ctx, fileName, nil)
		if err == nil {
			t.Fatal("error should not be nil")
		}
	})

	t.Run("should open file and call updateFunc", func(t *testing.T) {
		createTestFile(fileName, true)
		defer removeTestFile(fileName)
		wasCalled := false
		err := ParseAndCallUpdateFunc(ctx, fileName, func(context.Context, *pb.Port, ...grpc.CallOption) (*pb.Response, error) {
			wasCalled = true
			return nil, nil
		})
		if err != nil {
			t.Fatalf("error should be nil, got: %v", err)
		}
		if !wasCalled {
			t.Fatal("wasCalled should be true, got false")
		}
	})

}

func removeTestFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		log.Fatal(err)
	}
}

func createTestFile(fileName string, withContent bool) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if !withContent {
		return
	}
	_, err = f.WriteString(`{
   "TEST":{
      "name":"TestName",
      "city":"TestCity",
      "country":"TestCountry",
      "alias":[],
      "regions":[],
      "coordinates":[
         55.5136433,
         25.4052165
      ],
      "province":"TestProvince",
      "timezone":"Test/TimeZone",
      "unlocs":[
         "TestUnlocs"
      ],
      "code":"00000"
   }
}`)
	if err != nil {
		log.Fatal(err)
	}
}
