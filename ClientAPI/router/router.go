package router

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	pb "ciklum_test/protobuf"

	"github.com/gorilla/mux"
)

func NewRouter(ctx context.Context, grpcClient pb.TransporterClient) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ports", getPorts(ctx, grpcClient))
	return router
}

func getPorts(ctx context.Context, grpcClient pb.TransporterClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitStr := r.URL.Query().Get("limit")
		limit, err := strconv.ParseUint(limitStr, 10, 64)
		if err != nil {
			limit = 10
		}
		offsetStr := r.URL.Query().Get("offset")
		offset, err := strconv.ParseUint(offsetStr, 10, 64)
		if err != nil {
			offset = 0
		}
		response, err := grpcClient.GetPorts(ctx, &pb.PortsPage{Limit: limit, Offset: offset})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err = json.NewEncoder(w).Encode(&response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
