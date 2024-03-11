package post_voting

import (
	"context"
	"strconv"

	"vin.go_lang.redis/lib"
)

func getPostKey(postId int) string {
	return "post:" + strconv.Itoa(postId) + ":votes"
}

func UpVote(postId int) int {

	postKey := getPostKey(postId)

	client := lib.GetRedisClient()

	ctx := context.TODO()

	latestVotes := client.Incr(ctx, postKey).Val()

	return int(latestVotes)
}

func DownVote(postId int) int {
	postKey := getPostKey(postId)

	client := lib.GetRedisClient()

	ctx := context.TODO()

	latestVotes := client.Decr(ctx, postKey).Val()

	return int(latestVotes)
}
