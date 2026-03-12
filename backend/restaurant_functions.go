package main

import (
	"database/sql"
	"fmt"
)

var Restaurants map[int]*Restaurant

func (r *Restaurant) loadMenus(db *sql.DB) error {
	if r.Menu == nil {
		r.Menu = make(map[int]*MenuItem)
	}

	rows, err := db.Query("SELECT id, restaurant_id, name, category, prep_time, price, stock, is_active FROM menus WHERE restaurant_id = ? ORDER BY category ASC, name ASC", r.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		m := &MenuItem{}
		if err := rows.Scan(&m.ID, &m.RestaurantID, &m.Name, &m.Category, &m.PrepTime, &m.Price, &m.Stock, &m.IsActive); err != nil {
			return err
		}
		r.Menu[m.ID] = m
	}

	return rows.Err()
}

func (r *Restaurant) loadTables() error {

	r.Tables = make(map[int]*Table)
	TableAmount := r.Capacity

	if TableAmount <= 0 {
		return fmt.Errorf("Masa sayısı geçersiz!")
	}

	for i := 1; i <= TableAmount; i++ {
		t := &Table{Number: i}
		t.resetUncontrolled_unsafe()

		r.Tables[i] = t
	}

	return nil
}

func initRestaurants(db *sql.DB) error {
	rows, err := db.Query("SELECT id, name, capacity FROM restaurants ORDER BY id")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		r := &Restaurant{}
		if err := rows.Scan(&r.ID, &r.Name, &r.Capacity); err != nil {
			return err
		}
		Restaurants[r.ID] = r
	}

	for _, r := range Restaurants {
		if err := r.loadMenus(db); err != nil {
			return err
		}
		if err := r.loadTables(); err != nil {
			return err
		}
	}

	return rows.Err()
}
