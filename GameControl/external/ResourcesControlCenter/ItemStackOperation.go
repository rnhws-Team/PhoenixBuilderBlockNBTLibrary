package external_resources

import (
	"phoenixbuilder/GameControl/GlobalAPI"
	"phoenixbuilder/GameControl/ResourcesControlCenter"
	"phoenixbuilder/GameControl/external"
	"phoenixbuilder/minecraft/protocol"
)

type ContainerID ResourcesControlCenter.ContainerID

type StackRequestContainerInfo struct {
	WindowID     uint32                          `json:"window_id"`
	ChangeResult map[uint8]protocol.ItemInstance `json:"change_result"`
}

/*
Function List

func (*itemStackReuqestWithResponce).AwaitResponce(key int32)
func (*itemStackReuqestWithResponce).GetCurrentRequestID() int32
func (*itemStackReuqestWithResponce).GetNewItemData(newItem protocol.ItemInstance, resp protocol.StackResponseSlotInfo) (protocol.ItemInstance, error)
func (*itemStackReuqestWithResponce).GetNewRequestID() int32
func (*itemStackReuqestWithResponce).LoadResponceAndDelete(key int32) (protocol.ItemStackResponse, error)
func (*itemStackReuqestWithResponce).SetItemName(item *protocol.ItemInstance, newItemName string) error
func (*itemStackReuqestWithResponce).TestRequest(key int32) bool
func (*itemStackReuqestWithResponce).TestResponce(key int32) bool
func (*itemStackReuqestWithResponce).WriteRequest(key int32, datas map[ContainerID]StackRequestContainerInfo) error
*/

// ------------------------- TestRequest -------------------------

type ItemStackOperation_TestRequest struct {
	Key int32 `json:"key"`
}

type ItemStackOperation_TestRequest_Return struct {
	Exist bool `json:"exist"`
}

func (i *ItemStackOperation_TestRequest) Run(
	env *GlobalAPI.GlobalAPI,
) external.Return {
	resp := env.Resources.ItemStackOperation.TestRequest(i.Key)
	return ItemStackOperation_TestRequest_Return{Exist: resp}
}

// ------------------------- TestResponce -------------------------

type ItemStackOperation_TestResponce struct {
	Key int32 `json:"key"`
}

type ItemStackOperation_TestResponce_Return struct {
	Exist bool `json:"exist"`
}

func (i *ItemStackOperation_TestResponce) Run(
	env *GlobalAPI.GlobalAPI,
) external.Return {
	resp := env.Resources.ItemStackOperation.TestResponce(i.Key)
	return ItemStackOperation_TestResponce_Return{Exist: resp}
}

// ------------------------- WriteRequest -------------------------

type ItemStackOperation_WriteRequest struct {
	Key   int32                                     `json:"key"`
	Datas map[ContainerID]StackRequestContainerInfo `json:"datas"`
}

type ItemStackOperation_WriteRequest_Return struct {
	Error error `json:"error"`
}

func (i *ItemStackOperation_WriteRequest) Run(
	env *GlobalAPI.GlobalAPI,
) external.Return {
	tmp := map[ResourcesControlCenter.ContainerID]ResourcesControlCenter.StackRequestContainerInfo{}
	for key, value := range i.Datas {
		tmp[ResourcesControlCenter.ContainerID(key)] = ResourcesControlCenter.StackRequestContainerInfo(value)
	}
	resp := env.Resources.ItemStackOperation.WriteRequest(i.Key, tmp)
	return ItemStackOperation_WriteRequest_Return{Error: resp}
}

// ------------------------- LoadResponceAndDelete -------------------------

type ItemStackOperation_LoadResponceAndDelete struct {
	Key int32 `json:"key"`
}

type ItemStackOperation_LoadResponceAndDelete_Return struct {
	Responce protocol.ItemStackResponse `json:"item_info"`
	Error    error                      `json:"error"`
}

func (i *ItemStackOperation_LoadResponceAndDelete) Run(
	env *GlobalAPI.GlobalAPI,
) external.Return {
	resp, err := env.Resources.ItemStackOperation.LoadResponceAndDelete(i.Key)
	return ItemStackOperation_LoadResponceAndDelete_Return{
		Responce: resp,
		Error:    err,
	}
}

// ------------------------- AwaitResponce -------------------------

type ItemStackOperation_AwaitResponce struct {
	Key int32 `json:"key"`
}

type ItemStackOperation_AwaitResponce_Return struct{}

func (i *ItemStackOperation_AwaitResponce) Run(
	env *GlobalAPI.GlobalAPI,
) external.Return {
	env.Resources.ItemStackOperation.AwaitResponce(i.Key)
	return ItemStackOperation_AwaitResponce_Return{}
}

// ------------------------- GetCurrentRequestID -------------------------

type ItemStackOperation_GetCurrentRequestID struct{}

type ItemStackOperation_GetCurrentRequestID_Return struct {
	CurrentRequestID int32 `json:"current_request_id"`
}

func (i *ItemStackOperation_GetCurrentRequestID) Run(
	env *GlobalAPI.GlobalAPI,
) external.Return {
	resp := env.Resources.ItemStackOperation.GetCurrentRequestID()
	return ItemStackOperation_GetCurrentRequestID_Return{CurrentRequestID: resp}
}

// ------------------------- GetNewRequestID -------------------------

type ItemStackOperation_GetNewRequestID struct{}

type ItemStackOperation_GetNewRequestID_Return struct {
	NewRequestID int32 `json:"new_request_id"`
}

func (i *ItemStackOperation_GetNewRequestID) Run(
	env *GlobalAPI.GlobalAPI,
) external.Return {
	resp := env.Resources.ItemStackOperation.GetNewRequestID()
	return ItemStackOperation_GetNewRequestID_Return{NewRequestID: resp}
}

// ------------------------- SetItemName -------------------------

type ItemStackOperation_SetItemName struct {
	ItemInfo    *protocol.ItemInstance `json:"item_info"`
	NewItemName string                 `json:"new_item_name"`
}

type ItemStackOperation_SetItemName_Return struct {
	NewItemInfo protocol.ItemInstance `json:"new_item_info"`
	Error       error                 `json:"error"`
}

func (i *ItemStackOperation_SetItemName) Run(
	env *GlobalAPI.GlobalAPI,
) external.Return {
	resp := env.Resources.ItemStackOperation.SetItemName(i.ItemInfo, i.NewItemName)
	return ItemStackOperation_SetItemName_Return{
		NewItemInfo: *i.ItemInfo,
		Error:       resp,
	}
}

// ------------------------- GetNewItemData -------------------------

type ItemStackOperation_GetNewItemData struct {
	ItemInfo           protocol.ItemInstance          `json:"item_info"`
	ResponceFromServer protocol.StackResponseSlotInfo `json:"responce_from_server"`
}

type ItemStackOperation_GetNewItemData_Return struct {
	NewItemInfo protocol.ItemInstance `json:"new_item_info"`
	Error       error                 `json:"error"`
}

func (i *ItemStackOperation_GetNewItemData) Run(
	env *GlobalAPI.GlobalAPI,
) external.Return {
	newItemInfo, err := env.Resources.ItemStackOperation.GetNewItemData(
		i.ItemInfo,
		i.ResponceFromServer,
	)
	return ItemStackOperation_GetNewItemData_Return{
		NewItemInfo: newItemInfo,
		Error:       err,
	}
}

// ------------------------- END -------------------------
