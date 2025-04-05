package data

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-hclog"
	"golang.org/x/text/encoding/charmap"
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

func (e *ExchangeRates) GetRate(base, dest string) (float64, error) {
	br, ok := e.rates[base]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", base)
	}
	dr, ok := e.rates[dest]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", base)
	}


	return dr / br, nil
}

func (e *ExchangeRates) getRates() error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://www.cbr.ru/scripts/XML_daily.asp", nil)
	if err != nil {
		e.log.Error("Failed request")
		return err
	}

	// avoid problems with security
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("Accept", "text/xml")

	resp, err := client.Do(req)
	if err != nil {
		e.log.Error("Reading data error", err)
		return err
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		e.log.Error("Unknown encoding")
		return nil, fmt.Errorf("Unknown encoding: %s", charset)
	}

	var valCurs ValCurs
	err = decoder.Decode(&valCurs)
	if err != nil {
		e.log.Error("Parse XML error", err)
		return err
	}

	for _, valute := range valCurs.Valutes {
		e.rates[valute.CharCode], err = strconv.ParseFloat(strings.Replace(valute.Value, ",", ".", -1), 64)
		if err != nil {
			e.log.Error("Incorrect rate value", valute.Value)
		}
	}

	e.rates["RUB"] = 1

	return nil
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	XMLName  xml.Name `xml:"Valute"`
	ID       string   `xml:"ID,attr"`
	NumCode  string   `xml:"NumCode"`
	CharCode string   `xml:"CharCode"`
	Nominal  int      `xml:"Nominal"`
	Name     string   `xml:"Name"`
	Value    string   `xml:"Value"`
}
