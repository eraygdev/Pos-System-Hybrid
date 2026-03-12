package main

import (
	"sync"
)

type Restaurant struct {
	Name     string
	ID       int
	Capacity int
	Menu     map[int]*MenuItem
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

	ID           int
	RestaurantID int
	Name         string

	// --- DETAILS ---

	Category int
	PrepTime int
	Price    int

	// --- STATUS ---

	Stock    int
	IsActive bool
}
