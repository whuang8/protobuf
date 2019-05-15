package main

import (
	"github.com/whuang8/protobuf/src/simple"
	"fmt"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simplepb.SimpleMessage {
		Id: 12345,
		IsSimple: true,
		Name: "My simple message",
		SampleList: []int32{1,2,3,4,5},
	}

	fmt.Println(sm)
}
