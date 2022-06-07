package database

import (
	"github.com/giansalex/archway-gastracker/types"
)

// SaveMetadata allows to save the metadata into the database.
func (db Db) SaveMetadataHistory(meta types.Metadata) error {
	stmt := `
INSERT INTO codes (contract, developer, reward_addr, gas_rebate, premium, premium_percent, creation_time, index, height) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := db.Sql.Exec(stmt, meta.Contract, meta.DeveloperAddress, meta.RewardAddress, meta.GasRebateToUser, meta.CollectPremium, meta.PremiumPercentageCharged, meta.CreatedTime, meta.Index, meta.Height)
	return err
}
