package workhours

import "time"

type Report struct {
	TotalGrand      int               `json:"total_grand"`
	TotalBillable   int               `json:"total_billable"`
	TotalCount      int               `json:"total_count"`
	PerPage         int               `json:"per_page"`
	TotalCurrencies []TotalCurrencies `json:"total_currencies"`
	Data            []Data            `json:"data"`

	Start time.Time
	End time.Time
}

type TotalCurrencies struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type Data struct {
	ID          int         `json:"id"`
	Pid         int         `json:"pid"`
	Tid         interface{} `json:"tid"`
	UID         int         `json:"uid"`
	Description string      `json:"description"`
	Start       string      `json:"start"`
	End         string      `json:"end"`
	Updated     string      `json:"updated"`
	Dur         int         `json:"dur"`
	User        string      `json:"user"`
	UseStop     bool        `json:"use_stop"`
	Client      string      `json:"client"`
	Project     string      `json:"project"`
	Task        interface{} `json:"task"`
	Billable    float64     `json:"billable"`
	IsBillable  bool        `json:"is_billable"`
	Cur         string      `json:"cur"`
	Tags        []string    `json:"tags"`
}
