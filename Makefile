BINARY := sortfile

clean:
	rm -f $(BINARY)

build: clean
	CGO_ENABLED=0 go build -ldflags="-s -w" .

run: build
	./$(BINARY) -h
