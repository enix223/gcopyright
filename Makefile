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
	GOS=windows GARCH=amd64 go build -o bin/gcopyright-windows-amd64.exe main.go

linux-amd64: init
	GOS=linux GARCH=amd64 go build -o bin/gcopyright-linux-amd64 main.go

darwin-amd64: init
	GOS=darwin GARCH=amd64 go build -o bin/gcopyright-darwin-amd64 main.go

# x86

windows-x86: init
	GOS=windows GARCH=x86 go build -o bin/gcopyright-windows-x86.exe main.go

linux-x86: init
	GOS=linux GARCH=x86 go build -o bin/gcopyright-linux-x86 main.go

darwin-x86: init
	GOS=darwin GARCH=x86 go build -o bin/gcopyright-darwin-x86 main.go


all: windows-amd64 linux-amd64 darwin-amd64 windows-x86 linux-x86 darwin-x86
