package logs

import (
	"reflect"
	"testing"

	"github.com/Umiiii/raydium-go/types"
	"github.com/gagliardetto/solana-go"
)

func TestGetRaydiumCLMMLogsFromBase64Log(t *testing.T) {
	type args struct {
		logData string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "PoolSwapEvent",
			args: args{
				logData: "Program data: QMbN6CYIceIpVWpkzMLdriyW1P//w/S8BkpgzdDUa8y8cOO1G2aLKJ3La/kb0g5TeqkcAi/tFCPUU4htJMatqAsQ5ckpVx1gbL4LD1VSRnp0XAp6NZN1P7v6oAziHi4VoKqM4SI1x9y6NSm57hF3tOKleoXF8CpJ47WO/IJCd9zvlNr3H4OzTQR/cGEDAAAAAAAAAAAAAABypM2QAAAAAAAAAAAAAAAAAQAUsSi6a7doAAAAAAAAAADxPLmLEEMAAAAAAAAAAAAAKLr//w==",
			},
			want: &types.ClmmSwapEvent{
				PoolState:       solana.MustPublicKeyFromBase58("3nMFwZXwY1s1M5s8vYAHqd4wGs4iSxXE4LRoUMMYqEgF"),
				Sender:          solana.MustPublicKeyFromBase58("Bcxv7jCiVCfdzgZghGQpZBDDNGxRBLgDthmu469jHrCX"),
				TokenAccount0:   solana.MustPublicKeyFromBase58("8KV7BNJSeAppE6R7XmvWev2qHceQfRPwFsDevZ5CtRX9"),
				TokenAccount1:   solana.MustPublicKeyFromBase58("DXsrGPzKLqN78Bmai8D5F6rdQRAYpMb1iRggxz6EE8CU"),
				Amount0:         14519664388,
				TransferFee0:    0,
				Amount1:         2429396082,
				TransferFee1:    0,
				ZeroForOne:      true,
				SqrtPriceX64:    7545618147973338112,
				LiquidityBefore: 0,
				LiquidityAfter:  73738342710513,
				Tick:            0,
			},
			wantErr: false,
		},
		{
			name: "PoolCreatedEvent",
			args: args{
				logData: "Program data: GV5LL3BjNT8Gm4hX/quBhPtof2NGGMA12sQ53BrrO1WYoPAAAAAAATxAC5Lho3BivvSqHd1opmKh+SHf03DO77chQFijFPN/eAB+wn62iylTTb7dKCrlYq7tKXMZwEACmNAc8FafeQaCkoQFj5DvFLESOQAAAAAAAAD6OwEAqQWJ6MMnGbxO3RtiO7H2m0/3VtsdjctcDkxxeDwJMkI1oW4983FaaYJCvWFpZqSDRDsJ2Qd4t+bz/5+g/cGqmw==",
			},
			want: &types.ClmmPoolCreatedEvent{
				TokenMint0:   solana.MustPublicKeyFromBase58("So11111111111111111111111111111111111111112"),
				TokenMint1:   solana.MustPublicKeyFromBase58("54C726ZerpSMYkmks6gGKFxsnbpziNkR13TqDdnzpump"),
				TickSpacing:  120,
				PoolState:    solana.MustPublicKeyFromBase58("9XpPAv79tumf7VDxqPGBjrz7DueZXqj2PuCpzHLxMCTs"),
				SqrtPriceX64: 1346880782715520388,
				Tick:         57,
				TokenVault0:  solana.MustPublicKeyFromBase58("11113aGsyuSXvG9adwywdr5mWVSdqAsTTQ6BvJQeToq"),
				TokenVault1:  solana.MustPublicKeyFromBase58("xpKFB83Ay9egnd9Uy1Uf4LixEjCLtPQWdzmKdottBRf"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRaydiumCLMMLogsFromBase64Log(tt.args.logData)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRaydiumCLMMLogsFromBase64Log() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRaydiumCLMMLogsFromBase64Log() = %v, want %+v", got, tt.want)
			}
		})
	}
}
