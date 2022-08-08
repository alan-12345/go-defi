package main

import (
	"fmt"
	"go_defi/addresses/polygon"
	"go_defi/contracts/uniswap/query"
	"go_defi/utils/array"
	"log"
	"math/big"
	"time"

	"github.com/ALTree/bigfloat"
	"github.com/ethereum/go-ethereum/common"
)

type pair_data struct {
	Tokens   tokens
	Reserves reserves_data
	Prices   prices_data
	Weights  weights
}

type tokens struct {
	Token0     common.Address
	Token1     common.Address
}

type reserves_data struct {
	Reserve0     *big.Int
	Reserve1     *big.Int
}

type prices_data struct {
	Price0 *big.Float
	Price1 *big.Float
}

type weights struct {
	Weight0 *big.Float
	Weight1 *big.Float
}

type edge struct {
	Source common.Address
	Dest   common.Address
	Pair   common.Address
	Weight *big.Float
	Price  *big.Float
}

var (
	a          = polygon_addresses.TOKEN_ADDRS
	rev_a      = polygon_addresses.REVERSE_NAMING
	a_decimals = polygon_addresses.TOKEN_DECIMALS
	factories  = polygon_addresses.FACTORY_ADDRESSES
	all_tokens = polygon_addresses.TRADABLE_TOKENS
	inf        = new(big.Float).SetInf(false)
	neg_one    = new(big.Float).SetFloat64(-1)
	one        = new(big.Float).SetFloat64(1)
	zero       = new(big.Float).SetFloat64(0)
)

func generate_all_pairs() [][]common.Address {
	var pairs [][]common.Address
	for i := 0; i < len(all_tokens); i++ {
		for j := 0; j < len(all_tokens); j++ {
			if i < j {
				if all_tokens[i].Hash().Big().Cmp(all_tokens[j].Hash().Big()) == -1 {
					pairs = append(pairs, []common.Address{all_tokens[i], all_tokens[j]})
				} else {
					pairs = append(pairs, []common.Address{all_tokens[j], all_tokens[i]})
				}
			}
		}
	}
	return pairs
}

