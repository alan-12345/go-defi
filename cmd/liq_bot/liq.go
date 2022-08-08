package main

import (
	"context"
	"flag"
	"fmt"
	"go_defi/addresses/ethereum"
	"go_defi/addresses/fantom"
	"go_defi/addresses/polygon"
	"log"
	"math/big"

	"github.com/ALTree/bigfloat"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type network_data struct {
	rpc string
}

var (
	network = flag.String("network", "ethereum", "Network")
	configs = map[string]network_data{
		"ethereum": {
			rpc: ethereum_addresses.RPC_URL,
		},
		"polygon": {
			rpc: polygon_addresses.RPC_URL,
		},
		"fantom": {
			rpc: fantom_addresses.RPC_URL,
		},
	}
	config network_data
)

type transaction struct {
	Hash      common.Hash
	From      common.Address
	To        common.Address
	Data      string
	Signature string
	GasPrice  *big.Int
}

func process_pending_tx(raw_tx *types.Transaction) {
	from, err := types.Sender(types.NewEIP155Signer(raw_tx.ChainId()), raw_tx)
	if err != nil {
		log.Println(err)
		return
	}

	tx_data := hexutil.Encode(raw_tx.Data())
	signature := tx_data
	if tx_data != "0x" {
		signature = tx_data[0:10]
	}
	gas_price := new(big.Float).SetInt(raw_tx.GasPrice())
	precision := bigfloat.Pow(big.NewFloat(10), big.NewFloat(9))
	gwei_price := new(big.Float).Quo(gas_price, precision)

	tx := transaction{
		Hash:      raw_tx.Hash(),
		From:      from,
		To:        *raw_tx.To(),
		Data:      tx_data,
		Signature: signature,
		GasPrice:  raw_tx.GasPrice(),
	}

	fmt.Println("Hash:", tx.Hash)
	fmt.Println("From:", tx.From)
	fmt.Println("To:", tx.To)
	fmt.Println("Data:", tx_data)
	fmt.Println("Signature:", signature)
	fmt.Println("Gas Price:", tx.GasPrice, "(", gwei_price, "Gwei )")
	fmt.Println("------------------------------")
}

func start_bot() {
	config = configs[*network]

	rpc_client, _ := rpc.Dial(config.rpc)
	client, err := ethclient.DialContext(context.Background(), config.rpc)
	ch := make(chan common.Hash, 2000)
	sub, err := rpc_client.EthSubscribe(context.Background(), ch, "newPendingTransactions")
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case txHash := <-ch:
			go func(txHash common.Hash) {
				defer func() {
					if r := recover(); r != nil {
						fmt.Println("recovered from panic", r)
					}
				}()
				tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
				_ = err
				// if err != nil {
				// 	fmt.Println("tx %s TransactionByHash error: %s\n", txHash.String(), err.Error())
				// 	return
				// }
				// if !isPending {
				// 	receipt, err := client.TransactionReceipt(context.Background(), txHash)
				// 	if err != nil {
				// 		fmt.Println("tx %s TransactionReceipt error: %s\n", txHash.String(), err.Error())
				// 		return
				// 	}
				// 	if receipt.Status == types.ReceiptStatusFailed {
				// 		fmt.Println("tx failed")
				// 	} else if receipt.Status == types.ReceiptStatusSuccessful {
				// 		fmt.Println("tx success")
				// 	} else {
				// 		fmt.Println("unknown tx status")
				// 	}
				// 	fmt.Println("blockNumber: ", receipt.BlockNumber)
				// }
				// if err != nil {
				// 	fmt.Println("get from address: ", err)
				// 	return
				// }
				if isPending {
					process_pending_tx(tx)
				}
			}(txHash)
		case err := <-sub.Err():
			fmt.Println(err)
		}
	}
}
