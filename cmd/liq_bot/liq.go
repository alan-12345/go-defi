package main

import (
	"context"
	"flag"
	"fmt"
	"go_defi/contracts/compound/comptroller"
	"go_defi/contracts/multicall"
	"go_defi/networks/fantom"
	"go_defi/utils"
	"go_defi/utils/constants"
	"go_defi/utils/crypto"
	"go_defi/utils/decimal"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"golang.org/x/sync/errgroup"
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
			Multicaller:           fantom_addresses.MULTICALL_ADDR,
			CompoundLikeProtocols: fantom_addresses.COMPOUND_LIKE_PROTOCOLS,
		}
	}

	db, err := leveldb.OpenFile(*DB_PATH, nil)
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
	Prefix        string
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

type BotData struct {
	CompoundBots []CompoundBot
	AaveBots     []AaveBot
}

func store_event(prefix string, account common.Address) error {
	fmt.Println("PUT", prefix, account.String())

	wo := &opt.WriteOptions{
		Sync: false,
	}

	var key []byte
	key = append(key, []byte(prefix)...)
	key = append(key, account[:]...)

	return GLOBAL.DB.Put(key, account[:], wo)
}

func read_db() {
	iter := GLOBAL.DB.NewIterator(nil, nil)
	count := 0
	for iter.Next() {
		fmt.Println(string(iter.Key()), ":", hexutil.Encode(iter.Value()))
		count++
	}
	iter.Release()
	fmt.Println("# entries:", count)
}

func fetch_comp_events(bot CompoundBot) error {
	client, err := ethclient.Dial("https://rpcapi.fantom.network")
	if err != nil {
		return err
	}

	current_block, err := client.BlockNumber(context.Background())
	if err != nil {
		return err
	}

	troller, err := comptroller.NewComptroller(bot.Protocol.Unitroller, client)
	if err != nil {
		return err
	}

	interval := 500000
	for start_block := bot.Protocol.StartBlock.Uint64(); start_block < current_block; start_block += uint64(interval) {
		end_block := start_block + uint64(interval)
		if end_block >= current_block {
			end_block = current_block
		}
		fmt.Println(bot.Prefix+"loop", start_block, end_block)
		opts := &bind.FilterOpts{
			Start: start_block,
			End:   &end_block,
		}

		iter_market_entered, err := troller.FilterMarketEntered(opts)
		if err != nil {
			return err
		}

		for iter_market_entered.Next() {
			store_event(bot.Prefix, iter_market_entered.Event.Account)
		}

		iter_market_exited, err := troller.FilterMarketExited(opts)
		if err != nil {
			return err
		}

		for iter_market_exited.Next() {
			store_event(bot.Prefix, iter_market_exited.Event.Account)
		}
	}
	return nil
}

func listen_comp_events(bot CompoundBot) error {
	fmt.Println("listening for events", bot.Prefix)
	for {
		select {
		case err := <-bot.Subscriptions.MarketEntered.Err():
			return err
		case err := <-bot.Subscriptions.MarketExited.Err():
			return err
		case err := <-bot.Subscriptions.DistributedBorrowerComp.Err():
			return err
		case err := <-bot.Subscriptions.DistributedSupplierComp.Err():
			return err
		case payload := <-bot.Incoming.MarketEntered:
			store_event(bot.Prefix, payload.Account)
		case payload := <-bot.Incoming.MarketExited:
			store_event(bot.Prefix, payload.Account)
		case payload := <-bot.Incoming.DistributedBorrowerComp:
			store_event(bot.Prefix, payload.Borrower)
		case payload := <-bot.Incoming.DistributedSupplierComp:
			store_event(bot.Prefix, payload.Supplier)
		}
	}
}

func create_comp_subs(bots *[]CompoundBot) error {
	for i, bot := range *bots {

		market_entered_sub, err := bot.Comptroller.WatchMarketEntered(nil, bot.Incoming.MarketEntered)
		if err != nil {
			return err
		}

		market_exited_sub, err := bot.Comptroller.WatchMarketExited(nil, bot.Incoming.MarketExited)
		if err != nil {
			return err
		}

		market_comp_borrower, err := bot.Comptroller.WatchDistributedBorrowerComp(
			nil, bot.Incoming.DistributedBorrowerComp, nil, nil,
		)
		if err != nil {
			return err
		}

		market_comp_supplier, err := bot.Comptroller.WatchDistributedSupplierComp(
			nil, bot.Incoming.DistributedSupplierComp, nil, nil,
		)
		if err != nil {
			return err
		}

		bot.Subscriptions = CompEventSubscriptions{
			MarketEntered:           market_entered_sub,
			MarketExited:            market_exited_sub,
			DistributedBorrowerComp: market_comp_borrower,
			DistributedSupplierComp: market_comp_supplier,
		}
		(*bots)[i] = bot
	}

	return nil
}

