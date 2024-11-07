package logs

import (
	"strings"
	"testing"
)

func TestParseProgramLogs(t *testing.T) {
	testData := `Program ComputeBudget111111111111111111111111111111 invoke [1]
Program ComputeBudget111111111111111111111111111111 success
Program ComputeBudget111111111111111111111111111111 invoke [1]
Program ComputeBudget111111111111111111111111111111 success
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]
Program log: CreateIdempotent
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 11838 of 446700 compute units
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success
Program 6m2CDdhRgxpH4WjvdzxAYbGxwdGUz5MziiL5jek2kBma invoke [1]
Program log: Instruction: Swap2
Program log: 2qEHjDLDLbuBgRYvsxhc5D6uDWAivNFZGan56P1tpump
Program log: 9PR7nCP9DpcUotnDPVLUBUZKu5WAYkwrCUx9wDnSpump
Program log: before_source_balance: 313351340557, before_destination_balance: 0, amount_in: 153351340557, expect_amount_out: 724709663869, min_return: 717462567231
Program log: Dex::RaydiumSwap amount_in: 141083233312, offset: 0
Program 675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8 invoke [2]
Program log: ray_log: AyBMN9kgAAAAAQAAAAAAAAABAAAAAAAAAA0CMvVIAAAAtMsa/gIMAAAUxMNZ/g4AAJgIGQcaAAAA
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4645 of 374218 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4736 of 366592 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program 675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8 consumed 32275 of 393075 compute units
Program 675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8 success
Program data: QMbN6CYIceIEIEw32SAAAACYCBkHGgAAAA==
Program log: SwapEvent { dex: RaydiumSwap, amount_in: 141083233312, amount_out: 111788230808 }
Program log: G4Df5nFcecNmoKhkFTxm2EmsiAzXQ1PQUmJtvDjYmiV2
Program log: 2rikd7tzPbmowhUJzPNVtX7fuUGcnBa8jqJnx6HbtHeE
Program log: Dex::MeteoraDlmm amount_in: 9201080433, offset: 19
Program LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo invoke [2]
Program log: Instruction: Swap
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: TransferChecked
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6147 of 301290 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: TransferChecked
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6238 of 291710 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo invoke [3]
Program LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo consumed 2136 of 282041 compute units
Program LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo success
Program LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo consumed 62635 of 341007 compute units
Program LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo success
Program data: QMbN6CYIceINcVhtJAIAAABBffGzAQAAAA==
Program log: SwapEvent { dex: MeteoraDlmm, amount_in: 9201080433, amount_out: 7313915201 }
Program log: G4Df5nFcecNmoKhkFTxm2EmsiAzXQ1PQUmJtvDjYmiV2
Program log: 2rikd7tzPbmowhUJzPNVtX7fuUGcnBa8jqJnx6HbtHeE
Program log: Dex::WhirlpoolV2 amount_in: 3067026812, offset: 37
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc invoke [2]
Program log: Instruction: SwapV2
Program log: fee_growth: 382812256871403
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: TransferChecked
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6147 of 215901 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: TransferChecked
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6238 of 205777 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc consumed 60979 of 257710 compute units
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc success
Program data: QMbN6CYIceIQfB3PtgAAAABUDkGRAAAAAA==
Program log: SwapEvent { dex: WhirlpoolV2, amount_in: 3067026812, amount_out: 2436959828 }
Program log: G4Df5nFcecNmoKhkFTxm2EmsiAzXQ1PQUmJtvDjYmiV2
Program log: 2rikd7tzPbmowhUJzPNVtX7fuUGcnBa8jqJnx6HbtHeE
Program log: Dex::RaydiumClmmSwapV2 amount_in: 121539105837, offset: 53
Program CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK invoke [2]
Program log: Instruction: SwapV2
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: TransferChecked
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6238 of 47238 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: TransferChecked
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6147 of 37130 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program data: QMbN6CYIceLmfZ0ogtTVX+YeGXpEByU8O6bYcUaIU9bjQw65PXpVo/TnnWLPabh8F9TFvyB2wHgkKjclvP1voW9r2pZM385WG5hbGEk06WfdwETeTk2VW+QfSkuCB1qa/6jIw+AHLocSE66saVg1lbUsqUNU7+lrFNghb5j5NtLkHIHMpRu7ZS2US0wcAAAAAAAAAAAAAABfhDpRpwAAAAAAAAAAAAAAAV1QACHi8QBvAgAAAAAAAACPsSoWMg4AAAAAAAAAAAAAfEUAAA==
Program CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK consumed 151833 of 176295 compute units
Program CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK success
Program data: QMbN6CYIceILLZRLTBwAAABfhDpRpwAAAA==
Program log: SwapEvent { dex: RaydiumClmmSwapV2, amount_in: 121539105837, amount_out: 718622327903 }
Program log: 2rikd7tzPbmowhUJzPNVtX7fuUGcnBa8jqJnx6HbtHeE
Program log: 2DZkFYNAZBZXCVvyZ9d3BYoEzX593sZd7teiuYWQig1r
Program log: after_source_balance: 160000000000, after_destination_balance: 718622327903, source_token_change: 153351340557, destination_token_change: 718622327903
Program 6m2CDdhRgxpH4WjvdzxAYbGxwdGUz5MziiL5jek2kBma consumed 420946 of 434862 compute units
Program return: 6m2CDdhRgxpH4WjvdzxAYbGxwdGUz5MziiL5jek2kBma X4Q6UacAAAA=
Program 6m2CDdhRgxpH4WjvdzxAYbGxwdGUz5MziiL5jek2kBma success`

	logs := ParseProgramLogs(strings.Split(testData, "\n"))
	for _, program := range logs {
		t.Logf("%+v", program)
	}
}
