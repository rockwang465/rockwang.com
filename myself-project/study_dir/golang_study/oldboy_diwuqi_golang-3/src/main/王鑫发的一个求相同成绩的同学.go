package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileObj, err := os.Open("./score.txt")
	if err != nil {
		fmt.Println("Error : open file failed , err = ", err)
		return
	}
	buf := make([]byte, 10*1024)
	n, err := fileObj.Read(buf)
	if err != nil {
		fmt.Println("Error : file read failed, err = ", err)
	}
	res := string(buf[:n])
	//fmt.Println(res)

	scoreList := strings.Split(res, "\r\n")
	fmt.Printf("%#v\n", scoreList) // [李一 93 王二 83 王三 93 李四 60 王五 75 马六 61 孙七 75 刘八 75]

	var sameScore = make(map[string]int)

	for i := 0; i < len(scoreList); i++ {
		for j := 0; j < len(scoreList); j++ {
			if j > i {
				// 单个slice再split
				subListI := strings.Split(scoreList[i], " ")
				subListJ := strings.Split(scoreList[j], " ")
				//fmt.Println(subListI, subListJ)
				//fmt.Println(subListI[len(subListI)-1], subListJ[len(subListJ)-1])
				//fmt.Printf("%T, %T\n", subListI[len(subListI)-1], subListJ[len(subListJ)-1])
				intI, err := strconv.Atoi(subListI[len(subListI)-1])
				if err != nil {
					fmt.Println("Error : strconv.Atoi 1 err = ", err)
					return
				}
				intJ, err := strconv.Atoi(subListJ[len(subListJ)-1])
				if err != nil {
					fmt.Println("Error : strconv.Atoi 2 err = ", err)
					return
				}

				if intI == intJ {
					//fmt.Println(intI, intJ)
					//fmt.Printf("%v == %v\n", scoreList[i], scoreList[j])
					sameScore[subListI[0]] = intI
					sameScore[subListJ[0]] = intJ
				}
			}
		}
	}
	fmt.Println(sameScore)
}
