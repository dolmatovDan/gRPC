package server

import (
	"context"
	"io"
	"time"

	"github.com/dolmatovDan/gRPC/currency"
	"github.com/dolmatovDan/gRPC/data"
	"google.golang.org/grpc"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	rates *data.ExchangeRates
	log   hclog.Logger
	currency.UnimplementedCurrencyServer
}

func (c *Currency) GetRate(ctx context.Context, rr *currency.RateRequest) (*currency.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &currency.RateResponse{Rate: rate}, nil
}

func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{r, l, currency.UnimplementedCurrencyServer{}}
}

func (c *Currency) SubscribeRates(src grpc.BidiStreamingServer[currency.RateRequest, currency.RateResponse]) error {

	go func() {
		for {
			rr, err := src.Recv()
			if err == io.EOF {
				c.log.Info("Clieant has closed connection")
				break
			}
			if err != nil {
				c.log.Error("Unable to read from client", "error", err)
				break
			}

			c.log.Info("Handle client request", "request", rr)
		}
	}()

	for {
		err := src.Send(&currency.RateResponse{Rate: 12.1})
		if err != nil {
			return err
		}

		time.Sleep(5 * time.Second)
	}
}
