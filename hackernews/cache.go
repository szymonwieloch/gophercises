package main

import (
	"log"
	"sync"
	"time"

	"github.com/szymonwieloch/gophercises/hackernews/client"
)

// Common interface for all cache strategies
type cacheStrategy interface {
	getTopStories() ([]client.Story, error)
}

// No cache, pull the API each time
type cacheNone struct {
	count int
}

func (cn cacheNone) getTopStories() ([]client.Story, error) {
	return getTopStories(cn.count)
}

// assert that it implements the interface
var _ cacheStrategy = cacheNone{}

// Refresh the cache on demand in a periodic way
// Once for a while the cache request takes a long time
type cacheRefresh struct {
	count       int
	period      time.Duration
	lastRefresh time.Time
	mutex       sync.Mutex
	cached      []client.Story
}

func (cr *cacheRefresh) getTopStories() ([]client.Story, error) {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()
	if time.Since(cr.lastRefresh) > cr.period {
		cached, err := getTopStories(cr.count)
		if err != nil {
			return nil, err
		}
		cr.cached = cached
		cr.lastRefresh = time.Now()
	}
	return cr.cached, nil
}

var _ cacheStrategy = (*cacheRefresh)(nil)

// Cache gets refreshed by a background routine
// It is always fast but causes overhead when no input request are comming
type cacheBackground struct {
	mutex  sync.Mutex
	cached []client.Story
}

func newBackgroundCache(count int, period time.Duration) *cacheBackground {
	cb := &cacheBackground{}
	go func() {
		for {
			stories, err := getTopStories(count)
			if err != nil {
				log.Println("Error when refreshing stories in the background: ", err)
			} else {
				cb.mutex.Lock()
				cb.cached = stories
				cb.mutex.Unlock()
			}
			time.Sleep(period)
		}
	}()
	return cb
}

func (cb *cacheBackground) getTopStories() ([]client.Story, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	// in this version we always return something without error, even if the backgorund job fails
	// alternative approach would be to cach error and return it
	return cb.cached, nil
}

var _ cacheStrategy = (*cacheBackground)(nil)
