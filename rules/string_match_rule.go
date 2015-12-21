package rules

import (
	"time"
)

type StringMatch struct {
	Name       string
	ToMatch    string
	MatchCount int
	Gt         bool
	Lt         bool
	//rename
	EarliestTime    time.Duration
	LatestTime      time.Duration
	IsLatestNow       bool
	ServiceName     string
	EnvironmentName string
}

func (sm *StringMatch) MatchesString(keyword string) Rule {
	sm.ToMatch = keyword
	return sm
}

func (sm *StringMatch) MatchCountGreaterThan(keyword string, count int) Rule {
	sm.ToMatch = keyword
	sm.MatchCount = count
	sm.Gt = true
	return sm
}

func (sm *StringMatch) MatchCountLessThan(keyword string, count int) Rule {
	sm.ToMatch = keyword
	sm.MatchCount = count
	sm.Lt = true
	return sm
}

func (sm *StringMatch) Earliest(count int, unit time.Duration) Rule {
	sm.EarliestTime = time.Duration(count) * unit
	return sm
}

func (sm *StringMatch) Latest(count int, unit time.Duration) Rule {
	sm.LatestTime = time.Duration(count) * unit
	return sm
}

func (sm *StringMatch) LatestNow() Rule {
	sm.IsLatestNow = true
	return sm
}
func (sm *StringMatch) ForService(serviceName string) Rule {
	sm.ServiceName = serviceName
	return sm
}

func (sm *StringMatch) InEnvironment(environmentName string) Rule {
	sm.EnvironmentName = environmentName
	return sm
}

func (sm *StringMatch) Named(name string) Rule {
	sm.Name = name
	return sm
}
