package main

import (
	"app/constant"
	"app/infra/config"
	"app/infra/logger"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/mmcdole/gofeed"
	"go.uber.org/zap"
)

type Feed struct {
	Title string
	URL   string
}

var conf *config.Config

func main() {
	// initialize
	logger.Init()
	conf = config.Init()

	feeds := searchFeedConcurrently()
	logger.Info(fmt.Sprintf("total feed count: %d", len(feeds)))

	client, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		logger.Panic("failed to connect line bot")
	}

	if err := pushFeed2Line(client, feeds); err != nil {
		logger.Panic(fmt.Sprintf("failed to push to line: %s", err.Error()))
	}
	logger.Info("fin")
}

func pushFeed2Line(client *linebot.Client, feeds []Feed) error {
	for _, feed := range feeds {
		message := linebot.NewTextMessage(feed.Title + "\n\n" + feed.URL)

		if _, err := client.PushMessage(os.Getenv("CHANNEL_ID"), message).Do(); err != nil {
			return err
		}
	}
	return nil
}

func searchFeedConcurrently() []Feed {
	var wg sync.WaitGroup
	c := make(chan []Feed, len(constant.SITE_LIST))

	for _, site := range constant.SITE_LIST {
		logger.Info("Search new feed", zap.String("サイト名", site.Name))

		go func(s constant.Site) {
			defer wg.Done()
			feeds, err := SearchFeed(s, time.Duration(conf.Range)*time.Hour)
			if err != nil {
				logger.Error(err.Error())
			}
			c <- feeds
		}(site)

		wg.Add(1)
	}
	wg.Wait()
	close(c)

	feeds := make([]Feed, 0)
	for feed := range c {
		feeds = append(feeds, feed...)
	}
	return feeds
}

func SearchFeed(site constant.Site, searchRange time.Duration) ([]Feed, error) {
	fp := gofeed.NewParser()
	fd, err := fp.ParseURL(site.URL)
	if err != nil {
		return nil, err
	}

	feeds := make([]Feed, 0, len(fd.Items))
	for _, item := range fd.Items {
		publishedTime, err := time.Parse(site.DateTimeFormat, item.Published)
		if err != nil {
			continue
		}
		if publishedTime.Add(searchRange).After(time.Now()) {
			feeds = append(feeds, Feed{
				Title: item.Title,
				URL:   item.Link,
			})
		}
	}
	return feeds, nil
}
