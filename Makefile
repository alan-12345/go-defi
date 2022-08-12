NETWORK := ethereum

run_arb:
	go run cmd/arb_bot/*.go -network ${NETWORK}

run_liq:
	go run cmd/liq_bot/*.go -network ${NETWORK}

tidy:
	go mod tidy

clean:
	go clean