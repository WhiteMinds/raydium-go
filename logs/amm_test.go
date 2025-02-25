package logs

import (
	"reflect"
	"testing"

	"github.com/WhiteMinds/raydium-go/types"
	"github.com/gagliardetto/solana-go"
)

// TODO: Add more tests
func TestGetRaydiumAmmLogsFromRayLog(t *testing.T) {
	m, _ := solana.PublicKeyFromBase58("8TrtVbRU93hbP7TwV4Dtuj1cximBfa5AkceUCFR3ZvLm")
	type args struct {
		rayLog string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Init",
			args: args{
				rayLog: "Program: ray_log: AAAAAAAAAAAABglkAAAAAAAAAICWmAAAAAAAAAgBqSy8AACRHBZlEgAAAG7jUV5vrTnToQOjEWYfoQjU/LCcxQ2ys5IotRcUlt8O",
			},
			// 0 0 6 9 100 10000000 206900000000000 79005359249 8TrtVbRU93hbP7TwV4Dtuj1cximBfa5AkceUCFR3ZvLm
			want: &types.RaydiumAmmInitLog{
				LogType:      0,
				Time:         0,
				PcDecimals:   6,
				CoinDecimals: 9,
				PcLotSize:    100,
				CoinLotSize:  10000000,
				PcAmount:     206900000000000,
				CoinAmount:   79005359249,
				Market:       m,
			},
			wantErr: false,
		},
		{
			name: "swapBaseIn",
			args: args{
				rayLog: "Program log: ray_log: AwCaWNYEAAAAAQAAAAAAAAACAAAAAAAAAABUnu8EAAAAppH/A+MBAABwpiSjzhYAAJs74L85AAAA",
			},
			want: &types.RaydiumAmmSwapBaseInLog{
				LogType:    3,
				AmountIn:   20776000000,
				MinimumOut: 1,
				Direction:  2,
				UserSource: 21200000000,
				PoolCoin:   2074536284582,
				PoolPc:     25076756162160,
				OutAmount:  248032279451,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRaydiumAmmLogsFromRayLog(tt.args.rayLog)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRaydiumAmmLogsFromRayLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRaydiumAmmLogsFromRayLog() got = %v, want %v", got, tt.want)
			}
		})
	}
}
