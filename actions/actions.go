package actions

import "github.com/Financial-Times/coco-alerting-system/rules"

type Action interface {
	Execute(rule rules.Rule, parameters string) string
}
