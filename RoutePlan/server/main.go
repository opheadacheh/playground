package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"route_plan/server/service"
	"route_plan/server/utils"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	servicepb "route_plan/server/proto/service"
)

var (
	CertFilePath = "/root/server-cert.pem"
	KeyFilePath  = "/root/server-key.pem"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gaodeClient, err := utils.NewGaodeHttpClient(nil)
	if err != nil {
		log.Fatalf("failed to init gaode client: %v", err)
	}

	s := grpc.NewServer()
	servicepb.RegisterRoutePlanServer(s, &service.Server{
		GaodeClient: gaodeClient,
	})

	log.Println("Serving gRPC on localhost:50051")
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:50051",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = servicepb.RegisterRoutePlanHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// load tls certificates
	serverTLSCert, err := tls.LoadX509KeyPair(CertFilePath, KeyFilePath)
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverTLSCert},
	}
	gwServer := &http.Server{
		Addr:      "0.0.0.0:8080",
		Handler:   gwmux,
		TLSConfig: tlsConfig,
	}

	log.Println("Serving gRPC-Gateway on localhost:8080")
	if err := gwServer.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("failed to serve gRPC-gateway: %v", err)
	}
}
