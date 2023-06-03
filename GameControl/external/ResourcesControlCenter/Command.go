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
func (*commandRequestWithResponce).TestResponce(key uuid.UUID) bool
func (*commandRequestWithResponce).WriteRequest(key uuid.UUID) error
*/

// ------------------------- TestRequest -------------------------

type Command_TestRequest struct {
	Key uuid.UUID `json:"key"`
}

type Command_TestRequest_Return struct {
	Exist bool `json:"exist"`
}

func (c *Command_TestRequest) Run(env *GlobalAPI.GlobalAPI) external.Return {
	resp := env.Resources.Command.TestRequest(c.Key)
	return Command_TestRequest_Return{Exist: resp}
}

// ------------------------- TestResponce -------------------------

type Command_TestResponce struct {
	Key uuid.UUID `json:"key"`
}

type Command_TestResponce_Return struct {
	Exist bool `json:"exist"`
}

func (c *Command_TestResponce) Run(env *GlobalAPI.GlobalAPI) external.Return {
	resp := env.Resources.Command.TestResponce(c.Key)
	return Command_TestResponce_Return{Exist: resp}
}

// ------------------------- WriteRequest -------------------------

type Command_WriteRequest struct {
	Key uuid.UUID `json:"key"`
}

type Command_WriteRequest_Return struct {
	Error error `json:"error"`
}

func (c *Command_WriteRequest) Run(env *GlobalAPI.GlobalAPI) external.Return {
	resp := env.Resources.Command.WriteRequest(c.Key)
	return Command_WriteRequest_Return{Error: resp}
}

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
