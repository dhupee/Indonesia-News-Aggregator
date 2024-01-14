default: start

start: 
	go run main.go

install-deps:
	go mod download

install-playwright-deps:
	sudo apt-get install libnss3 libnspr4 libatk-bridge2.0-0 libxkbcommon0 libatspi2.0-0 libasound2

