package main

import (
	"context"
	"flag"
	"fmt"
	"go_defi/contracts/compound/comptroller"
	"go_defi/networks/fantom"
	"go_defi/utils"
	"go_defi/utils/constants"
	"go_defi/utils/crypto"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/syndtr/goleveldb/leveldb"
)

type GlobalData struct {
	DB                    *leveldb.DB
	Multicaller           common.Address
	CompoundLikeProtocols map[string]constants.Compound
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

	case "polygon":

	case "fantom":
		NETWORK = constants.NetworkData{
			RPC: fantom_addresses.RPC_URL,
		}
		GLOBAL = GlobalData{
			CompoundLikeProtocols: fantom_addresses.COMPOUND_LIKE_PROTOCOLS,
		}
	}

	db, err := leveldb.RecoverFile(*DB_PATH, nil)
	if err != nil {
		log.Fatal(err)
	}

	GLOBAL.DB = db
	log.Println("Found and loaded DB")

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
	DistributedBorrowerComp chan *comptroller.ComptrollerDistributedBorrowerComp
	DistributedSupplierComp chan *comptroller.ComptrollerDistributedSupplierComp
}

type CompEventSubscriptions struct {
	MarketEntered           event.Subscription
	MarketExited            event.Subscription
	DistributedSupplierComp event.Subscription
	DistributedBorrowerComp event.Subscription
}

type CompoundBot struct {
	Prefix        []byte
	Client        *ethclient.Client
	Protocol      constants.Compound
	Comptroller   *comptroller.Comptroller
	Incoming      CompIncomingChans
	Subscriptions CompEventSubscriptions
}

type AaveBot struct {
	Prefix []byte
	Client *ethclient.Client
}

type Bot struct {
	Database     *leveldb.DB
	CompoundBots []CompoundBot
	AaveBots     []AaveBot
}

func listen_comp_events(bots []CompoundBot) {
	for _, bot := range bots {
		for {
			select {
			case err := <-bot.Subscriptions.MarketEntered.Err():
				log.Fatal(err)
			case err := <-bot.Subscriptions.MarketExited.Err():
				log.Fatal(err)
			case err := <-bot.Subscriptions.DistributedBorrowerComp.Err():
				log.Fatal(err)
			case err := <-bot.Subscriptions.DistributedSupplierComp.Err():
				log.Fatal(err)
			case payload := <-bot.Incoming.MarketEntered:
				fmt.Println(payload.Account)
			case payload := <-bot.Incoming.MarketExited:
				fmt.Println(payload.Account)
			case payload := <-bot.Incoming.DistributedBorrowerComp:
				fmt.Println(payload.Borrower)
			case payload := <-bot.Incoming.DistributedSupplierComp:
				fmt.Println(payload.Supplier)
			}
		}
	}
}

func create_comp_subs(bots *[]CompoundBot) {
	for i, compound_bot := range *bots {
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
			nil, compound_bot.Incoming.DistributedBorrowerComp, nil, nil,
		)
		if err != nil {
			log.Fatal(err)
		}

		market_comp_supplier, err := troller.WatchDistributedSupplierComp(
			nil, compound_bot.Incoming.DistributedSupplierComp, nil, nil,
		)
		if err != nil {
			log.Fatal(err)
		}

		compound_bot.Comptroller = troller

		compound_bot.Subscriptions = CompEventSubscriptions{
			MarketEntered:           market_entered_sub,
			MarketExited:            market_exited_sub,
			DistributedBorrowerComp: market_comp_borrower,
			DistributedSupplierComp: market_comp_supplier,
		}
		(*bots)[i] = compound_bot
	}
}

func create_aave_subs(bots *[]AaveBot) {}

func create_bots() Bot {
	var bots Bot
	var compound_bots []CompoundBot
	var aave_bots []AaveBot
	for name, protocol := range GLOBAL.CompoundLikeProtocols {
		client, err := ethclient.Dial(NETWORK.RPC)
		if err != nil {
			log.Fatal(err)
		}

		compound_bots = append(compound_bots, CompoundBot{
			Prefix:   []byte(name + "user" + "-"),
			Client:   client,
			Protocol: protocol,
			Incoming: CompIncomingChans{
				MarketEntered:           make(chan *comptroller.ComptrollerMarketEntered),
				MarketExited:            make(chan *comptroller.ComptrollerMarketExited),
				DistributedBorrowerComp: make(chan *comptroller.ComptrollerDistributedBorrowerComp),
				DistributedSupplierComp: make(chan *comptroller.ComptrollerDistributedSupplierComp),
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
	listen_comp_events(bots.CompoundBots)
}

func start_bot() {
	utils.PrintDashed()
	log.Println("Running liq_bot")
	fmt.Println("Account:", crypto.GetPublicAddress())
	utils.PrintDashed()

	setup_global_data()

	find_liqs()
}
