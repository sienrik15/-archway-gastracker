package types

// Metadata type
type Reward struct {
	ContractAddress  string
	GasConsumed      uint64
	InflationRewards int64
	ContractRewards  int64

	Height int64
}
