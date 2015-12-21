package watchers

import (
	"time"

	"fmt"

	"github.com/Financial-Times/coco-alerting-system/actions"
	"github.com/Financial-Times/coco-alerting-system/rules"
	"github.com/Financial-Times/coco-alerting-system/sources"
)

type StringMatchWatcher struct{}

func (smw *StringMatchWatcher) Watch(rule rules.StringMatch, source sources.Source, actions []actions.Action) {
	tickerChan := time.NewTicker(rule.Earliest)
	for {
		queryString := "Generate from rule - maybe the source can do this?"
		result := source.ExecuteQuery(queryString)
		//interpret result
		fmt.Println(result)

		for _, action := range actions {
			//figure out what to send to the result
			action.Execute(result)
		}
		select {
		case <-tickerChan.C:
			continue
		}
	}

}
