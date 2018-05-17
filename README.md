# Word Counter
A simple application that let users send text and get the top ten words with how many times they occur in the text

## How to Run
```
1. Install Go
2. Set $GOPATH
3. Run `go get github.com/gorilla/mux`
4. Run `go get github.com/johannesridho/word-counter`
5. Run `go run $GOPATH/src/github.com/johannesridho/word-counter/main.go`
6. Server will run on port 8000
```

## Endpoint
```
Request :

POST /
Body (Content-Type application/json):
{
	"text": "The quick brown fox jumps over the lazy dog"
}
```
```
Response :
[
    {
        "word": "the",
        "count": 2
    },
    {
        "word": "lazy",
        "count": 1
    },
    {
        "word": "dog",
        "count": 1
    },
    {
        "word": "quick",
        "count": 1
    },
    {
        "word": "brown",
        "count": 1
    },
    {
        "word": "fox",
        "count": 1
    },
    {
        "word": "jumps",
        "count": 1
    },
    {
        "word": "over",
        "count": 1
    }
]
```