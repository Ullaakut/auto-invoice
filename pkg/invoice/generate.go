package invoice

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/Ullaakut/auto-invoice/pkg/workhours"
	"github.com/jinzhu/now"
	"github.com/tealeg/xlsx"
)

func Generate(config Config) error {
	report, err := workhours.Get(config.WorkHoursSource)
	if err != nil {
		return fmt.Errorf("unable to fetch work hours from source: %w", err)
	}

	report.Start = time.Date(config.Year, time.Month(config.Month), 0, 0, 0, 0, 0, time.UTC)
	report.End = now.With(report.Start).EndOfMonth()

	tmpl, err := xlsx.OpenFile(config.TemplatePath)
	if err != nil {
		return fmt.Errorf("unable to open XLSX template: %w", err)
	}

	fmap := sprig.TxtFuncMap()

	sheet := tmpl.Sheets[0]

	for row := 0; row < sheet.MaxRow; row++ {
		for col := 0; col < sheet.MaxCol; col++ {
			cell, _ := sheet.Cell(row, col)
			if cell.String() == "" {
				continue
			}

			t := template.Must(template.New("invoice").Funcs(fmap).Parse(cell.String()))
			var buf bytes.Buffer

			err := t.Execute(&buf, map[string]interface{}{
				"Client":  config.Client,
				"Self":    config.Self,
				"Invoice": report,
			})
			if err != nil {
				return fmt.Errorf("unable to render template: %w", err)
			}

			cell.SetString(buf.String())
		}
	}

	err = tmpl.Save("output/" + sheet.Name + ".xlsx")
	if err != nil {
		return fmt.Errorf("unable to save template to path %q: %w", "output/fineko.xlsx", err)
	}

	return nil
}
