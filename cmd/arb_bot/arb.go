package main

import (
	"context"
	"flag"
	"fmt"
	"go_defi/addresses/ethereum"
	"go_defi/utils"

	// "go_defi/addresses/fantom"
	"go_defi/addresses/polygon"
	"go_defi/contracts/bundler"
	"go_defi/contracts/curve/crypto-swap"
	"go_defi/contracts/curve/stable-swap"
	"go_defi/contracts/multicall"
	"go_defi/contracts/uniswap/v2/pair"
	"go_defi/contracts/uniswap/v3/quoter"
	"go_defi/utils/array"
	"go_defi/utils/constants"
	"go_defi/utils/crypto"
	"go_defi/utils/decimal"
	"log"
	"math/big"
	"time"

	"github.com/ALTree/bigfloat"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type GlobalData struct {
	Multicaller common.Address
	V3Quoter    common.Address
	Bundler     common.Address
	Pools       map[common.Address]constants.Pool
	Tokens      map[string]constants.Token
	LookUp      map[common.Address]string
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
		NETWORK = constants.NetworkData{
			RPC: ethereum_addresses.RPC_URL,
		}
		GLOBAL = GlobalData{
			Multicaller: ethereum_addresses.MULTICALL_ADDR,
			V3Quoter:    ethereum_addresses.UNISWAP_V3_QUOTER_ADDR,
			Pools:       ethereum_addresses.ALL_POOLS,
			Tokens:      ethereum_addresses.TOKENS,
			LookUp:      ethereum_addresses.LOOKUP,
		}
	case "polygon":
		NETWORK = constants.NetworkData{
			RPC: polygon_addresses.RPC_URL,
		}
		GLOBAL = GlobalData{
			Multicaller: polygon_addresses.MULTICALL_ADDR,
			V3Quoter:    polygon_addresses.UNISWAP_V3_QUOTER_ADDR,
			Bundler:     polygon_addresses.BUNDLER_ADDR,
			Pools:       polygon_addresses.ALL_POOLS,
			Tokens:      polygon_addresses.TOKENS,
			LookUp:      polygon_addresses.LOOKUP,
		}
	case "fantom":

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

type Edge struct {
	Source   constants.Token
	Dest     constants.Token
	Price    *big.Float
	Weight   *big.Float
	Pool     common.Address
	PoolData constants.Pool
}

func generate_edges(edges *[]Edge) {
	start := time.Now()

	generated_edges := []Edge{}
	for address, pool := range GLOBAL.Pools {
		var pairs [][]constants.Token
		tokens := pool.Tokens
		for i := 0; i < len(tokens); i++ {
			for j := 0; j < len(tokens); j++ {
				if i < j {
					pairs = append(pairs, []constants.Token{tokens[i], tokens[j]})
					pairs = append(pairs, []constants.Token{tokens[j], tokens[i]})
				}
			}
		}
		for _, pair := range pairs {
			new_edge := Edge{
				Source:   pair[0],
				Dest:     pair[1],
				Pool:     address,
				PoolData: pool,
			}
			generated_edges = append(generated_edges, new_edge)
		}
	}
	*edges = generated_edges

	end := time.Now()
	log.Println("Generated", len(*edges), "edges in", end.Sub(start).String())
	utils.PrintDashed()
}

func generate_calls(edges []Edge, calls *[]multicall.Multicall2Call) {
	start := time.Now()

	var generated_calls []multicall.Multicall2Call
	for _, edge := range edges {
		pool := edge.PoolData
		switch impl := pool.Implementation; impl {
		case "UniswapV2":
			encoded_args := crypto.EncodeArgs(pair.UniswapPairMetaData.ABI, "getReserves")
			generated_calls = append(generated_calls, multicall.Multicall2Call{
				Target:   edge.Pool,
				CallData: encoded_args,
			})
		case "CurveStableSwap":
			i := big.NewInt(int64(array.TokenIndexOf(edge.Source, pool.Tokens)))
			j := big.NewInt(int64(array.TokenIndexOf(edge.Dest, pool.Tokens)))
			encoded_args := crypto.EncodeArgs(stableswap.CurveStableSwapMetaData.ABI, "get_dy_underlying", i, j, edge.Source.Size)
			generated_calls = append(generated_calls, multicall.Multicall2Call{
				Target:   edge.Pool,
				CallData: encoded_args,
			})
		case "CurveCryptoSwap":
			i := big.NewInt(int64(array.TokenIndexOf(edge.Source, pool.Tokens)))
			j := big.NewInt(int64(array.TokenIndexOf(edge.Dest, pool.Tokens)))
			encoded_args := crypto.EncodeArgs(cryptoswap.CurveCryptoSwapMetaData.ABI, "get_dy", i, j, edge.Source.Size)
			generated_calls = append(generated_calls, multicall.Multicall2Call{
				Target:   edge.Pool,
				CallData: encoded_args,
			})
		case "UniswapV3":
			fee_fmt := new(big.Float).Mul(pool.Fee, big.NewFloat(1e6))
			fee_int64, _ := fee_fmt.Int64()
			fee_int := big.NewInt(fee_int64)
			encoded_args := crypto.EncodeArgs(quoter.UniswapV3QuoterMetaData.ABI,
				"quoteExactInputSingle",
				edge.Source.Address,
				edge.Dest.Address,
				fee_int,
				edge.Source.Size,
				big.NewInt(0),
			)
			generated_calls = append(generated_calls, multicall.Multicall2Call{
				Target:   GLOBAL.V3Quoter,
				CallData: encoded_args,
			})
		}
	}

	*calls = generated_calls

	end := time.Now()
	log.Println("Generated", len(*calls), "network calls in", end.Sub(start).String())
	utils.PrintDashed()
}

func populate_edges(calls []multicall.Multicall2Call, edges *[]Edge) {
	start := time.Now()
	encoded_calls := crypto.EncodeArgs(multicall.MulticallMetaData.ABI, "aggregate", calls)
	encoded_output := crypto.StaticCall(NETWORK.Client, GLOBAL.Multicaller, encoded_calls)
	decoded_output := (crypto.DecodeData(multicall.MulticallMetaData.ABI, "aggregate", encoded_output)[1]).([][]byte)
	populated_edges := []Edge{}
	for i, edge := range *edges {
		call := decoded_output[i]
		source_prec := edge.Source.Precision
		dest_prec := edge.Dest.Precision
		var price *big.Float
		switch impl := edge.PoolData.Implementation; impl {
		case "UniswapV2":
			decoded_data := crypto.DecodeData(pair.UniswapPairMetaData.ABI, "getReserves", call)
			reserve_0 := new(big.Float).SetInt(decoded_data[0].(*big.Int))
			reserve_1 := new(big.Float).SetInt(decoded_data[1].(*big.Int))
			var source_reserve, dest_reserve *big.Float
			if edge.PoolData.Tokens[0] == edge.Source {
				source_reserve = reserve_0
				dest_reserve = reserve_1
			} else {
				source_reserve = reserve_1
				dest_reserve = reserve_0
			}
			source_reserve_fmt := new(big.Float).Quo(source_reserve, new(big.Float).SetInt(source_prec))
			dest_reserve_fmt := new(big.Float).Quo(dest_reserve, new(big.Float).SetInt(dest_prec))
			price = new(big.Float).Quo(dest_reserve_fmt, source_reserve_fmt)
		case "CurveStableSwap":
			decoded_data := crypto.DecodeData(stableswap.CurveStableSwapMetaData.ABI, "get_dy_underlying", call)
			amount_out := new(big.Float).SetInt(decoded_data[0].(*big.Int))
			amount_out_fmt := new(big.Float).Quo(amount_out, new(big.Float).SetInt(dest_prec))
			amount_in := new(big.Float).SetInt(edge.Source.Size)
			amount_in_fmt := new(big.Float).Quo(amount_in, new(big.Float).SetInt(source_prec))
			price = new(big.Float).Quo(amount_out_fmt, amount_in_fmt)
		case "CurveCryptoSwap":
			decoded_data := crypto.DecodeData(cryptoswap.CurveCryptoSwapMetaData.ABI, "get_dy", call)
			amount_out := new(big.Float).SetInt(decoded_data[0].(*big.Int))
			amount_out_fmt := new(big.Float).Quo(amount_out, new(big.Float).SetInt(dest_prec))
			amount_in := new(big.Float).SetInt(edge.Source.Size)
			amount_in_fmt := new(big.Float).Quo(amount_in, new(big.Float).SetInt(source_prec))
			price = new(big.Float).Quo(amount_out_fmt, amount_in_fmt)
		case "UniswapV3":
			decoded_data := crypto.DecodeData(quoter.UniswapV3QuoterMetaData.ABI, "quoteExactInputSingle", call)
			amount_out := new(big.Float).SetInt(decoded_data[0].(*big.Int))
			amount_out_fmt := new(big.Float).Quo(amount_out, new(big.Float).SetInt(dest_prec))
			amount_in := new(big.Float).SetInt(edge.Source.Size)
			amount_in_fmt := new(big.Float).Quo(amount_in, new(big.Float).SetInt(source_prec))
			price = new(big.Float).Quo(amount_out_fmt, amount_in_fmt)
		}
		edge.Price = price
		lg := bigfloat.Log(price)
		weight := lg.Mul(lg, constants.NegOne)
		edge.Weight = weight
		populated_edges = append(populated_edges, edge)
	}
	*edges = populated_edges

	end := time.Now()
	log.Println("Fetched and populated", len(*edges), "edge weights in", end.Sub(start).String())
	utils.PrintDashed()
}

func filter_duplicate_edges(edges *[]Edge) {
	start := time.Now()

	var seen_edges []Edge
	for _, edge := range *edges {
		seen := false
		for i, seen_edge := range seen_edges {
			if seen_edge.Source == edge.Source && seen_edge.Dest == edge.Dest {
				seen = true
				if edge.Weight.Cmp(seen_edge.Weight) == -1 {
					seen_edges[i] = edge
				}
				break
			}
		}
		if !seen {
			seen_edges = append(seen_edges, edge)
		}
	}
	*edges = seen_edges

	end := time.Now()
	log.Println("Filtered and reduced to", len(*edges), "unique edges in", end.Sub(start).String())
	utils.PrintDashed()
}

func run_bellman_ford(edges []Edge, paths *[][]Edge) {
	start := time.Now()

	var nodes []constants.Token
	for _, token := range GLOBAL.Tokens {
		nodes = append(nodes, token)
	}

	n := len(nodes)
	distance := make(map[constants.Token]*big.Float)
	predecessor := make(map[constants.Token]constants.Token)

	for _, node := range nodes {
		distance[node] = constants.Inf
	}
	source := nodes[0]
	distance[source] = constants.Zero

	for i := 0; i < n-1; i++ {
		for _, edge := range edges {
			source := edge.Source
			dest := edge.Dest
			lhs := new(big.Float).Add(distance[source], edge.Weight)
			rhs := distance[dest]
			if lhs.Cmp(rhs) == -1 {
				distance[dest] = lhs
				predecessor[dest] = source
			}
		}
	}

	var edge_paths [][]Edge
	for _, edge := range edges {
		source := edge.Source
		dest := edge.Dest
		lhs := new(big.Float).Add(distance[source], edge.Weight)
		rhs := distance[dest]
		if lhs.Cmp(rhs) == -1 {
			node_path := []constants.Token{dest, source}

			for array.TokenIndexOf(predecessor[source], node_path) == -1 {
				node_path = append(node_path, predecessor[source])
				source = predecessor[source]
			}
			node_path = append(node_path, predecessor[source])

			for i, j := 0, len(node_path)-1; i < j; i, j = i+1, j-1 {
				node_path[i], node_path[j] = node_path[j], node_path[i]
			}

			if node_path[0] != node_path[len(node_path)-1] {
				continue
			}

			var edge_path []Edge
			for i := 0; i < len(node_path)-1; i++ {
				source := node_path[i]
				dest := node_path[i+1]
				for _, edge := range edges {
					if edge.Source == source && edge.Dest == dest {
						edge_path = append(edge_path, edge)
					}
				}
			}
			edge_paths = append(edge_paths, edge_path)

			// theoretical_profitability := big.NewFloat(1.0)
			// for _, edge := range edge_path {
			// theoretical_profitability = new(big.Float).Mul(theoretical_profitability, edge.Price)
			// fmt.Println(
			// 	GLOBAL.LookUp[edge.Source.Address], "=>",
			// 	GLOBAL.LookUp[edge.Dest.Address],
			// 	"(", edge.PoolData.Protocol, ":", edge.PoolData.Name, ")",
			// )
			// }
			// theoretical_profitability = new(big.Float).Sub(theoretical_profitability, big.NewFloat(1))
			// theoretical_profitability = new(big.Float).Mul(theoretical_profitability, big.NewFloat(100))
			// fmt.Println("Theoretical profitability:", theoretical_profitability, "%")
			// fmt.Println()
		}
	}
	*paths = edge_paths

	end := time.Now()
	log.Println("Detected", len(*paths), "negative cycles in", end.Sub(start).String())
	utils.PrintDashed()
}

func filter_profitable_paths(paths *[][]Edge) {
	start := time.Now()

	var calls []multicall.Multicall2Call
	iterations := 5

	for _, edge_path := range *paths {
		start_size := edge_path[0].Source.Size
		for i := 0; i < iterations; i++ {
			var swap_calls []bundler.SwapCall
			for _, edge := range edge_path {
				swap_calls = append(swap_calls, bundler.SwapCall{
					Target:   edge.Pool,
					SwapType: edge.PoolData.SwapType,
					TokenIn:  edge.Source.Address,
					TokenOut: edge.Dest.Address,
				})
			}
			encoded_args := crypto.EncodeArgs(bundler.BundlerMetaData.ABI, "getAmountsOut", swap_calls, start_size)
			calls = append(calls, multicall.Multicall2Call{
				Target:   GLOBAL.Bundler,
				CallData: encoded_args,
			})
			start_size = new(big.Int).Mul(start_size, big.NewInt(10))
		}
	}

	encoded_calls := crypto.EncodeArgs(multicall.MulticallMetaData.ABI, "aggregate", calls)
	encoded_output := crypto.StaticCall(NETWORK.Client, GLOBAL.Multicaller, encoded_calls)
	decoded_output := (crypto.DecodeData(multicall.MulticallMetaData.ABI, "aggregate", encoded_output)[1]).([][]byte)

	var arbs_to_take [][]Edge

	for i, edge_path := range *paths {
		token_decimals := edge_path[0].Source.Decimals
		max_profit := decimal.ZeroDec()
		optimal_size := decimal.ZeroDec()
		for j := 0; j < iterations; j++ {
			call := decoded_output[iterations*i+j]
			decoded_data := crypto.DecodeData(bundler.BundlerMetaData.ABI, "getAmountsOut", call)[0].([]*big.Int)
			start_amount := decimal.NewDecFromBigIntWithPrec(decoded_data[0], token_decimals)
			end_amount := decimal.NewDecFromBigIntWithPrec(decoded_data[len(decoded_data)-1], token_decimals)
			net := end_amount.Sub(start_amount)
			if net.GT(max_profit) {
				max_profit = net
				optimal_size = start_amount
			}
			// for k, edge := range edge_path {
			// 	amount_in := decimal.NewDecFromBigIntWithPrec(decoded_data[k], edge.Source.Decimals)
			// 	amount_out := decimal.NewDecFromBigIntWithPrec(decoded_data[k+1], edge.Dest.Decimals)

			// 	fmt.Println(
			// 		amount_in,
			// 		GLOBAL.LookUp[edge.Source.Address], "=>", amount_out,
			// 		GLOBAL.LookUp[edge.Dest.Address],
			// 		"(", edge.PoolData.Protocol, ":", edge.PoolData.Name, ")",
			// 	)
			// }
			// fmt.Println("Net Profit:", net)
			// fmt.Println()
		}
		if max_profit.GT(decimal.ZeroDec()) {
			for _, edge := range edge_path {
				fmt.Println(
					GLOBAL.LookUp[edge.Source.Address], "=>",
					GLOBAL.LookUp[edge.Dest.Address],
					"(", edge.PoolData.Protocol, ":", edge.PoolData.Name, ")",
				)
			}
			arbs_to_take = append(arbs_to_take, edge_path)
			fmt.Println("Optimal trade size:", optimal_size, GLOBAL.LookUp[edge_path[0].Source.Address])
			fmt.Println("Net profit:", max_profit, GLOBAL.LookUp[edge_path[0].Source.Address])
		}
	}
	*paths = arbs_to_take

	end := time.Now()
	log.Println("Calculated", len(*paths), "profitable arbs in", end.Sub(start).String())
	utils.PrintDashed()
}

func execute_arbs(paths [][]Edge) {}

func find_arbs() {
	var edges []Edge
	generate_edges(&edges)

	var calls []multicall.Multicall2Call
	generate_calls(edges, &calls)

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

			edges := edges

			populate_edges(calls, &edges)

			filter_duplicate_edges(&edges)

			var paths [][]Edge

			run_bellman_ford(edges, &paths)

			filter_profitable_paths(&paths)

			execute_arbs(paths)

			end := time.Now()
			log.Println("Total time elapsed (since block):", end.Sub(start).String())
			utils.PrintDashed()
		}
	}
}

func start_bot() {
	utils.PrintDashed()
	log.Println("Running arb_bot")
	fmt.Println("Account:", crypto.GetPublicAddress())
	utils.PrintDashed()

	setup_global_data()

	find_arbs()
}
