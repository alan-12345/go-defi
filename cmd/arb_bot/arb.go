package main

import (
	"context"
	"flag"
	"fmt"
	"go_defi/addresses/ethereum"
	// "go_defi/addresses/fantom"
	// "go_defi/addresses/polygon"
	"go_defi/contracts/curve/crypto-swap"
	"go_defi/contracts/curve/stable-swap"
	"go_defi/contracts/multicall"
	"go_defi/contracts/uniswap/v2/pair"
	"go_defi/utils/array"
	"go_defi/utils/constants"
	"go_defi/utils/crypto"
	"log"
	"math/big"
	"time"

	"github.com/ALTree/bigfloat"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NetworkData struct {
	Client      *ethclient.Client
	RPC         string
	Multicaller common.Address
	Pools       map[common.Address]constants.Pool
	Tokens      map[string]constants.Token
	LookUp      map[common.Address]string
}

var (
	SELECTED_NETWORK = flag.String("network", "ethereum", "Network")
	GLOBAL           NetworkData
)

func setup_network_data() {
	fmt.Println("Selected Network:", *SELECTED_NETWORK)
	switch *SELECTED_NETWORK {
	case "ethereum":
		GLOBAL = NetworkData{
			RPC:         ethereum_addresses.RPC_URL,
			Multicaller: ethereum_addresses.MULTICALL_ADDR,
			Pools:       ethereum_addresses.ALL_POOLS,
			Tokens:      ethereum_addresses.TOKENS,
			LookUp:      ethereum_addresses.LOOKUP,
		}
	case "polygon":

	case "fantom":

	}

	client, err := ethclient.Dial(GLOBAL.RPC)
	if err != nil {
		log.Fatal(err)
	}
	GLOBAL.Client = client
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
	log.Println("Generated", len(generated_edges), "edges in", end.Sub(start).String())
	constants.PrintDashed()
}

func populate_edges(edges *[]Edge) {
	start := time.Now()

	var calls []multicall.Multicall2Call
	for _, edge := range *edges {
		switch impl := edge.PoolData.Implementation; impl {
		case "UniswapV2":
			encoded_args := crypto.EncodeArgs(pair.UniswapPairMetaData.ABI, "getReserves")
			calls = append(calls, multicall.Multicall2Call{
				Target:   edge.Pool,
				CallData: encoded_args,
			})
		case "CurveStableSwap":
			i := big.NewInt(int64(array.TokenIndexOf(edge.Source, edge.PoolData.Tokens)))
			j := big.NewInt(int64(array.TokenIndexOf(edge.Dest, edge.PoolData.Tokens)))
			encoded_args := crypto.EncodeArgs(stableswap.CurveStableSwapMetaData.ABI, "get_dy_underlying", i, j, edge.Source.Size)
			calls = append(calls, multicall.Multicall2Call{
				Target:   edge.Pool,
				CallData: encoded_args,
			})
		case "CurveCryptoSwap":
			i := big.NewInt(int64(array.TokenIndexOf(edge.Source, edge.PoolData.Tokens)))
			j := big.NewInt(int64(array.TokenIndexOf(edge.Dest, edge.PoolData.Tokens)))
			encoded_args := crypto.EncodeArgs(cryptoswap.CurveCryptoSwapMetaData.ABI, "get_dy", i, j, edge.Source.Size)
			calls = append(calls, multicall.Multicall2Call{
				Target:   edge.Pool,
				CallData: encoded_args,
			})
		}
	}
	encoded_calls := crypto.EncodeArgs(multicall.MulticallMetaData.ABI, "aggregate", calls)
	encoded_output := crypto.StaticCall(GLOBAL.Client, GLOBAL.Multicaller, encoded_calls)
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
		}
		edge.Price = price
		lg := bigfloat.Log(price)
		weight := lg.Mul(lg, constants.NegOne)
		edge.Weight = weight
		populated_edges = append(populated_edges, edge)
	}
	*edges = populated_edges

	end := time.Now()
	log.Println("Fetched and populated", len(calls), "edge weights in", end.Sub(start).String())
	constants.PrintDashed()
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
	log.Println("Filtered through", len(*edges), "edges in", end.Sub(start).String())
	constants.PrintDashed()
}

