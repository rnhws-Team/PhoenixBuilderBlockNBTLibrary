package GlobalAPI

import (
	"phoenixbuilder/ResourcesControlCenter"
	"phoenixbuilder/minecraft/protocol/packet"
)

// 描述了一个通用型的 API ，它可以被用在任何地方，但主要被用于 Bdump/blockNBT
type GlobalAPI struct {
	WritePacket        func(packet.Packet) error                  // 用于向租赁服发送数据包的函数
	BotName            string                                     // 客户端的游戏昵称
	BotIdentity        string                                     // 客户端的唯一标识符 [当前还未使用]
	BotUniqueID        int64                                      // 客户端的唯一 ID [当前还未使用]
	BotRunTimeID       uint64                                     // 客户端的运行时 ID
	PacketHandleResult *ResourcesControlCenter.PacketHandleResult // 保存包处理结果；由外部实现实时更新
}
