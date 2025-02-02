package main

import (
	"time"

	"github.com/aaradhyakul/rssagg/internal/databases"
	"github.com/google/uuid"
)

type User struct{
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	APIKey string `json:"api_key"`
}

type Feed struct{
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed databases.Feed) Feed{
	return Feed{
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []databases.Feed) []Feed{
	feeds := []Feed{}
	for _,dbFeed := range dbFeeds {
		feeds = append(feeds,databaseFeedToFeed(dbFeed))
	}
	return feeds
}


func databaseUserToUser(dbUser databases.User) User{
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name :dbUser.Name,
		APIKey: dbUser.ApiKey,
	}
}
type FeedFollow struct{
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID 	uuid.UUID	`json:"feed_id"`	
}

func databaseFeedFollowToFeedFollow(dbFeedFollow databases.FeedFollow)FeedFollow{
	return FeedFollow{
		ID: dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID: dbFeedFollow.UserID,
		FeedID: dbFeedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []databases.FeedFollow)[]FeedFollow{
	feedFollows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows{
		feedFollows = append(feedFollows,databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time	`json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string `json:"title"`
	Description *string `json:"description"` //marshalling using *string defaults to NULL(as in json)
	PublishedAt time.Time`json:"published_at"`
	Url         string`json:"url"`
	FeedID      uuid.UUID`json:"feed_id"`
}

func databasePostToPost(dbPost databases.Post)Post{
	var description *string
	if dbPost.Description.Valid{
		description=&dbPost.Description.String
	}
	return Post{
		ID: dbPost.ID,
		CreatedAt: dbPost.CreatedAt,
		UpdatedAt: dbPost.UpdatedAt,
		Title: dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url: dbPost.Url,
		FeedID: dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []databases.Post)[]Post{
	posts := []Post{}
	for _,dbPost := range dbPosts{
		posts = append(posts, databasePostToPost(dbPost) )
	}
	return posts
}