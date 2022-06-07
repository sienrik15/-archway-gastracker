package types

// Metadata type
type Metadata struct {
	Contract                 string
	DeveloperAddress         string
	RewardAddress            string
	GasRebateToUser          bool
	CollectPremium           bool
	PremiumPercentageCharged uint64

	CreatedTime string
	Index       int
	Height      int64
}
