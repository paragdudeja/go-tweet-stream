package tweets

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	hashtag "go-tweet-stream/hashtags"
	"go-tweet-stream/queries"
	"go-tweet-stream/rules"
	"log"
	"net/http"
	"os"
	"time"
)

type Data struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Text      string `json:"text"`
}

type TweetResponse struct {
	Data Data `json:"data"`
}

func GetTweets(duration int, query string, db *sql.DB) int {
	var url string
	if len(query) > 0 {
		url = "https://api.twitter.com/2/tweets/search/stream?tweet.fields=created_at,geo&expansions=author_id,geo.place_id"
		// reset rules
		var ids []string = rules.GetRules()
		if len(ids) > 0 {
			rules.DeleteRules(ids)
		}
		log.Println("Query :", query)
		rules.AddRules(query)
	} else {
		url = "https://api.twitter.com/2/tweets/sample/stream?tweet.fields=created_at,geo"
	}

	client := &http.Client{Timeout: time.Duration(duration) * time.Minute}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	req.Header.Add("Authorization", rules.Bearer())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer res.Body.Close()

	log.Println("Opening stream")
	reader := bufio.NewReader(res.Body)
	count := 0

	queries.ClearDB(db)
	for {
		body, err := reader.ReadBytes('\n')
		if err != nil {
			if os.IsTimeout(err) {
				log.Println("Closing Stream")
			} else {
				fmt.Println(err)
			}
			break
		}

		var jsonData TweetResponse
		err = json.Unmarshal(body, &jsonData)
		if err != nil {
			fmt.Println(err)
		}

		hashtags := hashtag.ExtractHashtags(jsonData.Data.Text)

		go queries.AddToDB(db, jsonData.Data.Id, jsonData.Data.CreatedAt, jsonData.Data.Text, hashtags)
		count++
	}

	return count
}
