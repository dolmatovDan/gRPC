package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/go-hclog"
)

type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

func NewRates(l hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{log: l, rates: map[string]float64{}}

	err := er.getRates()

	return er, err
}

func (e *ExchangeRates) getRates() error {
	req, _ := http.NewRequest("GET", "https://www.cbr.ru/scripts/XML_daily.asp", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return fmt.Errorf("Expected error code 200 got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	md := &ValCurs{}
	err = xml.NewDecoder(resp.Body).Decode(md)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(len(md.Rates))

	for _, c := range md.Rates {
		r, err := strconv.ParseFloat(c.Value, 64)
		if err != nil {
			return err
		}
		e.rates[c.Name] = r
	}
	return nil
}

type ValCurs struct {
	Date  string   `xml:"Date,attr"`
	Rates []Valute `xml:"Valute"`
}

type Valute struct {
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}
