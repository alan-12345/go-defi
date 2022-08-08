//SPDX-License-Identifier: UNLICENSED
pragma solidity 0.6.12;

pragma experimental ABIEncoderV2;

interface IUniswapV2Pair {
    function token0() external view returns (address);
    function token1() external view returns (address);
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
}

abstract contract UniswapV2Factory  {
    mapping(address => mapping(address => address)) public getPair;
}

struct PairQuery {
    UniswapV2Factory factory;
    address tokenA;
    address tokenB;
}

contract UniswapQuery {
    function getPairs(PairQuery[] calldata queries) external view returns (address[] memory) {
        address[] memory result = new address[](queries.length);
        for (uint i = 0; i < queries.length; i++) {
            result[i] = getPair(queries[i]);
        }
        return result;
    }

    function getPair(PairQuery memory query) public view returns (address) {
        return query.factory.getPair(query.tokenA, query.tokenB);
    }

    function getReservesByPairs(IUniswapV2Pair[] calldata _pairs) external view returns (uint[2][] memory) {
        uint[2][] memory result = new uint[2][](_pairs.length);
        for (uint i = 0; i < _pairs.length; i++) {
            IUniswapV2Pair pair = _pairs[i];
            (uint112 r0, uint112 r1, ) = pair.getReserves();
            (result[i][0], result[i][1]) = (r0, r1);
        }
        return result;
    }
}