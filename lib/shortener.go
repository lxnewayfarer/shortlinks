package lib

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/lxnewayfarer/shortlinks/storage"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func shortlink(path string) string {
	appUrl := os.Getenv("APP_URL")

	return appUrl + "/l/" + path
}

func ShortenLink(ctx context.Context, rdb storage.RedisClient, link string) (string, error) {
	cached, err := rdb.Get(ctx, link).Result()
	if err != nil && err != storage.Nil() {
		return "", err
	}

	if err == nil {
		return shortlink(cached), nil
	}

	path := randSeq(8)
	for {
		_, err := rdb.Get(ctx, path).Result()
		if err != nil {
			if err == storage.Nil() {
				break
			}
			return "", err
		}
	}

	pipe := rdb.TxPipeline()
	pipe.Set(ctx, path, link, 24*time.Hour)
	pipe.Set(ctx, link, path, 24*time.Hour)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return "", err
	}

	return shortlink((path)), nil
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
