package ResourcesControlCenter

import (
	"phoenixbuilder/minecraft/protocol/packet"

	"github.com/google/uuid"
)

/*
占用客户端的容器资源。

当 tryMode 为真时，将尝试占用资源并返回占用结果，此对应返回值 bool 部分。
若 tryMode 为假，则返回值 bool 部分永远为真。

无论 tryMode 的真假性如何，当且仅当容器资源被成功占用时，
返回值的字符串的长度才会大于 0 ，且表达为 uuid 形式，否则为空字符串
*/
func (c *container) Occupy(tryMode bool) (bool, string) {
	newUUID, err := uuid.NewUUID()
	if err != nil || newUUID == uuid.Nil {
		return c.Occupy(tryMode)
	}
	uniqueId := newUUID.String()
	// spawn a new uuid to set the holder of container resources
	if tryMode {
		success := c.isUsing.lockDown.TryLock()
		if success {
			c.isUsing.holder = uniqueId
			// set the holder of container resources
			return true, uniqueId
			// return
		} else {
			return false, ""
			// return
		}
	}
	// if is try mode
	c.isUsing.lockDown.Lock()
	// lock down resources
	c.isUsing.holder = uniqueId
	// set the holder of container resources
	return true, uniqueId
	// return
}

// 释放客户端的容器资源，返回值代表执行结果。
// holder 指代容器资源的占用者，当且仅当填写的占用者
// 可以与内部记录的占用者对应时才会成功释放容器资源
func (c *container) Release(holder string) bool {
	if c.isUsing.holder == holder && c.isUsing.holder != "" {
		c.isUsing.holder = ""
		c.isUsing.lockDown.Unlock()
		return true
	}
	return false
}

// 用于在 打开/关闭 容器前执行，
// 便于后续调用 AwaitResponceAfterSendPacket 以阻塞程序的执行从而
// 达到等待租赁服响应容器操作的目的
func (c *container) AwaitResponceBeforeSendPacket() {
	c.awaitChanges.Lock()
}

// 用于在 打开/关闭 容器后执行。
// 用于等待租赁服响应容器的打开或关闭操作。
// 在调用此函数后，会持续阻塞直到相关操作所对应的互斥锁被释放
func (c *container) AwaitResponceAfterSendPacket() {
	c.awaitChanges.Lock()
	c.awaitChanges.Unlock()
}

// 释放 c.awaitChanges 中关于容器操作的互斥锁。如果互斥锁未被锁定，程序也仍不会发生惊慌。
// 当且仅当租赁服确认客户端的容器操作时，此函数才会被调用。
// 属于私有实现
func (c *container) responceContainerOperation() {
	c.awaitChanges.TryLock()
	c.awaitChanges.Unlock()
}

// 将 datas 写入 c.containerOpen.datas ，属于私有实现
func (c *container) writeContainerOpenDatas(datas *packet.ContainerOpen) {
	c.containerOpen.lockDown.Lock()
	defer c.containerOpen.lockDown.Unlock()
	// init
	c.containerOpen.datas = datas
	// set values
}

// 取得当前已打开容器的数据。
// 如果容器未被打开或已被关闭，则会返回 nil 。
// 返回值虽然是一个地址，但它所指向的实际是一个副本
func (c *container) GetContainerOpenDatas() *packet.ContainerOpen {
	c.containerOpen.lockDown.RLock()
	defer c.containerOpen.lockDown.RUnlock()
	// init
	if c.containerOpen.datas == nil {
		return nil
	} else {
		new := *c.containerOpen.datas
		return &new
	}
	// return
}

// 将 datas 写入 c.containerClose.datas ，属于私有实现
func (c *container) writeContainerCloseDatas(datas *packet.ContainerClose) {
	c.containerClose.lockDown.Lock()
	defer c.containerClose.lockDown.Unlock()
	// init
	c.containerClose.datas = datas
	// set values
}

// 取得上次关闭容器时租赁服的响应数据。
// 如果现在有容器已被打开或容器从未被关闭，则会返回 nil 。
// 返回值虽然是一个地址，但它所指向的实际是一个副本
func (c *container) GetContainerCloseDatas() *packet.ContainerClose {
	c.containerClose.lockDown.RLock()
	defer c.containerClose.lockDown.RUnlock()
	// init
	if c.containerClose.datas == nil {
		return nil
	} else {
		new := *c.containerClose.datas
		return &new
	}
	// return
}
