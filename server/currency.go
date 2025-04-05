package server

import (
	"context"

	"github.com/dolmatovDan/gRPC/currency"
	"github.com/dolmatovDan/gRPC/data"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	rates *data.ExchangeRates
	log hclog.Logger
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

func NewCurrency(r *data.ExchangeRates,l hclog.Logger) *Currency {
	return &Currency{r, l, currency.UnimplementedCurrencyServer{}}
}
