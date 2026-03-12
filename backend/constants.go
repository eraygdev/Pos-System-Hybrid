package main

const (
	Default_WaiterID   = 0     // 0
	Default_IsBusy     = false // false
	Default_GuestCount = 0     // 0
	Default_Total      = 0     // 0
)

var (
	Default_Orders = []int(nil)              // []int(nil)
	Default_Tables = make(map[int]*Table)    // make(map[int]*Table)
	Default_Menus  = make(map[int]*MenuItem) // make(map[int]*MenuItem)
)
