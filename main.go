package main

import (
	"encoding/json"
	"fmt"
	"phoenixbuilder/minecraft/protocol/packet"
)

func main() {
	var pkt packet.InventoryTransaction
	m := `{
		"LegacyRequestID": 0,
		"LegacySetItemSlots": null,
		"Actions": [],
		"TransactionData": {
			"LegacyRequestID": 0,
			"LegacySetItemSlots": null,
			"Actions": null,
			"ActionType": 0,
			"BlockPosition": [
				23,
				23,
				23
			],
			"BlockFace": 0,
			"HotBarSlot": 3,
			"HeldItem": {
				"StackNetworkID": 0,
				"Stack": {
					"NetworkID": 0,
					"MetadataValue": 0,
					"BlockRuntimeID": 0,
					"Count": 0,
					"NBTData": null,
					"CanBePlacedOn": null,
					"CanBreak": null,
					"HasNetworkID": false
				}
			},
			"Position": [
				0,
				0,
				0
			],
			"ClickedPosition": [
				0,
				0,
				0
			],
			"BlockRuntimeID": 0
		}
	}`
	err := json.Unmarshal([]byte(m), &pkt)
	fmt.Printf("%#v\n", pkt)
	fmt.Println(err)
}
