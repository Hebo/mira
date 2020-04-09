package models

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func BenchmarkCreateSubreddit(b *testing.B) {
	data, _ := ioutil.ReadFile("./tests/subreddit.json")
	subredditExampleJSON := string(data)
	for i := 0; i < b.N; i++ {
		sub := Subreddit{}
		json.Unmarshal([]byte(subredditExampleJSON), &sub)
	}
}
