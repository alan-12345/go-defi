package main

import (
	"context"
	"flag"
	"fmt"
	"go_defi/contracts/aave-v2/lending-pool"
	"go_defi/contracts/compound/cerc20"
	"go_defi/contracts/compound/comptroller"
	"go_defi/networks/ethereum"
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
	EventFetchSize        int
	Multicaller           common.Address
	CompoundLikeProtocols map[string]constants.LendingProtocol
	AaveV2LikeProtocols   map[string]constants.LendingProtocol
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
		GLOBAL = GlobalData{
			EventFetchSize:        ethereum_addresses.EVENT_FETCH_SIZE,
			Multicaller:           ethereum_addresses.MULTICALL_ADDR,
			CompoundLikeProtocols: ethereum_addresses.COMPOUND_LIKE_PROTOCOLS,
			AaveV2LikeProtocols:   ethereum_addresses.AAVE_V2_LIKE_PROTOCOLS,
		}
	case "polygon":

	case "fantom":
		NETWORK = constants.NetworkData{
			RPC: fantom_addresses.RPC_URL,
		}
		GLOBAL = GlobalData{
			EventFetchSize:        fantom_addresses.EVENT_FETCH_SIZE,
			Multicaller:           fantom_addresses.MULTICALL_ADDR,
			CompoundLikeProtocols: fantom_addresses.COMPOUND_LIKE_PROTOCOLS,
		}
	}

	db, err := leveldb.OpenFile(*DB_PATH, nil)
	if err != nil {
		log.Fatal(err)
	}

	GLOBAL.DB = db

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

type BotData struct {
	CompoundBots []CompoundBot
	AaveV2Bots   []AaveV2Bot
}

type CompoundBot struct {
	Name          string
	Client        *ethclient.Client
	Protocol      constants.LendingProtocol
	Comptroller   *comptroller.Comptroller
	Incoming      CompIncomingChannels
	Subscriptions CompEventSubscriptions
	Shutdown      chan struct{}
}

