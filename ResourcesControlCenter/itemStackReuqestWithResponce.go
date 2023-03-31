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
// 并占用(锁定)此请求对应的互斥锁。
// datas 指代相应槽位的变动结果，这用于更新本地库存数据
func (i *itemStackReuqestWithResponce) WriteRequest(
	key int32,
	datas []StackRequestContainerInfo,
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
// 并释放此请求对应的互斥锁
func (i *itemStackReuqestWithResponce) DeleteRequest(key int32) error {
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
// 并释放 i.itemStackRequest.datas 中对应的互斥锁，属于私有实现
func (i *itemStackReuqestWithResponce) writeResponce(key int32, resp protocol.ItemStackResponse) error {
	i.itemStackResponce.lockDown.Lock()
	defer i.itemStackResponce.lockDown.Unlock()
	// init
	i.itemStackResponce.datas[key] = resp
	// send item stack responce
	err := i.DeleteRequest(key)
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

func (i *itemStackReuqestWithResponce) GetNewItemData(
	key int32,
	resp protocol.ItemStackResponse,
) {

}
