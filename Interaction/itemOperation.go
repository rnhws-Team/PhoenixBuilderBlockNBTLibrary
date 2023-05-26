package GlobalAPI

import (
	"fmt"
	"phoenixbuilder/ResourcesControlCenter"
	"phoenixbuilder/minecraft/protocol"
	"phoenixbuilder/minecraft/protocol/packet"
)

// 空物品数据使用它来描述
var AirItem protocol.ItemInstance = protocol.ItemInstance{
	StackNetworkID: 0,
	Stack: protocol.ItemStack{
		ItemType: protocol.ItemType{
			NetworkID:     0,
			MetadataValue: 0,
		},
		BlockRuntimeID: 0,
		Count:          0,
		NBTData:        map[string]interface{}(nil),
		CanBePlacedOn:  []string(nil),
		CanBreak:       []string(nil),
		HasNetworkID:   false,
	},
}

// 在描述物品移动操作时所使用的结构体
type MoveItemDatas struct {
	WindowID    int16 // 物品所在库存的窗口 ID
	ContainerID uint8 // 物品所在库存的库存类型 ID
	Slot        uint8 // 物品所在的槽位
}

// 将库存编号为 source 所指代的物品移动到 destination 所指代的槽位
// 且只移动 moveCount 个物品。
// details 指代相应槽位的预期变动结果，它将作为更新本地库存数据的依据。
// 当且仅当物品操作得到租赁服的响应后，此函数才会返回物品操作结果。
func (g *GlobalAPI) MoveItem(
	source MoveItemDatas,
	destination MoveItemDatas,
	details ItemChangeDetails,
	moveCount uint8,
) ([]protocol.ItemStackResponse, error) {
	placeStackRequestAction := protocol.PlaceStackRequestAction{}
	// 初始化
	itemOnSource, err := g.Resources.Inventory.GetItemStackInfo(uint32(source.WindowID), source.Slot)
	if err != nil {
		return []protocol.ItemStackResponse{}, fmt.Errorf("MoveItem: %v", err)
	}
	itemOnDestination, _ := g.Resources.Inventory.GetItemStackInfo(uint32(destination.WindowID), destination.Slot)
	// 取得 source 和 destination 处的物品信息
	if moveCount <= uint8(itemOnSource.Stack.Count) || source.WindowID == -1 {
		placeStackRequestAction.Count = moveCount
	} else {
		placeStackRequestAction.Count = uint8(itemOnSource.Stack.Count)
	}
	// 得到欲移动的物品数量
	placeStackRequestAction.Source = protocol.StackRequestSlotInfo{
		ContainerID:    source.ContainerID,
		Slot:           source.Slot,
		StackNetworkID: itemOnSource.StackNetworkID,
	}
	placeStackRequestAction.Destination = protocol.StackRequestSlotInfo{
		ContainerID:    destination.ContainerID,
		Slot:           destination.Slot,
		StackNetworkID: itemOnDestination.StackNetworkID,
	}
	// 构造 placeStackRequestAction 结构体
	ans, err := g.SendItemStackRequestWithResponce(
		&packet.ItemStackRequest{
			Requests: []protocol.ItemStackRequest{
				{
					Actions: []protocol.StackRequestAction{
						&placeStackRequestAction,
					},
					FilterStrings: []string{},
				},
			},
		},
		[]ItemChangeDetails{details},
	)
	if err != nil {
		return []protocol.ItemStackResponse{}, fmt.Errorf("MoveItem: %v", err)
	}
	// 发送物品操作请求
	return ans, nil
	// 返回值
}

