package external_resources

import (
	"phoenixbuilder/GameControl/GlobalAPI"
	"phoenixbuilder/GameControl/external"
	"phoenixbuilder/minecraft/protocol"
)

/*
Function List

func (*inventoryContents).GetInventoryInfo(windowID uint32) (map[uint8]protocol.ItemInstance, error)
func (*inventoryContents).GetItemStackInfo(windowID uint32, slotLocation uint8) (protocol.ItemInstance, error)
func (*inventoryContents).ListSlot(windowID uint32, filter *[]int32) ([]uint8, error)
func (*inventoryContents).ListWindowID() []uint32
*/

// ------------------------- ListWindowID -------------------------

type Inventory_ListWindowID struct {
}

type Inventory_ListWindowID_Return struct {
	ListResult []uint32 `json:"list_result"`
}

func (i *Inventory_ListWindowID) Run(env *GlobalAPI.GlobalAPI) external.Return {
	resp := env.Resources.Inventory.ListWindowID()
	return Inventory_ListWindowID_Return{ListResult: resp}
}

// ------------------------- ListSlot -------------------------

type Inventory_ListSlot struct {
	WindowID uint32   `json:"window_id"`
	Filter   *[]int32 `json:"filter"`
}

type Inventory_ListSlot_Return struct {
	ListResult []uint8 `json:"list_result"`
	Error      error   `json:"error"`
}

func (i *Inventory_ListSlot) Run(env *GlobalAPI.GlobalAPI) external.Return {
	resp, err := env.Resources.Inventory.ListSlot(i.WindowID, i.Filter)
	return Inventory_ListSlot_Return{
		ListResult: resp,
		Error:      err,
	}
}

// ------------------------- GetInventoryInfo -------------------------

type Inventory_GetInventoryInfo struct {
	WindowID uint32 `json:"window_id"`
}

type Inventory_GetInventoryInfo_Return struct {
	Inventory map[uint8]protocol.ItemInstance `json:"inventory"`
	Error     error                           `json:"error"`
}

func (i *Inventory_GetInventoryInfo) Run(env *GlobalAPI.GlobalAPI) external.Return {
	inventory, err := env.Resources.Inventory.GetInventoryInfo(i.WindowID)
	return Inventory_GetInventoryInfo_Return{
		Inventory: inventory,
		Error:     err,
	}
}

// ------------------------- GetItemStackInfo -------------------------

type Inventory_GetItemStackInfo struct {
	WindowID     uint32 `json:"window_id"`
	SlotLocation uint8  `json:"slot_location"`
}

type Inventory_GetItemStackInfo_Return struct {
	ItemInfo protocol.ItemInstance `json:"item_info"`
	Error    error                 `json:"error"`
}

func (i *Inventory_GetItemStackInfo) Run(env *GlobalAPI.GlobalAPI) external.Return {
	itemInfo, err := env.Resources.Inventory.GetItemStackInfo(
		i.WindowID,
		i.SlotLocation,
	)
	return Inventory_GetItemStackInfo_Return{
		ItemInfo: itemInfo,
		Error:    err,
	}
}

// ------------------------- END -------------------------
