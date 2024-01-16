package urlgenerator

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

type RedisClientProvider interface {
	ShortenURL(ctx context.Context, url string) (string, error)
	RedirectToURL(ctx context.Context, shortURL string) (string, error)
}

// RedisClient represents a Redis client connection
type RedisClient struct {
	client *redis.Client
	mu     *sync.Mutex
}

// NewRedisClient creates a new instance of RedisClient
func NewRedisClient() *RedisClient {
	return &RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
			Password: "",
			DB:       0,
		}),
		mu: &sync.Mutex{},
	}
}

// ShortenURL generates a short URL and stores it in Redis
func (rc *RedisClient) ShortenURL(ctx context.Context, url string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(url))
	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)
	shortURL := hashString[:8]

	rc.mu.Lock()
	defer rc.mu.Unlock()

	// Check if the key already exists in Redis
	exists, err := rc.client.Exists(ctx, shortURL).Result()
	if err != nil {
		return "", err
	}

	// If the key exists, append a counter until a unique key is found
	counter := 1
	for exists > 0 {
		shortURL = fmt.Sprintf("%s-%d", hashString[:8], counter)

		// Check again if the new key exists
		exists, err = rc.client.Exists(ctx, shortURL).Result()
		if err != nil {
			return "", err
		}

		counter++
	}

	// Store the mapping in Redis
	if err := rc.SaveUrl(ctx, shortURL, url); err != nil {
		return "", err
	}

	return shortURL, nil
}

// RedirectToURL retrieves the original URL from the short URL in Redis
func (rc *RedisClient) RedirectToURL(ctx context.Context, shortURL string) (string, error) {
	url, err := rc.client.Get(context.Background(), shortURL).Result()
	if err != nil {
		return "", err
	}
	return url, nil
}

func (rc *RedisClient) SaveUrl(ctx context.Context, shortUrl string, originalUrl string) error {
	err := rc.client.Set(ctx, shortUrl, originalUrl, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// MockRedisClient represents a mock Redis client for testing
type MockRedisClient struct{}

func (m *MockRedisClient) ShortenURL(ctx context.Context, destination string) (string, error) {
	// Mock the ShortenURL method
	// You can implement your own logic for testing
	return "mockShortURL", nil
}

func (m *MockRedisClient) RedirectToURL(ctx context.Context, shortURL string) (string, error) {
	// Mock the RedirectToURL method
	// You can implement your own logic for testing
	return "https://example.com", nil
}

func NewMockRedisClient() *MockRedisClient {
	return &MockRedisClient{}
}
