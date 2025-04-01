package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"github.com/dolmatovDan/gRPC/currency"
	"github.com/dolmatovDan/gRPC/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if false {
		log := hclog.Default()

		gs := grpc.NewServer()
		cs := server.NewCurrency(log)

		currency.RegisterCurrencyServer(gs, cs)

		reflection.Register(gs)

		l, err := net.Listen("tcp", ":9092")
		if err != nil {
			log.Error("Unable to listen", "error", err)
			os.Exit(1)
		}

		gs.Serve(l)
	}

	req, _ := http.NewRequest("GET", "https://www.cbr.ru/scripts/XML_daily.asp", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(body)
	fmt.Println("Status:", resp.Status)
	fmt.Println("Body length:", len(body))
}
