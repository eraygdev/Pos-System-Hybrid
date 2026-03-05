package main

import (
	"fmt"
)

func lockBoth(t1, t2 *Table) (unlock func()) {
	if t1 == t2 {
		t1.mu.Lock()
		return t1.mu.Unlock
	}

	if t1.Number < t2.Number {
		t1.mu.Lock()
		t2.mu.Lock()
	} else {
		t2.mu.Lock()
		t1.mu.Lock()
	}

	return func() {
		t1.mu.Unlock()
		t2.mu.Unlock()
	}
}

func (t *Table) resetUncontrolled_unsafe() {
	t.WaiterID = Default_WaiterID
	t.IsBusy = Default_IsBusy
	t.GuestCount = Default_GuestCount
	t.Orders = Default_Orders
	t.Total = Default_Total
}

func (t *Table) isBusy_unsafe() bool {
	return t.IsBusy
}

func (t *Table) openTable(waiter int, guestCount int) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.IsBusy {
		return fmt.Errorf("Masa %d zaten dolu!", t.Number)
	}

	t.WaiterID = waiter
	t.IsBusy = true
	t.GuestCount = guestCount

	t.Orders = Default_Orders
	return nil
}

func (t *Table) reset_unsafe() error {
	if !t.isBusy_unsafe() {
		return fmt.Errorf("Masa zaten boş!")
	}

	t.resetUncontrolled_unsafe()
	return nil
}

func moveTable(fromTable, toTable *Table) error {
	unlock := lockBoth(fromTable, toTable)
	defer unlock()

	if !fromTable.isBusy_unsafe() {
		return fmt.Errorf("Taşımaya çalıştığınız masa boş!")
	}

	if toTable.isBusy_unsafe() {
		return fmt.Errorf("Masa %d dolu! Taşıma işlemi gerçekleşmedi", toTable.Number)
	}

	toTable.WaiterID = fromTable.WaiterID
	toTable.IsBusy = true
	toTable.GuestCount = fromTable.GuestCount
	toTable.Orders = append([]int(nil), fromTable.Orders...)
	toTable.Total = fromTable.Total

	fromTable.resetUncontrolled_unsafe()
	return nil
}

func (t *Table) addOrder(OrderID int) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.isBusy_unsafe() {
		return fmt.Errorf("Masa %d boş olduğu için sipariş girilemez!", t.Number)
	}

	t.Orders = append(t.Orders, OrderID)
	return nil
}

func (t *Table) removeOrder(OrderIndex int) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.isBusy_unsafe() {
		return fmt.Errorf("Masa %d boş olduğu için sipariş girilemez!", t.Number)
	}

	if OrderIndex < 0 || OrderIndex >= len(t.Orders) {
		return fmt.Errorf("Geçersiz sipariş index'i!")
	}

	t.Orders = append(t.Orders[:OrderIndex], t.Orders[OrderIndex+1:]...)
	return nil
}

func moveOrder(fromTable, toTable *Table, orderIndex int) error {
	unlock := lockBoth(fromTable, toTable)
	defer unlock()

	if !fromTable.isBusy_unsafe() {
		return fmt.Errorf("Masa %d boş, taşıma yapılamaz!", fromTable.Number)
	}

	if !toTable.isBusy_unsafe() {
		return fmt.Errorf("Hedef masa (%d) henüz açılmadığı için taşıma yapılamaz!", toTable.Number)
	}

	if orderIndex < 0 || orderIndex >= len(fromTable.Orders) {
		return fmt.Errorf("Geçersiz siparış index'i!")
	}

	orderToMove := fromTable.Orders[orderIndex]

	toTable.Orders = append(toTable.Orders, orderToMove)
	fromTable.Orders = append(fromTable.Orders[:orderIndex], fromTable.Orders[orderIndex+1:]...)
	return nil
}

func (t *Table) updateTotal(Price float64) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.isBusy_unsafe() {
		return fmt.Errorf("Masa %d boş olduğu için ödeme eklenemez!", t.Number)
	}

	t.Total += Price
	return nil
}

func (t *Table) getBill(r *Restaurant) (int, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.isBusy_unsafe() {
		return 0, fmt.Errorf("Masa %d boş!", t.Number)
	}

	var total int
	for _, ID := range t.Orders {
		if item, ok := r.Menu[ID]; ok {
			total += item.Price
		} else {
			return 0, fmt.Errorf("Sipariş listesindeki %d ID'li ürün menüde bulunamadı", ID)
		}
	}
	return total, nil
}
