package rules

import "time"

type Rule interface {
	Named(name string) Rule
	MatchesString(keyword string) Rule
	MatchCountGreaterThan(keyword string, count int) Rule
	MatchCountLessThan(keyword string, count int) Rule
	Earliest(count int, unit time.Duration) Rule
	Latest(count int, unit time.Duration) Rule
	LatestNow() Rule
	ForService(serviceName string) Rule
	InEnvironment(environmentName string) Rule
}
