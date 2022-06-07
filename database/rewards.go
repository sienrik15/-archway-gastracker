package database

import (
	"github.com/giansalex/archway-gastracker/types"
)

// SaveBlockContractRewards allows to save the metadata into the database.
func (db Db) SaveBlockContractRewards(meta types.Reward) error {
	stmt := `
INSERT INTO block_rewards (contract, gas_consumed, inflation, rewards, height) 
VALUES ($1, $2, $3, $4, $5)`
	// TODO fix gas_consumed field, wrong big value from archway.
	_, err := db.Sql.Exec(stmt, meta.ContractAddress, uint64(0), meta.InflationRewards, meta.ContractRewards, meta.Height)
	return err
}

// ExistContractRewards verify contract in total rewards
func (db Db) ExistContractRewards(contract string) (bool, error) {
	stmt := `SELECT count(1) FROM total_rewads WHERE contract = $1`

	var count int
	err := db.Sqlx.Get(&count, stmt, contract)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// AddContractRewards adds the rewards to the total rewards.
func (db Db) AddContractRewards(contract string, rewards int64) error {
	if exist, err := db.ExistContractRewards(contract); err != nil {
		return err
	} else if !exist {
		stmt := `
INSERT INTO total_rewads (contract, rewards) 
VALUES ($1, $2)`
		_, err := db.Sql.Exec(stmt, contract, rewards)
		return err
	} else {
		stmt := `
UPDATE total_rewads SET rewards=rewards+$2 WHERE contract = $1`
		_, err := db.Sql.Exec(stmt, contract, rewards)
		return err
	}
}
