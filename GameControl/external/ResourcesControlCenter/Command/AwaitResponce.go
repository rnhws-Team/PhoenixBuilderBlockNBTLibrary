package external_resources_command

import (
	"phoenixbuilder/GameControl/GlobalAPI"
	"phoenixbuilder/GameControl/external"
	"phoenixbuilder/encoding"

	"github.com/google/uuid"
)

// func (*commandRequestWithResponce).AwaitResponce(key uuid.UUID)

type AwaitResponce struct {
	Key uuid.UUID `json:"key"`
}

type AwaitResponce_Return struct{}

func (c *AwaitResponce) Marshal(io encoding.IO) {
	external.TestError(io.UUID(&c.Key))
}

func (c *AwaitResponce_Return) Marshal(io encoding.IO) {}

func (c *AwaitResponce) Run(env *GlobalAPI.GlobalAPI) external.Return {
	env.Resources.Command.AwaitResponce(c.Key)
	return &TestResponce_Return{}
}