// 将已放入铁砧第一格(注意是第一格)的物品的物品名称修改为 name 并返还到背包中的 slot 处。
// 当且仅当租赁服回应操作结果后此函数再返回值。
//
// 当遭遇改名失败时，将尝试撤销名称修改请求。
// 如果原有物品栏已被占用，则对应的物品将被直接丢出，
// 此对应返回的第二个布尔值，当为真时代表物品已被丢出，
// 否则物品将会被还原到背包的原始位置。
// 改名成功时此参数将始终返回 nil
//
// 部分情况下此函数可能会遇见无法处理的错误，届时程序将抛出严重错误(panic)
func (g *GlobalAPI) ChangeItemName(
	name string,
	slot uint8,
) (bool, *bool, error) {
	containerOpenDatas := g.Resources.Container.GetContainerOpenDatas()
	if containerOpenDatas == nil {
		return false, nil, fmt.Errorf("ChangeItemName: Anvil has not opened")
	}
	// 如果铁砧未被打开
	itemDatas, err := g.Resources.Inventory.GetItemStackInfo(
		uint32(containerOpenDatas.WindowID),
		1,
	)
	if err != nil {
		return false, nil, fmt.Errorf("ChangeItemName: %v", err)
	}
	// 取得已放入铁砧的物品的物品数据
	revertFunc := func() (bool, *bool, error) {
		filterAns, err := g.Resources.Inventory.ListSlot(0, &[]int32{0})
		if err != nil {
			panic(fmt.Sprintf("ChangeItemName: %v", err)) // 这个错误理论上是不可能发生的
		}
		// 筛选出背包中还未被占用实际物品的物品栏
		successStates := false
		for _, value := range filterAns {
			if value == slot {
				successStates = true
			}
		}
		// 我们需要确定槽位 slot 是否已经被占用了。
		// successStates 显示了占用结果。
		// 当被占用时，改名失败的物品会被直接丢出。
		if len(filterAns) <= 0 || !successStates {
			if err != nil {
				panic(fmt.Sprintf("ChangeItemName: %v", err)) // 发生这个错误只有一种可能，那就是铁砧被拆了什么的，但这是用户导致的
			}
			_, err = g.DropItemAll(protocol.StackRequestSlotInfo{
				ContainerID:    0,
				Slot:           1,
				StackNetworkID: itemDatas.StackNetworkID,
			}, uint32(containerOpenDatas.WindowID))
			if err != nil {
				panic(fmt.Sprintf("ChangeItemName: %v", err)) // 发生了一些未知错误
			}
			revertMethod := true
			return true, &revertMethod, nil
		}
		// 如果背包中没有空出的物品栏，或者槽位 slot 占用了，
		// 则将未能成功改名的物品直接丢出来
		source := MoveItemDatas{
			WindowID:    int16(containerOpenDatas.WindowID),
			ContainerID: 0,
			Slot:        1,
		}
		destination := MoveItemDatas{
			WindowID:    0,
			ContainerID: 0xc,
			Slot:        slot,
		}
		ans, err := g.MoveItem(
			source,
			destination,
			ItemChangeDetails{
				map[ResourcesControlCenter.ContainerID]ResourcesControlCenter.StackRequestContainerInfo{
					0: {
						WindowID: uint32(containerOpenDatas.WindowID),
						ChangeResult: map[uint8]protocol.ItemInstance{
							1: AirItem,
						},
					},
					0xc: {
						WindowID: 0,
						ChangeResult: map[uint8]protocol.ItemInstance{
							slot: itemDatas,
						},
					},
				},
			},
			uint8(itemDatas.Stack.Count),
		)
		if err != nil {
			panic(fmt.Sprintf("ChangeItemName: %v", err))
		}
		if ans[0].Status != protocol.ItemStackResponseStatusOK {
			return false, nil, nil
		}
		// 看起来一切都很顺利，
		// 我们成功将未能成功改名的物品送回了 slot 槽位
		revertMethod := false
		return true, &revertMethod, nil
		// return
	}
	// 构造一个函数用于处理改名失败时的善后处理
	// 返回的第一个布尔值代表善后处理的结果，如果为假，
	// 那么我们将持续善后处理，直到成功或者程序惊慌。
	// 返回的第二个布尔值代表善后处理的方式，
	// 为真时代表使用丢出法处理物品，否则采用正常方法处理物品
	revertFuncRuner := func() (bool, *bool, error) {
		for {
			successStates, revertMethod, err := revertFunc()
			if err != nil {
				panic(fmt.Sprintf("ChangeItemName: %v", err))
			}
			if !successStates {
				continue
			}
			return false, revertMethod, nil
		}
	}
	// 构造善后处理函数的包装函数
	newRequestID := g.Resources.ItemStackOperation.GetNewRequestID()
	// 请求一个新的 RequestID 用于 ItemStackRequest
	placeStackRequestAction := protocol.PlaceStackRequestAction{}
	placeStackRequestAction.Count = byte(itemDatas.Stack.Count)
	placeStackRequestAction.Source = protocol.StackRequestSlotInfo{
		ContainerID:    0x3c,
		Slot:           0x32,
		StackNetworkID: newRequestID,
	}
	placeStackRequestAction.Destination = protocol.StackRequestSlotInfo{
		ContainerID:    0xc,
		Slot:           slot,
		StackNetworkID: 0,
	}
	// 构造一个新的 PlaceStackRequestAction 结构体
	newItemStackRequest := packet.ItemStackRequest{
		Requests: []protocol.ItemStackRequest{
			{
				RequestID: newRequestID,
				Actions: []protocol.StackRequestAction{
					&protocol.CraftRecipeOptionalStackRequestAction{
						RecipeNetworkID:   0,
						FilterStringIndex: 0,
					},
					&protocol.ConsumeStackRequestAction{
						DestroyStackRequestAction: protocol.DestroyStackRequestAction{
							Count: byte(itemDatas.Stack.Count),
							Source: protocol.StackRequestSlotInfo{
								ContainerID:    0,
								Slot:           1,
								StackNetworkID: itemDatas.StackNetworkID,
							},
						},
					},
					&placeStackRequestAction,
				},
				FilterStrings: []string{name},
			},
		},
	}
	// 构造一个新的 ItemStackRequest 结构体
	err = g.Resources.ItemStackOperation.SetItemName(
		&itemDatas,
		name,
	)
	if err != nil {
		return revertFuncRuner()
	}
	// 更新物品数据中的名称字段以用于更新本地库存数据
	_, ok := itemDatas.Stack.NBTData["RepairCost"]
	if !ok {
		itemDatas.Stack.NBTData["RepairCost"] = int32(0)
	}
	// 更新物品数据中的 RepairCost 字段用于更新本地库存数据
	err = g.Resources.ItemStackOperation.WriteRequest(
		newRequestID,
		map[ResourcesControlCenter.ContainerID]ResourcesControlCenter.StackRequestContainerInfo{
			0xc: {
				WindowID: 0,
				ChangeResult: map[uint8]protocol.ItemInstance{
					slot: itemDatas,
				},
			},
			0x0: {
				WindowID: uint32(containerOpenDatas.WindowID),
				ChangeResult: map[uint8]protocol.ItemInstance{
					1: AirItem,
				},
			},
			0x1: {
				WindowID: uint32(containerOpenDatas.WindowID),
				ChangeResult: map[uint8]protocol.ItemInstance{
					2: AirItem,
				},
			},
			0x3c: {
				WindowID: uint32(containerOpenDatas.WindowID),
				ChangeResult: map[uint8]protocol.ItemInstance{
					0x32: AirItem,
				},
			},
		},
	)
	if err != nil {
		return revertFuncRuner()
	}
	// 写入请求到等待队列
	err = g.WritePacket(&newItemStackRequest)
	if err != nil {
		panic(fmt.Sprintf("ChangeItemName: %v", err))
	}
	// 发送物品操作请求
	g.Resources.ItemStackOperation.AwaitResponce(newRequestID)
	// 等待租赁服响应物品操作请求
	ans, err := g.Resources.ItemStackOperation.LoadResponceAndDelete(newRequestID)
	if err != nil || ans.Status != protocol.ItemStackResponseStatusOK {
		return revertFuncRuner()
	}
	// 当改名失败时尝试将物品恢复到背包中对应的位置
	return true, nil, nil
	// 返回值
}

// 将 source 所指代的槽位中的全部物品丢出。
// windowID 指代被丢出物品所在库存的窗口 ID
func (g *GlobalAPI) DropItemAll(
	source protocol.StackRequestSlotInfo,
	windowID uint32,
) (bool, error) {
	ans, err := g.SendItemStackRequestWithResponce(
		&packet.ItemStackRequest{
			Requests: []protocol.ItemStackRequest{
				{
					Actions: []protocol.StackRequestAction{
						&protocol.DropStackRequestAction{
							Count:    64,
							Source:   source,
							Randomly: false,
						},
					},
				},
			},
		},
		[]ItemChangeDetails{
			{
				map[ResourcesControlCenter.ContainerID]ResourcesControlCenter.StackRequestContainerInfo{
					ResourcesControlCenter.ContainerID(source.ContainerID): {
						WindowID: windowID,
						ChangeResult: map[uint8]protocol.ItemInstance{
							source.Slot: AirItem,
						},
					},
				},
			},
		},
	)
	if err != nil {
		return false, fmt.Errorf("DropItemAll: %v", err)
	}
	if ans[0].Status != protocol.ItemStackResponseStatusOK {
		return false, nil
	}
	// 发送物品丢掷请求
	return true, nil
	// 返回值
}
