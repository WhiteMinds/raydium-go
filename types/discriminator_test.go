package types

import "testing"

func TestCheckDiscriminator(t *testing.T) {
	tests := []struct {
		name          string
		discriminator [8]byte
		eventName     string
	}{
		{
			name:          "ConfigChangeEvent",
			discriminator: ClmmConfigChangeEventDiscriminator,
			eventName:     ClmmEventConfigChange,
		},
		{
			name:          "CollectPersonalFeeEvent",
			discriminator: ClmmCollectPersonalFeeEventDiscriminator,
			eventName:     ClmmEventCollectPersonalFee,
		},
		{
			name:          "CollectProtocolFeeEvent",
			discriminator: ClmmCollectProtocolFeeEventDiscriminator,
			eventName:     ClmmEventCollectProtocolFee,
		},
		{
			name:          "CreatePersonalPositionEvent",
			discriminator: ClmmCreatePersonalPositionEventDiscriminator,
			eventName:     ClmmEventCreatePersonalPosition,
		},
		{
			name:          "DecreaseLiquidityEvent",
			discriminator: ClmmDecreaseLiquidityEventDiscriminator,
			eventName:     ClmmEventDecreaseLiquidity,
		},
		{
			name:          "IncreaseLiquidityEvent",
			discriminator: ClmmIncreaseLiquidityEventDiscriminator,
			eventName:     ClmmEventIncreaseLiquidity,
		},
		{
			name:          "LiquidityCalculateEvent",
			discriminator: ClmmLiquidityCalculateEventDiscriminator,
			eventName:     ClmmEventLiquidityCalculate,
		},
		{
			name:          "LiquidityChangeEvent",
			discriminator: ClmmLiquidityChangeEventDiscriminator,
			eventName:     ClmmEventLiquidityChange,
		},
		{
			name:          "SwapEvent",
			discriminator: ClmmSwapEventDiscriminator,
			eventName:     ClmmEventSwap,
		},
		{
			name:          "PoolCreatedEvent",
			discriminator: ClmmPoolCreatedEventDiscriminator,
			eventName:     ClmmEventPoolCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !CheckDiscriminator(tt.discriminator, tt.eventName) {
				t.Errorf("Want: %s, Got: %+v, Actual: %+v", tt.eventName, tt.discriminator, GetDiscriminator(tt.eventName))
				t.Errorf("CheckDiscriminator() failed for %s", tt.name)
			}
		})
	}
}

func TestGenerateDiscriminator(t *testing.T) {
	// [233 146 209 142 207 104 64 188]
	discriminator := GetInstructionDiscriminator("CreatePool")
	t.Logf("%+v", discriminator)
}
