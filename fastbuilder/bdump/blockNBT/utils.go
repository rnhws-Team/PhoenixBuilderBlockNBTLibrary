package blockNBT

import (
	"fmt"
	"phoenixbuilder/fastbuilder/mcstructure"
	"phoenixbuilder/fastbuilder/types"
	"strings"
)

// 从 SupportBlocksPool 检查这个方块实体是否已被支持。
// 如果尚未被支持，则返回空字符串，否则返回这种方块的类型。
// 以告示牌为例，所有的告示牌都可以写作为 Sign
func checkIfIsEffectiveNBTBlock(blockName string) string {
	value, ok := SupportBlocksPool[blockName]
	if ok {
		return value
	}
	return ""
}

// 将 types.Module 转换为 blockNBT_depends.GeneralBlock
func parseBlockModule(singleBlock *types.Module) (GeneralBlock, error) {
	// init var
	got, err := mcstructure.ParseStringNBT(singleBlock.Block.BlockStates, true)
	if err != nil {
		return GeneralBlock{}, fmt.Errorf("parseBlockModule: Could not parse block states; singleBlock.Block.BlockStates = %#v", singleBlock.Block.BlockStates)
	}
	blockStates, normal := got.(map[string]interface{})
	if !normal {
		return GeneralBlock{}, fmt.Errorf("parseBlockModule: The target block states is not map[string]interface{}; got = %#v", got)
	}
	// get block states
	return GeneralBlock{
		Name:   strings.Replace(strings.ToLower(strings.ReplaceAll(*singleBlock.Block.Name, " ", "")), "minecraft:", "", 1),
		States: blockStates,
		NBT:    singleBlock.NBTMap,
	}, nil
	// return
}

// 取得用于放置目标方块实体的 接口/方法
func getMethod(pack Package) GeneralBlockNBT {
	switch pack.Datas.Type {
	case "CommandBlock":
		return &CommandBlock{Package: &pack, NeedToPlaceBlock: true}
		// 命令方块
	case "Container":
		return &Container{Package: &pack}
		// 容器
	case "Sign":
		return &Sign{Package: &pack}
		// 告示牌
	default:
		return &Default{Package: &pack}
		// 其他尚且未被支持的方块实体
	}
	// 返回值
}
