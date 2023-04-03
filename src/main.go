package main

import (
	"app/constant"
	"app/infra/config"
	"app/infra/logger"
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
	"go.uber.org/zap"
)

type Feed struct {
	Title string
	URL   string
}

func main() {
	// initialize
	logger.Init()
	conf := config.Init()

	var feed []Feed
	for _, Site := range constant.SITE_LIST {
		logger.Info("Search new feed", zap.String("サイト名", Site.Name))
		f := SearchFeed(Site, time.Duration(-1*conf.Range)*time.Hour)
		logger.Info(fmt.Sprintf("new feed count: %d", len(f)), zap.String("サイト名", Site.Name))
		feed = append(feed, f...)
	}
	logger.Info(fmt.Sprintf("total feed count: %d", len(feed)))
}

func SearchFeed(site constant.Site, searchRange time.Duration) []Feed {
	var feed []Feed

	fp := gofeed.NewParser()
	fd, _ := fp.ParseURL(site.URL)

	for _, item := range fd.Items {
		t, _ := time.Parse(site.DateTimeFormat, item.Published)
		if time.Now().Add(searchRange).Before(t) {
			feed = append(feed, Feed{
				Title: item.Title,
				URL:   item.Link,
			})
		}
	}

	return feed
}
