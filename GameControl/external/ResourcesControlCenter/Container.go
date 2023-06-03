package external_resources

import (
	"phoenixbuilder/GameControl/GlobalAPI"
	"phoenixbuilder/GameControl/external"
)

/*
Function List

func (*resourcesOccupy).GetOccupyStates() bool
func (*resourcesOccupy).Occupy() string
func (*resourcesOccupy).Release(holder string) bool
func (*container).AwaitResponceAfterSendPacket()
func (*container).AwaitResponceBeforeSendPacket()
func (*container).GetContainerCloseDatas() *packet.ContainerClose
func (*container).GetContainerOpenDatas() *packet.ContainerOpen
*/

// ------------------------- GetOccupyStates -------------------------

type Container_GetOccupyStates struct{}

type Container_GetOccupyStates_Return struct {
	OccupyStates bool `json:"occupy_states"`
}

func (i *Container_GetOccupyStates) Run(env *GlobalAPI.GlobalAPI) external.Return {
	resp := env.Resources.Container.GetOccupyStates()
	return Container_GetOccupyStates_Return{OccupyStates: resp}
}

// ------------------------- END -------------------------
