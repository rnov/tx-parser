.PHONY: build run request-block request-transaction request-subscribe

build:
	@mkdir -p tx-parser/out
	go build -o ./out/build tx-parser/cmd/app/parser

run:
	./out/build

clean:
	rm -rf ./out

all: build run

get-block:
	curl --location "http://127.0.0.1:8080/block"

get-txs:
	curl --location "http://127.0.0.1:8080/transactions/$(address)"

subscribe:
	curl --location "http://127.0.0.1:8080/subscribe" \
	--header 'Content-Type: application/json' \
	--data '{"address": "$(address)"}'
