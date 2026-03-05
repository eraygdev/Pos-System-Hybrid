package main

type RestaurantManager interface {
	// --- TABLE MANAGEMENT ---

	OpenTable(tableID int, waiterID int, guestCount int) error // Initialize Table
	ClearTable(tableID int) error                              // Reset Table
	MoveTable(fromID, toID int) error                          // Transfer Table

	// --- ORDER MANAGEMENT ---

	AddOrder(tableID int, orderID int) error           // Append Order
	RemoveOrder(tableID int, orderIndex int) error     // Delete Order
	ShiftOrder(fromID, toID int, orderIndex int) error // Transfer Order

	// --- FINANCIALS ---

	UpdateManualTotal(tableID int, amount float64) error // Adjust Balance
	CalculateBill(tableID int) (int, error)              // Calculate Invoice
}
