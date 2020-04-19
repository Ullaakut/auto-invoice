package invoice

import (
	"io/ioutil"
	"fmt"

	"github.com/Ullaakut/auto-invoice/pkg/workhours"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Client          Client           `yaml:"client"`
	Self            Self             `yaml:"self"`
	TemplatePath    string           `yaml:"template_path"`
	WorkHoursSource workhours.Source `yaml:"work_hours_source"`

	Year  int
	Month int
}

type Client struct {
	Company         Company `yaml:"company"`
	TaxNumber       string  `yaml:"tax_number"`
	VATNumber       string  `yaml:"vat_number"`
	MaxPaymentDelay string  `yaml:"max_payment_delay"`
}

type Company struct {
	Name      string  `yaml:"name"`
	Address   Address `yaml:"address"`
	TaxNumber string  `yaml:"tax_number"`
	VATNumber string  `yaml:"vat_number"`
	Bank      Bank    `yaml:"bank"`
}

type Self struct {
	Company Company `yaml:"company"`
	Service Service `yaml:"service"`
}

type Address struct {
	Country     string `yaml:"country"`
	State       string `yaml:"state"`
	City        string `yaml:"city"`
	PostalCode  string `yaml:"postal_code"`
	Street      string `yaml:"street"`
	HouseNumber string `yaml:"house_number"`
}

type Bank struct {
	Name  string `yaml:"name"`
	BIC   string `yaml:"bic"`
	IBAN  string `yaml:"iban"`
	SWIFT string `yaml:"swift"`
}

type Service struct {
	Name     string `yaml:"name"`
	Currency string `yaml:"currency"`
	PerHour  int    `yaml:"per_hour"`
	PerMonth int    `yaml:"per_month"`
}

func ParseConfig(path string) (Config, error) {
	var cfg Config

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("unable to open configuration: %w", err)
	}

	err = yaml.Unmarshal(contents, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("unable to parse configuration: %w", err)
	}

	return cfg, nil
}
