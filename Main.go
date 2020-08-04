package main

import (
	"net/url"
	"os"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

func initAnaconda() *anaconda.TwitterApi {
	return anaconda.NewTwitterApiWithCredentials(
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_TOKEN_SECRET"),
		os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_KEY_SECRET"))
}

func main() {

	api := initAnaconda()

	params := url.Values{}
	params.Set("count", "200")
	params.Set("exclude_replies", "false")
	params.Set("include_rts", "false")
	params.Set("trim_user", "false")

	var cursor int64 = -1 // Initial value, which is the first page
	for cursor != 0 {
		params.Set("cursor", strconv.FormatInt(cursor, 10))
		tweets, err := api.GetUserTimeline(params)
		if err != nil {
			panic(err)
		}

		for _, tweet := range tweets {
			if tweet.FavoriteCount == 0 && tweet.RetweetCount == 0 {
				api.DeleteTweet(tweet.Id, true)
			}
		}
		cursor = 0
	}
}
