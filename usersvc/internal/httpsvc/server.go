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

func NewHTTPSvc(port int, otherport int, isGrpcClient bool) *http.Server {

	h := httpsvc{otherport: otherport, isGrpcClient: isGrpcClient}

	if isGrpcClient {
		var err error
		h.grpcClientConn, err = grpc.Dial(fmt.Sprintf("127.0.0.1:%d", otherport), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed to create gRPC client connection: %s", err.Error())
		}
	}

	mux := mux.NewRouter()
	mux.HandleFunc("/user", h.userDataHandler).Methods("POST")

	return &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		Handler: mux,
	}
}
