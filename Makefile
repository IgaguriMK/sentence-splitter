.PHONY: build
build:
	go build splitter.go

.PHONY: clean
clean:
	- rm splitter
	- rm *.exe
