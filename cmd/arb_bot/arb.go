package main

import (
	"context"
	"flag"
	"fmt"
	"go_defi/networks/ethereum"
	"go_defi/utils"
	"strconv"

	// "go_defi/networks/fantom"
	"go_defi/networks/polygon"
	"go_defi/contracts/bundler"
	"go_defi/contracts/multicall"
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
			Bundler:     ethereum_addresses.BUNDLER_ADDR,
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
	Price    decimal.Dec
	Weight   decimal.Dec
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
		var swap_calls []bundler.SwapCall
		i := big.NewInt(int64(array.TokenIndexOf(edge.Source, edge.PoolData.Tokens)))
		j := big.NewInt(int64(array.TokenIndexOf(edge.Dest, edge.PoolData.Tokens)))
		swap_calls = append(swap_calls, bundler.SwapCall{
			Pool:     edge.Pool,
			SwapType: edge.PoolData.SwapType,
			TokenIn:  edge.Source.Address,
			TokenOut: edge.Dest.Address,
			I:        i,
			J:        j,
		})
		encoded_args := crypto.EncodeArgs(bundler.BundlerMetaData.ABI, "getAmountsOut", swap_calls, edge.Source.Size)
		generated_calls = append(generated_calls, multicall.Multicall2Call{
			Target:   GLOBAL.Bundler,
			CallData: encoded_args,
		})
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
		decoded_data := crypto.DecodeData(bundler.BundlerMetaData.ABI, "getAmountsOut", call)[0].([]*big.Int)
		amount_in := decimal.NewDecFromBigIntWithPrec(decoded_data[0], edge.Source.Decimals)
		amount_out := decimal.NewDecFromBigIntWithPrec(decoded_data[1], edge.Dest.Decimals)
		price := amount_out.Quo(amount_in)
		edge.Price = price
		lg := bigfloat.Log(big.NewFloat(price.MustFloat64()))
		weight := lg.Mul(lg, constants.NegOne)
		weight_float64, _ := weight.Float64()
		dec_weight := decimal.MustNewDecFromStr(strconv.FormatFloat(weight_float64, 'f', 15, 64))
		edge.Weight = dec_weight
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
				if edge.Weight.LT(seen_edge.Weight) {
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
	distance := make(map[constants.Token]decimal.Dec)
	predecessor := make(map[constants.Token]constants.Token)

	for _, node := range nodes {
		distance[node] = decimal.NewDecFromBigInt(constants.ReallyBigInt)
	}
	source := nodes[0]
	distance[source] = decimal.ZeroDec()

	for i := 0; i < n-1; i++ {
		for _, edge := range edges {
			source := edge.Source
			dest := edge.Dest
			lhs := distance[source].Add(edge.Weight)
			rhs := distance[dest]
			if lhs.LT(rhs) {
				distance[dest] = lhs
				predecessor[dest] = source
			}
		}
	}

	var edge_paths [][]Edge
	for _, edge := range edges {
		source := edge.Source
		dest := edge.Dest
		lhs := distance[source].Add(edge.Weight)
		rhs := distance[dest]
		if lhs.LT(rhs) {
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

			theoretical_profitability := decimal.NewDec(1)
			for _, edge := range edge_path {
				theoretical_profitability = theoretical_profitability.Mul(edge.Price)
				fmt.Println(
					GLOBAL.LookUp[edge.Source.Address], "=>",
					GLOBAL.LookUp[edge.Dest.Address],
					"(", edge.PoolData.Protocol, ":", edge.PoolData.Name, ")",
				)
			}
			theoretical_profitability = theoretical_profitability.Sub(decimal.NewDec(1))
			theoretical_profitability = theoretical_profitability.Mul(decimal.NewDec(100))
			fmt.Println("Theoretical profitability:", theoretical_profitability, "%")
			fmt.Println()
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
	iterations := 7

	for _, edge_path := range *paths {
		start_size := edge_path[0].Source.Size
		for i := 0; i < iterations; i++ {
			var swap_calls []bundler.SwapCall
			for _, edge := range edge_path {
				i := big.NewInt(int64(array.TokenIndexOf(edge.Source, edge.PoolData.Tokens)))
				j := big.NewInt(int64(array.TokenIndexOf(edge.Dest, edge.PoolData.Tokens)))
				swap_calls = append(swap_calls, bundler.SwapCall{
					Pool:     edge.Pool,
					SwapType: edge.PoolData.SwapType,
					TokenIn:  edge.Source.Address,
					TokenOut: edge.Dest.Address,
					I:        i,
					J:        j,
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
			arbs_to_take = append(arbs_to_take, edge_path)

			arb_details := "Network: " + *SELECTED_NETWORK + "\n"
			for _, edge := range edge_path {
				arb_details += GLOBAL.LookUp[edge.Source.Address] + " => " +
					GLOBAL.LookUp[edge.Dest.Address] +
					" (" + edge.PoolData.Protocol + ": " + edge.PoolData.Name + ")" + "\n"
			}

			arb_details += "\nOptimal trade size: " + optimal_size.String() + " " + GLOBAL.LookUp[edge_path[0].Source.Address] + "\n"
			arb_details += "Net profit: " + max_profit.String() + " " + GLOBAL.LookUp[edge_path[0].Source.Address]
			fmt.Println(arb_details)

			utils.SendTelegramMessage(arb_details)
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

			// execute_arbs(paths)

			end := time.Now()
			log.Println("Time elapsed (since block):", end.Sub(start).String())
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
