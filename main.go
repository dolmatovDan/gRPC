package main

import (
	"net"
	"os"

	"github.com/dolmatovDan/gRPC/currency"
	"github.com/dolmatovDan/gRPC/data"
	"github.com/dolmatovDan/gRPC/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	rates, err := data.NewRates(log)
	if err != nil {
		log.Error("Unable to get rates", "error", err)
		os.Exit(1)
	}

	gs := grpc.NewServer()
	c := server.NewCurrency(rates, log)

	currency.RegisterCurrencyServer(gs, c)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	log.Info("Listening on port :9092")

	gs.Serve(l)
}
