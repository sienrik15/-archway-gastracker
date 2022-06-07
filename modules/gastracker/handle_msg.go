package gastracker

import (
	gstTypes "github.com/archway-network/archway/x/gastracker/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
	"github.com/giansalex/archway-gastracker/types"
)

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *gstTypes.MsgSetContractMetadata:
		return m.handleMsgSetContractMetadata(tx, index, cosmosMsg)
	}

	return nil
}

func (m *Module) handleMsgSetContractMetadata(tx *juno.Tx, index int, msg *gstTypes.MsgSetContractMetadata) error {
	meta := types.Metadata{
		Contract:                 msg.ContractAddress,
		DeveloperAddress:         msg.Metadata.DeveloperAddress,
		RewardAddress:            msg.Metadata.RewardAddress,
		GasRebateToUser:          msg.Metadata.GasRebateToUser,
		CollectPremium:           msg.Metadata.CollectPremium,
		PremiumPercentageCharged: msg.Metadata.PremiumPercentageCharged,
		CreatedTime:              tx.Timestamp,
		Index:                    index,
		Height:                   tx.Height,
	}

	return m.db.SaveMetadataHistory(meta)
}
