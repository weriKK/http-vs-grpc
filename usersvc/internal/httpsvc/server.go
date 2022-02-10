package httpsvc

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type httpsvc struct {
	otherport      int
	isGrpcClient   bool
	grpcClientConn *grpc.ClientConn
}

func NewHTTPSvc(port int, otherport int, grpcClient bool) *http.Server {

	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", otherport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC client connection: %s", err.Error())
	}

	h := httpsvc{otherport: otherport, isGrpcClient: grpcClient, grpcClientConn: conn}

	mux := mux.NewRouter()
	mux.HandleFunc("/user", h.userDataHandler).Methods("POST")

	return &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		Handler: mux,
	}
}
