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
	router.HandleFunc("/{text}", GetWordCounts).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetWordCounts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["text"])

	words := strings.Fields(params["text"])
	wordCountsMap := make(map[string]int)

	for _, word := range words {
		wordCountsMap[strings.ToLower(word)]++
	}

	type WordCount struct {
		Word  string
		Count int
	}

	var wordCounts []WordCount
	for w, c := range wordCountsMap {
		wordCounts = append(wordCounts, WordCount{w, c})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})

	if len(wordCounts) < 10  {
		json.NewEncoder(w).Encode(wordCounts)
	} else {
		topTenWordCounts := append(wordCounts[:0], wordCounts[:10]...)
		json.NewEncoder(w).Encode(topTenWordCounts)
	}
}