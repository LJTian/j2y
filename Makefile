all: build
	@echo "build ok, look bin/"
build: clean
	go build -o bin/j2y cmd/main.go
build-all: clean
        GOOS=linux GOARCH=amd64 go build -o bin/j2y-linux-amd64 cmd/main.go
        GOOS=linux GOARCH=arm64 go build -o bin/j2y-linux-arm64 cmd/main.go
        GOOS=darwin GOARCH=arm64 go build -o bin/j2y-darwin-arm64 cmd/main.go
        GOOS=darwin GOARCH=amd64 go build -o bin/j2y-darwin-amd64 cmd/main.go
        GOOS=windows GOARCH=amd64 go build -o bin/j2y-windows-amd64.exe cmd/main.go
clean: 
	@rm -rf bin/*
install: build
	@echo "install /usr/bin/"
	cp bin/j2y /usr/bin/
