package ResourcesControlCenter

import (
	"fmt"
	"phoenixbuilder/minecraft/protocol"
	"sync"
	"sync/atomic"
)

// 测定请求 ID 为 key 的物品操作请求是否在 i.itemStackRequest.datas 中。
// 如果存在，那么返回真，否则返回假
func (i *itemStackReuqestWithResponce) TestRequest(key int32) bool {
	i.itemStackRequest.lockDown.RLock()
	defer i.itemStackRequest.lockDown.RUnlock()
	// init
	_, ok := i.itemStackRequest.datas[key]
	return ok
	// return
}

// 测定请求 ID 为 key 的物品操作请求 key 是否在 i.itemStackResponce.datas 中。
// 如果存在，那么返回真，否则返回假
func (i *itemStackReuqestWithResponce) TestResponce(key int32) bool {
	i.itemStackResponce.lockDown.RLock()
	defer i.itemStackResponce.lockDown.RUnlock()
	// init
	_, ok := i.itemStackResponce.datas[key]
	return ok
	// return
}

// 将物品请求 ID 为 key 的物品操作放入 i.itemStackRequest.datas ，
// 并占用(锁定)此请求对应的互斥锁 i.itemStackRequest.datas[key].lockDown 。
// datas 指代相应槽位的变动结果，这用于更新本地库存数据
func (i *itemStackReuqestWithResponce) WriteRequest(
	key int32,
	datas map[ContainerID]StackRequestContainerInfo,
) error {
	if i.TestRequest(key) {
		return fmt.Errorf("WriteRequest: %v is already exist in i.itemStackRequest.datas", key)
	}
	// if key is already exist
	i.itemStackRequest.lockDown.Lock()
	defer i.itemStackRequest.lockDown.Unlock()
	// lock down resources
	i.itemStackRequest.datas[key] = singleItemStackRequest{
		lockDown: &sync.Mutex{},
		datas:    datas,
	}
	i.itemStackRequest.datas[key].lockDown.Lock()
	// lock down item stack request
	return nil
	// return
}

// 加载请求 ID 为 key 的物品操作请求
func (i *itemStackReuqestWithResponce) LoadRequest(key int32) (singleItemStackRequest, error) {
	if !i.TestRequest(key) {
		return singleItemStackRequest{}, fmt.Errorf("LoadRequest: %v is not recorded in i.itemStackRequest.datas", key)
	}
	// if key is not exist
	i.itemStackRequest.lockDown.RLock()
	defer i.itemStackRequest.lockDown.RUnlock()
	// lock down resources
	return i.itemStackRequest.datas[key], nil
	// return
}

// 将请求 ID 为 key 的物品操作请求从 i.itemStackRequest.datas 中移除
// 并释放此请求对应的互斥锁 i.itemStackRequest.datas[key].lockDown 。
// 属于私有实现
func (i *itemStackReuqestWithResponce) deleteRequest(key int32) error {
	if !i.TestRequest(key) {
		return fmt.Errorf("DeleteRequest: %v is not recorded in i.itemStackRequest.datas", key)
	}
	// if key is not exist
	i.itemStackRequest.lockDown.Lock()
	defer i.itemStackRequest.lockDown.Unlock()
	// lock down resources
	tmp := i.itemStackRequest.datas[key].lockDown
	// get tmp of the current resources
	delete(i.itemStackRequest.datas, key)
	newMap := map[int32]singleItemStackRequest{}
	for k, value := range i.itemStackRequest.datas {
		newMap[k] = value
	}
	i.itemStackRequest.datas = newMap
	// remove the key and values from i.itemStackRequest.datas
	tmp.Unlock()
	// unlock item stack request
	return nil
	// return
}

// 将请求 ID 为 key 的物品操作请求的返回值写入 i.itemStackResponce.datas
// 并释放 i.itemStackRequest.datas[key].lockDown 中对应的互斥锁，属于私有实现
func (i *itemStackReuqestWithResponce) writeResponce(key int32, resp protocol.ItemStackResponse) error {
	i.itemStackResponce.lockDown.Lock()
	defer i.itemStackResponce.lockDown.Unlock()
	// init
	i.itemStackResponce.datas[key] = resp
	// send item stack responce
	err := i.deleteRequest(key)
	if err != nil {
		return fmt.Errorf("writeResponce: %v", err)
	}
	// remove item stack reuqest from i.itemStackRequest.datas
	return nil
	// return
}

// 从 i.itemStackResponce.datas 读取请求 ID 为 key 的物品操作请求的返回值
// 并将此返回值从 i.itemStackResponce.datas 移除
func (i *itemStackReuqestWithResponce) LoadResponceAndDelete(key int32) (protocol.ItemStackResponse, error) {
	if !i.TestResponce(key) {
		return protocol.ItemStackResponse{}, fmt.Errorf("LoadResponceAndDelete: %v is not recorded in i.itemStackResponce.datas", key)
	}
	// if key is not exist
	i.itemStackResponce.lockDown.Lock()
	defer i.itemStackResponce.lockDown.Unlock()
	// lock down resources
	ans := i.itemStackResponce.datas[key]
	newMap := map[int32]protocol.ItemStackResponse{}
	for k, value := range i.itemStackResponce.datas {
		newMap[k] = value
	}
	i.itemStackResponce.datas = newMap
	// get responce and remove the key and values from i.itemStackResponce.datas
	return ans, nil
	// return
}

