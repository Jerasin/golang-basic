package pkg

import (
	"fmt"
	"learn/helper"

	"github.com/goforj/godump"
)

type Test struct {
	Name string
	Age  int
}

func forLoopEx1() {
	for {
		fmt.Println("Infinite Loop")
	}
}

func forLoopEx2() {
	arrNum1 := []int{1, 2, 3, 4, 5, 6}
	for index, value := range arrNum1 {
		fmt.Printf("index = %d value = %+v \n", index, value)
	}
}

func forLoopEx3() {
	isExist := 0
	for isExist < 5 {
		fmt.Printf("while loop = %d \n", isExist)
		isExist++
	}
}

func forLoopEx4() {
	arrMap := []map[string]any{
		{"name": "Ali", "age": 20},
		{"name": "Reza", "age": 30},
		{"name": "Sara", "age": 25},
	}

	var testList []Test
	for _, value := range arrMap {
		testList = append(testList, Test{
			Name: value["name"].(string),
			Age:  value["age"].(int),
		})
	}

	godump.Dump(testList)
}

// Convert SliceMap To Struct
func forLoopEx5() {
	arrMap := []map[string]any{
		{"name": "Ali", "age": 20},
		{"name": "Reza", "age": 30},
		{"name": "Sara", "age": 25},
	}

	var testList []Test
	err := helper.ConvertMapToStruct(&testList, arrMap)

	if err != nil {
		fmt.Println("Error converting map to struct:", err)
		return
	}

	var test Test
	err2 := helper.ConvertMapToStruct(&test, arrMap[0])
	if err2 != nil {
		fmt.Println("Error converting map to struct:", err2)
		return
	}
	godump.Dump(test)
}

func forLoopEx6() {
	item := []string{"item1", "item2", "item3"}
	data := helper.AppendSlice(item, "new item")
	godump.Dump(data)
}

func forLoopEx7() {
	item := []string{"item1", "item2", "item3"}
	helper.AppendSlice2(&item, "new item")
	godump.Dump(item)
}

func RunForLoop() {
	// forLoopEx1()
	// forLoopEx2()
	// forLoopEx3()
	// forLoopEx4()
	// forLoopEx6()
	forLoopEx7()
}
