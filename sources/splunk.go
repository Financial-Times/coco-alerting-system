package sources

type (
	splunkSource struct {
	}
)

func NewSplunkSource() Source {
	return &splunkSource{}
}

func (source *splunkSource) Execute(query interface{}) *Result {
	return &Result{0, "No results"}
}
