package redisservice

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"
	"time"

	"backend.app/configs"
	"github.com/go-redis/redis/v8"
)

// Redis -.
type Redis struct {
	Client *redis.Client
}

func NewConnection(config *configs.Config) Redis {
	parsed, err := url.Parse(config.RedisUri)
	if err != nil {
		return Redis{Client: redis.NewClient(&redis.Options{Addr: config.RedisUri})}
	}
	password, _ := parsed.User.Password()
	return Redis{Client: redis.NewClient(&redis.Options{Addr: parsed.Host, Password: password})}
}

func IsOpen(ctx context.Context, r Redis) bool {
	return strings.EqualFold(r.Client.Ping(ctx).Val(), "PONG")
}
func (r Redis) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}
func (r Redis) GetValue(ctx context.Context, key string) string { return r.Client.Get(ctx, key).Val() }
func (r Redis) GetIntValue(ctx context.Context, key string) int {
	value, _ := r.Client.Get(ctx, key).Int()
	return value
}
func (r Redis) KeyExists(ctx context.Context, key string) int64 {
	return r.Client.Exists(ctx, key).Val()
}
func (r Redis) Delete(ctx context.Context, key string) error { return r.Client.Del(ctx, key).Err() }
func (r Redis) DeleteByPattern(ctx context.Context, pattern string) error {
	keys, err := r.Client.Keys(ctx, pattern).Result()
	if err != nil || len(keys) == 0 {
		return err
	}
	return r.Client.Del(ctx, keys...).Err()
}
func (r Redis) AddToSet(ctx context.Context, key, member string) error {
	return r.Client.SAdd(ctx, key, member).Err()
}
func (r Redis) RemoveFromSet(ctx context.Context, key, member string) error {
	return r.Client.SRem(ctx, key, member).Err()
}
func (r Redis) JsonSet(ctx context.Context, key string, value map[string]interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Client.Do(ctx, "JSON.SET", key, "$", data).Err()
}

// IsMemberOfSet checks if a member exists in a set
func (r Redis) IsMemberOfSet(ctx context.Context, key, member string) bool {
	result, _ := r.Client.SIsMember(ctx, key, member).Result()
	return result
}
