package main

type RestaurantManager interface {
	// TABLE STATUS
	isTableBusy(t *Table) (bool, error) // Meşguliyet kontrolü
	resetTable(t *Table) error          // Masayı sıfırlama

	// ORDER MANAGEMENT
	addOrder(t *Table, id int) error                      // Sipariş ekleme
	removeOrder(t *Table, idx int) error                  // Sipariş silme
	moveOrder(fromTable, toTable *Table, index int) error // Sipariş kaydırma

	// FINANCIALS
	updateTableTotal(t *Table, p float64) error // Manuel ödeme ekleme
	getTableBill(t *Table) (int, error)         // Toplam hesap çıkarma
}
