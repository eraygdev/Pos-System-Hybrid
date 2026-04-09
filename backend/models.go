package main

import (
	"sync"
)

type Restaurant struct {
	Name     string `json:"name"`
	ID       int    `json:"id"`
	STATE    bool
	Capacity int               `json:"capacity"`
	Menu     map[int]*MenuItem `json:"menu"`
	Tables   map[int]*Table
	mu       sync.Mutex
}

type Table struct {
	// --- IDENTITY ---

	ID           int
	RestaurantID int
	Number       int // Table ID
	WaiterID     int // Waiter ID

	// --- STATUS ---

	IsBusy     bool // Occupancy status
	GuestCount int  // Customer count

	// --- FINANCIAL ---

	Orders []int   // Item list
	Total  float64 // Bill total

	// --- TECHNICAL ---

	mu sync.Mutex // Thread safety
}

type ORDER struct {
	// --- IDENTITY ---

	ID int
}

type MenuItem struct {
	// --- IDENTITY ---

	ID           int `json:"id"`
	RestaurantID int
	Name         string `json:"name"`

	// --- DETAILS ---

	Category int     `json:"category"`
	PrepTime int     `json:"prep_time"`
	Price    float64 `json:"price"`

	// --- STATUS ---

	Stock    int  `json:"stock"`
	IsActive bool `json:"is_active"`
}

type MenuOption struct {
	Label  string
	Action func() error
}
