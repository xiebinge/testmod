package testmod

import (
	"fmt"
)

func Hi(name string) string { //go用到参数的地方，要考虑数据类型
	return fmt.Sprintf("hi,%s", name)
}
