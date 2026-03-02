package main

import (
	"encoding/json"
	"fmt"
)

// Product: Kütüphanedeki Book yapısı gibi
type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Table: Masalarımızı temsil eder
type Table struct {
	ID     int       `json:"id"`
	IsFull bool      `json:"is_full"`
	Orders []Product `json:"orders"`
}

func main() {
	fmt.Println("eraygdev POS Sistemi - Backend v0.1")

	// Örnek bir masa oluşturalım
	masa1 := Table{
		ID:     1,
		IsFull: true,
		Orders: []Product{
			{Name: "Pizza", Price: 350.0},
			{Name: "Ayran", Price: 40.0},
		},
	}

	// JSON'a dönüştürüp ekrana basalım (C'de bunu yapmak çok zordu!)
	jsonData, _ := json.MarshalIndent(masa1, "", "  ")
	fmt.Println(string(jsonData))
}
