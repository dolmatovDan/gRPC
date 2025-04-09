package server

import (
	"context"
	"io"
	"time"

	"github.com/dolmatovDan/gRPC/currency"
	"github.com/dolmatovDan/gRPC/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	rates         *data.ExchangeRates
	log           hclog.Logger
	subscriptions map[currency.Currency_SubscribeRatesServer][]*currency.RateRequest
	currency.UnimplementedCurrencyServer
}

func (c *Currency) GetRate(ctx context.Context, rr *currency.RateRequest) (*currency.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	if rr.Base == rr.Destination {
		err := status.Newf(
			codes.InvalidArgument,
			"Base currency %s can not bet the same as the destination currency %s",
			rr.Base.String(),
			rr.Destination.String(),
		)

		err, wde := err.WithDetails(rr)
		if wde != nil {
			return nil, wde
		}

		return nil, err.Err()
	}

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &currency.RateResponse{Rate: rate, Base: rr.Base, Destination: rr.Destination}, nil
}

func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	c := &Currency{r, l, make(map[currency.Currency_SubscribeRatesServer][]*currency.RateRequest), currency.UnimplementedCurrencyServer{}}
	go c.handleUpdates()

	return c
}

func (c *Currency) handleUpdates() {
	ru := c.rates.MonitorRates(5 * time.Second)

	for range ru {
		c.log.Info("Get updated rates")

		// loop over subscribed clients
		for k, v := range c.subscriptions {

			// loop over rates
			for _, rr := range v {
				rate, err := c.rates.GetRate(rr.Base.String(), rr.Destination.String())
				if err != nil {
					c.log.Error("Unable to get updated rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
					return
				}

				err = k.Send(&currency.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: rate})
				if err != nil {
					c.log.Error("Unable to send updated rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
					return
				}
			}
		}
	}
}

func (c *Currency) SubscribeRates(src grpc.BidiStreamingServer[currency.RateRequest, currency.RateResponse]) error {

	for {
		rr, err := src.Recv()
		if err == io.EOF {
			c.log.Info("Clieant has closed connection")
			break
		}
		if err != nil {
			c.log.Error("Unable to read from client", "error", err)
			return err
		}

		c.log.Info("Handle client request", "request", rr)
		rrs, ok := c.subscriptions[src]
		if !ok {
			rrs = []*currency.RateRequest{}
		}

		rrs = append(rrs, rr)
		c.subscriptions[src] = rrs
	}

	return nil
}
