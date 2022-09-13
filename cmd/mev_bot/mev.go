package main

import (
	"context"
	"flag"
	"fmt"
	ethereum_addresses "go_defi/networks/ethereum"
	fantom_addresses "go_defi/networks/fantom"
	"go_defi/utils"
	"go_defi/utils/constants"
	"go_defi/utils/crypto"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/sync/errgroup"
)

type ClientData struct {
	Client       *ethclient.Client
	Channel      chan common.Hash
	Subscription *rpc.ClientSubscription
}

type GlobalData struct {
	Clients []ClientData
}

var (
	SELECTED_NETWORK = flag.String("network", "ethereum", "Network")
	DB_PATH          = flag.String("db", "networks/fantom/db", "Path to DB")
	NETWORK          constants.NetworkData
	GLOBAL           GlobalData
)

func setup_global_data() {
	fmt.Println("Selected network:", *SELECTED_NETWORK)
	switch *SELECTED_NETWORK {
	case "ethereum":
		NETWORK = constants.NetworkData{
			RPC: ethereum_addresses.RPC_URL,
		}
	case "polygon":

	case "fantom":
		NETWORK = constants.NetworkData{
			RPC: fantom_addresses.RPC_URL,
		}
	}

	for i := 0; i < 10; i++ {
		rpc_client, err := rpc.Dial(NETWORK.RPC)
		if err != nil {
			log.Fatal(err)
		}

		client, err := ethclient.DialContext(context.Background(), NETWORK.RPC)
		if err != nil {
			log.Fatal(err)
		}

		ch := make(chan common.Hash, 2000)
		sub, err := rpc_client.EthSubscribe(context.Background(), ch, "newPendingTransactions")
		if err != nil {
			log.Fatal(err)
		}
		GLOBAL.Clients = append(GLOBAL.Clients, ClientData{
			Client:       client,
			Channel:      ch,
			Subscription: sub,
		})
	}

	log.Println("Connected to RPC node")
	utils.PrintDashed()
}

type Transaction struct {
	Hash      string
	From      common.Address
	To        common.Address
	Value     string
	Data      string
	Signature string
	Gas       uint64
	GasPrice  uint64
}

func process_pending_tx(raw_tx *types.Transaction) {
	from, err := types.Sender(types.NewEIP155Signer(raw_tx.ChainId()), raw_tx)
	if err != nil {
		return
	}

	tx_data := hexutil.Encode(raw_tx.Data())
	signature := tx_data
	if tx_data != "0x" {
		signature = tx_data[0:10]
	}

	tx := Transaction{
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

	if tx.From.String() == "0x8881AcF21D569ec8182441eD381ff1945BFeC5E5" {
		fmt.Println("Hash:", tx.Hash)
		fmt.Println("From:", tx.From)
		fmt.Println("To:", tx.To)
		fmt.Println("Value:", tx.Value)
		fmt.Println("Data:", tx.Data)
		fmt.Println("Signature:", tx.Signature)
		fmt.Println("Gas:", tx.Gas)
		fmt.Println("Gas Price:", tx.GasPrice)
		fmt.Println("------------------------------")
	}
}

func run_bot() error {
	var g errgroup.Group
	for {
		for _, raw_client := range GLOBAL.Clients {
			client := raw_client
			g.Go(func() error {
				for {
					select {
					case txHash := <-client.Channel:
						go func(txHash common.Hash) {
							defer func() {
								if r := recover(); r != nil {
									fmt.Println("recovered from panic", r)
								}
							}()

							tx, isPending, err := client.Client.TransactionByHash(context.Background(), txHash)
							_ = err

							if isPending {
								process_pending_tx(tx)
							}
						}(txHash)
					case err := <-client.Subscription.Err():
						log.Fatal(err)
					}

				}
			})
		}
	}
}

func start_bot() {
	utils.PrintDashed()
	log.Println("Running mev_bot")
	fmt.Println("Account:", crypto.GetPublicAddress())
	utils.PrintDashed()

	setup_global_data()

	if err := run_bot(); err != nil {
		log.Println("something wrong with running the bot", err)
	}
}
