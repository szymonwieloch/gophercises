package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := parseArgs()
	keys := readKeys(cfg.Keys)
	client, err := auth(keys)
	if err != nil {
		panic(err)
	}
	users, err := readUsers(cfg.Users)
	if err != nil {
		panic(err)
	}
	users = runLoop(users, client, cfg)
	winners := pickWinners(users, cfg.Winners)
	if len(winners) > 0 {
		fmt.Println("Winners:")
		for _, winner := range winners {
			fmt.Println(winner)
		}
	}

}

func runLoop(users []string, client *http.Client, cfg args) []string {
	sigs := make(chan os.Signal, 1)
	ticker := time.NewTicker(cfg.Period)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	for {
		retweeters, err := retweets(client, cfg.Tweet)
		if err != nil {
			log.Println("Error when getting retweets: ", err)
		}
		users = mergeUsers(users, retweeters)
		err = saveUsers(cfg.Users, users)
		if err != nil {
			log.Println("Failed to persist users: ", err)
		}
		select {
		case <-sigs:
			return users
		case <-ticker.C:
		}
	}
}

func pickWinners(retweeters []string, count uint) []string {
	if count == 0 {
		return []string{}
	}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	indexes := rnd.Perm(min(int(count), len(retweeters)))
	var result []string
	for _, idx := range indexes {
		result = append(result, retweeters[idx])
	}
	return result
}
