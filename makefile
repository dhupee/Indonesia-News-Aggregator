default: start

start: 
	go run main.go
	
start-container:
	docker run --env-file .env -p 8080:8080 indonesian-news-aggregator

build-container:
	docker build . -t "indonesian-news-aggregator" --file Dockerfile

install-deps:
	go mod download

install-playwright-deps:
	sudo apt-get install libnss3 libnspr4 libatk-bridge2.0-0 libxkbcommon0 libatspi2.0-0 libasound2