func create_aave_subs(bots *[]AaveBot) error {
	return nil
}

func create_bots(bots *BotData) error {
	var compound_bots []CompoundBot
	var aave_bots []AaveBot
	for name, protocol := range GLOBAL.CompoundLikeProtocols {
		client, err := ethclient.Dial(NETWORK.RPC)
		if err != nil {
			return err
		}
		troller, err := comptroller.NewComptroller(protocol.Unitroller, client)
		if err != nil {
			return err
		}
		compound_bots = append(compound_bots, CompoundBot{
			Prefix:      name + "-user-",
			Client:      client,
			Protocol:    protocol,
			Comptroller: troller,
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

	return nil
}

func listen_blocks() {
	for {
		select {
		case err := <-NETWORK.Subscription.Err():
			log.Fatal(err)
		case header := <-NETWORK.Headers:
			block, err := NETWORK.Client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			start := time.Now()

			log.Println("New block #", block.Number().Uint64())
			utils.PrintDashed()

			end := time.Now()
			log.Println("Time elapsed (since block):", end.Sub(start).String())
			utils.PrintDashed()
		}
	}
}

func run_bot() error {
	var bots BotData
	if err := create_bots(&bots); err != nil {
		return err
	}

	if err := create_comp_subs(&bots.CompoundBots); err != nil {
		return err
	}

	if err := create_aave_subs(&bots.AaveBots); err != nil {
		return err
	}

	var g errgroup.Group

	g.Go(func() error {
		interrupt := make(chan os.Signal, 1)
		defer signal.Stop(interrupt)
		defer close(interrupt)
		signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-interrupt
		time.Sleep(time.Second)
		fmt.Println("\nBot Stopped")
		return nil
	})

	for _, bot := range bots.CompoundBots {
		closed_bot := bot
		// g.Go(func() error {
		// 	return fetch_comp_events(closed_bot)
		// })

		g.Go(func() error {
			return listen_comp_events(closed_bot)
		})

		g.Go(func() error {
			var addresses []common.Address
			iter := GLOBAL.DB.NewIterator(util.BytesPrefix([]byte(closed_bot.Prefix)), nil)
			for iter.Next() {
				account := common.HexToAddress(hexutil.Encode(iter.Value()))
				addresses = append(addresses, account)
			}
			iter.Release()

			var shortfall_accounts []common.Address
			max_calls := 400
			for start_idx := 0; start_idx < len(addresses); start_idx += max_calls {
				var calls []multicall.Multicall2Call

				end_idx := start_idx + max_calls
				if end_idx > len(addresses) {
					end_idx = len(addresses)
				}

				for _, address := range addresses[start_idx:end_idx] {
					encoded_args := crypto.EncodeArgs(comptroller.ComptrollerMetaData.ABI, "getAccountLiquidity", address)
					calls = append(calls, multicall.Multicall2Call{
						Target:   closed_bot.Protocol.Unitroller,
						CallData: encoded_args,
					})
				}
				encoded_calls := crypto.EncodeArgs(multicall.MulticallMetaData.ABI, "aggregate", calls)
				encoded_output := crypto.StaticCall(closed_bot.Client, GLOBAL.Multicaller, encoded_calls)
				decoded_output := (crypto.DecodeData(multicall.MulticallMetaData.ABI, "aggregate", encoded_output)[1]).([][]byte)
				for j, call := range decoded_output {
					address := addresses[start_idx+j]

					oops := decimal.NewDecFromBigIntWithPrec(crypto.DecodeData(comptroller.ComptrollerMetaData.ABI, "getAccountLiquidity", call)[0].(*big.Int), 18)
					shortfall := decimal.NewDecFromBigIntWithPrec(crypto.DecodeData(comptroller.ComptrollerMetaData.ABI, "getAccountLiquidity", call)[2].(*big.Int), 18)

					if !oops.IsZero() {
						log.Println("error fetching account liquidity for", address)
						continue
					}

					if !shortfall.IsZero() {
						shortfall_accounts = append(shortfall_accounts, address)
					}
				}
			}

			fmt.Println(closed_bot.Prefix, "# of shortfall accounts", len(shortfall_accounts))
			return nil
		})
	}

	// listen_blocks()
	return g.Wait()
}

func start_bot() {
	utils.PrintDashed()
	log.Println("Running liq_bot")
	fmt.Println("Account:", crypto.GetPublicAddress())
	utils.PrintDashed()

	setup_global_data()

	if err := run_bot(); err != nil {
		log.Println("something wrong with running the bot", err)
	}
}
