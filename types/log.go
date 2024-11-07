package types

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"

	"github.com/near/borsh-go"
)

const (
	RaydiumAmmLogTypeInitialize  = 0
	RaydiumAmmLogTypeDeposit     = 1
	RaydiumAmmLogTypeWithdraw    = 2
	RaydiumAmmLogTypeSwapBaseIn  = 3
	RaydiumAmmLogTypeSwapBaseOut = 4
)

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

func GetRaydiumAmmLogsFromRayLog(rayLog string) (interface{}, error) {
	// Extract the base64 encoded data from various possible log formats:
	// 1. "Program log: ray_log: <base64>"
	// 2. "ray_log: <base64>"
	// 3. "<base64>"
	parts := strings.Split(rayLog, ": ")
	var encodedData string

	switch len(parts) {
	case 1:
		encodedData = parts[0]
	case 2:
		encodedData = parts[1]
	case 3:
		encodedData = parts[2]
	default:
		return nil, fmt.Errorf("invalid rayLog format")
	}

	// Clean up any potential whitespace
	encodedData = strings.TrimSpace(encodedData)

	// Convert base64 string to bytes
	data, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 string: %w", err)
	}

	// First byte is the discriminator
	if len(data) == 0 {
		return nil, fmt.Errorf("empty data")
	}
	discriminator := data[0]

	var result interface{}
	switch discriminator {
	case RaydiumAmmLogTypeInitialize:
		logs := new(RaydiumAmmInitLog)
		err = borsh.Deserialize(logs, data)
		result = logs
	case RaydiumAmmLogTypeDeposit:
		logs := new(RaydiumAmmDepositLog)
		err = borsh.Deserialize(logs, data)
		result = logs
	case RaydiumAmmLogTypeWithdraw:
		logs := new(RaydiumAmmWithdrawLog)
		err = borsh.Deserialize(logs, data)
		result = logs
	case RaydiumAmmLogTypeSwapBaseIn:
		logs := new(RaydiumAmmSwapBaseInLog)
		err = borsh.Deserialize(logs, data)
		result = logs
	case 4:
		logs := new(RaydiumAmmSwapBaseOutLog)
		err = borsh.Deserialize(logs, data)
		result = logs
	default:
		return nil, fmt.Errorf("unknown discriminator: %d", discriminator)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to decode borsh data: %w", err)
	}

	return result, nil
}
