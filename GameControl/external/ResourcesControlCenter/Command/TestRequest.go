package external_resources_command

import (
	"phoenixbuilder/GameControl/GlobalAPI"
	"phoenixbuilder/GameControl/external"
	"phoenixbuilder/encoding"

	"github.com/google/uuid"
)

// func (*commandRequestWithResponce).TestRequest(key uuid.UUID) bool

type TestRequest struct {
	Key uuid.UUID `json:"key"`
}

type TestRequest_Return struct {
	Exist bool `json:"exist"`
}

func (c *TestRequest) Marshal(io encoding.IO) {
	external.TestError(io.UUID(&c.Key))
}

func (c *TestRequest_Return) Marshal(io encoding.IO) {
	external.TestError(io.Bool(&c.Exist))
}

func (c *TestRequest) Run(env *GlobalAPI.GlobalAPI) external.Return {
	resp := env.Resources.Command.TestRequest(c.Key)
	return &TestRequest_Return{Exist: resp}
}
