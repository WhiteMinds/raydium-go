package types

import "math/big"

func GetBigIntFromUint128(bytes [16]byte) *big.Int {
	return GetBigIntFromLittleEndianBytes(bytes[:])
}

func GetBigIntFromLittleEndianBytes(bytes []byte) *big.Int {
	reversed := make([]byte, len(bytes))
	for i := 0; i < len(bytes); i++ {
		reversed[len(bytes)-1-i] = bytes[i]
	}

	return new(big.Int).SetBytes(reversed)
}

// TODO: find the discriminator for this instruction

type ClmmCreatePoolArgs struct {
	Discriminator [8]byte
	SqrtPriceX64  [16]byte
	OpenTime      uint64
}

type ClmmSwapV2Args struct {
	Discriminator        [8]byte
	Amount               uint64
	OtherAmountThreshold uint64
	SqrtPriceLimitX64    [16]byte
	IsBaseInput          bool
}

// 5824288018642792
// 7545618147973338112
// 25015126562555429946130432

// f8c69e91e17587c800743ba40b00000000000000000000000014b128ba6bb768000000000000000001
// f8c69e91e17587c800
// 743ba40b - 8
// 00000000 - 8
// 0000000000000000 - 16
// 14b128ba6bb76800 - 16
// 00000000 - 8
// 00000001 - 8

// amount: ba43b7400 - 50000000000
// otherAmountThreshold: 0
// sqrtPriceLimitX64: 68B76BBA28B11400
// isBaseInput: 1
