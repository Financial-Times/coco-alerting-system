package watchers

import (
	"time"

	"github.com/Financial-Times/coco-alerting-system/actions"
	"github.com/Financial-Times/coco-alerting-system/rules"
	"github.com/Financial-Times/coco-alerting-system/sources"
)

type StringMatchWatcher struct{}

func (smw *StringMatchWatcher) Watch(rule rules.Rule, source sources.Source, actions []actions.Action) {
	tickerChan := time.NewTicker(time.Duration(rule.Interval) * time.Minute)
	for {
		result := source.Execute(rule.Query)
		if result != nil && result.Hits >= rule.Threshold {
			for _, action := range actions {
				action.Execute(rule, result.Body)
			}
		}
		select {
		case <-tickerChan.C:
			continue
		}
	}

}
