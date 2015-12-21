package sources

type (
	splunkSource struct {
		connectionParams ConnectionParams
	}
)

func NewSplunkSource(connectionParams ConnectionParams) Source {
	return &splunkSource{connectionParams: connectionParams}
}

func (source *splunkSource) ExecuteQuery(query string) string {
	return "this is the result for query: " + query
}
