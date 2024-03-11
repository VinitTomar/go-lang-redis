package post_voting

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"vin.go_lang.redis/lib"
)

type VotingTestSuite struct {
	suite.Suite
	postId       int
	initialVotes int
}

func (suite *VotingTestSuite) SetupTest() {
	suite.postId = 123
	suite.initialVotes = 5

	client := lib.GetRedisClient()

	ctx := context.TODO()

	client.Set(ctx, getPostKey(suite.postId), suite.initialVotes, 0)
}

func (suite *VotingTestSuite) TestUpVote() {
	assert.Equal(suite.T(), suite.initialVotes+1, UpVote(suite.postId), "It should increase vote count by one")
}

func (suite *VotingTestSuite) TestDownVote() {
	assert.Equal(suite.T(), suite.initialVotes-1, DownVote(suite.postId), "It should increase vote count by one")
}

func TestVotingTestSuite(t *testing.T) {
	suite.Run(t, new(VotingTestSuite))
}
