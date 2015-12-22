package main

import (
	"os"

	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"

	"github.com/Financial-Times/coco-alerting-system/actions"
	"github.com/Financial-Times/coco-alerting-system/rules"
	"github.com/Financial-Times/coco-alerting-system/sources"
	"github.com/Financial-Times/coco-alerting-system/watchers"
	"github.com/kr/pretty"
	"os/signal"
	"syscall"
)

const logPattern = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC

var configFileName = flag.String("config", "", "Path to configuration file")
var info *log.Logger
var warn *log.Logger
var sourceList []sources.Source

type Config struct {
	Slack         SlackConfig   `json:"slack"`
	ElasticSearch ElasticSearch `json:"elastic"`
	Rules         []rules.Rule  `json:"rules"`
}
type SlackConfig struct {
	Username  string `json:"username"`
	IconEmoji string `json:"emoji"`
	Hook      string `json:"hook"`
}

type ElasticSearch struct {
	Host   string `json:"host"`
	Cookie string `json:"cookie"`
}

func main() {
	initLogs(os.Stdout, os.Stdout, os.Stderr)
	flag.Parse()

	var err error
	appConfig, err := ParseConfig(*configFileName)
	if err != nil {
		log.Printf("Cannot load configuration: [%v]", err)
		return
	}

	elasticSource := sources.NewElasticSource(appConfig.ElasticSearch.Host, appConfig.ElasticSearch.Cookie)
	sourceList = append(sourceList, elasticSource)

	info.Printf("Using sources: %# v\n", pretty.Formatter(sourceList))
	slackMessage := actions.NewSlackMessage(appConfig.Slack.Username, appConfig.Slack.IconEmoji, appConfig.Slack.Hook)
	watcher := watchers.StringMatchWatcher{}
	for _, r := range appConfig.Rules {
		go watcher.Watch(r, elasticSource, []actions.Action{slackMessage})
	}
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}

// ParseConfig opens the file at configFileName and unmarshals it into an AppConfig.
func ParseConfig(configFileName string) (*Config, error) {
	file, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Printf("Error reading configuration file [%v]: [%v]", configFileName, err.Error())
		return nil, err
	}

	var conf Config
	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Printf("Error unmarshalling configuration file [%v]: [%v]", configFileName, err.Error())
		return nil, err
	}

	info.Printf("Using configuration: %# v", pretty.Formatter(conf))
	return &conf, nil
}

func initLogs(infoHandle io.Writer, warnHandle io.Writer, panicHandle io.Writer) {
	//to be used for INFO-level logging: info.Println("foo is now bar")
	info = log.New(infoHandle, "INFO  - ", logPattern)
	//to be used for WARN-level logging: warn.Println("foo is now bar")
	warn = log.New(warnHandle, "WARN  - ", logPattern)

	log.SetFlags(logPattern)
	log.SetPrefix("ERROR - ")
	log.SetOutput(panicHandle)
}
