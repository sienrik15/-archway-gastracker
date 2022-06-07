package gastracker

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	gstTypes "github.com/archway-network/archway/x/gastracker/types"
	"github.com/forbole/juno/v3/modules/messages"
)

func GasTrackerMessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {
	case *gstTypes.MsgSetContractMetadata:
		return []string{msg.Sender}, nil
	}

	return nil, messages.MessageNotSupported(cosmosMsg)
}
