package pkg

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Location struct {
	Lat float64
	Lng float64
}

type Address struct {
	Province string
	Amphure  string
	Location Location
}

type People struct {
	Name    string
	Age     uint
	Address Address
}

type PeopleV2 struct {
	Name    string
	Age     uint
	Address *Address
}

func (p People) setName(name string) {
	p.Name = name
}

func (p *People) setNameV2(name string) {
	p.Name = name
}

func structCase1() {
	address := Address{
		Province: "กรุงเทพ",
		Amphure:  "ลาดพร้าว",
	}

	people := People{
		Name:    "Jin",
		Age:     30,
		Address: address,
	}

	fmt.Printf("people1 = %+v \n", people)

	address.Province = "เชียงใหม่"

	fmt.Printf("people2 = %+v \n", people)
	fmt.Printf("address = %+v \n", address)
}

func structCase2() {
	address := Address{
		Province: "กรุงเทพ",
		Amphure:  "ลาดพร้าว",
	}

	people := PeopleV2{
		Name:    "Jin",
		Age:     30,
		Address: &address,
	}

	fmt.Printf("people1 = %+v \n", people)

	address.Province = "เชียงใหม่"

	fmt.Printf("people2 = %+v \n", people)
	jsonData, _ := json.Marshal(people)
	fmt.Printf("people json = %s \n", string(jsonData))
	fmt.Printf("address = %+v \n", address)
}

func structCase3() {
	address := Address{
		Province: "กรุงเทพ",
		Amphure:  "ลาดพร้าว",
	}

	people := People{
		Name:    "Jin",
		Age:     30,
		Address: address,
	}

	fmt.Printf("people1 = %+v \n", people)

	address.Province = "เชียงใหม่"

	fmt.Printf("people2 = %+v \n", people)
	fmt.Printf("address = %+v \n", address)
}

func structCase4() {
	address := Address{
		Province: "กรุงเทพ",
		Amphure:  "ลาดพร้าว",
	}

	people := People{
		Name:    "Jin",
		Age:     30,
		Address: address,
	}

	fmt.Printf("people 1 = %+v \n", people)
	people.setName("Jame")
	fmt.Printf("people 2 = %+v \n", people)
	people.setNameV2("Jame")
	fmt.Printf("people 3 = %+v \n", people)
}

func structCase5() {
	location := Location{
		Lat: 1.1,
		Lng: 2.2,
	}

	address := Address{
		Province: "กรุงเทพ",
		Amphure:  "ลาดพร้าว",
		Location: location,
	}

	people := People{
		Name:    "Jin",
		Age:     30,
		Address: address,
	}

	t := reflect.TypeOf(people)
	v := reflect.ValueOf(people)

	for item := 0; item < t.NumField(); item++ {
		fmt.Println("item:", item)
		fmt.Println("Field name:", t.Field(item).Name)
		fmt.Println("Field type:", t.Field(item).Type)
		fmt.Println("Field value:", v.Field(item))
		fmt.Println("------")
	}

	// godump.Dump(people)

	// output := godump.DumpStr(people)
	// fmt.Println("str", output)

}

func RunStruct() {
	// structCase1()
	// structCase2()
	// structCase3()
	// structCase4()
	structCase5()
}
