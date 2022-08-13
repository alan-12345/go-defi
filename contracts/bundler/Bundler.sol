//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

pragma experimental ABIEncoderV2;

enum SwapType{ UniswapV2, UniswapV3, CurveStableBase, CurveStableUnderlying, CurveCryptoBase, CurveCryptoUnderlying }

struct SwapCall {
    address target;
    SwapType swapType;
    address tokenIn;
    address tokenOut;
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

interface ICurveStableBase {
    function coins(uint) external view returns (address);
    function get_dy(int128 i, int128 j, uint dx) external view returns (uint);
}

interface ICurveStableUnderlying {
    function underlying_coins(uint) external view returns (address);
    function get_dy_underlying(int128 i, int128 j, uint dx) external view returns (uint);
}

interface ICurveCryptoBase {
    function coins(uint) external view returns (address);
    function get_dy(uint i, uint j, uint dx) external view returns (uint);
}

interface ICurveCryptoUnderlying {
    function underlying_coins(uint) external view returns (address);
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
                (uint112 r0, uint112 r1,) = IUniswapV2(call.target).getReserves();
                address token0 = IUniswapV2(call.target).token0();
                (uint reserveIn, uint reserveOut) = call.tokenIn == token0 ? (r0, r1) : (r1, r0);
                amountOut = getUniswapV2AmountOut(amountIn, reserveIn, reserveOut);  
            } else if (call.swapType == SwapType.UniswapV3) {
                uint24 fee = IUniswapV3(call.target).fee();
                amountOut = IQuoter(UNISWAP_QUOTER).quoteExactInputSingle(call.tokenIn, call.tokenOut, fee, amountIn, 0);
            } else if (call.swapType == SwapType.CurveStableBase) {
                int128 i_idx = int128(int(getCurveIndex(call.target, call.tokenIn)));
                int128 j_idx = int128(int(getCurveIndex(call.target, call.tokenOut)));
                amountOut = ICurveStableBase(call.target).get_dy(i_idx, j_idx, amountIn);
            } else if (call.swapType == SwapType.CurveStableUnderlying) {
                int128 i_idx = int128(int(getCurveIndexUnderlying(call.target, call.tokenIn)));
                int128 j_idx = int128(int(getCurveIndexUnderlying(call.target, call.tokenOut)));
                amountOut = ICurveStableUnderlying(call.target).get_dy_underlying(i_idx, j_idx, amountIn);
            } else if (call.swapType == SwapType.CurveCryptoBase) {
                uint i_idx = getCurveIndex(call.target, call.tokenIn);
                uint j_idx = getCurveIndex(call.target, call.tokenOut);
                amountOut = ICurveCryptoBase(call.target).get_dy(i_idx, j_idx, amountIn);
            } else if (call.swapType == SwapType.CurveCryptoUnderlying) {
                uint i_idx = getCurveIndexUnderlying(call.target, call.tokenIn);
                uint j_idx = getCurveIndexUnderlying(call.target, call.tokenOut);
                amountOut = ICurveCryptoUnderlying(call.target).get_dy_underlying(i_idx, j_idx, amountIn);
            }
            
            amountsOut[i+1] = amountOut;
            amountIn = amountOut;
        }
    }

    function getUniswapV2AmountOut(uint amountIn, uint reserveIn, uint reserveOut) public pure returns (uint amountOut) {
        uint amountInWithFee = amountIn * 997;
        uint numerator = amountInWithFee * reserveOut;
        uint denominator = (reserveIn * 1000) + amountInWithFee;
        amountOut = numerator / denominator;
    }

    function getCurveIndex(address pool, address coin) public view returns (uint) {
        for (uint i = 0; i < 20; i++) {
            address ret_coin = ICurveStableBase(pool).coins(i);
            if (ret_coin == coin) {
                return i;
            }
        }
        
        return 0;
    }

    function getCurveIndexUnderlying(address pool, address coin) public view returns (uint) {
        for (uint i = 0; i < 20; i++) {
            address ret_coin = ICurveStableUnderlying(pool).underlying_coins(i);
            if (ret_coin == coin) {
                return i;
            }
        }
        
        return 0;
    }
}