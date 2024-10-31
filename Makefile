EXE_NAME     = tsetse

all: build

build: build_mac build_win

build_mac:
	go build -o $(EXE_NAME) -ldflags "-s -w" cmd/cli/main.go

build_win:
	GOOS=windows GOARCH=amd64 go build -o $(EXE_NAME).exe -ldflags "-s -w" cmd/cli/main.go

clean:
	rm $(EXE_NAME) $(EXE_NAME).exe