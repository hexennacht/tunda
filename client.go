package tunda

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"

	"github.com/redis/rueidis"
)

// Create function to run the job
type TundaClient interface {
	Create(ctx context.Context, kerjaan *Kerjaan) (KerjaanID, error)
}

type tundaClient struct {
	redis RedisRepository
}

func NewTundaClient() TundaClient {
	return &tundaClient{redis: NewRedisRepository()}
}

// Create creates a new job in the tunda system.
// It encodes and hashes the provided `kerjaan` data,
// then enqueues the job using the encoded data and returns the generated `KerjaanID`.
// If an error occurs during encoding, hashing, or enqueueing the job, it returns an empty `KerjaanID` and the error.
func (t *tundaClient) Create(ctx context.Context, kerjaan *Kerjaan) (KerjaanID, error) {
	data, key, err := encodeAndHashData(kerjaan)
	if err != nil {
		return "", err
	}

	kerjaanID, err := t.enqueueJob(ctx, key, data)
	if err != nil {
		return "", err
	}

	return kerjaanID, nil
}

// enqueueJob enqueues a job with the specified key and data into the tunda client.
// If a job with the same key already exists, it returns the existing job ID.
// Otherwise, it creates a new job with the given key and data, and returns the new job ID.
// The function returns an error if there is an issue with the Redis operations.
func (t *tundaClient) enqueueJob(ctx context.Context, key string, data []byte) (KerjaanID, error) {
	value, err := t.redis.Get(ctx, JOB_PREFIX+key)
	if err, ok := rueidis.IsRedisErr(err); ok {
		return "", err
	}

	if value != nil {
		return KerjaanID(JOB_PREFIX + key), nil
	}

	err = t.redis.Set(ctx, JOB_PREFIX+key, data)
	if err, ok := rueidis.IsRedisErr(err); ok {
		return "", err
	}

	err = t.redis.RPush(ctx, KEY_LIST_OF_JOBS, JOB_PREFIX+key)
	if err, ok := rueidis.IsRedisErr(err); ok {
		return "", err
	}

	return KerjaanID(JOB_PREFIX + key), nil
}

func encodeAndHashData(kerjaan *Kerjaan) ([]byte, string, error) {
	data, err := json.Marshal(kerjaan)
	if err != nil {
		return nil, "", err
	}

	hash := sha256.New()

	_, err = hash.Write([]byte(string(data)))
	if err != nil {
		return nil, "", err
	}

	key := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	
	return data, key, nil
}
