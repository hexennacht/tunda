package tunda

import (
	"context"
	"net"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/redis/rueidis"
)

type RedisRepository interface {
	Get(ctx context.Context, key string) ([]byte, error)
	MGet(ctx context.Context, keys ...string) ([]byte, error)
	LRange(ctx context.Context, key string, start int64, end int64) ([]byte, error)
	Set(ctx context.Context, key string, value []byte) error
	RPush(ctx context.Context, key string, element ...string) error
	LPush(ctx context.Context, key string, element ...string) error
	Delete(ctx context.Context, keys ...string) error
}

type repository struct {
	client rueidis.Client
}

func NewRedisRepository() RedisRepository {
	var config configuration

	envconfig.MustProcess("", &config)

	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress:      strings.Split(config.RedisAddress, ","),
		SelectDB:         config.RedisDB,
		Username:         config.RedisUsername,
		Password:         config.RedisPassword,
		ConnWriteTimeout: time.Duration(config.RedisWriteTimeout) * time.Second,
		Dialer: net.Dialer{
			Timeout:   time.Duration(config.RedisTimeout) * time.Second,
			KeepAlive: time.Duration(config.RedisKeepAlive) * time.Second,
		},
	})
	if err != nil {
		panic(err)
	}

	return &repository{client: client}
}

// Delete implements RedisRepository.
func (r *repository) Delete(ctx context.Context, keys ...string) error {
	cmd := r.client.B().Del().Key(keys...).Build()

	return r.client.Do(ctx, cmd).Error()
}

// Get implements RedisRepository.
func (r *repository) Get(ctx context.Context, key string) ([]byte, error) {
	get := r.client.B().Get().Key(string(key)).Build()

	return r.client.Do(ctx, get).AsBytes()
}

// MGet implements RedisRepository.
func (r *repository) MGet(ctx context.Context, keys ...string) ([]byte, error) {
	get := r.client.B().Mget().Key(keys...).Build()

	return r.client.Do(ctx, get).AsBytes()
}

// Set implements RedisRepository.
func (r *repository) Set(ctx context.Context, key string, value []byte) error {
	cmd := r.client.B().Set().Key(key).Value(string(value)).Build()

	return r.client.Do(ctx, cmd).Error()
}

// LRange implements RedisRepository.
func (r *repository) LRange(ctx context.Context, key string, start int64, end int64) ([]byte, error) {
	cmd := r.client.B().Lrange().Key(key).Start(start).Stop(end).Build()

	return r.client.Do(ctx, cmd).AsBytes()
}

// RPush implements RedisRepository.
func (r *repository) RPush(ctx context.Context, key string, element ...string) error {
	cmd := r.client.B().Rpush().Key(key).Element(element...).Build()

	return r.client.Do(ctx, cmd).Error()
}

// LPush implements RedisRepository.
func (r *repository) LPush(ctx context.Context, key string, element ...string) error {
	cmd := r.client.B().Lpush().Key(key).Element(element...).Build()

	return r.client.Do(ctx, cmd).Error()
}
