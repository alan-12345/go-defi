NETWORK := ethereum

build_arb:
	go build -o arb_bot cmd/arb_bot/*.go

run_arb:
	go run cmd/arb_bot/*.go -network ${NETWORK}

build_liq:
	go build -o liq_bot cmd/liq_bot/*.go

run_liq:
	go run cmd/liq_bot/*.go -network ${NETWORK}

tidy:
	go mod tidy

clean:
	go clean