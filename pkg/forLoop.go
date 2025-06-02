package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/goforj/godump"
)

type Test struct {
	Name string
	Age  int
}

func ConvertMapToStruct(structType any, data any) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &structType)
	if err != nil {
		return err
	}

	return nil

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
		fmt.Printf("isExist = %d \n", isExist)
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

func forLoopEx5() {
	arrMap := []map[string]any{
		{"name": "Ali", "age": 20},
		{"name": "Reza", "age": 30},
		{"name": "Sara", "age": 25},
	}

	// แปลง JSON ไปเป็น []Test
	var testList []Test
	err := ConvertMapToStruct(&testList, arrMap)

	if err != nil {
		fmt.Println("Error converting map to struct:", err)
		return
	}

	var test Test
	err2 := ConvertMapToStruct(&test, arrMap[0])
	if err2 != nil {
		fmt.Println("Error converting map to struct:", err2)
		return
	}
	godump.Dump(test)
}

func RunForLoop() {
	// forLoopEx1()
	// forLoopEx2()
	// forLoopEx3()
	// forLoopEx4()
	forLoopEx5()
}
