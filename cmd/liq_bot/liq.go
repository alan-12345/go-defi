package main

import (
	"context"
	"flag"
	"fmt"
	"go_defi/addresses/ethereum"
	"go_defi/addresses/fantom"
	"go_defi/addresses/polygon"
	"go_defi/utils/crypto"
	"log"

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
	network = flag.String("network", "fantom", "Network")
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
	Hash      string
	From      common.Address
	To        common.Address
	Value     string
	Data      string
	Signature string
	Gas       uint64
	GasPrice  uint64
}

func process_pending_tx(client *ethclient.Client, raw_tx *types.Transaction) {
	from, err := types.Sender(types.NewEIP155Signer(raw_tx.ChainId()), raw_tx)
	if err != nil {
		return
	}

	tx_data := hexutil.Encode(raw_tx.Data())
	signature := tx_data
	if tx_data != "0x" {
		signature = tx_data[0:10]
	}

	tx := transaction{
		Hash:      raw_tx.Hash().Hex(),
		From:      from,
		To:        *raw_tx.To(),
		Value:     raw_tx.Value().String(),
		Data:      tx_data,
		Signature: signature,
		Gas:       raw_tx.Gas(),
		GasPrice:  raw_tx.GasPrice().Uint64(),
	}

	fmt.Println("Hash:", tx.Hash, tx.Signature)

	if tx.Signature == "0xf8a8fd6d" {
		fmt.Println("Hash:", tx.Hash)
		fmt.Println("From:", tx.From)
		fmt.Println("To:", tx.To)
		fmt.Println("value:", tx.Value)
		fmt.Println("Data:", tx.Data)
		fmt.Println("Signature:", tx.Signature)
		fmt.Println("Gas:", tx.Gas)
		fmt.Println("Gas Price:", tx.GasPrice)
		fmt.Println("------------------------------")
		front_run(client, tx)
	}
}

func front_run(client *ethclient.Client, tx transaction) {
	// _, err := dummycontract.NewDummy(fantom_addresses.DUMMY_ADDR, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// function_data := crypto.GetFunctionData("test()")

	// opts := crypto.GetOpts(client)
	// gasLimit := crypto.GetGasLimit(client, fantom_addresses.DUMMY_ADDR, function_data)
	// fmt.Println(gasLimit)
	// opts.GasLimit = gasLimit
	// tx, err := dummy_contract.Test(opts)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870


}

func start_bot() {
	fmt.Println("Running liq_bot (", *network, ")")
	fmt.Println("Account:", crypto.GetPublicAddress())

	config = configs[*network]

	rpc_client, err := rpc.Dial(config.rpc)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.DialContext(context.Background(), config.rpc)
	if err != nil {
		log.Fatal(err)
	}

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
					process_pending_tx(client, tx)
				}
			}(txHash)
		case err := <-sub.Err():
			log.Fatal(err)
		}
	}
}
