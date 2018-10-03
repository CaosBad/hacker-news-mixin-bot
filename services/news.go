package services

import (
	"context"
	"encoding/base64"
	"log"
	"sort"

	bot "github.com/MixinNetwork/bot-api-go-client"
	"github.com/crossle/hacker-news-mixin-bot/config"
	"github.com/crossle/hacker-news-mixin-bot/models"
	"github.com/jasonlvhit/gocron"
	h "github.com/qube81/hackernews-api-go"
)

type NewsService struct{}

type Stats struct {
	prevStoryId int // lastest news id
}

func (self Stats) getPrevTopStoryId() int {
	return self.prevStoryId
}

func (self *Stats) updatePrevTopStoryId(id int) {
	self.prevStoryId = id
}

func getTopStoryId() int {
	topStories, _ := h.GetStories("top") // get top item
	return topStories[0]
}

func getTopTenStories() []int {
	stories, _ := h.GetStories("top") // hacker news api get news
	topTen := stories[:10]
	sort.Ints(topTen)
	return topTen
}

func sendTopStoryToChannel(ctx context.Context, stats *Stats) {
	prevStoryId := stats.getPrevTopStoryId()
	topTenStories := getTopTenStories()
	for _, storyId := range topTenStories {
		if storyId > prevStoryId {
			story, _ := h.GetItem(storyId) // get story by id
			log.Printf("Sending top story to channel...")
			stats.updatePrevTopStoryId(story.ID)
			subscribers, _ := models.FindSubscribers(ctx) // query subscribers
			for _, subscriber := range subscribers 
				conversationId := bot.UniqueConversationId(config.MixinClientId, subscriber.UserId) // get mixin conversationId 
				data := base64.StdEncoding.EncodeToString([]byte(story.Title + " " + story.URL)) // assembling news
				// 
				bot.PostMessage(ctx, conversationId, subscriber.UserId, bot.UuidNewV4().String(), "PLAIN_TEXT", data, config.MixinClientId, config.MixinSessionId, config.MixinPrivateKey)

			}
		} else {
			log.Printf("Same top story ID: %d, no message sent.", storyId)
		}
	}
}
func (service *NewsService) Run(ctx context.Context) error {
	stats := &Stats{getTopStoryId()}
	gocron.Every(5).Minutes().Do(sendTopStoryToChannel, ctx, stats)
	<-gocron.Start() // start cron 
	return nil
}
