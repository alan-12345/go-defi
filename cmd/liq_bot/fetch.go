package main

import (
	"context"
	"fmt"
	"go_defi/contracts/compound/comptroller"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func fetch_comp_markets(unitroller common.Address) []common.Address {
	
	client, err := ethclient.Dial("https://rpc.ftm.tools")
	if err != nil {
		log.Fatal(err)
	}
	unitroller_contract, err := comptroller.NewComptroller(unitroller, client)
	if err != nil {
		log.Fatal(err)
	}

	markets, err := unitroller_contract.GetAllMarkets(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(markets)
	return markets
}

func query_comp_events(unitroller common.Address, start_block *big.Int) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(44876162),
		ToBlock: nil,
		Addresses: []common.Address{common.HexToAddress("0x8d9aed9882b4953a0c9fa920168fa1fdfa0ebe75")},
	}

	client, err := ethclient.Dial("https://rpc.ftm.tools")
	if err != nil {
		log.Fatal(err)
	}

	logs, err := client.FilterLogs(context.Background(), query)
	fmt.Println(logs)
    if err != nil {
        log.Fatal(err)
    }

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
	}
}