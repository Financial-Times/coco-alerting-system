package sources

type (
	elasticSource struct {
		connectionParams ConnectionParams
	}
)

func NewElasticSource(connectionParams ConnectionParams) Source {
	return &elasticSource{connectionParams: connectionParams}
}

func (source *elasticSource) ExecuteQuery(query string) string {
	return "this is the result for query: " + query
}
