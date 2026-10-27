package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/p7rz/blog_aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>\n", cmd.Name)
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{ID: uuid.New(), Name: feedName, Url: feedUrl, CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID})
	if err != nil {
		return fmt.Errorf("Couldn't create feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return fmt.Errorf("Couldn't create feed: %w", err)
	}

	fmt.Println("Feed created!")
	printFeed(feed, user)
	fmt.Println()
	fmt.Printf("Feed followed successfully:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("=====================================")
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeed(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't get feeds: %w\n", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		userName, err := s.db.GetName(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("Couldn't get user: %w", err)
		}
		printFeed(feed, userName)
		fmt.Println("=====================================")
	}

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
	fmt.Printf("* LastFetchedAt: %v\n", feed.LastFetchedAt.Time)
}
