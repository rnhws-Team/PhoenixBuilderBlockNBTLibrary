package external_resources_command

import (
	"phoenixbuilder/GameControl/GlobalAPI"
	"phoenixbuilder/GameControl/external"
	"phoenixbuilder/encoding"

	"github.com/google/uuid"
)

// func (*commandRequestWithResponce).TestResponce(key uuid.UUID) bool

type TestResponce struct {
	Key uuid.UUID `json:"key"`
}

type TestResponce_Return struct {
	Exist bool `json:"exist"`
}

func (c *TestResponce) Marshal(io encoding.IO) {
	external.TestError(io.UUID(&c.Key))
}

func (c *TestResponce_Return) Marshal(io encoding.IO) {
	external.TestError(io.Bool(&c.Exist))
}

func (c *TestResponce) Run(env *GlobalAPI.GlobalAPI) external.Return {
	resp := env.Resources.Command.TestResponce(c.Key)
	return &TestResponce_Return{Exist: resp}
}
