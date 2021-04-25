package main

import (
	"fmt"
	"strings"
)

func main()  {
	s := "http://www.baidu.com/v1/api/"

	s1 := strings.TrimSuffix(s,"/")
	fmt.Println(s1)
	Versions("senseguard-tools")
}

const namedChartVersionList = "%s/api/charts/%s"

// Returns a named chart's version list.
func Versions(name string){
	//var out []*ChartVersion

	addr := "http://10.151.3.75:8080"
	uri := fmt.Sprintf(namedChartVersionList, addr, name)
	//err := c.get(uri, &out)
	//return out, err
	fmt.Println(uri)
}