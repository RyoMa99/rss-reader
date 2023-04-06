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
		f, err := SearchFeed(Site, time.Duration(conf.Range)*time.Hour)
		if err != nil {
			logger.Error("failed to fetch rss", zap.String("サイト名", Site.Name))
		}
		logger.Info(fmt.Sprintf("new feed count: %d", len(f)), zap.String("サイト名", Site.Name))
		feed = append(feed, f...)
	}
	logger.Info(fmt.Sprintf("total feed count: %d", len(feed)))
}

func SearchFeed(site constant.Site, searchRange time.Duration) ([]Feed, error) {
	fp := gofeed.NewParser()
	fd, err := fp.ParseURL(site.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	feed := make([]Feed, 0, len(fd.Items))
	for _, item := range fd.Items {
		publishedTime, err := time.Parse(site.DateTimeFormat, item.Published)
		if err != nil {
			continue
		}
		if publishedTime.Add(searchRange).After(time.Now()) {
			feed = append(feed, Feed{
				Title: item.Title,
				URL:   item.Link,
			})
		}
	}

	return feed, nil
}
