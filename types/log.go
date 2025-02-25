package types

import (
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

const (
	RaydiumAmmLogTypeInitialize  = 0
	RaydiumAmmLogTypeDeposit     = 1
	RaydiumAmmLogTypeWithdraw    = 2
	RaydiumAmmLogTypeSwapBaseIn  = 3
	RaydiumAmmLogTypeSwapBaseOut = 4
)

// Amm Types
// https://github.com/raydium-io/raydium-amm/blob/master/program/src/log.rs

// RaydiumAmmInitLog represents Initialize/Initialize2 logs
type RaydiumAmmInitLog struct {
	LogType      uint8            `borsh:"logType"`
	Time         uint64           `borsh:"time"`
	PcDecimals   uint8            `borsh:"pcDecimals"`
	CoinDecimals uint8            `borsh:"coinDecimals"`
	PcLotSize    uint64           `borsh:"pcLotSize"`
	CoinLotSize  uint64           `borsh:"coinLotSize"`
	PcAmount     uint64           `borsh:"pcAmount"`
	CoinAmount   uint64           `borsh:"coinAmount"`
	Market       solana.PublicKey `borsh:"market"`
}

// RaydiumAmmDepositLog represents Deposit logs
type RaydiumAmmDepositLog struct {
	LogType    uint8  `borsh:"logType"`
	MaxCoin    uint64 `borsh:"maxCoin"`
	MaxPc      uint64 `borsh:"maxPc"`
	Base       uint64 `borsh:"base"`
	PoolCoin   uint64 `borsh:"poolCoin"`
	PoolPc     uint64 `borsh:"poolPc"`
	PoolLp     uint64 `borsh:"poolLp"`
	CalcPnlX   uint64 `borsh:"calcPnlX"`
	CalcPnlY   uint64 `borsh:"calcPnlY"`
	DeductCoin uint64 `borsh:"deductCoin"`
	DeductPc   uint64 `borsh:"deductPc"`
	MintLp     uint64 `borsh:"mintLp"`
}

// RaydiumAmmWithdrawLog represents Withdraw logs
type RaydiumAmmWithdrawLog struct {
	LogType    uint8  `borsh:"logType"`
	WithdrawLp uint64 `borsh:"withdrawLp"`
	UserLp     uint64 `borsh:"userLp"`
	PoolCoin   uint64 `borsh:"poolCoin"`
	PoolPc     uint64 `borsh:"poolPc"`
	PoolLp     uint64 `borsh:"poolLp"`
	CalcPnlX   uint64 `borsh:"calcPnlX"`
	CalcPnlY   uint64 `borsh:"calcPnlY"`
	OutCoin    uint64 `borsh:"outCoin"`
	OutPc      uint64 `borsh:"outPc"`
}

// RaydiumAmmSwapBaseInLog represents SwapBaseIn logs
type RaydiumAmmSwapBaseInLog struct {
	LogType    uint8  `borsh:"logType"`
	AmountIn   uint64 `borsh:"amountIn"`
	MinimumOut uint64 `borsh:"minimumOut"`
	Direction  uint64 `borsh:"direction"`
	UserSource uint64 `borsh:"userSource"`
	PoolCoin   uint64 `borsh:"poolCoin"`
	PoolPc     uint64 `borsh:"poolPc"`
	OutAmount  uint64 `borsh:"outAmount"`
}

// RaydiumAmmSwapBaseOutLog represents SwapBaseOut logs
type RaydiumAmmSwapBaseOutLog struct {
	LogType    uint8  `borsh:"logType"`
	MaxIn      uint64 `borsh:"maxIn"`
	AmountOut  uint64 `borsh:"amountOut"`
	Direction  uint64 `borsh:"direction"`
	UserSource uint64 `borsh:"userSource"`
	PoolCoin   uint64 `borsh:"poolCoin"`
	PoolPc     uint64 `borsh:"poolPc"`
	DirectIn   uint64 `borsh:"directIn"`
}

// CLMM Types
// https://github.com/raydium-io/raydium-clmm/blob/master/programs/amm/src/states/pool.rs
// Event structs
type ClmmConfigChangeEvent struct {
	Owner       solana.PublicKey `borsh:"owner"`
	FeeRate     uint32           `borsh:"fee_rate"`
	ProtocolFee uint32           `borsh:"protocol_fee"`
	FundFee     uint32           `borsh:"fund_fee"`
}

type ClmmCollectPersonalFeeEvent struct {
	OwnerTokenVault0 solana.PublicKey `borsh:"owner_token_vault_0"`
	OwnerTokenVault1 solana.PublicKey `borsh:"owner_token_vault_1"`
	Amount0          uint64           `borsh:"amount_0"`
	Amount1          uint64           `borsh:"amount_1"`
}

type ClmmCollectProtocolFeeEvent struct {
	ProtocolFeeTokenVault0 solana.PublicKey `borsh:"protocol_fee_token_vault_0"`
	ProtocolFeeTokenVault1 solana.PublicKey `borsh:"protocol_fee_token_vault_1"`
	Amount0                uint64           `borsh:"amount_0"`
	Amount1                uint64           `borsh:"amount_1"`
}

type ClmmCreatePersonalPositionEvent struct {
	PoolState        solana.PublicKey `borsh:"pool_state"`
	Minter           solana.PublicKey `borsh:"minter"`
	PositionNftMint  solana.PublicKey `borsh:"position_nft_mint"`
	TickLowerIndex   int32            `borsh:"tick_lower_index"`
	TickUpperIndex   int32            `borsh:"tick_upper_index"`
	Liquidity        uint64           `borsh:"liquidity"`
	FeeGrowthInside0 uint64           `borsh:"fee_growth_inside_0"`
	FeeGrowthInside1 uint64           `borsh:"fee_growth_inside_1"`
	TokenFees0       uint64           `borsh:"token_fees_0"`
	TokenFees1       uint64           `borsh:"token_fees_1"`
}

type ClmmDecreaseLiquidityEvent struct {
	PoolState       solana.PublicKey `borsh:"pool_state"`
	PositionNftMint solana.PublicKey `borsh:"position_nft_mint"`
	Liquidity       uint64           `borsh:"liquidity"`
	Amount0         uint64           `borsh:"amount_0"`
	Amount1         uint64           `borsh:"amount_1"`
}

type ClmmIncreaseLiquidityEvent struct {
	PoolState       solana.PublicKey `borsh:"pool_state"`
	PositionNftMint solana.PublicKey `borsh:"position_nft_mint"`
	Liquidity       uint64           `borsh:"liquidity"`
	Amount0         uint64           `borsh:"amount_0"`
	Amount1         uint64           `borsh:"amount_1"`
}

type ClmmLiquidityCalculateEvent struct {
	PoolState        solana.PublicKey `borsh:"pool_state"`
	CurrentTick      int32            `borsh:"current_tick"`
	CurrentLiquidity uint64           `borsh:"current_liquidity"`
	Amount0          uint64           `borsh:"amount_0"`
	Amount1          uint64           `borsh:"amount_1"`
	StartTime        uint64           `borsh:"start_time"`
	Status           uint8            `borsh:"status"`
}

type ClmmLiquidityChangeEvent struct {
	PoolState       solana.PublicKey `borsh:"pool_state"`
	TickIndex       int32            `borsh:"tick_index"`
	CurrentTick     int32            `borsh:"current_tick"`
	LiquidityBefore uint64           `borsh:"liquidity_before"`
	LiquidityAfter  uint64           `borsh:"liquidity_after"`
}

// src/states/pool.rs
type ClmmSwapEvent struct {
	PoolState      solana.PublicKey `borsh:"pool_state"`
	Sender         solana.PublicKey `borsh:"sender"`
	TokenAccount0  solana.PublicKey `borsh:"token_account_0"`
	TokenAccount1  solana.PublicKey `borsh:"token_account_1"`
	Amount0        uint64           `borsh:"amount_0"`
	TransferFee0   uint64           `borsh:"transfer_fee_0"`
	Amount1        uint64           `borsh:"amount_1"`
	TransferFee1   uint64           `borsh:"transfer_fee_1"`
	ZeroForOne     bool             `borsh:"zero_for_one"`
	SqrtPriceX64   bin.Uint128      `borsh:"sqrt_price_x64"`
	LiquidityAfter bin.Uint128      `borsh:"liquidity_after"`
	Tick           int32            `borsh:"tick"`
}

type ClmmPoolCreatedEvent struct {
	TokenMint0   solana.PublicKey `borsh:"token_mint_0"`
	TokenMint1   solana.PublicKey `borsh:"token_mint_1"`
	TickSpacing  uint16           `borsh:"tick_spacing"`
	PoolState    solana.PublicKey `borsh:"pool_state"`
	SqrtPriceX64 uint64           `borsh:"sqrt_price_x64"`
	Tick         int32            `borsh:"tick"`
	TokenVault0  solana.PublicKey `borsh:"token_vault_0"`
	TokenVault1  solana.PublicKey `borsh:"token_vault_1"`
}
