package pkg

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/goforj/godump"
)

var Name string
var Age uint = 1
var Number int = -1
var Boolean bool = false
var Map = map[string]any{
	"Name": "Jame",
}
var pi float64 = 7.41

const (
	GET    = "get"
	POST   = "post"
	PUT    = "put"
	DELETE = "delete"
)

func sliceCase1() {
	numbers := []uint{1, 2, 3, 4}
	sliceMap := []map[string]any{{
		"name": "Jin",
		"age":  10,
	},
		{
			"name": "Jame",
			"age":  20,
		},
	}

	fmt.Println("numbers1", numbers)
	numbers = append(numbers, 6)

	fmt.Println("numbers2", numbers)

	fmt.Println("sliceMap", sliceMap)
}

func sliceCase2() {
	type People struct {
		Name string
		Age  int16
	}
	sliceMap := []map[string]any{{
		"name": "Jin",
		"age":  10,
	},
		{
			"name": "Jame",
			"age":  20,
		},
	}

	x, err := json.Marshal(sliceMap)

	if err != nil {
		panic(err)
	}
	var people []People

	if err := json.Unmarshal(x, &people); err != nil {
		panic(err)
	}

	godump.Dump((people))

}

func mapCase1() {
	newMap := map[string]any{
		"name": "Jin",
	}

	for i := range newMap {
		fmt.Printf("key = %s , value = %v \n", i, newMap[i])
	}
}

func convertStrToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func convertStrToFloat(value int) float64 {
	return float64(value)
}

func removeSymbol(str string) {
	newStr := strings.Split(str, ",")
	joinStr := strings.Join(newStr, " ")
	replaceStr := strings.Replace(str, ",", "", 1)
	replaceAllStr := strings.ReplaceAll(str, ",", "")
	fmt.Printf("string Split = %s \n", newStr)
	fmt.Printf("stringsJoin  = %s \n", joinStr)
	fmt.Printf("replaceStr  = %s \n", replaceStr)
	fmt.Printf("replaceAllStr  = %s \n", replaceAllStr)
}

func Run() {
	// address := Address{
	// 	Province: "กรุงเทพ",
	// 	Amphure:  "ลาดพร้าว",
	// }
	// fmt.Printf("Address 1 = %+v \n", address)

	// address.Province = "เชียงใหม่"

	// fmt.Printf("Address 2 = %+v \n", address)

	// fmt.Printf("Map = %+v \n", Map)

	// fmt.Printf("Print value = %v \n", GET)
	// fmt.Printf("Print struct = %+v \n", address)
	// fmt.Printf("Print Type = %T \n", GET)
	// fmt.Printf("Type Name = %s \n", GET)

	// msg := fmt.Sprintf("Address = %+v, GET = %s", address, GET)
	// fmt.Printf("Sprintf = %+v \n", msg)

	// sliceCase1()
	// mapCase1()

	// num1 := "1"
	// value, err := convertStrToInt(num1)
	// if err != nil {
	// 	panic("error")
	// }
	// fmt.Printf("value = %+v \n", value)

	// const str = "Hello,World"
	// removeSymbol(str)
	sliceCase2()
	// fmt.Printf("float64 = %f \n", convertStrToFloat(1234))
}
