package sources

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type elasticSource struct {
	host       string
	authCookie string
	httpClient http.Client
}

type esResult struct {
	Hits Hits `json:"hits"`
}

type Hits struct {
	Total   int     `json:"total"`
	Entries []Entry `json:"hits"`
}

type Entry struct {
	Fields interface{} `json:"fields"`
}

func NewElasticSource(host string, authCookie string) Source {
	return &elasticSource{host, authCookie, http.Client{}}
}

func (source *elasticSource) Execute(query interface{}) *Result {
	q, _ := json.Marshal(query)
	r, err := http.NewRequest("GET", source.host+"/_search", bytes.NewReader(q))
	if err != nil {
		log.Printf("Error creating request: %s", err.Error())
		return nil
	}
	r.Header.Add("Cookie", source.authCookie)
	resp, err := source.httpClient.Do(r)
	if err != nil {
		log.Printf("Error executing request: %s", err.Error())
		return nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Unexpected response status %d. Expected: 200.", resp.StatusCode)
		return nil
	}

	rawResult, _ := ioutil.ReadAll(resp.Body)
	var esr esResult
	err = json.Unmarshal(rawResult, &esr)

	body := ""
	for _, e := range esr.Hits.Entries {
		fields, _ := json.Marshal(e.Fields)
		body += string(fields) + "\n"
	}
	return &Result{esr.Hits.Total, body}
}
