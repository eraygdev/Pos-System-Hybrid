package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Table struct {
	ID     int       `json:"id"`
	IsFull bool      `json:"is_full"`
	Orders []Product `json:"orders"`
}

func main() {
	
}

