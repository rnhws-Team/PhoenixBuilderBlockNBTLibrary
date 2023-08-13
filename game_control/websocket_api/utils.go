package external

import "fmt"

// 检测 err 是否非空。
// 当非空时尝试惊慌程序
func TestError(err error) {
	if err != nil {
		panic(fmt.Sprintf("TestError: %v", err))
	}
}