type CompIncomingChannels struct {
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

type CToken struct {
	Address      common.Address
	Amount       *big.Int
	ExchangeRate *big.Int
}

type AaveV2Bot struct {
	Name          string
	Client        *ethclient.Client
	Protocol      constants.LendingProtocol
	LendingPool   *lendingpool.LendingPool
	Incoming      AaveV2IncomingChannels
	Subscriptions AaveV2EventSubscriptions
	Shutdown      chan struct{}
}

type AaveV2IncomingChannels struct {
	Deposit   chan *lendingpool.LendingPoolDeposit
	Withdraw  chan *lendingpool.LendingPoolWithdraw
	Borrow    chan *lendingpool.LendingPoolBorrow
	Repay     chan *lendingpool.LendingPoolRepay
	Swap      chan *lendingpool.LendingPoolSwap
	Liquidate chan *lendingpool.LendingPoolLiquidationCall
}

type AaveV2EventSubscriptions struct {
	Deposit   event.Subscription
	Withdraw  event.Subscription
	Borrow    event.Subscription
	Repay     event.Subscription
	Swap      event.Subscription
	Liquidate event.Subscription
}

func store_event(protocol string, account common.Address) error {
	var key []byte
	key = append(key, []byte(protocol+"-user-")...)
	key = append(key, account[:]...)

	return store_in_db(key, account[:])
}

func store_last_fetched_block(protocol string, block_number uint64) error {
	key := []byte(protocol + "-last-block")
	value := big.NewInt(int64(block_number)).Bytes()

	return store_in_db(key, value)
}

func store_in_db(key []byte, value []byte) error {
	wo := &opt.WriteOptions{
		Sync: false,
	}

	return GLOBAL.DB.Put(key, value, wo)
}

func fetch_comp_events(bot CompoundBot) error {
	start := time.Now()

	current_block, err := NETWORK.Client.BlockNumber(context.Background())
	if err != nil {
		return err
	}

	last_fetched_block_bytes, err := GLOBAL.DB.Get([]byte(bot.Name+"-last-block"), nil)
	var last_fetched_block uint64
	if err != nil {
		last_fetched_block = bot.Protocol.StartBlock.Uint64()
	} else {
		last_fetched_block = big.NewInt(0).SetBytes(last_fetched_block_bytes).Uint64()
	}

	for start_block := last_fetched_block; start_block < current_block; start_block += uint64(GLOBAL.EventFetchSize) {
		end_block := start_block + uint64(GLOBAL.EventFetchSize)
		if end_block >= current_block {
			end_block = current_block
		}
		opts := &bind.FilterOpts{
			Start: start_block,
			End:   &end_block,
		}

		var g errgroup.Group
		g.Go(func() error {
			iter_market_entered, err := bot.Comptroller.FilterMarketEntered(opts)
			if err != nil {
				return err
			}

			for iter_market_entered.Next() {
				store_event(bot.Name, iter_market_entered.Event.Account)
			}
			iter_market_entered.Close()

			return nil
		})

		g.Go(func() error {
			iter_market_exited, err := bot.Comptroller.FilterMarketExited(opts)
			if err != nil {
				return err
			}

			for iter_market_exited.Next() {
				store_event(bot.Name, iter_market_exited.Event.Account)
			}
			iter_market_exited.Close()

			return nil
		})

		if err := g.Wait(); err != nil {
			return err
		}

		store_last_fetched_block(bot.Name, end_block)
	}

	end := time.Now()
	log.Println(bot.Name, "Synced events in", end.Sub(start).String())
	utils.PrintDashed()

	return nil
}

func fetch_aave_v2_events(bot AaveV2Bot) error {
	start := time.Now()

	current_block, err := NETWORK.Client.BlockNumber(context.Background())
	if err != nil {
		return err
	}

	last_fetched_block_bytes, err := GLOBAL.DB.Get([]byte(bot.Name+"-last-block"), nil)
	var last_fetched_block uint64
	if err != nil {
		last_fetched_block = bot.Protocol.StartBlock.Uint64()
	} else {
		last_fetched_block = big.NewInt(0).SetBytes(last_fetched_block_bytes).Uint64()
	}

	for start_block := last_fetched_block; start_block < current_block; start_block += uint64(GLOBAL.EventFetchSize) {
		end_block := start_block + uint64(GLOBAL.EventFetchSize)
		if end_block >= current_block {
			end_block = current_block
		}

		opts := &bind.FilterOpts{
			Start: start_block,
			End:   &end_block,
		}

		var g errgroup.Group
		g.Go(func() error {
			iter_deposit, err := bot.LendingPool.FilterDeposit(opts, nil, nil, nil)
			if err != nil {
				return err
			}

			for iter_deposit.Next() {
				store_event(bot.Name, iter_deposit.Event.User)
			}
			iter_deposit.Close()

			return nil
		})

		g.Go(func() error {
			iter_withdraw, err := bot.LendingPool.FilterWithdraw(opts, nil, nil, nil)
			if err != nil {
				return err
			}
			for iter_withdraw.Next() {
				store_event(bot.Name, iter_withdraw.Event.User)
			}
			iter_withdraw.Close()

			return nil
		})

		g.Go(func() error {
			iter_borrow, err := bot.LendingPool.FilterBorrow(opts, nil, nil, nil)
			if err != nil {
				return err
			}

			for iter_borrow.Next() {
				store_event(bot.Name, iter_borrow.Event.User)
			}
			iter_borrow.Close()

			return nil
		})

		g.Go(func() error {
			iter_repay, err := bot.LendingPool.FilterRepay(opts, nil, nil, nil)
			if err != nil {
				return err
			}

			for iter_repay.Next() {
				store_event(bot.Name, iter_repay.Event.User)
			}
			iter_repay.Close()

			return nil
		})

		g.Go(func() error {
			iter_swap, err := bot.LendingPool.FilterSwap(opts, nil, nil)
			if err != nil {
				return err
			}

			for iter_swap.Next() {
				store_event(bot.Name, iter_swap.Event.User)
			}
			iter_swap.Close()

			return nil
		})

		g.Go(func() error {
			iter_liquidate, err := bot.LendingPool.FilterLiquidationCall(opts, nil, nil, nil)
			if err != nil {
				return err
			}

			for iter_liquidate.Next() {
				store_event(bot.Name, iter_liquidate.Event.User)
				store_event(bot.Name, iter_liquidate.Event.Liquidator)
			}
			iter_liquidate.Close()

			return nil
		})

		if err := g.Wait(); err != nil {
			return err
		}

		store_last_fetched_block(bot.Name, end_block)
	}

	end := time.Now()
	log.Println(bot.Name, "Synced events in", end.Sub(start).String())
	utils.PrintDashed()

	return nil
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

func create_aave_v2_subs(bots *[]AaveV2Bot) error {
	for i, bot := range *bots {
		deposit, err := bot.LendingPool.WatchDeposit(nil, bot.Incoming.Deposit, nil, nil, nil)
		if err != nil {
			return err
		}

		withdraw, err := bot.LendingPool.WatchWithdraw(nil, bot.Incoming.Withdraw, nil, nil, nil)
		if err != nil {
			return err
		}

		borrow, err := bot.LendingPool.WatchBorrow(nil, bot.Incoming.Borrow, nil, nil, nil)
		if err != nil {
			return err
		}

		repay, err := bot.LendingPool.WatchRepay(nil, bot.Incoming.Repay, nil, nil, nil)
		if err != nil {
			return err
		}

		swap, err := bot.LendingPool.WatchSwap(nil, bot.Incoming.Swap, nil, nil)
		if err != nil {
			return err
		}

		liquidate, err := bot.LendingPool.WatchLiquidationCall(nil, bot.Incoming.Liquidate, nil, nil, nil)
		if err != nil {
			return err
		}

		bot.Subscriptions = AaveV2EventSubscriptions{
			Deposit:   deposit,
			Withdraw:  withdraw,
			Borrow:    borrow,
			Repay:     repay,
			Swap:      swap,
			Liquidate: liquidate,
		}
		(*bots)[i] = bot
	}
	return nil
}

func create_compound_bots(bots *BotData) error {
	var compound_bots []CompoundBot
	for name, protocol := range GLOBAL.CompoundLikeProtocols {
		client, err := ethclient.Dial(NETWORK.RPC)
		if err != nil {
			return err
		}
		troller, err := comptroller.NewComptroller(protocol.Address, client)
		if err != nil {
			return err
		}
		compound_bots = append(compound_bots, CompoundBot{
			Name:        name,
			Client:      client,
			Protocol:    protocol,
			Comptroller: troller,
			Incoming: CompIncomingChannels{
				MarketEntered:           make(chan *comptroller.ComptrollerMarketEntered),
				MarketExited:            make(chan *comptroller.ComptrollerMarketExited),
				DistributedBorrowerComp: make(chan *comptroller.ComptrollerDistributedBorrowerComp),
				DistributedSupplierComp: make(chan *comptroller.ComptrollerDistributedSupplierComp),
			},
			Shutdown: make(chan struct{}),
		})
	}
	bots.CompoundBots = compound_bots

	return nil
}

func create_aave_v2_bots(bots *BotData) error {
	var aave_v2_bots []AaveV2Bot
	for name, protocol := range GLOBAL.AaveV2LikeProtocols {
		client, err := ethclient.Dial(NETWORK.RPC)
		if err != nil {
			return err
		}
		lending_pool, err := lendingpool.NewLendingPool(protocol.Address, client)
		if err != nil {
			return err
		}
		aave_v2_bots = append(aave_v2_bots, AaveV2Bot{
			Name:        name,
			Client:      client,
			Protocol:    protocol,
			LendingPool: lending_pool,
			Incoming: AaveV2IncomingChannels{
				Deposit:   make(chan *lendingpool.LendingPoolDeposit),
				Withdraw:  make(chan *lendingpool.LendingPoolWithdraw),
				Borrow:    make(chan *lendingpool.LendingPoolBorrow),
				Repay:     make(chan *lendingpool.LendingPoolRepay),
				Swap:      make(chan *lendingpool.LendingPoolSwap),
				Liquidate: make(chan *lendingpool.LendingPoolLiquidationCall),
			},
			Shutdown: make(chan struct{}),
		})
	}
	bots.AaveV2Bots = aave_v2_bots

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

	if err := create_compound_bots(&bots); err != nil {
		return err
	}

	if err := create_aave_v2_bots(&bots); err != nil {
		return err
	}

	if err := create_comp_subs(&bots.CompoundBots); err != nil {
		return err
	}

	if err := create_aave_v2_subs(&bots.AaveV2Bots); err != nil {
		return err
	}

	fmt.Println("Finished creating all bots")
	utils.PrintDashed()

	var g1 errgroup.Group

	g1.Go(func() error {
		interrupt := make(chan os.Signal, 1)
		defer signal.Stop(interrupt)
		defer close(interrupt)
		defer GLOBAL.DB.Close()
		signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-interrupt
		time.Sleep(time.Second)
		fmt.Println("\nBot Stopped")
		return nil
	})

	g1.Go(func() error {
		for _, bot := range bots.CompoundBots {
			var g2 errgroup.Group

			g2.Go(func() error {
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
						store_event(bot.Name, payload.Account)
					case payload := <-bot.Incoming.MarketExited:
						store_event(bot.Name, payload.Account)
					case payload := <-bot.Incoming.DistributedBorrowerComp:
						store_event(bot.Name, payload.Borrower)
					case payload := <-bot.Incoming.DistributedSupplierComp:
						store_event(bot.Name, payload.Supplier)
					}
				}
			})

			bot2 := bot

			g2.Go(func() error {
				err := fetch_comp_events(bot2)
				if err != nil {
					return err
				}

				var g3 errgroup.Group

				iter := GLOBAL.DB.NewIterator(util.BytesPrefix([]byte(bot2.Name+"-user-")), nil)
				for iter.Next() {
					account := common.HexToAddress(hexutil.Encode(iter.Value()))
					g3.Go(func() error {
						raw_err_code, _, raw_shortfall, err := bot2.Comptroller.GetAccountLiquidity(nil, account)
						if err != nil {
							return err
						}

						err_code := decimal.NewDecFromBigIntWithPrec(raw_err_code, 18)
						shortfall := decimal.NewDecFromBigIntWithPrec(raw_shortfall, 18)
						if err_code.GT(decimal.ZeroDec()) {
							fmt.Println("error code??? :", err_code)
							return fmt.Errorf(account.String() + " getAccountLiquidity error: " + err_code.String())
						}

						if shortfall.GT(decimal.ZeroDec()) {
							c_tokens, err := bot2.Comptroller.GetAssetsIn(nil, account)
							if err != nil {
								return err
							}

							var (
								borrows     []CToken
								collaterals []CToken
							)
							for _, c_token_addr := range c_tokens {
								c_token, err := cerc20.NewCErc20(c_token_addr, bot2.Client)
								if err != nil {
									return err
								}

								borrowed_amount, err := c_token.BorrowBalanceStored(nil, account)
								if err != nil {
									return err
								}

								exchange_rate, err := c_token.ExchangeRateStored(nil)
								if err != nil {
									return err
								}

								if borrowed_amount.Cmp(constants.OneInt) == 1 {
									borrows = append(borrows, CToken{
										Address:      c_token_addr,
										Amount:       borrowed_amount,
										ExchangeRate: exchange_rate,
									})
								}

								c_token_balance, err := c_token.BalanceOf(nil, account)
								if err != nil {
									return err
								}

								if c_token_balance.Cmp(constants.OneInt) == 1 {
									collaterals = append(collaterals, CToken{
										Address:      c_token_addr,
										Amount:       c_token_balance,
										ExchangeRate: exchange_rate,
									})
								}
							}

							if len(borrows) > 0 {
								fmt.Println(account)
								fmt.Println("Shortfall", shortfall)
								fmt.Println("Collaterals", collaterals)
								fmt.Println("Borrows", borrows)
								utils.PrintDashed()
							}
						}

						return nil
					})
				}
				iter.Release()
				if err := g3.Wait(); err != nil {
					return err
				}

				fmt.Println(bot2.Name, "Done scanning positions")

				return nil
			})
		}

		return nil
	})

	g1.Go(func() error {
		for _, bot := range bots.AaveV2Bots {
			var g2 errgroup.Group

			g2.Go(func() error {
				for {
					select {
					case err := <-bot.Subscriptions.Deposit.Err():
						return err
					case err := <-bot.Subscriptions.Withdraw.Err():
						return err
					case err := <-bot.Subscriptions.Borrow.Err():
						return err
					case err := <-bot.Subscriptions.Repay.Err():
						return err
					case err := <-bot.Subscriptions.Swap.Err():
						return err
					case err := <-bot.Subscriptions.Liquidate.Err():
						return err
					case payload := <-bot.Incoming.Deposit:
						store_event(bot.Name, payload.User)
					case payload := <-bot.Incoming.Withdraw:
						store_event(bot.Name, payload.User)
					case payload := <-bot.Incoming.Borrow:
						store_event(bot.Name, payload.User)
					case payload := <-bot.Incoming.Repay:
						store_event(bot.Name, payload.User)
					case payload := <-bot.Incoming.Swap:
						store_event(bot.Name, payload.User)
					case payload := <-bot.Incoming.Liquidate:
						store_event(bot.Name, payload.User)
						store_event(bot.Name, payload.Liquidator)
					}
				}
			})

			bot_copy := bot

			g2.Go(func() error {
				return fetch_aave_v2_events(bot_copy)
			})
		}

		return nil
	})

	return g1.Wait()
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
