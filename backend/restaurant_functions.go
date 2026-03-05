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

func initRestaurants() error {
	resdata, err := os.ReadFile("restaurants_data.json")
	if err != nil {
		return err
	}
	menudata, err0 := os.ReadFile("menus_data.json")
	if err0 != nil {
		return err0
	}

	var captured map[int]*Restaurant
	err1 := json.Unmarshal(resdata, &captured)
	if err1 != nil {
		return err1
	}
	var capturedmenu map[int]map[int]*MenuItem
	err2 := json.Unmarshal(menudata, &capturedmenu)
	if err2 != nil {
		return err2
	}

	for id, rData := range captured {
		if r, ok := Restaurants[id]; ok {
			// **Place restaurant data**
			r.Name = rData.Name
			r.ID = rData.ID
			r.Capacity = rData.Capacity

			// **Load child elements**
			if err := r.loadTables(); err != nil {
				return err
			}
		}
	}

	for resid, items := range capturedmenu {
		if r, ok := Restaurants[resid]; ok {
			for itemid, item := range items {
				item.ID = itemid
				r.Menu[itemid] = item
			}
		}
	}
	return nil
}
