package instructions

import (
	"encoding/hex"
	"math/big"
	"reflect"
	"testing"

	"github.com/Umiiii/raydium-go/types"
)

const (
	SwapV2DataHex = "f8c69e91e17587c800743ba40b00000000000000000000000014b128ba6bb768000000000000000001"
	InitDataHex   = "e992d18ecf6840bc84058f90ef14b1123900000000000000094a246700000000"
	// 1052811292984159962500, 1730431497
)

func LittleEndianToBigInt(data []byte) *big.Int {
	reversed := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		reversed[len(data)-1-i] = data[i]
	}

	return new(big.Int).SetBytes(reversed)
}

func TestDecodeRaydiumData(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		want    interface{}
		wantErr bool
	}{
		{
			name: "decode SwapV2 data",
			data: SwapV2DataHex,
			want: types.ClmmSwapV2Args{
				//	Unused:               0,
				Amount:               50000000000,
				OtherAmountThreshold: 0,
				SqrtPriceLimitX64:    [16]byte{0, 20, 177, 40, 186, 107, 183, 104, 0, 0, 0, 0, 0, 0, 0, 0},
				IsBaseInput:          true,
			},
			wantErr: false,
		},
		{
			name: "decode Init data",
			data: InitDataHex,
			want: types.ClmmCreatePoolArgs{
				//		Unused:       0,
				SqrtPriceX64: [16]byte{132, 5, 143, 144, 239, 20, 177, 18, 57, 0, 0, 0, 0, 0, 0, 0},
				OpenTime:     1730431497,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := hex.DecodeString(tt.data)
			if err != nil && !tt.wantErr {
				t.Errorf("Failed to decode hex string: %v", err)
				return
			}

			switch v := tt.want.(type) {
			case types.ClmmSwapV2Args:
				got, err := DecodeRaydiumData[types.ClmmSwapV2Args](data)
				if (err != nil) != tt.wantErr {
					t.Errorf("DecodeRaydiumData() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && !reflect.DeepEqual(got, v) {
					t.Errorf("DecodeRaydiumData() = %v, want %v", got, v)
				}
			case types.ClmmCreatePoolArgs:
				got, err := DecodeRaydiumData[types.ClmmCreatePoolArgs](data)
				if (err != nil) != tt.wantErr {
					t.Errorf("DecodeRaydiumData() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && !reflect.DeepEqual(got, v) {
					t.Errorf("DecodeRaydiumData() = %v, want %v", got, v)
				}
			default:
				t.Errorf("Unsupported type: %T", v)
			}
		})
	}
}
