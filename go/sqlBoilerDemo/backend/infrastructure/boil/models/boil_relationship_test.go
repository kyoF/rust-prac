// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("LikeToUserUsingUser", testLikeToOneUserUsingUser)
	t.Run("LikeToTweetUsingTweet", testLikeToOneTweetUsingTweet)
	t.Run("TweetToUserUsingUser", testTweetToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("TweetToLikes", testTweetToManyLikes)
	t.Run("UserToLikes", testUserToManyLikes)
	t.Run("UserToTweets", testUserToManyTweets)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("LikeToUserUsingLikes", testLikeToOneSetOpUserUsingUser)
	t.Run("LikeToTweetUsingLikes", testLikeToOneSetOpTweetUsingTweet)
	t.Run("TweetToUserUsingTweets", testTweetToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("TweetToLikes", testTweetToManyAddOpLikes)
	t.Run("UserToLikes", testUserToManyAddOpLikes)
	t.Run("UserToTweets", testUserToManyAddOpTweets)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}
