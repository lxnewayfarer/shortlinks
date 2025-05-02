package main

import (
	"testing"

	"github.com/lxnewayfarer/shortlinks/lib"
	"github.com/lxnewayfarer/shortlinks/storage"
)

func TestShortenerWithCache(t *testing.T) {
	link := "https://example.com/"
	shortlink := "qwertysz"
	mockRandom := lib.MockRandomInstance{Value: shortlink}

	rdb, mock := storage.InitMockRedis()
	mock.ExpectGet(link).SetVal(shortlink)

	res, err := lib.ShortenLink(t.Context(), rdb, link, mockRandom)
	if err != nil || res != "/l/" + shortlink {
		t.Fatalf("ShortenLink() failed")
	}
}

func TestShortenerWithNoCache(t *testing.T) {
	link := "https://example.com/"
	shortlink := "123"
	mockRandom := lib.MockRandomInstance{Value: shortlink}
	rdb, mock := storage.InitMockRedis()

	mock.ExpectGet(link).RedisNil()
	mock.ExpectGet(shortlink).RedisNil()

	mock.ExpectTxPipeline()
	mock.ExpectSet(shortlink, link, 0).SetVal("ok")
	mock.ExpectSet(link, shortlink, 0).SetVal("ok")
	mock.ExpectTxPipelineExec()

	res, err := lib.ShortenLink(t.Context(), rdb, link, mockRandom)
	if err != nil || res != "/l/" + shortlink {
		t.Fatalf("ShortenLink() failed")
	}
}