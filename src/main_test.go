package main

import (
	"app/constant"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestSearchFeed(t *testing.T) {
	mockResponse := `
<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
	<channel>
		<title>Test Channel</title>
		<link>http://example.com/</link>
		<description>Test Channel Description</description>
		<item>
			<title>Test Title 1</title>
			<link>http://example.com/1</link>
			<description>Test Description 1</description>
			<pubDate>` + time.Now().Add(-30*time.Minute).Format(time.RFC1123Z) + `</pubDate>
		</item>
		<item>
			<title>Test Title 2</title>
			<link>http://example.com/2</link>
			<description>Test Description 2</description>
			<pubDate>` + time.Now().Add(-50*time.Minute).Format(time.RFC1123Z) + `</pubDate>
		</item>
		<item>
			<title>Test Title 3</title>
			<link>http://example.com/3</link>
			<description>Test Description 3</description>
			<pubDate>` + time.Now().Add(-2*time.Hour).Format(time.RFC1123Z) + `</pubDate>
		</item>
	</channel>
</rss>
`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/rss+xml")
		fmt.Fprintln(w, mockResponse)
	}))
	defer ts.Close()

	testSite := constant.Site{
		Name:           "mock site",
		URL:            ts.URL,
		DateTimeFormat: time.RFC1123Z,
	}

	cases := []struct {
		testName     string
		searchRange  time.Duration
		expectedFeed []Feed
	}{
		{
			testName:    "searchFeedの第2引数の時間範囲内に追加されたコンテンツのみ取得する。",
			searchRange: time.Hour,
			expectedFeed: []Feed{
				{
					Title: "Test Title 1",
					URL:   "http://example.com/1",
				},
				{
					Title: "Test Title 2",
					URL:   "http://example.com/2",
				},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.testName, func(t *testing.T) {

			actualFeed, err := SearchFeed(testSite, time.Hour)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actualFeed, tt.expectedFeed) {
				t.Errorf("actualFeed = %#v, expectedFeed = %#v", actualFeed, tt.expectedFeed)
			}
		})
	}
}
