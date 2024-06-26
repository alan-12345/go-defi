NETWORK := polygon
PATH_TO_DB := networks/${NETWORK}/db

build_arb:
	go build -o arb_bot cmd/arb_bot/*.go

run_arb:
	go run cmd/arb_bot/*.go -network ${NETWORK}

build_liq:
	go build -o liq_bot cmd/liq_bot/*.go

run_liq:
	go run cmd/liq_bot/*.go -network ${NETWORK} -db ${PATH_TO_DB}

run_mev:
	go run cmd/mev_bot/*.go -network ${NETWORK}

run_cex:
	go run cmd/cex_bot/*.go

tidy:
	go mod tidy

clean:
	go clean