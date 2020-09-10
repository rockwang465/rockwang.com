package common

import "fmt"

func TransferUint(id interface{}) (uint, error) {
	switch v := id.(type) {
	case int:
		fmt.Println("int: v, id:", v, id)
		return id.(uint), nil
	case string:
		fmt.Println("string: v, id:", v, id)
		return id.(uint), nil
	default:
		return 0, fmt.Errorf("不属于可转换类型")
	}
}
