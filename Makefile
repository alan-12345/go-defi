NETWORK := ethereum
PATH_TO_DB := networks/${NETWORK}/db

build_arb:
	go build -o arb_bot cmd/arb_bot/*.go

run_arb:
	go run cmd/arb_bot/*.go -network ${NETWORK}

build_liq:
	go build -o liq_bot cmd/liq_bot/*.go

run_liq:
	go run cmd/liq_bot/*.go -network ${NETWORK} -db ${PATH_TO_DB}

tidy:
	go mod tidy

clean:
	go clean