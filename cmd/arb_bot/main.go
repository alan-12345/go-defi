package main

import (
	"context"
	"go_defi/addresses/polygon"
	"go_defi/contracts/uniswap/query"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	rpc := "wss://ws-matic-mainnet.chainstacklabs.com"
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}

	query_contract, err := query.NewUniswapQuery(polygon_addresses.UNISWAP_QUERY_ADDR, client)
	if err != nil {
		log.Fatal(err)
	}

	pairs := generate_all_pairs()
	pair_addrs := fetch_pair_addrs(query_contract, pairs)

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			log.Println("New block #", block.Number().Uint64())

			arbitrage(query_contract, pair_addrs)
		}
	}
}
