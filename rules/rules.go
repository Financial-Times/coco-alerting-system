package rules

type Rule struct {
	Name      string      `json: "name"`
	Query     interface{} `json: "query"`
	Interval  int         `json: "interval"`
	Threshold int         `json:"threshold"`
}
