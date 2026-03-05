package main

import (
	"sync"
)

type Restaurant struct {
	Name     string `json:"name"`
	ID       int    `json:"id"`
	Capacity int    `json:"capacity"`
	Menu     map[int]*MenuItem
	Tables   map[int]*Table
	mu       sync.Mutex
}

type Table struct {
	// --- IDENTITY ---

	Number   int // Table ID
	WaiterID int // Waiter ID

	// --- STATUS ---

	IsBusy     bool // Occupancy status
	GuestCount int  // Customer count

	// --- FINANCIAL ---

	Orders []int   // Item list
	Total  float64 // Bill total

	// --- TECHNICAL ---

	mu sync.Mutex // Thread safety
}

type MenuItem struct {
	// --- IDENTITY ---

	ID   int    `json:"id"`
	Name string `json:"name"`

	// --- DETAILS ---

	Category int `json:"category"`
	PrepTime int `json:"prep_time"`
	Price    int `json:"price"`

	// --- STATUS ---

	Stock    int  `json:"stock"`
	IsActive bool `json:"is_active"`
}
