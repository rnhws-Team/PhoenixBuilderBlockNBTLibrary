package external_resources_command

import (
	"phoenixbuilder/GameControl/GlobalAPI"
	"phoenixbuilder/GameControl/external"
	"phoenixbuilder/encoding"
	"phoenixbuilder/minecraft/protocol/packet"

	"github.com/google/uuid"
)

// func (*commandRequestWithResponce).LoadResponceAndDelete(key uuid.UUID) (packet.CommandOutput, error)

type LoadResponceAndDelete struct {
	Key uuid.UUID `json:"key"`
}

type LoadResponceAndDelete_Return struct {
	CommandResponce packet.CommandOutput `json:"command_responce"`
	Error           string               `json:"error"` // error
}

func (c *LoadResponceAndDelete) Marshal(io encoding.IO) {
	external.TestError(io.UUID(&c.Key))
}

func (c *LoadResponceAndDelete_Return) Marshal(io encoding.IO) {
	{
		external.TestError(io.Uint32(&c.CommandResponce.CommandOrigin.Origin))
		external.TestError(io.UUID(&c.CommandResponce.CommandOrigin.UUID))
		external.TestError(io.String(&c.CommandResponce.CommandOrigin.RequestID))
		external.TestError(io.Int64(&c.CommandResponce.CommandOrigin.PlayerUniqueID))
		// CommandOrigin
		external.TestError(io.Uint8(&c.CommandResponce.OutputType))
		// OutputType
		external.TestError(io.Uint32(&c.CommandResponce.SuccessCount))
		// SuccessCount
		external.TestError(io.CommandOutputMessageSlice(&c.CommandResponce.OutputMessages))
		// OutputMessages
		external.TestError(io.String(&c.CommandResponce.DataSet))
		// DataSet
	}
	// CommandResponce
	external.TestError(io.String(&c.Error))
	// Error
}

func (c *LoadResponceAndDelete) Run(env *GlobalAPI.GlobalAPI) external.Return {
	resp, err := env.Resources.Command.LoadResponceAndDelete(c.Key)
	return &LoadResponceAndDelete_Return{
		CommandResponce: resp,
		Error:           err.Error(),
	}
}
