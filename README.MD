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
TBD
### Parse Parameters

```go

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