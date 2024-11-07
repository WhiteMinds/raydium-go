package logs

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/Umiiii/raydium-go/types"
	"github.com/near/borsh-go"
)

func GetRaydiumCLMMLogsFromBase64Log(logData string) (interface{}, error) {
	// "Program data: <base64>"
	// "<base64>"
	// Split the log string if it contains prefix "Program data: "
	parts := strings.Split(logData, ": ")
	var base64Data string
	if len(parts) == 2 {
		base64Data = parts[1]
	} else {
		base64Data = logData
	}
	// Decode base64 string
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, fmt.Errorf("invalid base64 data: %w", err)
	}

	// Check data length
	if len(data) < 8 {
		return nil, fmt.Errorf("data too short: expected at least 8 bytes, got %d", len(data))
	}

	// Get discriminator (first 8 bytes)
	var discriminator [8]byte
	copy(discriminator[:], data[:8])
	// Remaining data for borsh deserialization
	remainingData := data[8:]

	// Match discriminator and deserialize accordingly
	var result interface{}
	switch discriminator {

	case types.ClmmConfigChangeEventDiscriminator:
		event := new(types.ClmmConfigChangeEvent)
		err = borsh.Deserialize(event, remainingData)
		result = event
	case types.ClmmCollectPersonalFeeEventDiscriminator:
		event := new(types.ClmmCollectPersonalFeeEvent)
		err = borsh.Deserialize(event, remainingData)
		result = event
	case types.ClmmCollectProtocolFeeEventDiscriminator:
		event := new(types.ClmmCollectProtocolFeeEvent)
		err = borsh.Deserialize(event, remainingData)
		result = event
	case types.ClmmCreatePersonalPositionEventDiscriminator:
		event := new(types.ClmmCreatePersonalPositionEvent)
		err = borsh.Deserialize(event, remainingData)
		result = event
	case types.ClmmDecreaseLiquidityEventDiscriminator:
		event := new(types.ClmmDecreaseLiquidityEvent)
		err = borsh.Deserialize(event, remainingData)
		result = event
	case types.ClmmIncreaseLiquidityEventDiscriminator:
		event := new(types.ClmmIncreaseLiquidityEvent)
		err = borsh.Deserialize(event, remainingData)
		result = event
	case types.ClmmLiquidityCalculateEventDiscriminator:
		event := new(types.ClmmLiquidityCalculateEvent)
		err = borsh.Deserialize(event, remainingData)
		result = event
	case types.ClmmLiquidityChangeEventDiscriminator:
		event := new(types.ClmmLiquidityChangeEvent)
		err = borsh.Deserialize(event, remainingData)
		result = event
	case types.ClmmSwapEventDiscriminator:
		event := new(types.ClmmSwapEvent)
		err = borsh.Deserialize(event, remainingData)
		result = event
	case types.ClmmPoolCreatedEventDiscriminator:
		event := new(types.ClmmPoolCreatedEvent)
		err = borsh.Deserialize(event, remainingData)
		result = event
	default:
		return nil, fmt.Errorf("unknown discriminator: %v", discriminator)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to decode borsh data: %w", err)
	}
	fmt.Printf("result: %+v\n", result)

	return result, nil
}
