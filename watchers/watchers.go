package watchers

import (
	"github.com/Financial-Times/coco-alerting-system/actions"
	"github.com/Financial-Times/coco-alerting-system/rules"
	"github.com/Financial-Times/coco-alerting-system/sources"
)

type Watcher interface {
	//has to be run concurrently
	Watch(rule rules.Rule, source sources.Source, actions []actions.Action)
}
