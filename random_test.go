package main

import (
	"regexp"
	"testing"

	"github.com/lxnewayfarer/shortlinks/lib"
)

func TestRandom(t *testing.T) {
	random := lib.RandomInstance{}

	if match, _ := regexp.MatchString(`[a-zA-Z0-9]`, random.RandSeq(8)); !match {
		t.Fatalf("Random() regexp match failed")
	}
}

func TestMock(t *testing.T) {
	mockValue := "123"
	mockRandom := lib.MockRandomInstance{Value: mockValue}

	if mockRandom.RandSeq(8) != mockValue {
		t.Fatalf("Random() mock failed")
	}
}