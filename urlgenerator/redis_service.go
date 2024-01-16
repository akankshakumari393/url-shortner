package urlgenerator

import (
	"context"
)

type RedisClientProvider interface {
	ShortenURL(ctx context.Context, url string) (string, error)
	RedirectToURL(ctx context.Context, shortURL string) (string, error)
}
