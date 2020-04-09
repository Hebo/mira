package models

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func BenchmarkCreateSubmission(b *testing.B) {
	data, _ := ioutil.ReadFile("./tests/submission.json")
	submissionExampleJSON := string(data)
	for i := 0; i < b.N; i++ {
		sub := Submission{}
		json.Unmarshal([]byte(submissionExampleJSON), &sub)
	}
}
