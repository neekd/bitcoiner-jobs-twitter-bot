package main

import (
	"os"
	"time"

	"github.com/swan-bitcoin/bitcoiner-jobs-twitter-bot/pkg/config"
	"github.com/swan-bitcoin/bitcoiner-jobs-twitter-bot/pkg/tweet"

	"github.com/mmcdole/gofeed"
	log "github.com/sirupsen/logrus"
)

var (
	conf *config.Config
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	var err error
	conf, err = config.Load()
	if err != nil {
		log.Fatalf("Unable to parse config: %s\n", err.Error())
	}

	if err := os.Setenv("GOTWI_API_KEY", conf.ClientKey); err != nil {
		log.Fatalf("Unable to set env var: %s\n", err.Error())
	}
	if err := os.Setenv("GOTWI_API_KEY_SECRET", conf.ClientKeySecret); err != nil {
		log.Fatalf("Unable to set env var: %s\n", err.Error())
	}
}

func main() {
	agent, err := tweet.NewAgent(conf.OAuthToken, conf.OAuthSecret)
	if err != nil {
		log.Fatalf("Unable to load the tweet agent: %s\n", err.Error())
	}

	var lastItemGUID string
	for {
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(conf.RSSFeedURL)
		if err != nil {
			log.Errorf("Unable to read feed: %s", err.Error())
			continue
		}
		currentItem := feed.Items[0]
		if currentItem.GUID != lastItemGUID {
			fTweet := tweet.FormatFeedItem(currentItem, conf.HashTags)
			if err := agent.SendTweet(fTweet); err != nil {
				log.Errorf("Unable to send tweet: %s\n", err.Error())
			}
		}
		lastItemGUID = currentItem.GUID
		time.Sleep(time.Minute * 5)
	}
}
