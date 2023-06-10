package external_resources

import (
	"phoenixbuilder/GameControl/GlobalAPI"
	"phoenixbuilder/GameControl/external"
	"phoenixbuilder/minecraft/protocol/packet"

	"github.com/google/uuid"
)

/*
Function List

func (*commandRequestWithResponce).AwaitResponce(key uuid.UUID)
func (*commandRequestWithResponce).LoadResponceAndDelete(key uuid.UUID) (packet.CommandOutput, error)
func (*commandRequestWithResponce).TestRequest(key uuid.UUID) bool
*/

// ------------------------- TestRequest -------------------------

// ------------------------- TestResponce -------------------------

// ------------------------- WriteRequest -------------------------

// ------------------------- LoadResponceAndDelete -------------------------

type Command_LoadResponceAndDelete struct {
	Key uuid.UUID `json:"key"`
}

type Command_LoadResponceAndDelete_Return struct {
	CommandResponce packet.CommandOutput `json:"command_responce"`
	Error           error                `json:"error"`
}

func (c *Command_LoadResponceAndDelete) Run(env *GlobalAPI.GlobalAPI) external.Return {
	resp, err := env.Resources.Command.LoadResponceAndDelete(c.Key)
	return Command_LoadResponceAndDelete_Return{
		CommandResponce: resp,
		Error:           err,
	}
}

// ------------------------- AwaitResponce -------------------------

type Command_AwaitResponce struct {
	Key uuid.UUID `json:"key"`
}

type Command_AwaitResponce_Return struct{}

func (c *Command_AwaitResponce) Run(env *GlobalAPI.GlobalAPI) external.Return {
	env.Resources.Command.LoadResponceAndDelete(c.Key)
	return Command_AwaitResponce_Return{}
}

// ------------------------- END -------------------------
