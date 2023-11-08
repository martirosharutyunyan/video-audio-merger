default: run

run:
	go run ./cmd/video-audio-merger

build:
	go build ./cmd/video-audio-merger

install:
	go install ./cmd/video-audio-merger

tidy:
	go mod tidy

update:
	go get -u ./...