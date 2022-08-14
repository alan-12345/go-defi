//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

pragma experimental ABIEncoderV2;

enum SwapType{ UniswapV2, UniswapV3, CurveStableBase, CurveStableUnderlying, CurveCryptoBase, CurveCryptoUnderlying, CurveMetaPool }

struct SwapCall {
    address pool;
    SwapType swapType;
    address tokenIn;
    address tokenOut;
    uint i;
    uint j;
}

interface IUniswapV2 {
    function getReserves() external view returns (uint112 r0, uint112 r1, uint32);
    function token0() external view returns (address);
}

interface IUniswapV3 {
    function fee() external view returns (uint24);
}

interface IQuoter {
    function quoteExactInputSingle(
        address tokenIn,
        address tokenOut,
        uint24 fee,
        uint256 amountIn,
        uint160 sqrtPriceLimitX96
    ) external returns (uint256 amountOut);
}

interface ICurve {
    function get_dy(int128 i, int128 j, uint dx) external view returns (uint);
    function get_dy(uint i, uint j, uint dx) external view returns (uint);

    function get_dy_underlying(int128 i, int128 j, uint dx) external view returns (uint);
    function get_dy_underlying(uint i, uint j, uint dx) external view returns (uint);
}

contract Bundler {
    address public constant UNISWAP_QUOTER = address(0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6);
    
    function getAmountsOut(SwapCall[] memory calls, uint amountIn) public returns (uint[] memory amountsOut) {
        amountsOut = new uint[](calls.length + 1);
        amountsOut[0] = amountIn;

        for (uint i = 0; i < calls.length; i++) {
            SwapCall memory call = calls[i];
            uint amountOut;
            if (call.swapType == SwapType.UniswapV2){
                amountOut = getUniswapV2AmountOut(call.pool, call.tokenIn, amountIn);  
            } else if (call.swapType == SwapType.UniswapV3) {
                amountOut = getUniswapV3AmountOut(call.pool, call.tokenIn, call.tokenOut, amountIn);
            } else if (call.swapType == SwapType.CurveStableBase) {
                amountOut = ICurve(call.pool).get_dy(int128(int(call.i)), int128(int(call.j)), amountIn);
            } else if (call.swapType == SwapType.CurveStableUnderlying) {
                amountOut = ICurve(call.pool).get_dy_underlying(int128(int(call.i)), int128(int(call.j)), amountIn);
            } else if (call.swapType == SwapType.CurveCryptoBase) {
                amountOut = ICurve(call.pool).get_dy(call.i, call.j, amountIn);
            } else if (call.swapType == SwapType.CurveCryptoUnderlying) {
                amountOut = ICurve(call.pool).get_dy_underlying(call.i, call.j, amountIn);
            } else if (call.swapType == SwapType.CurveMetaPool) {
                amountOut = ICurve(call.pool).get_dy_underlying(int128(int(call.i)), int128(int(call.j)), amountIn);
            }
            
            amountsOut[i+1] = amountOut;
            amountIn = amountOut;
        }
    }

    function getUniswapV2AmountOut(address pool, address tokenIn, uint amountIn) public view returns (uint amountOut) {
        (uint112 r0, uint112 r1,) = IUniswapV2(pool).getReserves();
        address token0 = IUniswapV2(pool).token0();
        (uint reserveIn, uint reserveOut) = tokenIn == token0 ? (r0, r1) : (r1, r0);

        uint amountInWithFee = amountIn * 997;
        uint numerator = amountInWithFee * reserveOut;
        uint denominator = (reserveIn * 1000) + amountInWithFee;
        amountOut = numerator / denominator;
    }

    function getUniswapV3AmountOut(address pool, address tokenIn, address tokenOut, uint amountIn) public returns (uint amountOut) {
        uint24 fee = IUniswapV3(pool).fee();
        amountOut = IQuoter(UNISWAP_QUOTER).quoteExactInputSingle(tokenIn, tokenOut, fee, amountIn, 0);
    }
}