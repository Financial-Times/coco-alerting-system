package sources

type Source interface {
	Execute(query interface{}) *Result
}

type Result struct {
	Hits int
	Body string
}
