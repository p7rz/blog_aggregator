package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/p7rz/blog_aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>`\n", cmd.Name)
	}
	url := cmd.Args[0]
	feed, err := s.db.GetUrl(context.Background(), url)
	if err != nil {
		return err
	}
	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Feed has been followed.\nFeed name: %v\nUser Name: %v\n", follow.FeedName, follow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, follow := range following {
		fmt.Println(follow.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>`\n", cmd.Name)
	}
	url := cmd.Args[0]
	feed, err := s.db.GetUrl(context.Background(), url)
	if err != nil {
		return err
	}
	unfollow, err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{FeedID: feed.ID, UserID: user.ID})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Feed has been unfollowed.\nFeedID: %v\nUserID: %v\n", unfollow.FeedID, unfollow.UserID)
	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
