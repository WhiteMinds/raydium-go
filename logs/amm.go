package logs

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/Umiiii/raydium-go/types"
	"github.com/near/borsh-go"
)

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
	case types.RaydiumAmmLogTypeInitialize:
		logs := new(types.RaydiumAmmInitLog)
		err = borsh.Deserialize(logs, data)
		result = logs
	case types.RaydiumAmmLogTypeDeposit:
		logs := new(types.RaydiumAmmDepositLog)
		err = borsh.Deserialize(logs, data)
		result = logs
	case types.RaydiumAmmLogTypeWithdraw:
		logs := new(types.RaydiumAmmWithdrawLog)
		err = borsh.Deserialize(logs, data)
		result = logs
	case types.RaydiumAmmLogTypeSwapBaseIn:
		logs := new(types.RaydiumAmmSwapBaseInLog)
		err = borsh.Deserialize(logs, data)
		result = logs
	case types.RaydiumAmmLogTypeSwapBaseOut:
		logs := new(types.RaydiumAmmSwapBaseOutLog)
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
