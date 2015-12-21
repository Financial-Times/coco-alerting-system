package main

import (
	"os"

	"time"

	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"

	"github.com/Financial-Times/coco-alerting-system/actions"
	"github.com/Financial-Times/coco-alerting-system/rules"
	"github.com/Financial-Times/coco-alerting-system/sources"
	"github.com/kr/pretty"
)

const logPattern = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC

var configFileName = flag.String("config", "", "Path to configuration file")
var info *log.Logger
var warn *log.Logger

var sourceList []sources.Source

type Config struct {
	Slack SlackConfig `json:"slack"`
}
type SlackConfig struct {
	Username  string `json:"username"`
	IconEmoji string `json:"emoji"`
	Hook      string `json:"hook"`
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

	connParams := sources.ConnectionParams{
		Url:      "http://yes",
		User:     "foo",
		Password: "bar",
	}

	elasticSource := sources.NewElasticSource(connParams)
	splunkSource := sources.NewSplunkSource(connParams)
	sourceList = append(sourceList, elasticSource)
	sourceList = append(sourceList, splunkSource)

	info.Printf("Using sources: %# v\n", pretty.Formatter(sourceList))

	binaryIngesterErrors := rules.StringMatch{}
	binaryIngesterErrors.
		Named("bineryIngesterErrors").
		InEnvironment("prod-us").
		ForService("binary-ingester").
		Earliest(10, time.Minute).
		LatestNow().
		MatchCountGreaterThan("error", 5)

	info.Printf("Using rule: %# v\n", pretty.Formatter(binaryIngesterErrors))

	action := actions.NewSendEmailAction("server", "recipient", "sender", "subject", "body")
	info.Printf("Using action: %# v\n", pretty.Formatter(action))

	slackMessage := actions.NewSlackMessage(appConfig.Slack.Username, appConfig.Slack.IconEmoji, appConfig.Slack.Hook)
	actionResult := slackMessage.Execute("Demo message from app")
	info.Printf("Result of slack message send: [%v]", actionResult)

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
