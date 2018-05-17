package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strings"
	"sort"
)

type WordCount struct {
	word  string
	count int
}

type WordCountsRequest struct {
	text string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetWordCounts).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetWordCounts(writer http.ResponseWriter, request *http.Request) {
	var wordCountsRequest WordCountsRequest
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&wordCountsRequest); err != nil {
		createErrorResponse(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()

	text := wordCountsRequest.text

	words := strings.Fields(text)
	wordCountsMap := make(map[string]int)

	for _, word := range words {
		wordCountsMap[strings.ToLower(word)]++
	}

	var wordCounts []WordCount
	for w, c := range wordCountsMap {
		wordCounts = append(wordCounts, WordCount{w, c})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].count > wordCounts[j].count
	})

	if len(wordCounts) < 10  {
		createJsonResponse(writer, http.StatusOK, wordCounts)
	} else {
		topTenWordCounts := append(wordCounts[:0], wordCounts[:10]...)
		createJsonResponse(writer, http.StatusOK, topTenWordCounts)
	}
}

func createErrorResponse(writer http.ResponseWriter, code int, message string) {
	createJsonResponse(writer, code, map[string]string{"error": message})
}

func createJsonResponse(writer http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(response)
}
