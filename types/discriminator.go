package types

import "crypto/sha256"

// The discriminators are typically derived using Anchor's #[event] macro,
// which automatically generates an 8-byte discriminator for each event.
// The discriminator is calculated using a hash of the event's name.

func GetInstructionDiscriminator(name string) [8]byte {
	// Create the namespace string "instruction:{name}"
	namespace := "global:" + name

	// Calculate SHA256
	hash := sha256.Sum256([]byte(namespace))

	// Take first 8 bytes
	var result [8]byte
	copy(result[:], hash[:8])

	return result
}

func GetDiscriminator(name string) [8]byte {
	// Create the namespace string "event:{name}"
	namespace := "event:" + name

	// Calculate SHA256
	hash := sha256.Sum256([]byte(namespace))

	// Take first 8 bytes
	var result [8]byte
	copy(result[:], hash[:8])

	return result
}

func CheckDiscriminator(discriminator [8]byte, name string) bool {
	return discriminator == GetDiscriminator(name)
}

var (
	ClmmConfigChangeEventDiscriminator           = [8]byte{247, 189, 7, 119, 106, 112, 95, 151}
	ClmmCollectPersonalFeeEventDiscriminator     = [8]byte{166, 174, 105, 192, 81, 161, 83, 105}
	ClmmCollectProtocolFeeEventDiscriminator     = [8]byte{206, 87, 17, 79, 45, 41, 213, 61}
	ClmmCreatePersonalPositionEventDiscriminator = [8]byte{100, 30, 87, 249, 196, 223, 154, 206}
	ClmmDecreaseLiquidityEventDiscriminator      = [8]byte{58, 222, 86, 58, 68, 50, 85, 56}
	ClmmIncreaseLiquidityEventDiscriminator      = [8]byte{49, 79, 105, 212, 32, 34, 30, 84}
	ClmmLiquidityCalculateEventDiscriminator     = [8]byte{237, 112, 148, 230, 57, 84, 180, 162}
	ClmmLiquidityChangeEventDiscriminator        = [8]byte{126, 240, 175, 206, 158, 88, 153, 107}
	ClmmSwapEventDiscriminator                   = [8]byte{64, 198, 205, 232, 38, 8, 113, 226}
	ClmmPoolCreatedEventDiscriminator            = [8]byte{25, 94, 75, 47, 112, 99, 53, 63}

	ClmmEventConfigChange           = "ConfigChangeEvent"
	ClmmEventCollectPersonalFee     = "CollectPersonalFeeEvent"
	ClmmEventCollectProtocolFee     = "CollectProtocolFeeEvent"
	ClmmEventCreatePersonalPosition = "CreatePersonalPositionEvent"
	ClmmEventDecreaseLiquidity      = "DecreaseLiquidityEvent"
	ClmmEventIncreaseLiquidity      = "IncreaseLiquidityEvent"
	ClmmEventLiquidityCalculate     = "LiquidityCalculateEvent"
	ClmmEventLiquidityChange        = "LiquidityChangeEvent"
	ClmmEventSwap                   = "SwapEvent"
	ClmmEventPoolCreated            = "PoolCreatedEvent"
)