func fetch_pair_addrs(query_contract *query.UniswapQuery, pairs [][]common.Address) [][]common.Address {
	start := time.Now()

	var queries = []query.PairQuery{}
	for _, factory := range factories {
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
		pair_addrs = append(pair_addrs, []common.Address{pair_addr, queries[i].TokenA, queries[i].TokenB})
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

func get_amount_out(reserve_in *big.Int, reserve_out *big.Int, amount_in *big.Int) *big.Int {
	amount_in_w_fee := new(big.Int).Mul(amount_in, big.NewInt(997))
	numerator := new(big.Int).Mul(amount_in_w_fee, reserve_out)
	denominator := new(big.Int).Mul(reserve_in, big.NewInt(1000))
	return new(big.Int).Div(numerator, denominator)
}

func compile_data(pair_addrs [][]common.Address, reserves [][2]*big.Int) map[common.Address]pair_data {
	var all_data = make(map[common.Address]pair_data)
	for i, pair_addr := range pair_addrs {
		token_0_decimals := a_decimals[rev_a[pair_addr[1]]]
		token_0_decimal_pow := bigfloat.Pow(big.NewFloat(10), new(big.Float).SetInt(token_0_decimals))
		token_1_decimals := a_decimals[rev_a[pair_addr[2]]]
		token_1_decimal_pow := bigfloat.Pow(big.NewFloat(10), new(big.Float).SetInt(token_1_decimals))

		reserve_0_fmt := new(big.Float).Quo(new(big.Float).SetInt(reserves[i][0]), token_0_decimal_pow)
		reserve_1_fmt := new(big.Float).Quo(new(big.Float).SetInt(reserves[i][1]), token_1_decimal_pow)
		spot_price_0 := get_spot_price(reserve_0_fmt, reserve_1_fmt)
		spot_price_1 := get_spot_price(reserve_1_fmt, reserve_0_fmt)
		weighted_price_0 := get_weighted_price(spot_price_0)
		weighted_price_1 := get_weighted_price(spot_price_1)

		// store all data
		all_data[pair_addr[0]] = pair_data{
			Tokens: tokens{
				Token0:     pair_addr[1],
				Token1:     pair_addr[2],
			},
			Reserves: reserves_data{
				Reserve0:     reserves[i][0],
				Reserve1:     reserves[i][1],
			},
			Prices: prices_data{
				Price0: spot_price_0,
				Price1: spot_price_1,
			},
			Weights: weights{
				Weight0: weighted_price_0,
				Weight1: weighted_price_1,
			},
		}
	}
	return all_data
}

func generate_nodes(all_tokens []common.Address) map[common.Address]int {
	var nodes = make(map[common.Address]int)
	for i, token := range all_tokens {
		nodes[token] = i
	}
	return nodes
}

func generate_edges(all_data map[common.Address]pair_data) []edge {
	var edges []edge
	for k, v := range all_data {
		tmp_edge := edge{
			Source: v.Tokens.Token0,
			Dest:   v.Tokens.Token1,
			Pair:   k,
			Weight: v.Weights.Weight0,
			Price:  v.Prices.Price0,
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
		tmp_edge = edge{
			Source: v.Tokens.Token1,
			Dest:   v.Tokens.Token0,
			Pair:   k,
			Weight: v.Weights.Weight1,
			Price:  v.Prices.Price1,
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

func bellman_ford(nodes map[common.Address]int, edges []edge) {
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

	for _, edge := range edges {
		source := nodes[edge.Source]
		dest := nodes[edge.Dest]
		lhs := new(big.Float).Add(distance[source], edge.Weight)
		rhs := distance[dest]
		if lhs.Cmp(rhs) == -1 {
			print_cycle := []int{dest, source}
			for array.IndexOf(predecessor[source], print_cycle) == -1 {
				print_cycle = append(print_cycle, predecessor[source])
				source = predecessor[source]
			}
			print_cycle = append(print_cycle, predecessor[source])

			for i, j := 0, len(print_cycle)-1; i < j; i, j = i+1, j-1 {
				print_cycle[i], print_cycle[j] = print_cycle[j], print_cycle[i]
			}
			start := print_cycle[0]
			if print_cycle[len(print_cycle)-1] != start {
				print_cycle = append(print_cycle, start)
			}

			if len(print_cycle) >= 3 {
				var path string
				for i, step := range print_cycle {
					path += rev_a[all_tokens[step]]
					if i < len(print_cycle)-1 {
						path += " -> "
					}
				}
				fmt.Println(path)

				// calculate profit
				profit_multiplier := big.NewFloat(1)
				for i := 0; i < len(print_cycle)-1; i++ {
					source := all_tokens[print_cycle[i]]
					dest := all_tokens[print_cycle[i+1]]
					for _, edge := range edges {
						if edge.Source == source && edge.Dest == dest {
							fmt.Println(rev_a[source], "->", rev_a[dest], edge.Price)
							profit_multiplier = new(big.Float).Mul(profit_multiplier, edge.Price)
							continue
						}
					}
				}
				profit_percentage := new(big.Float).Sub(profit_multiplier, big.NewFloat(1))
				profit_percentage = new(big.Float).Mul(profit_percentage, big.NewFloat(100))
				fmt.Println("Profit percentage:", profit_percentage, "%")
				fmt.Println("------------------------------------------")
			}
		}
	}
}

func arbitrage(query_contract *query.UniswapQuery, raw_pair_addrs [][]common.Address) {
	start := time.Now()
	var pair_addrs []common.Address
	for _, pair_addr := range raw_pair_addrs {
		pair_addrs = append(pair_addrs, pair_addr[0])
	}

	reserves := fetch_reserves(query_contract, pair_addrs)

	all_data := compile_data(raw_pair_addrs, reserves)

	edges := generate_edges(all_data)

	nodes := generate_nodes(all_tokens)

	bellman_ford(nodes, edges)

	end := time.Now()
	log.Println("Cycled in", end.Sub(start).String()+"\n")
}
