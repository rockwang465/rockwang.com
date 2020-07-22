package main

import "fmt"

func main() {
	nums := []int{2, 7, 2, 7, 11, 15}
	targe := 9

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			//if i != j && i >j{
			if i > j {
				sumNum := nums[i] + nums[j]
				if sumNum == targe {
					fmt.Println(i, j)
				}
			}
		}
	}
}
