package external_resources_command

import (
	"phoenixbuilder/GameControl/GlobalAPI"
	"phoenixbuilder/GameControl/external"
	"phoenixbuilder/lib/encoding"

	"github.com/google/uuid"
)

// func (*commandRequestWithResponce).WriteRequest(key uuid.UUID) error

type WriteRequest struct {
	Key uuid.UUID `json:"key"`
}

type WriteRequest_Return struct {
	Error string `json:"error"` // error
}

func (c *WriteRequest) Marshal(io encoding.IO) {
	external.TestError(io.UUID(&c.Key))
}

func (c *WriteRequest_Return) Marshal(io encoding.IO) {
	external.TestError(io.String(&c.Error))
}

func (c *WriteRequest) Run(env *GlobalAPI.GlobalAPI) external.Return {
	resp := env.Resources.Command.WriteRequest(c.Key)
	return &WriteRequest_Return{Error: resp.Error()}
}
