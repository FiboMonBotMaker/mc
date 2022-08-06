package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Car struct {
	Wheels int `json:"wheels"`
}

type Truck struct {
	Car
	Tons int `json:"tons"`
}

func main() {
	truck := Truck{}
	b, _ := os.ReadFile("tmp.json")
	json.Unmarshal(b, &truck)
	fmt.Printf("Wheels: %+v\n", truck.Wheels)
	fmt.Printf("Tons: %v\n", truck.Tons)
}