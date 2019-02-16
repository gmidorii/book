BIN=bs
APP_NAME=book-server
APP_PORT=8088
AUTO_PORT=3333
WIRE_BIN=wire
TAG=latest

wire:
	$(WIRE_BIN) github.com/gmidorii/book/server/injector

build: wire
	go build -o $(BIN)

run: build
	./$(BIN)

run-reload: build
	gin -a $(APP_PORT) -p $(AUTO_PORT)

d-build:
	docker build -t $(APP_NAME):$(TAG) .

d-run:
	docker run --name $(APP_NAME) -p 8088:8088 -d $(APP_NAME):$(TAG)

d-rm:
	docker rm -f $(APP_NAME)