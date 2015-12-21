package main

import (
	"fmt"

	"time"

	"github.com/Financial-Times/coco-alerting-system/actions"
	"github.com/Financial-Times/coco-alerting-system/rules"
	"github.com/Financial-Times/coco-alerting-system/sources"
	"github.com/kr/pretty"
)

var sourceList []sources.Source

func main() {
	connParams := sources.ConnectionParams{
		Url:      "http://yes",
		User:     "foo",
		Password: "bar",
	}

	elasticSource := sources.NewElasticSource(connParams)
	splunkSource := sources.NewSplunkSource(connParams)
	sourceList = append(sourceList, elasticSource)
	sourceList = append(sourceList, splunkSource)

	fmt.Printf("Using sources: %# v\n", pretty.Formatter(sourceList))

	binaryIngesterErrors := rules.StringMatch{}
	binaryIngesterErrors.
		Named("bineryIngesterErrors").
		InEnvironment("prod-us").
		ForService("binary-ingester").
		Earliest(10, time.Minute).
		LatestNow().
		MatchCountGreaterThan("error", 5)

	fmt.Printf("Using rule: %# v\n", pretty.Formatter(binaryIngesterErrors))

	action := actions.NewSendEmailAction("server", "recipient", "sender", "subject", "body")
	fmt.Printf("Using action: %# v\n", pretty.Formatter(action))
	
	

}
