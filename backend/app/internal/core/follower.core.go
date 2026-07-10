package core

import (
	"context"

	"github.com/google/uuid"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/common/messages"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/pkg/response"
)

// FollowUser method used to follow a user
func (c *Core) FollowUser(ctx context.Context, userId uuid.UUID, user *models.User) *dtos.ResponseObject {
	// check if user is already followed
	followed, err := c.repo.GetFollowerByField(ctx, helpers.Map{"user_id": userId, "follower_id": user.ID})
	if err != nil && err != messages.ErrFollowerNotFound {
		return response.ServerErrorResponse(err)
	}

	if followed != nil {
		return response.BadRequestResponse(messages.ErrUserAlreadyFollowed)
	}

	// check if you're blocked by user
	blocked, err := c.repo.GetBlockedUserByField(ctx, helpers.Map{"user_id": userId, "blocked_user_id": user.ID})
	if err != nil && err != messages.ErrBlockedUserNotFound {
		return response.ServerErrorResponse(err)
	}
	if blocked != nil {
		return response.BadRequestResponse(messages.ErrUserBlockedYou)
	}

	newFollower := &models.Follower{
		UserId:     userId,
		FollowerId: user.ID,
	}
	// follow user
	follower, err := c.repo.CreateFollower(ctx, newFollower)
	if err != nil {
		return response.ServerErrorResponse(err)
	}
	// increase follower
	if err = c.incrementFollowing(ctx, userId, user.ID); err != nil {
		response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.UserFollowedSuccessfully, follower)
}

// UnFollowUser method used to un-follow a user
func (c *Core) UnFollowUser(ctx context.Context, userId uuid.UUID, user *models.User) *dtos.ResponseObject {
	// unfollow user
	if err := c.repo.DeleteFollower(ctx, helpers.Map{"user_id": userId, "follower_id": user.ID}); err != nil {
		return response.ServerErrorResponse(err)
	}
	// increase follower
	if err := c.decrementFollowing(ctx, userId, user.ID); err != nil {
		response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.UserUnFollowedSuccessfully, nil)
}

// incrementFollowing method that increments a user's following
func (c Core) incrementFollowing(ctx context.Context, userId, actorId uuid.UUID) error {
	if err := c.repo.IncrementUserFields(ctx, userId, []*models.Incrementor{
		{
			Field:    "followers",
			Operator: "+",
			Value:    1,
		},
	}); err != nil {
		return err
	}
	// increase following
	return c.repo.IncrementUserFields(ctx, actorId, []*models.Incrementor{
		{
			Field:    "following",
			Operator: "+",
			Value:    1,
		},
	})
}

// decrementFollowing method that decrements a user's following
func (c Core) decrementFollowing(ctx context.Context, userId, actorId uuid.UUID) error {
	if err := c.repo.IncrementUserFields(ctx, userId, []*models.Incrementor{
		{
			Field:    "followers",
			Operator: "-",
			Value:    1,
		},
	}); err != nil {
		return err
	}
	// increase following
	return c.repo.IncrementUserFields(ctx, actorId, []*models.Incrementor{
		{
			Field:    "following",
			Operator: "-",
			Value:    1,
		},
	})
}

// GetAllFollowers used to get all user's followers
func (c *Core) GetAllFollowers(ctx context.Context, user *models.User, query *dtos.APIPagingDto) *dtos.ResponseObject {
	result, err := c.repo.GetAllFollowers(ctx, user, query)
	if err != nil {
		response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.FollowersFetchedSuccessfully, result)
}

// GetSingleFollower gets a single follower
func (c *Core) GetSingleFollower(ctx context.Context, userId uuid.UUID, user *models.User) *dtos.ResponseObject {
	follower, err := c.repo.GetFollowerByField(ctx, helpers.Map{"user_id": user.ID, "follower_id": userId})
	if err != nil {
		return response.ServerErrorResponse(err)
	}
	follower.Following = c.isFollowing(ctx, user.ID, user.ID)
	return response.SuccessResponse(constants.FollowerSuccessFullyFetched, follower)
}

// BlockUser blocks a user
func (c *Core) BlockUser(ctx context.Context, userId uuid.UUID, user *models.User) *dtos.ResponseObject {
	alreadyBlocked, err := c.repo.GetBlockedUserByField(ctx, helpers.Map{"user_id": user.ID, "blocked_user_id": userId})
	if err != nil && err != messages.ErrBlockedUserNotFound {
		response.ServerErrorResponse(err)
	}
	if alreadyBlocked != nil {
		return response.SuccessResponse(constants.UserAlreadyBlocked, nil)
	}
	// block user proper
	newBlock := models.Blocked{
		UserId:        user.ID,
		BlockedUserId: userId,
	}

	block, err := c.repo.CreateBlock(ctx, &newBlock)
	if err != nil {
		response.ServerErrorResponse(err)
	}

	return response.SuccessResponse(constants.UserBlockedSuccessfully, block)
}

// GetBlockedUser gets a single blocked user
func (c *Core) GetBlockedUser(ctx context.Context, userId uuid.UUID, user *models.User) *dtos.ResponseObject {
	Blocked, err := c.repo.GetBlockedUserByField(ctx, helpers.Map{"user_id": user.ID, "blocked_user_id": userId})
	if err != nil && err != messages.ErrBlockedUserNotFound {
		response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.UserBlockedSuccessfully, Blocked)
}

// GetAllBlockedUsers gets all blocked users by a particular user
func (c *Core) GetAllBlockedUsers(ctx context.Context, user *models.User, query *dtos.APIPagingDto) *dtos.ResponseObject {
	result, err := c.repo.GetAllBlocked(ctx, user, query)
	if err != nil {
		response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.BlockedUsersFetchedSuccessfully, result)
}

// UnBlockUser unblock removes a blocked user from the blocked list
// thereby unblocking the user
func (c *Core) UnBlockUser(ctx context.Context, userId uuid.UUID, user *models.User) *dtos.ResponseObject {
	if err := c.repo.DeleteBlocked(ctx, helpers.Map{"user_id": user.ID, "blocked_user_id": userId}); err != nil && err != messages.ErrBlockedUserNotFound {
		response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.UserUnBlockedSuccessfully, nil)
}

// GetFriends gets a user's friends
func (c *Core) GetFriends(ctx context.Context, user *models.User, query *dtos.APIPagingDto) *dtos.ResponseObject {
	result, err := c.repo.GetAllFriends(ctx, user, query)
	if err != nil {
		response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.FriendsFetchedSuccessfully, result)
}

// isFollowing checks if user is following the follower
func (c *Core) isFollowing(ctx context.Context, userId, followerId uuid.UUID) bool {
	follower, _ := c.repo.GetFollowerByField(ctx, helpers.Map{"user_id": followerId, "follower_id": userId})
	return follower != nil
}
