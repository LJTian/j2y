all: build
	@echo "build ok, look bin/"
build: clean
	go build -o bin/j2y cmd/main.go	
clean: 
	@rm bin/*
install: build
	@echo "install /usr/bin/"
	cp bin/j2y /usr/bin/
