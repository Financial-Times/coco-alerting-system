package sources

type Source interface {
	ExecuteQuery(query string) string
}

type ConnectionParams struct {
	Url      string
	User     string
	Password string
}
