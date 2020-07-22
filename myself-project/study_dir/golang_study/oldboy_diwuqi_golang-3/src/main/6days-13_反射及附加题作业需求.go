package main

import (
	"fmt"
	"reflect"
)

//type person struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}
//
//func main() {
//	var p person
//	jsonP := `{"name": "rock", "age": 18}`
//	json.Unmarshal([]byte(jsonP), &p)
//	fmt.Println(p.Name, p.Age)
//}

type cat struct {
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("Type : %v\n", v)
	fmt.Printf("Type name: %v, Type kind: %v\n", v.Name(), v.Kind())
}
func main() {
	var a float32 = 3.14
	reflectType(a) // Type : float32  , 而Type name: float32, Type kind: float32

	var c = cat{}
	reflectType(c) // Type : main.cat , 但是 Type name: cat, Type kind: struct
}
