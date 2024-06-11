package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

func auth(key keyCfg) (*http.Client, error) {
	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(key.ConsumerKey, key.ConsumerSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	var t oauth2.Token
	err = dec.Decode(&t)
	if err != nil {
		return nil, err
	}
	oauthCfg := oauth2.Config{}
	tclient := oauthCfg.Client(context.Background(), &t)
	return tclient, nil
}

type retweetData struct {
	User struct {
		ScreenName string `json:"screen_name"`
	} `json:"user"`
}

func retweets(client *http.Client, id string) ([]string, error) {
	url := fmt.Sprintf("https://api.twitter.com/1.1/statuses/retweets/%s.json", id)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//dec := json.NewDecoder()
	var data []retweetData
	//err = dec.Decode(&data)
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Println("Invalid retweets response: ", string(content))
		return nil, err
	}

	var usernames []string
	for _, dt := range data {
		usernames = append(usernames, dt.User.ScreenName)
	}
	return usernames, nil
}
