package main

import (
	"context"
	"flag"
	"fmt"
	"go_defi/addresses/ethereum"
	"go_defi/addresses/polygon"
	"go_defi/contracts/uniswap/query"
	"go_defi/utils/array"
	"log"
	"math/big"
	"time"

	"github.com/ALTree/bigfloat"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type network_data struct {
	rpc           string
	QueryAddress  common.Address
	Factories     []common.Address
	Tokens        []common.Address
	TokenDecimals map[string]*big.Int
	RevLookup     map[common.Address]string
}

var (
	network = flag.String("network", "ethereum", "Network")
	configs = map[string]network_data{
		"ethereum": {
			rpc:           ethereum_addresses.RPC_URL,
			QueryAddress:  ethereum_addresses.UNISWAP_QUERY_ADDR,
			Factories:     ethereum_addresses.FACTORY_ADDRESSES,
			Tokens:        ethereum_addresses.TRADABLE_TOKENS,
			TokenDecimals: ethereum_addresses.TOKEN_DECIMALS,
			RevLookup:     ethereum_addresses.REVERSE_NAMING,
		},
		"polygon": {
			rpc:           polygon_addresses.RPC_URL,
			QueryAddress:  polygon_addresses.UNISWAP_QUERY_ADDR,
			Factories:     polygon_addresses.FACTORY_ADDRESSES,
			Tokens:        polygon_addresses.TRADABLE_TOKENS,
			TokenDecimals: polygon_addresses.TOKEN_DECIMALS,
			RevLookup:     polygon_addresses.REVERSE_NAMING,
		},
	}
	config network_data
)

type pair_data struct {
	Tokens   tokens
	Reserves reserves_data
	Prices   prices_data
	Weights  weights_data
}

type tokens struct {
	Token0 common.Address
	Token1 common.Address
}

type reserves_data struct {
	Reserve0      *big.Int
	Reserve0_Norm *big.Float
	Reserve1      *big.Int
	Reserve1_Norm *big.Float
}

type prices_data struct {
	Price0 *big.Float
	Price1 *big.Float
}

type weights_data struct {
	Weight0 *big.Float
	Weight1 *big.Float
}

type edge_data struct {
	Source        common.Address
	Dest          common.Address
	Pair          common.Address
	Weight        *big.Float
	Price         *big.Float
	ReserveSource *big.Float
	ReserveDest   *big.Float
}

var (
	zero_addr = common.HexToAddress("0x0000000000000000000000000000000000000000")
	inf       = new(big.Float).SetInf(false)
	neg_inf   = new(big.Float).SetInf(true)
	one       = new(big.Float).SetFloat64(1)
	neg_one   = new(big.Float).SetFloat64(-1)
	zero      = new(big.Float).SetFloat64(0)
)

func generate_all_pairs() [][]common.Address {
	config := config
	var pairs [][]common.Address
	for i := 0; i < len(config.Tokens); i++ {
		for j := 0; j < len(config.Tokens); j++ {
			if i < j {
				if config.Tokens[i].Hash().Big().Cmp(config.Tokens[j].Hash().Big()) == -1 {
					pairs = append(pairs, []common.Address{config.Tokens[i], config.Tokens[j]})
				} else {
					pairs = append(pairs, []common.Address{config.Tokens[j], config.Tokens[i]})
				}
			}
		}
	}
	return pairs
}

func fetch_pair_addrs(query_contract *query.UniswapQuery, pairs [][]common.Address) [][]common.Address {
	start := time.Now()

	var queries = []query.PairQuery{}
	for _, factory := range config.Factories {
		for _, pair := range pairs {
			queries = append(queries, query.PairQuery{Factory: factory, TokenA: pair[0], TokenB: pair[1]})
		}
	}
	raw_pair_addrs, err := query_contract.GetPairs(nil, queries)
	if err != nil {
		log.Fatal(err)
	}
	var pair_addrs [][]common.Address
	for i, pair_addr := range raw_pair_addrs {
		if pair_addr != zero_addr {
			pair_addrs = append(pair_addrs, []common.Address{pair_addr, queries[i].TokenA, queries[i].TokenB})
		}
	}

	end := time.Now()
	log.Println("Fetched", len(pair_addrs), "pairs in", end.Sub(start).String())
	return pair_addrs
}

func fetch_reserves(query_contract *query.UniswapQuery, pair_addrs []common.Address) [][2]*big.Int {
	start := time.Now()
	reserves, err := query_contract.GetReservesByPairs(nil, pair_addrs)
	if err != nil {
		log.Fatal(err)
	}

	end := time.Now()
	log.Println("Fetched", len(reserves), "reserves in", end.Sub(start).String())
	return reserves
}

func get_spot_price(reserve_a *big.Float, reserve_b *big.Float) *big.Float {
	return new(big.Float).Quo(reserve_a, reserve_b)
}

func get_weighted_price(price *big.Float) *big.Float {
	lg := bigfloat.Log(price)
	return lg.Mul(lg, neg_one)
}

func get_weighted_prices(prices [][]*big.Float) [][]*big.Float {
	var weighted_prices [][]*big.Float
	for _, price := range prices {
		weighted_price_0 := get_weighted_price(price[0])
		weighted_price_1 := get_weighted_price(price[1])
		weighted_prices = append(weighted_prices, []*big.Float{weighted_price_0, weighted_price_1})
	}
	return weighted_prices
}

func get_amount_out(reserve_in *big.Float, reserve_out *big.Float, amount_in *big.Float) *big.Float {
	amount_in_w_fee := new(big.Float).Mul(amount_in, big.NewFloat(997))
	numerator := new(big.Float).Mul(amount_in_w_fee, reserve_out)
	denominator := new(big.Float).Mul(reserve_in, big.NewFloat(1000))
	denominator = new(big.Float).Add(denominator, amount_in_w_fee)
	return new(big.Float).Quo(numerator, denominator)
}

func compile_data(pair_addrs [][]common.Address, reserves [][2]*big.Int) map[common.Address]pair_data {
	var all_data = make(map[common.Address]pair_data)
	for i, pair_addr := range pair_addrs {
		token_0_decimals := config.TokenDecimals[config.RevLookup[pair_addr[1]]]
		token_0_decimal_pow := bigfloat.Pow(big.NewFloat(10), new(big.Float).SetInt(token_0_decimals))
		token_1_decimals := config.TokenDecimals[config.RevLookup[pair_addr[2]]]
		token_1_decimal_pow := bigfloat.Pow(big.NewFloat(10), new(big.Float).SetInt(token_1_decimals))

		reserve_0_norm := new(big.Float).Quo(new(big.Float).SetInt(reserves[i][0]), token_0_decimal_pow)
		reserve_1_norm := new(big.Float).Quo(new(big.Float).SetInt(reserves[i][1]), token_1_decimal_pow)
		spot_price_0 := get_spot_price(reserve_0_norm, reserve_1_norm)
		spot_price_1 := get_spot_price(reserve_1_norm, reserve_0_norm)
		weighted_price_0 := get_weighted_price(spot_price_0)
		weighted_price_1 := get_weighted_price(spot_price_1)

		all_data[pair_addr[0]] = pair_data{
			Tokens: tokens{
				Token0: pair_addr[1],
				Token1: pair_addr[2],
			},
			Reserves: reserves_data{
				Reserve0:      reserves[i][0],
				Reserve0_Norm: reserve_0_norm,
				Reserve1:      reserves[i][1],
				Reserve1_Norm: reserve_1_norm,
			},
			Prices: prices_data{
				Price0: spot_price_0,
				Price1: spot_price_1,
			},
			Weights: weights_data{
				Weight0: weighted_price_0,
				Weight1: weighted_price_1,
			},
		}
	}
	return all_data
}

func generate_nodes() map[common.Address]int {
	var nodes = make(map[common.Address]int)
	for i, token := range config.Tokens {
		nodes[token] = i
	}
	return nodes
}

func generate_edges(all_data map[common.Address]pair_data) []edge_data {
	var edges []edge_data
	for k, v := range all_data {
		tmp_edge := edge_data{
			Source:        v.Tokens.Token0,
			Dest:          v.Tokens.Token1,
			Pair:          k,
			Weight:        v.Weights.Weight0,
			Price:         v.Prices.Price0,
			ReserveSource: v.Reserves.Reserve0_Norm,
			ReserveDest:   v.Reserves.Reserve1_Norm,
		}

		idx := -1
		for i, edge := range edges {
			if edge.Source == tmp_edge.Source && edge.Dest == tmp_edge.Dest {
				idx = i
			}
		}
		if idx >= 0 {
			dup_edge := edges[idx]
			if tmp_edge.Weight.Cmp(dup_edge.Weight) == -1 {
				edges[idx] = tmp_edge
			}
		} else {
			edges = append(edges, tmp_edge)
		}

		idx = -1
		tmp_edge = edge_data{
			Source:        v.Tokens.Token1,
			Dest:          v.Tokens.Token0,
			Pair:          k,
			Weight:        v.Weights.Weight1,
			Price:         v.Prices.Price1,
			ReserveSource: v.Reserves.Reserve1_Norm,
			ReserveDest:   v.Reserves.Reserve0_Norm,
		}

		for i, edge := range edges {
			if edge.Source == tmp_edge.Source && edge.Dest == tmp_edge.Dest {
				idx = i
			}
		}
		if idx >= 0 {
			dup_edge := edges[idx]
			if tmp_edge.Weight.Cmp(dup_edge.Weight) == -1 {
				edges[idx] = tmp_edge
			}
		} else {
			edges = append(edges, tmp_edge)
		}
	}
	return edges
}

func run_bellman_ford(nodes map[common.Address]int, edges []edge_data) [][]edge_data {
	start := time.Now()

	n := len(nodes)
	distance := make([]*big.Float, n)
	predecessor := make([]int, n)

	for i := range distance {
		distance[i] = inf
		predecessor[i] = -1
	}
	source := 0
	distance[source] = zero

	for i := 0; i < n-1; i++ {
		for _, edge := range edges {
			source := nodes[edge.Source]
			dest := nodes[edge.Dest]
			lhs := new(big.Float).Add(distance[source], edge.Weight)
			rhs := distance[dest]
			if lhs.Cmp(rhs) == -1 {
				distance[dest] = lhs
				predecessor[dest] = source
			}
		}
	}

	var edge_paths [][]edge_data
	for _, edge := range edges {
		source := nodes[edge.Source]
		dest := nodes[edge.Dest]
		lhs := new(big.Float).Add(distance[source], edge.Weight)
		rhs := distance[dest]
		if lhs.Cmp(rhs) == -1 {
			node_path := []int{dest, source}
			for array.IndexOf(predecessor[source], node_path) == -1 {
				node_path = append(node_path, predecessor[source])
				source = predecessor[source]
			}
			node_path = append(node_path, predecessor[source])

			for i, j := 0, len(node_path)-1; i < j; i, j = i+1, j-1 {
				node_path[i], node_path[j] = node_path[j], node_path[i]
			}
			start := node_path[0]
			if node_path[len(node_path)-1] != start {
				node_path = append(node_path, start)
			}

			var path string
			for i, step := range node_path {
				path += config.RevLookup[config.Tokens[step]]
				if i < len(node_path)-1 {
					path += " -> "
				}
			}
			fmt.Println(path + "\n")

			var edge_path []edge_data
			for i := 0; i < len(node_path)-1; i++ {
				source := config.Tokens[node_path[i]]
				dest := config.Tokens[node_path[i+1]]
				for _, edge := range edges {
					if edge.Source == source && edge.Dest == dest {
						edge_path = append(edge_path, edge)
					}
				}
			}
			edge_paths = append(edge_paths, edge_path)
		}
	}
	end := time.Now()
	log.Println("Found", len(edge_paths), "negative cycles in", end.Sub(start).String())
	return edge_paths
}

func find_best_paths(edge_paths [][]edge_data) [][]edge_data {
	start := time.Now()

	var best_paths [][]edge_data
	for _, edge_path := range edge_paths {
		swap_sizes := []*big.Float{
			big.NewFloat(0.001),
			big.NewFloat(0.01),
			big.NewFloat(0.1),
			big.NewFloat(1),
			big.NewFloat(10),
			big.NewFloat(100),
			big.NewFloat(1000),
		}
		for _, swap_amount := range swap_sizes {
			start_amount := swap_amount
			for _, edge := range edge_path {
				amount_out := get_amount_out(edge.ReserveSource, edge.ReserveDest, swap_amount)
				swap_amount = amount_out
			}
			net := new(big.Float).Sub(swap_amount, start_amount)
			if net.Cmp(zero) == 1 {
				best_paths = append(best_paths, edge_path)
			}
		}
	}

	for _, best_path := range best_paths {
		for _, step := range best_path {
			fmt.Println(config.RevLookup[step.Source], "->", config.RevLookup[step.Dest])
		}
		fmt.Println("---------------------")
		fmt.Println()
	}

	end := time.Now()
	log.Println("Found", len(best_paths), "profitable paths in", end.Sub(start).String())
	return best_paths
}

func execute_arbs(best_paths [][]edge_data) {}

func find_arbs(query_contract *query.UniswapQuery, raw_pair_addrs [][]common.Address) {
	start := time.Now()

	var pair_addrs []common.Address
	for _, pair_addr := range raw_pair_addrs {
		pair_addrs = append(pair_addrs, pair_addr[0])
	}

	reserves := fetch_reserves(query_contract, pair_addrs)

	all_data := compile_data(raw_pair_addrs, reserves)

	edges := generate_edges(all_data)

	nodes := generate_nodes()

	edge_paths := run_bellman_ford(nodes, edges)

	best_paths := find_best_paths(edge_paths)

	execute_arbs(best_paths)

	end := time.Now()
	log.Println("Completed search in", end.Sub(start).String()+"\n")
}

func start_bot() {
	config = configs[*network]

	client, err := ethclient.Dial(config.rpc)
	if err != nil {
		log.Fatal(err)
	}

	query_contract, err := query.NewUniswapQuery(config.QueryAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	pairs := generate_all_pairs()

	raw_pair_addrs := fetch_pair_addrs(query_contract, pairs)

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			log.Println("New block #", block.Number().Uint64())

			find_arbs(query_contract, raw_pair_addrs)
		}
	}
}
