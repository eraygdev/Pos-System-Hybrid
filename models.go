package main

type Table struct {
	Number     int
	Busy       bool
	GuestCount int
	Orders     []int
	Total      float64
	Waiter     int
}

type MenuItem struct {
	ID       int
	Name     string
	Price    int
	Category int
	PrepTime int
	Stock    int
	isActive bool
}

var RestaurantTables []Table   // The tables in the restaurant
var Menu = map[int]*MenuItem{} // Menu
