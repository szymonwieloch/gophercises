package main

import (
	"log"
	"slices"

	"github.com/szymonwieloch/gophercises/hackernews/client"
)

func takeN[T any](in []T, n int) ([]T, []T) {
	take := min(n, len(in))
	return in[:take], in[take:]
}

type asyncResult struct {
	story client.Story
	err   error
}

func getBatch(storyIDs []client.StoryID) []client.Story {
	ch := make(chan asyncResult)
	for _, storyId := range storyIDs {
		go func(storyId client.StoryID) {
			story, err := client.GetStory(storyId)
			ch <- asyncResult{
				story: story,
				err:   err,
			}
		}(storyId)
	}
	var stories []client.Story
	for range storyIDs {
		ar := <-ch
		if ar.err != nil {
			log.Println("Error getting story: ", ar.err)
			continue
		}
		stories = append(stories, ar.story)
	}
	return stories
}

func getTopStories(count int) ([]client.Story, error) {
	topStories, err := client.TopStories()
	if err != nil {
		return nil, err
	}
	remaining := topStories
	var stories []client.Story
	for len(stories) < count && len(remaining) > 0 {
		missing := count - len(stories)
		batchSize := (missing * 5 / 4) + 3
		var batchIDs []client.StoryID
		batchIDs, remaining = takeN(remaining, batchSize)
		batchStories := getBatch(batchIDs)
		for _, story := range batchStories {
			if isStoryLink(story) {
				stories = append(stories, story)
			}
		}
	}
	// we may have too many stories
	stories = stories[:min(count, len(stories))]
	//sort
	storyIDToIdx := map[client.StoryID]int{}
	for idx, storyId := range topStories {
		storyIDToIdx[storyId] = idx
	}
	slices.SortFunc(stories, func(a, b client.Story) int {
		return storyIDToIdx[a.ID] - storyIDToIdx[b.ID]
	})
	return stories, nil
}

func isStoryLink(story client.Story) bool {
	return story.Type == "story" && story.URL != ""
}
