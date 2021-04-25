package main

import "fmt"

func main() {
	test()
}

type ERROR interface {
	Error() string
}

func Error() interface{} {
	return "Error func has an error"
}

func test() {
	//content := ""
	//errCode := 404

	//err := fmt.Errorf("has an error")
	err := Error()
	if err != "" {
		switch errType := err.(type) {
		case string:
			if errType == "Error func has an error" {
				fmt.Println("string error: ", errType)
				break // break 跳出当前switch
			}
			fmt.Println("not Error func has an error")
			panic(err)
		default:
			fmt.Println("in default")
			panic(err)
		}
	}

}
