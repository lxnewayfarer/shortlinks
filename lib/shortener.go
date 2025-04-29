package lib

import (
	"context"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/lxnewayfarer/shortlinks/storage"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func shortlink(path string) string {
	appUrl := os.Getenv("APP_URL")

	return appUrl + "/l/" + path
}

func linkTTL() time.Duration {
	ttl, _ := strconv.Atoi(os.Getenv("LINK_MINUTES_TTL"))
	
	return time.Duration(ttl) * time.Minute
}

func ShortenLink(ctx context.Context, rdb storage.RedisClient, link string) (string, error) {
	cached, err := rdb.Get(ctx, link).Result()
	if err != nil && err != storage.Nil() {
		return "", err
	}

	if err == nil {
		return shortlink(cached), nil
	}

	path, err := uniqSeq(ctx, rdb)
	if err != nil {
		return "", err
	}

	pipe := rdb.TxPipeline()
	pipe.Set(ctx, path, link, linkTTL())
	pipe.Set(ctx, link, path, linkTTL())
	_, err = pipe.Exec(ctx)
	if err != nil {
		return "", err
	}

	return shortlink((path)), nil
}

func uniqSeq(ctx context.Context, rdb storage.RedisClient) (string, error) {
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

	return path, nil
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
