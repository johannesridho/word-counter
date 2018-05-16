package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strings"
	"sort"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{text}", GetWordCount).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetWordCount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["text"])

	words := strings.Fields(params["text"])
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[strings.ToLower(word)]++
	}

	type wordcount struct {
		Word  string
		Count int
	}

	var wordCountKeyval []wordcount
	for w, c := range wordCount {
		wordCountKeyval = append(wordCountKeyval, wordcount{w, c})
	}

	sort.Slice(wordCountKeyval, func(i, j int) bool {
		return wordCountKeyval[i].Count > wordCountKeyval[j].Count
	})

	if len(wordCountKeyval) < 10  {
		json.NewEncoder(w).Encode(wordCountKeyval)
	} else {
		topTenWordCount := append(wordCountKeyval[:0], wordCountKeyval[:10]...)
		json.NewEncoder(w).Encode(topTenWordCount)
	}
}