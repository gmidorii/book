BIN=bs
APP_PORT=8088
AUTO_PORT=3333
WIRE_BIN=wire

wire:
	$(WIRE_BIN) github.com/gmidorii/book/server/injector

build: wire
	go build -o $(BIN)

run: build
	./$(BIN)

run-reload: build
	gin -a $(APP_PORT) -p $(AUTO_PORT)

