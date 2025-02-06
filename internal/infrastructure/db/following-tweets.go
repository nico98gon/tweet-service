package db

import (
	"fmt"
	"sort"
	"sync"
	"tweet-service/internal/domain/tweets"
)

const maxWorkers = 10

func GetFollowingTweets(userID string, cursor string) ([]*tweets.ReturnTweets, string, bool) {
	following, ok := GetFollowingFromUserService(userID)
	if !ok {
		return nil, "", false
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var allTweets []*tweets.ReturnTweets
	tweetChan := make(chan []*tweets.ReturnTweets, len(following))
	semaphore := make(chan struct{}, maxWorkers)

	for _, followingID := range following {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(id string) {
			defer wg.Done()
			defer func() { <-semaphore }()

			tweets, _, err := GetTweets(id, cursor)
			if err != nil {
				fmt.Printf("Error obteniendo tweets de %s: %v\n", id, err)
				return
			}
			tweetChan <- tweets
		}(followingID)
	}

	go func() {
		wg.Wait()
		close(tweetChan)
	}()

	for tweetsBatch := range tweetChan {
		mu.Lock()
		allTweets = append(allTweets, tweetsBatch...)
		mu.Unlock()
	}

	sortTweetsByDate(allTweets)

	var nextCursor string
	if len(allTweets) > 0 {
		nextCursor = allTweets[len(allTweets)-1].ID.Hex()
	}

	fmt.Println("Total de tweets obtenidos:", len(allTweets))

	return allTweets, nextCursor, true
}

func sortTweetsByDate(tweets []*tweets.ReturnTweets) {
	sort.Slice(tweets, func(i, j int) bool {
		return tweets[i].Date.After(tweets[j].Date)
	})
}
