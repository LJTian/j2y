all: build
	@echo "build ok, look bin/"
build:
	go build -o bin/j2y cmd/j2y.go
clean:
	@rm bin/*
