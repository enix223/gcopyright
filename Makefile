.PHONY: run-example init clean

init:
	mkdir -p bin

clean:
	rm -rf bin/

run-example:
	go run main.go \
		-o "/tmp/code.docx" \
		-p "./example/**/*.js" "./example/**/*.css"

# amd64

windows-amd64: init
	GOOS=windows GARCH=amd64 go build -o bin/gcopyright-windows-amd64.exe main.go

linux-amd64: init
	GOOS=linux GARCH=amd64 go build -o bin/gcopyright-linux-amd64 main.go

darwin-amd64: init
	GOOS=darwin GARCH=amd64 go build -o bin/gcopyright-darwin-amd64 main.go

darwin-arm64: init
	GOOS=darwin GARCH=arm64 go build -o bin/gcopyright-darwin-arm64 main.go

# 386

windows-386: init
	GOOS=windows GARCH=386 go build -o bin/gcopyright-windows-386.exe main.go

linux-386: init
	GOOS=linux GARCH=386 go build -o bin/gcopyright-linux-386 main.go

darwin-386: init
	GOOS=darwin GARCH=386 go build -o bin/gcopyright-darwin-386 main.go


all: windows-amd64 linux-amd64 darwin-amd64 windows-386 linux-386 darwin-386
