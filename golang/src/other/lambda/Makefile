all: main.js main.go
	GOOS=linux GOARCH=amd64 go build -o example ./main.go
	zip bin/example ./example ./main.js
