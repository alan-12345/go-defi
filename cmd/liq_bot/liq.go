package main

import (
	"context"
	"flag"
	"fmt"
	"go_defi/addresses/fantom"
	"go_defi/contracts/compound/comptroller"
	"go_defi/utils"
	"go_defi/utils/constants"
	"go_defi/utils/crypto"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

type GlobalData struct {
	Multicaller           common.Address
	CompoundLikeProtocols map[string]constants.Compound
}

var (
	SELECTED_NETWORK = flag.String("network", "ethereum", "Network")
	NETWORK          constants.NetworkData
	GLOBAL           GlobalData
)

func setup_global_data() {
	fmt.Println("Selected network:", *SELECTED_NETWORK)
	switch *SELECTED_NETWORK {
	case "ethereum":

	case "polygon":

	case "fantom":
		NETWORK = constants.NetworkData{
			RPC: fantom_addresses.RPC_URL,
		}
		GLOBAL = GlobalData{
			CompoundLikeProtocols: fantom_addresses.COMPOUND_LIKE_PROTOCOLS,
		}
	}

	client, err := ethclient.Dial(NETWORK.RPC)
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	NETWORK.Client = client
	NETWORK.Headers = headers
	NETWORK.Subscription = sub

	log.Println("Connected to RPC node")
	utils.PrintDashed()
}

type CompIncomingChans struct {
	MarketEntered           chan *comptroller.ComptrollerMarketEntered
	MarketExited            chan *comptroller.ComptrollerMarketExited
	CompDistributedBorrower chan *comptroller.ComptrollerDistributedBorrowerComp
	CompDistributedSupplier chan *comptroller.ComptrollerDistributedSupplierComp
}

type CompEventSubscriptions struct {
	MarketEntered           event.Subscription
	MarketExited            event.Subscription
	CompDistributedSupplier event.Subscription
	CompDistributedBorrower event.Subscription
}

type CompoundBot struct {
	Client        *ethclient.Client
	Protocol      constants.Compound
	Comptroller   *comptroller.Comptroller
	Incoming      CompIncomingChans
	Subscriptions CompEventSubscriptions
}

type AaveBot struct {
	Client *ethclient.Client
}

type Bot struct {
	CompoundBots []CompoundBot
	AaveBots     []AaveBot
}

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

func create_comp_subs(compound_bots *[]CompoundBot) {
	for i, compound_bot := range *compound_bots {
		troller, err := comptroller.NewComptroller(compound_bot.Protocol.Unitroller, compound_bot.Client)
		if err != nil {
			log.Fatal(err)
		}

		market_entered_sub, err := troller.WatchMarketEntered(nil, compound_bot.Incoming.MarketEntered)
		if err != nil {
			log.Fatal(err)
		}

		market_exited_sub, err := troller.WatchMarketExited(nil, compound_bot.Incoming.MarketExited)
		if err != nil {
			log.Fatal(err)
		}

		market_comp_borrower, err := troller.WatchDistributedBorrowerComp(
			nil, compound_bot.Incoming.CompDistributedBorrower, nil, nil,
		)
		if err != nil {
			log.Fatal(err)
		}

		market_comp_supplier, err := troller.WatchDistributedSupplierComp(
			nil, compound_bot.Incoming.CompDistributedSupplier, nil, nil,
		)
		if err != nil {
			log.Fatal(err)
		}

		compound_bot.Comptroller = troller

		compound_bot.Subscriptions = CompEventSubscriptions{
			MarketEntered:           market_entered_sub,
			MarketExited:            market_exited_sub,
			CompDistributedBorrower: market_comp_borrower,
			CompDistributedSupplier: market_comp_supplier,
		}
		(*compound_bots)[i] = compound_bot
	}
}

func create_aave_subs(aave_bots *[]AaveBot) {

}

func create_bots() Bot {
	var bots Bot
	var compound_bots []CompoundBot
	var aave_bots []AaveBot
	for _, protocol := range GLOBAL.CompoundLikeProtocols {
		client, err := ethclient.Dial(NETWORK.RPC)
		if err != nil {
			log.Fatal(err)
		}

		compound_bots = append(compound_bots, CompoundBot{
			Client:   client,
			Protocol: protocol,
			Incoming: CompIncomingChans{
				MarketEntered:           make(chan *comptroller.ComptrollerMarketEntered),
				MarketExited:            make(chan *comptroller.ComptrollerMarketExited),
				CompDistributedBorrower: make(chan *comptroller.ComptrollerDistributedBorrowerComp),
				CompDistributedSupplier: make(chan *comptroller.ComptrollerDistributedSupplierComp),
			},
		})
	}
	bots.CompoundBots = compound_bots

	bots.AaveBots = aave_bots

	return bots
}

func find_liqs() {
	bots := create_bots()
	create_comp_subs(&bots.CompoundBots)
	create_aave_subs(&bots.AaveBots)
}

func start_bot() {
	utils.PrintDashed()
	log.Println("Running liq_bot")
	fmt.Println("Account:", crypto.GetPublicAddress())
	utils.PrintDashed()

	setup_global_data()

	find_liqs()

	// rpc_client, err := rpc.Dial(config.rpc)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client, err := ethclient.DialContext(context.Background(), config.rpc)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ch := make(chan common.Hash, 2000)
	// sub, err := rpc_client.EthSubscribe(context.Background(), ch, "newPendingTransactions")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for {
	// 	select {
	// 	case txHash := <-ch:
	// 		go func(txHash common.Hash) {
	// 			defer func() {
	// 				if r := recover(); r != nil {
	// 					fmt.Println("recovered from panic", r)
	// 				}
	// 			}()
	// 			tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	// 			_ = err
	// 			// if err != nil {
	// 			// 	fmt.Println("tx %s TransactionByHash error: %s\n", txHash.String(), err.Error())
	// 			// 	return
	// 			// }
	// 			// if !isPending {
	// 			// 	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	// 			// 	if err != nil {
	// 			// 		fmt.Println("tx %s TransactionReceipt error: %s\n", txHash.String(), err.Error())
	// 			// 		return
	// 			// 	}
	// 			// 	if receipt.Status == types.ReceiptStatusFailed {
	// 			// 		fmt.Println("tx failed")
	// 			// 	} else if receipt.Status == types.ReceiptStatusSuccessful {
	// 			// 		fmt.Println("tx success")
	// 			// 	} else {
	// 			// 		fmt.Println("unknown tx status")
	// 			// 	}
	// 			// 	fmt.Println("blockNumber: ", receipt.BlockNumber)
	// 			// }
	// 			// if err != nil {
	// 			// 	fmt.Println("get from address: ", err)
	// 			// 	return
	// 			// }
	// 			if isPending {
	// 				process_pending_tx(client, tx)
	// 			}
	// 		}(txHash)
	// 	case err := <-sub.Err():
	// 		log.Fatal(err)
	// 	}
	// }
}
