package main

import (
	"carsvc/internal/grpcserver"
	"carsvc/internal/httpserver"
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	servicePort int
	grpcServer  bool
)

func init() {
	flag.IntVar(&servicePort, "port", 8081, "listening port of this service (CarSvc)")
	flag.BoolVar(&grpcServer, "grpcserver", false, "use grpc for communication. Default is http/1.1")
	flag.Parse()
}

func main() {

	if grpcServer {
		l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", servicePort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpcserver.NewGRPCServer()

		log.Printf("Listening on 127.0.0.1:%d using gRPC.\n", servicePort)
		log.Fatal(s.Serve(l))

	} else {
		s := httpserver.NewHTTPServer(servicePort)

		log.Printf("Listening on 127.0.0.1:%d using HTTP/1.1.\n", servicePort)
		log.Fatal(s.ListenAndServe())
	}
}
