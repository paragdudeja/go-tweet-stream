package queries

import (
	"database/sql"
	"fmt"
)

func AddToDB(db *sql.DB, tweetid string, timestamp string, tweet string, hashtags []string) {
	// tweets
	sqlStatement := `
INSERT INTO tweets(tweetid, timestamp, tweet)
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`
	_, err := db.Exec(sqlStatement, tweetid, timestamp, tweet)
	if err != nil {
		panic(err)
	}

	// hashtags
	sqlStatement = `
INSERT INTO hashtags(tweetid, hashtag)	
VALUES ($1, $2) ON CONFLICT DO NOTHING`

	for _, hashtag := range hashtags {
		_, err = db.Exec(sqlStatement, tweetid, hashtag)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func ClearDB(db *sql.DB) {
	sqlStatement := `
DELETE FROM hashtags`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

	sqlStatement = `
DELETE FROM tweets`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}