// 等待租赁服响应请求 ID 为 key 的物品操作请求。
// 在调用此函数后，会持续阻塞直到此物品操作请求所对应的互斥锁被释放
func (i *itemStackReuqestWithResponce) AwaitResponce(key int32) {
	if !i.TestRequest(key) {
		return
	}
	// if key is not exist
	i.itemStackRequest.lockDown.RLock()
	defer i.itemStackRequest.lockDown.RUnlock()
	// lock down resources
	tmp := i.itemStackRequest.datas[key].lockDown
	// get tmp of the current resources
	tmp.Lock()
	tmp.Unlock()
	// await responce
}

// 以原子操作获取上一次的请求 ID ，也就是 RequestID 。
// 如果从未进行过物品操作，则将会返回 1
func (i *itemStackReuqestWithResponce) GetCurrentRequestID() int32 {
	return atomic.LoadInt32(&i.currentRequestID)
}

// 以原子操作获取一个唯一的请求 ID ，也就是 RequestID
func (i *itemStackReuqestWithResponce) GetNewRequestID() int32 {
	return atomic.AddInt32(&i.currentRequestID, -2)
}

// 根据 newItem 中预期的新数据和租赁服返回的 resp ，
// 返回完整的新物品数据
func (i *itemStackReuqestWithResponce) GetNewItemData(
	newItem protocol.ItemInstance,
	resp protocol.StackResponseSlotInfo,
) (protocol.ItemInstance, error) {
	nbt := newItem.Stack.NBTData
	// 获取物品的旧 NBT 数据
	if resp.CustomName != "" {
		_, ok := nbt["tag"]
		if !ok {
			nbt["tag"] = map[string]interface{}{}
		}
		tag, normal := nbt["tag"].(map[string]interface{})
		if !normal {
			return protocol.ItemInstance{}, fmt.Errorf("getNewItemData: Failed to convert nbt[\"tag\"] into map[string]interface{}; nbt = %#v", nbt)
		}
		// tag
		_, ok = tag["display"]
		if !ok {
			tag["display"] = map[string]interface{}{}
			nbt["tag"].(map[string]interface{})["display"] = map[string]interface{}{}
		}
		_, normal = tag["display"].(map[string]interface{})
		if !normal {
			return protocol.ItemInstance{}, fmt.Errorf("getNewItemData: Failed to convert tag[\"display\"] into map[string]interface{}; tag = %#v", tag)
		}
		// display
		nbt["tag"].(map[string]interface{})["display"].(map[string]interface{})["Name"] = resp.CustomName
		// name
	}
	// set names
	newItem.Stack.NBTData = nbt
	newItem.Stack.Count = uint16(resp.Count)
	newItem.StackNetworkID = resp.StackNetworkID
	// update values
	return newItem, nil
	// return
}

// 根据租赁服返回的 resp 字段更新对应库存中对应槽位的物品数据。
// inventory 必须是一个指针，它指向了客户端唯一的库存数据。
// 属于私有实现
func (i *itemStackReuqestWithResponce) updateItemData(
	resp protocol.ItemStackResponse,
	inventory *inventoryContents,
) error {
	datas, err := i.LoadRequest(resp.RequestID)
	if err != nil {
		return fmt.Errorf("updateItemData: %v", err)
	}
	// get responce of the target item stack
	for _, value := range resp.ContainerInfo {
		if datas.datas == nil {
			panic("updateItemData: Attempt to send packet.ItemStackRequest without using ResourcesControlCenter")
		}
		_, ok := datas.datas[ContainerID(value.ContainerID)]
		if !ok {
			panic(fmt.Sprintf("updateItemData: item change result %v not found or not provided(packet.ItemStackRequest related); datas.datas = %#v, value = %#v", ContainerID(value.ContainerID), datas.datas, value))
		}
		currentChanges := datas.datas[ContainerID(value.ContainerID)].ChangeResult
		windowID := datas.datas[ContainerID(value.ContainerID)].WindowID
		for _, v := range value.SlotInfo {
			newItem, err := i.GetNewItemData(
				currentChanges[v.Slot],
				v,
			)
			if err != nil {
				panic(fmt.Sprintf("updateItemData: Failed to get new item data; currentChanges[v.Slot] = %#v, v = %#v", currentChanges[v.Slot], v))
			}
			inventory.writeItemStackInfo(windowID, v.Slot, newItem)
		}
	}
	// set item datas
	return nil
	// return
}
