package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/lxnewayfarer/shortlinks/lib"
	"github.com/lxnewayfarer/shortlinks/storage"
)

func TestShortener(t *testing.T) {
	rdb, _ := storage.InitRedis()
	res, err := lib.ShortenLink(context.Background(), rdb, "https://example.com")
	fmt.Println(res)
	fmt.Println(err)
}