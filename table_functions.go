package main

import "fmt"

func isTableBusy(t *Table) (bool, error) {
	if t.GuestCount <= 0 {
		return false, fmt.Errorf("Masa %d için misafir sayısı geçersiz!", t.Number)
	}

	return t.Busy, nil
}

func openTable() {

}

func resetTable(thetable *Table) error {
	busy, err := isTableBusy(thetable)

	if err != nil {
		return err
	}

	if !busy {
		return fmt.Errorf("Masa zaten boş!")
	}

	*thetable = Table{Number: thetable.Number}
	return nil
}

func addOrder(t *Table, OrderID int) error {
	busy, err := isTableBusy(t)
	if err != nil {
		return err
	}

	if !busy {
		return fmt.Errorf("Masa %d boş olduğu için sipariş girilemez!", t.Number)
	}

	t.Orders = append(t.Orders, OrderID)
	return nil
}

func removeOrder(t *Table, OrderIndex int) error {
	busy, err := isTableBusy(t)
	if err != nil {
		return err
	}

	if !busy {
		return fmt.Errorf("Masa %d boş olduğu için sipariş girilemez!", t.Number)
	}

	t.Orders = append(t.Orders[:OrderIndex], t.Orders[OrderIndex+1:]...)
	return nil
}

func moveOrder(fromTable *Table, toTable *Table, orderIndex int) error {
	busy, err := isTableBusy(toTable)
	if err != nil {
		return err
	}

	if !busy {
		return fmt.Errorf("Masa %d boş olduğu için buraya taşıma yapılamaz!", toTable.Number)
	}

	if orderIndex < 0 || orderIndex >= len(fromTable.Orders) {
		return fmt.Errorf("Geçersiz siparış index'i!")
	}

	orderToMove := fromTable.Orders[orderIndex]

	toTable.Orders = append(toTable.Orders, orderToMove)
	fromTable.Orders = append(fromTable.Orders[:orderIndex], fromTable.Orders[orderIndex+1:]...)
	return nil
}

func updateTableTotal(t *Table, Price float64) error {
	busy, err := isTableBusy(t)

	if err != nil {
		return err
	}

	if !busy {
		return fmt.Errorf("Masa %d boş olduğu için ödeme eklenemez!", t.Number)
	}

	t.Total += Price
	return nil
}

func getTableBill(t *Table) (int, error) {
	busy, err := isTableBusy(t)
	if err != nil {
		return 0, err
	}

	if !busy {
		return 0, fmt.Errorf("Masa %d boş!", t.Number)
	}

	var total int
	for _, ID := range t.Orders {
		if item, ok := Menu[ID]; ok {
			total += item.Price
		} else {
			return 0, fmt.Errorf("Sipariş listesindeki %d ID'li ürün menüde bulunamadı", ID)
		}
	}
	return total, nil
}
