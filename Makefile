#Makefile for setting up and building steamplaydb.

.PHONY: all setup build production clean

GO = go

TARGET_SRC = src/main.go

BIN_OUTPUT = steamplaydb
OUTPUT_PATH = bin


DEVELOPMENT_VARIABLES =
PRODUCTION_VARIABLES = GOOS=linux GOARCH=amd64

all: setup build

build:
	 $(DEVELOPMENT_VARIABLES) $(GO) build -o $(OUTPUT_PATH)/$(BIN_OUTPUT) $(TARGET_SRC)

production:
	$(PRODUCTION_VARIABLES) $(GO) build -o $(OUTPUT_PATH)/$(BIN_OUTPUT) $(TARGET_SRC)

setup:
	go get -u github.com/labstack/echo/...
	go get -u github.com/joho/godotenv

clean:
	rm -rf $(OUTPUT_PATH)
