package instructions

import (
	"bytes"
	"encoding/binary"

	"github.com/WhiteMinds/raydium-go/types"
)

// DecodeRaydiumData is a generic function that can decode any struct that follows
// the binary.Read format
func DecodeRaydiumData[T any](data []byte) (T, error) {
	buf := bytes.NewReader(data)
	var state T

	err := binary.Read(buf, binary.LittleEndian, &state)
	return state, err
}

func DecodeRaydiumCreatePoolData(data []byte) (types.ClmmCreatePoolArgs, error) {
	return DecodeRaydiumData[types.ClmmCreatePoolArgs](data)
}

func DecodeRaydiumSwapV2Data(data []byte) (types.ClmmSwapV2Args, error) {
	return DecodeRaydiumData[types.ClmmSwapV2Args](data)
}
