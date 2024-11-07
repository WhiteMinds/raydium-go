# Raydium-Go

Raydium-Go is a Go library for parsing the Raydium DEX events.

## Installation

go get github.com/Umiiii/raydium-go

## Supported Protocols

| Protocol | Repository | Status | Notes |
|----------|------------|--------|-------|
| OpenBook AMM (Raydium V4) | [raydium-amm](https://github.com/raydium-io/raydium-amm) | DONE | - |
| AMM Stable | [raydium-amm](https://github.com/raydium-io/raydium-amm) | NO PLAN | For stablecoins only |
| Concentrated Liquidity (CLMM/V3) | [raydium-clmm](https://github.com/raydium-io/raydium-clmm/) | WORKING | Solana Mainnet 2023-02-28 |
| Standard AMM (CP-Swap) | [raydium-cp-swap](https://github.com/raydium-io/raydium-cp-swap) | NO PLAN | Solana Mainnet 2024-05-15 |


## Usage 

### Parse Log Messages
```go
import (
    "github.com/Umiiii/raydium-go/logs"
)

...
    rayLog := "Program: ray_log: AAAAAAAAAAAABglkAAAAAAAAAICWmAAAAAAAAAgBqSy8AACRHBZlEgAAAG7jUV5vrTnToQOjEWYfoQjU/LCcxQ2ys5IotRcUlt8O"
    log, err := logs.GetRaydiumAmmLogsFromRayLog(rayLog)
    if err != nil {
        panic(err)
    }
    fmt.Println(log)
...
```

### Parse Instruction Data
```go
import (
    "github.com/Umiiii/raydium-go/instructions"
)
...
instructions.DecodeRaydiumCreatePoolData(outerInst.Data)
...
```
### Parse Parameters

```go
import (
    "github.com/Umiiii/raydium-go/parameters"
)
...
parameters.DecodeRaydiumCreatePoolParams(outerInst.Data)
...
```

### Mapping Account Address
``` go
import (
    "github.com/Umiiii/raydium-go/accounts"
    "github.com/Umiiii/raydium-go/types"    
)

....
acc := accounts.ParseAccountsIntoStruct[types.ClmmCreatePoolAccounts](outerInst.Accounts)
....

```


### Get Raydium Swap Details
#### CLMM

Actually you can infer the swap details from the log message ONLY, but it is quite tricky.

So a log function is provided to group the log message first.
```go
import (
    "github.com/Umiiii/raydium-go/logs"
)

func main() {
    logs := logs.ParseProgramLogs(parsedTransaction.Meta.LogMessages)
 
    for idx, singleLog := range allLogs {
		if singleLog.Address == CLMMAddress {
           	var swapRawData *types.ClmmSwapEvent
			var ok bool
			if swapRawData, ok = (logsData).(*types.ClmmSwapEvent); !ok {
				return swapEvent, errors.New("failed to parse clmm logs")
			}
            poolAddress := swapRawData.PoolAddress
            tokenMint0 := swapRawData.TokenMint0
            tokenMint1 := swapRawData.TokenMint1
            amount0 := swapRawData.Amount0
            amount1 := swapRawData.Amount1
            if swapRawData.ZeroForOne {
                // swap from token0 to token1
            } else {
                // swap from token1 to token0
            }
        }
    }
}
```


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