func run_bellman_ford(paths *[][]Edge, edges []Edge) {
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
			theoretical_profitability := big.NewFloat(1.0)
			for _, edge := range edge_path {
				theoretical_profitability = new(big.Float).Mul(theoretical_profitability, edge.Price)
				fmt.Println(
					GLOBAL.LookUp[edge.Source.Address], "=>",
					GLOBAL.LookUp[edge.Dest.Address],
					"(", edge.PoolData.Protocol, ":", edge.PoolData.Name, ")",
				)
			}
			edge_paths = append(edge_paths, edge_path)
			theoretical_profitability = new(big.Float).Sub(theoretical_profitability, big.NewFloat(1))
			theoretical_profitability = new(big.Float).Mul(theoretical_profitability, big.NewFloat(100))
			fmt.Println("Theoretical profitability:", theoretical_profitability, "%")
			fmt.Println()
		}
	}
	*paths = edge_paths

	end := time.Now()
	log.Println("Found", len(edge_paths), "negative cycles in", end.Sub(start).String())
	constants.PrintDashed()
}

func filter_profitable_paths(paths *[][]Edge) {
	start := time.Now()

	var best_paths [][]Edge
	for _, edge_path := range *paths {
		swap_sizes := []*big.Float{
			big.NewFloat(0.00001),
			big.NewFloat(0.0001),
			big.NewFloat(0.001),
			big.NewFloat(0.01),
			big.NewFloat(0.1),
			big.NewFloat(1),
			big.NewFloat(10),
			big.NewFloat(100),
			big.NewFloat(1000),
		}
		optimal_size := swap_sizes[0]
		max_profit := new(big.Float).SetInf(true)
		for _, swap_amount := range swap_sizes {
			start_amount := swap_amount
			for _, edge := range edge_path {
				amount_out := new(big.Float).Mul(swap_amount, edge.Price)
				swap_amount = amount_out
			}
			net := new(big.Float).Sub(swap_amount, start_amount)
			if net.Cmp(max_profit) == 1 {
				optimal_size = start_amount
				max_profit = net
			}
		}
		if max_profit.Cmp(constants.Zero) == 1 {
			best_paths = append(best_paths, edge_path)
			fmt.Println("Swap", optimal_size, GLOBAL.LookUp[edge_path[0].Source.Address])
			for _, edge := range edge_path {
				fmt.Println(
					GLOBAL.LookUp[edge.Source.Address], "=>",
					GLOBAL.LookUp[edge.Dest.Address],
					"(", edge.PoolData.Protocol, ":", edge.PoolData.Name, ")",
				)
			}
			fmt.Println("Profit:", max_profit, GLOBAL.LookUp[edge_path[0].Source.Address])
			fmt.Println()
		}
	}
	*paths = best_paths

	end := time.Now()
	log.Println("Found", len(best_paths), "profitable paths in", end.Sub(start).String())
	constants.PrintDashed()
}

func execute_arbs(best_paths [][]Edge) {}

func find_arbs() {
	start := time.Now()

	var edges []Edge

	generate_edges(&edges)

	populate_edges(&edges)

	filter_duplicate_edges(&edges)

	var paths [][]Edge

	run_bellman_ford(&paths, edges)

	filter_profitable_paths(&paths)

	// execute_arbs(best_paths)

	end := time.Now()
	log.Println("Completed search in", end.Sub(start).String())
	constants.PrintDashed()
}

func start_bot() {
	fmt.Println("Running arb_bot")
	fmt.Println("Account:", crypto.GetPublicAddress())

	setup_network_data()

	headers := make(chan *types.Header)
	sub, err := GLOBAL.Client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := GLOBAL.Client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("New block #", block.Number().Uint64())
			constants.PrintDashed()

			find_arbs()
		}
	}
}
