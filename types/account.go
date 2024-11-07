package types

import "github.com/gagliardetto/solana-go"

type AmmSwapBaseAccounts struct { // For SwapBaseIn and SwapBaseOut instructions
	Amm                         solana.PublicKey `idx:"1"`
	AmmAuthority                solana.PublicKey `idx:"2"`
	AmmOpenOrders               solana.PublicKey `idx:"3"`
	AmmTargetOrders             solana.PublicKey `idx:"4"`
	PoolCoinTokenAccount        solana.PublicKey `idx:"5"`
	PoolPcTokenAccount          solana.PublicKey `idx:"6"`
	SerumMarket                 solana.PublicKey `idx:"8"`
	SerumBids                   solana.PublicKey `idx:"9"`
	SerumAsks                   solana.PublicKey `idx:"10"`
	SerumEventQueue             solana.PublicKey `idx:"11"`
	SerumCoinVaultAccount       solana.PublicKey `idx:"12"`
	SerumPcVaultAccount         solana.PublicKey `idx:"13"`
	SerumVaultSigner            solana.PublicKey `idx:"14"`
	UserSourceTokenAccount      solana.PublicKey `idx:"15"`
	UserDestinationTokenAccount solana.PublicKey `idx:"16"`
	UserSourceOwner             solana.PublicKey `idx:"17"`
}

// ClmmSwapAccounts holds the accounts required for a swap operation
type ClmmSwapAccounts struct {
	// The user performing the swap
	Signer solana.PublicKey `idx:"0"`

	// The user token account for input token
	InputTokenAccount solana.PublicKey `idx:"1"`

	// The user token account for output token
	OutputTokenAccount solana.PublicKey `idx:"2"`

	// The vault token account for input token
	InputVault solana.PublicKey `idx:"3"`

	// The vault token account for output token
	OutputVault solana.PublicKey `idx:"4"`

	// SPL program for token transfers
	TokenProgram solana.PublicKey `idx:"5"`

	// The factory state to read protocol fees
	AmmConfig solana.PublicKey `idx:"6"`

	// The program account of the pool in which the swap will be performed
	PoolState solana.PublicKey `idx:"7"`

	// The tick_array account of current or next initialized
	TickArrayState solana.PublicKey `idx:"8"`

	// The program account for the oracle observation
	ObservationState solana.PublicKey `idx:"9"`
}

// ClmmCreatePoolAccounts holds the accounts required for creating a new pool
type ClmmCreatePoolAccounts struct {
	// Address paying to create the pool. Can be anyone
	PoolCreator solana.PublicKey `idx:"0"`

	// Which config the pool belongs to
	AmmConfig solana.PublicKey `idx:"1"`

	// Account to store the pool state
	PoolState solana.PublicKey `idx:"2"`

	// Token 0 mint (must be less than token_1 mint)
	TokenMint0 solana.PublicKey `idx:"3"`

	// Token 1 mint
	TokenMint1 solana.PublicKey `idx:"4"`

	// Token 0 vault for the pool
	TokenVault0 solana.PublicKey `idx:"5"`

	// Token 1 vault for the pool
	TokenVault1 solana.PublicKey `idx:"6"`

	// Account to store oracle observations
	ObservationState solana.PublicKey `idx:"7"`

	// Account to store tick array initialization status
	TickArrayBitmap solana.PublicKey `idx:"8"`

	// SPL token program for token 0
	TokenProgram0 solana.PublicKey `idx:"9"`

	// SPL token program for token 1
	TokenProgram1 solana.PublicKey `idx:"10"`

	// System program for creating new accounts
	SystemProgram solana.PublicKey `idx:"11"`

	// Rent sysvar
	Rent solana.PublicKey `idx:"12"`
}

// ClmmSwapSingleV2Accounts holds the accounts required for a single-sided swap operation
type ClmmSwapSingleV2Accounts struct {
	// The user performing the swap
	Payer solana.PublicKey `idx:"0"`

	// The factory state to read protocol fees
	AmmConfig solana.PublicKey `idx:"1"`

	// The program account of the pool in which the swap will be performed
	PoolState solana.PublicKey `idx:"2"`

	// The user token account for input token
	InputTokenAccount solana.PublicKey `idx:"3"`

	// The user token account for output token
	OutputTokenAccount solana.PublicKey `idx:"4"`

	// The vault token account for input token
	InputVault solana.PublicKey `idx:"5"`

	// The vault token account for output token
	OutputVault solana.PublicKey `idx:"6"`

	// The program account for the most recent oracle observation
	ObservationState solana.PublicKey `idx:"7"`

	// SPL program for token transfers
	TokenProgram solana.PublicKey `idx:"8"`

	// // SPL program 2022 for token transfers
	// TokenProgram2022 solana.PublicKey

	// // Memo program
	// MemoProgram solana.PublicKey

	// The ticker array account of current or next initialized
	TickArrayState solana.PublicKey `idx:"9"`

	// The mint of input token vault
	InputVaultMint solana.PublicKey `idx:"10"`

	// The mint of output token vault
	OutputVaultMint solana.PublicKey `idx:"11"`
}
