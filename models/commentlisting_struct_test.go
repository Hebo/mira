package models

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func BenchmarkCreateCommentListing(b *testing.B) {
	data, _ := ioutil.ReadFile("./tests/commentlisting.json")
	commentListingExampleJSON := string(data)
	for i := 0; i < b.N; i++ {
		sub := CommentListing{}
		json.Unmarshal([]byte(commentListingExampleJSON), &sub)
	}
}
