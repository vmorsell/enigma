.PHONY: gen test

gen:
	cd enigma && make gen

test:
	go test ./...
