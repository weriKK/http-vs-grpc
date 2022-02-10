package main

import (
	"flag"
	"log"
	"usersvc/internal/httpsvc"
)

var (
	servicePort      int
	otherServicePort int
	isGrpcClient     bool
)

func init() {
	flag.IntVar(&servicePort, "port", 8080, "listening port of this service (User)")
	flag.IntVar(&otherServicePort, "otherport", 8081, "listening port of other service (Car)")
	flag.BoolVar(&isGrpcClient, "grpcclient", false, "use grpc for connecting to Car service. Default is http/1.1")
	flag.Parse()
}

func main() {
	if isGrpcClient {
		log.Printf("Listening on 127.0.0.1:%d via HTTP/1.1. Car service on 127.0.0.1:%d via gRPC\n", servicePort, otherServicePort)
	} else {
		log.Printf("Listening on 127.0.0.1:%d via HTTP/1.1. Car service on 127.0.0.1:%d via HTTP/1.1\n", servicePort, otherServicePort)
	}

	s := httpsvc.NewHTTPSvc(servicePort, otherServicePort, isGrpcClient)
	log.Fatal(s.ListenAndServe())

}
