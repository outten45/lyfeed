package importer

import (
	"fmt"
	"lyfeed"

	"os"
	"time"

	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/jteeuwen/go-pkg-xmlx"
)

func Run(c *lyfeed.Context, feedURLs []string) {
	for _, f := range feedURLs {
		go PollFeed(c, f, 5, nil)
	}
}

func PollFeed(c *lyfeed.Context, uri string, timeout int, cr xmlx.CharsetFunc) {
	h := &RssHandler{Context: c}
	feed := rss.NewWithHandlers(timeout, true, h, h)

	for {
		if err := feed.Fetch(uri, cr); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %s: %s", uri, err)
			return
		}

		<-time.After(time.Duration(feed.SecondsTillUpdate() * 1e9))
	}
}

type RssHandler struct {
	Context *lyfeed.Context
}

func (mh *RssHandler) ProcessChannels(feed *rss.Feed, newchannels []*rss.Channel) {
	// fmt.Printf("----------------------------------------\nchannel: \n%+v\n\n\n", newchannels[0])
	fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func (mn *RssHandler) ProcessItems(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	fmt.Printf("%d new item(s) in %s\n", len(newitems), feed.Url)
	fmt.Printf("---------------------------------------------------------\n")
	fmt.Printf("%+v\n", newitems[0])
	fmt.Printf("---------------------------------------------------------\n")
}
