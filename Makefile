all: build
	@echo "build ok, look bin/"
build: clean
	go build -o bin/j2y cmd/main.go
build-all: clean
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/j2y-linux-amd64 cmd/main.go
        CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/j2y-linux-arm64 cmd/main.go
        CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/j2y-darwin-arm64 cmd/main.go
        CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/j2y-darwin-amd64 cmd/main.go
        CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/j2y-windows-amd64.exe cmd/main.go
clean: 
	@rm -rf bin/*
install: build
	@echo "install /usr/bin/"
	cp bin/j2y /usr/bin/
