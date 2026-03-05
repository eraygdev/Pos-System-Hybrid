package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var Restaurants map[int]*Restaurant

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

func (r *Restaurant) loadMenu() error {
	data, err := os.ReadFile("menu.json")
	if err != nil {
		return err
	}

	var captured map[int]map[int]*MenuItem
	err1 := json.Unmarshal(data, &captured)
	if err1 != nil {
		return err1
	}

	for _, items := range captured {
		if len(items) > 0 {
			for itemid, item := range items {
				if _, ok := r.Menu[itemid]; ok {
					r.Menu[itemid].ID = itemid
					r.Menu[itemid].Name = item.Name
					r.Menu[itemid].Category = item.Category
					r.Menu[itemid].PrepTime = item.PrepTime
					r.Menu[itemid].Price = item.Price
					r.Menu[itemid].Stock = item.Stock
					r.Menu[itemid].IsActive = item.IsActive
				}
			}
		}
	}
	return nil
}

func initRestaurants() error {
	data, err := os.ReadFile("restaurants_data.json")
	if err != nil {
		return err
	}

	var captured map[int][]Restaurant
	err1 := json.Unmarshal(data, &captured)
	if err1 != nil {
		return err1
	}

	for id, list := range captured {
		if len(list) > 0 {
			if r, ok := Restaurants[id]; ok {
				// **Place restaurant data**
				r.Name = list[id].Name
				r.ID = list[id].ID
				r.Capacity = list[id].Capacity

				// **Load child elements**
				if err := r.loadTables(); err != nil {
					return err
				}
				if err := r.loadMenu(); err != nil {
					return err
				}
				return nil
			}
		}
	}
	return nil
}
