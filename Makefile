EXE_NAME     = tsetse

all: build

build:
	go build -o $(EXE_NAME) cmd/cli/main.go