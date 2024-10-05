package handlers

import (
	"encoding/json"
	"go-tweet-stream/database"
	"go-tweet-stream/tweets"
	"log"
	"net/http"
	"strconv"
)

type JsonResponse struct {
	Type       string `json:"type"`
	TweetCount int    `json:"tweet_count"`
	Message    string `json:"message"`
}

type AirflowResponse struct {
	Type    string          `json:"type"`
	Message json.RawMessage `json:"message"`
}

// Handler for tweets
func HandleTweets(w http.ResponseWriter, r *http.Request) {

	duration := r.FormValue("duration")

	var response = JsonResponse{}
	if duration == "" {
		response = JsonResponse{Type: "error", Message: "You are missing duration parameter."}
	} else {
		dur_min, err := strconv.ParseInt(duration, 10, 64)

		if err != nil {
			log.Fatal(err)
			return
		}

		if dur_min < 1 {
			response = JsonResponse{Type: "error", Message: "Duratiion paramater must be an integer greater than or equal to 0"}
		} else {

			db := database.SetupDB()
			query := r.FormValue("query")
			tweetCount := tweets.GetTweets(int(dur_min), query, db)
			response = JsonResponse{
				Type: "success", 
				TweetCount: tweetCount,Message: "Stream successful"
			}
		}
		json.NewEncoder(w).Encode(response)
	}
}
