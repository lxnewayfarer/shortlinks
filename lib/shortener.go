package lib

import (
	"context"
	"os"

	"github.com/lxnewayfarer/shortlinks/storage"
)

func shortlink(path string) string {
	appUrl := os.Getenv("APP_URL")

	return appUrl + "/l/" + path
}

func ShortenLink(ctx context.Context, rdb storage.RedisClient, link string, r Random) (string, error) {
	cached, err := rdb.Get(ctx, link).Result()
	if err != nil && err != storage.Nil() {
		return "", err
	}

	if err == nil {
		return shortlink(cached), nil
	}

	path, err := uniqSeq(ctx, rdb, r)
	if err != nil {
		return "", err
	}

	pipe := rdb.TxPipeline()
	pipe.Set(ctx, path, link, 0)
	pipe.Set(ctx, link, path, 0)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return "", err
	}

	return shortlink((path)), nil
}

func uniqSeq(ctx context.Context, rdb storage.RedisClient, r Random) (string, error) {
	path := r.RandSeq(8)
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
