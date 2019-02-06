BIN=bs
APP_PORT=8088
AUTO_PORT=3333

build:
	go build -o $(BIN)

run: build
	./$(BIN)

run-reload: build
	gin -a $(APP_PORT) -p $(AUTO_PORT)

