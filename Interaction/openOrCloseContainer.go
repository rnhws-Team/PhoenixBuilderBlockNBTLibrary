package GlobalAPI

import (
	"fmt"
	"phoenixbuilder/minecraft/protocol/packet"
)

/*
打开 pos 处名为 blockName 且方块状态为 blockStates 的容器，且只有当打开完成后才会返回值。
slot 字段代表玩家此时手持的物品栏，因为打开容器实际上是一次方块点击事件。

当 needOccupyContainerResources 为真时，此函数会主动占用容器资源并在函数执行完成后释放。
一般情况下我建议此参数填 false ，因为打开容器仅仅是一系列容器操作的一个步骤，
因此此函数中不应该贸然修改容器资源，否则可能会造成潜在的问题。
*/
func (g *GlobalAPI) OpenContainer(
	pos [3]int32,
	blockName string,
	blockStates map[string]interface{},
	slot uint8,
	needOccupyContainerResources bool,
) (bool, error) {
	if needOccupyContainerResources {
		_, holder := g.Resources.Container.Occupy(false)
		defer g.Resources.Container.Release(holder)
	}
	// lock down resources
	g.Resources.Container.AwaitResponceBeforeSendPacket()
	// await responce before send packet
	err := g.UseItemOnBlocks(slot, pos, blockName, blockStates, false)
	if err != nil {
		return false, fmt.Errorf("OpenContainer: %v", err)
	}
	// open container
	g.Resources.Container.AwaitResponceAfterSendPacket()
	// wait changes
	if g.Resources.Container.GetContainerOpenDatas() == nil {

	}
	return true, nil
	// return
}

/*
关闭已经打开的容器，且只有当容器被关闭后才会返回值。您应该确保容器被关闭后，对应的容器公用资源被释放。
此函数并不会检查此前是否已经打开过容器，因此您需要确保关闭容器前已经打开了一个容器，否则可能会造成潜在的问题
*/
func (g *GlobalAPI) CloseContainer() error {
	g.PacketHandleResult.ContainerResources.AwaitResponceBeforeSendPacket()
	// await responce before send packet
	err := g.WritePacket(&packet.ContainerClose{
		WindowID:   g.PacketHandleResult.ContainerResources.GetContainerOpenDatas().WindowID,
		ServerSide: false,
	})
	if err != nil {
		return fmt.Errorf("CloseContainer: %v", err)
	}
	// close container
	g.PacketHandleResult.ContainerResources.AwaitResponceAfterSendPacket()
	// wait changes
	return nil
	// return
}
