package importer

import (
	"fmt"
	"os"
	"time"

	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/jteeuwen/go-pkg-xmlx"
)

func Run(feedURLs []string) {
	fmt.Println("running!!")
	for _, f := range feedURLs {
		go PollFeed(f, 5, nil)
	}

}

func PollFeed(uri string, timeout int, cr xmlx.CharsetFunc) {
	feed := rss.NewWithHandlers(timeout, true, chanHandler, itemHandler)

	for {
		if err := feed.Fetch(uri, cr); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %s: %s", uri, err)
			return
		}

		<-time.After(time.Duration(feed.SecondsTillUpdate() * 1e9))
	}
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	fmt.Printf("%d new item(s) in %s\n", len(newitems), feed.Url)
}
