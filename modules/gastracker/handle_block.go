package gastracker

import (
	gstTypes "github.com/archway-network/archway/x/gastracker/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
	"github.com/giansalex/archway-gastracker/types"
	"github.com/gogo/protobuf/proto"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

type rewardsSummary struct {
	gas       uint64
	inflation sdk.Dec
	rewards   sdk.Dec
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleBlock(block *tmctypes.ResultBlock, results *tmctypes.ResultBlockResults, txs []*juno.Tx, vals *tmctypes.ResultValidators) error {
	rewardEventType := proto.MessageName(&gstTypes.ContractRewardCalculationEvent{})

	rewardEvents := juno.FindEventsByType(results.BeginBlockEvents, rewardEventType)
	contracts := make(map[string]rewardsSummary)
	for _, e := range rewardEvents {
		protomsg, err := sdk.ParseTypedEvent(e)
		if err != nil {
			return err
		}

		event := protomsg.(*gstTypes.ContractRewardCalculationEvent)

		// TODO: multiple denom?
		reward := rewardsSummary{
			gas:       event.GetGasConsumed(),
			inflation: event.GetInflationRewards().Amount,
			rewards:   event.GetContractRewards()[0].Amount,
		}
		contractAddress := event.GetContractAddress()
		contract, ok := contracts[contractAddress]
		if !ok {
			contracts[contractAddress] = reward
		} else {
			contract.inflation = contract.inflation.Add(reward.inflation)
			contract.rewards = contract.rewards.Add(reward.rewards)
			contract.gas += reward.gas

			contracts[contractAddress] = contract
		}
	}

	// TODO: verify RewardDistributionEvent
	for contractAddress, v := range contracts {
		reward := types.Reward{
			ContractAddress:  contractAddress,
			GasConsumed:      v.gas,
			InflationRewards: v.inflation.RoundInt64(),
			ContractRewards:  v.rewards.RoundInt64(),
			Height:           block.Block.Height,
		}

		if err := m.db.SaveBlockContractRewards(reward); err != nil {
			return err
		}

		if err := m.db.AddContractRewards(contractAddress, reward.ContractRewards); err != nil {
			return err
		}
	}

	return nil
}
