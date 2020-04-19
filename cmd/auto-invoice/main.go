package main

import (
	"os"
	"time"

	"github.com/Ullaakut/disgo"
	"github.com/Ullaakut/disgo/style"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/Ullaakut/auto-invoice/pkg/invoice"
)

func parseArguments() error {
	defaultYear := time.Now().Year()
	defaultMonth := int(time.Now().Month()) -1
	if defaultMonth == 0 {
		defaultMonth = 12
		defaultYear = defaultYear - 1
	}

	pflag.StringP("config-file", "f", "./config/config.yaml", "Configuration file to use")
	pflag.IntP("year", "y", defaultYear, "Year for which to generate an invoice")
	pflag.IntP("month", "m", defaultMonth, "Month for which to generate an invoice (default is last month)")
	pflag.Int("override-price", 0, "Override the usual invoice price")
	pflag.BoolP("debug", "d", true, "Enable the debug logs")
	pflag.BoolP("help", "h", false, "Displays this help message")

	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return err
	}

	if viper.GetBool("help") {
		pflag.Usage()
		os.Exit(0)
	}

	return nil
}

func main() {
	err := parseArguments()
	if err != nil {
		printErr(err)
	}

	cfg, err := invoice.ParseConfig(viper.GetString("config-file"))
	if err != nil {
		printErr(err)
	}

	cfg.Year = viper.GetInt("year")
	cfg.Month = viper.GetInt("month")

	if viper.GetInt("override-price") != 0 {
		cfg.Self.Service.PerMonth = viper.GetInt("override-price")
	}

	err = invoice.Generate(cfg)
	if err != nil {
		printErr(err)
	}
}

func printErr(err error) {
	disgo.Errorln(style.Failure(style.SymbolCross), err)
	os.Exit(1)
}
