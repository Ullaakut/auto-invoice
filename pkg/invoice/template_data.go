package invoice

import "github.com/Ullaakut/auto-invoice/pkg/workhours"

type TemplateData struct {
	Client    Client           `json:"client"`
	Self      Self             `json:"self"`
	WorkHours workhours.Report `json:"work_hours"`
}

