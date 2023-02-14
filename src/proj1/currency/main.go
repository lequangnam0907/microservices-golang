package main

import (
	"fmt"
	"net"
	"os"

	protos "proj1/currency/protos/currency"
	"proj1/currency/server"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	// tao mot gRPC server moi
	gs := grpc.NewServer()

	c := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, c)

	reflection.Register(gs)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	// listen yeu cau
	gs.Serve(l)
}
