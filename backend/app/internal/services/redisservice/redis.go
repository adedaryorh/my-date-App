package redisservice

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Redis -.
type Redis struct {
	Client *redis.Client
}

// IsMemberOfSet checks if a member exists in a set
func (r Redis) IsMemberOfSet(ctx context.Context, key, member string) bool {
	result, _ := r.Client.SIsMember(ctx, key, member).Result()
	return result
}