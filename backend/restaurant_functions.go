package main

import (
	"encoding/json"
	"os"
)

var Restaurants map[int]*Restaurant

func (r *Restaurant) loadMenu() error {
	data, err := os.ReadFile("menu.json")
	if err != nil {
		return err
	}

	var captured map[int][]MenuItem
	error := json.Unmarshal(data, &captured)
	if error != nil {
		return err
	}

}

func (r *Restaurant) loadTables() {
	r.Tables = make(map[int]*Table)

	TableAmount := r.Capacity
	for i := 1; i <= TableAmount; i++ {
		t := &Table{Number: i}
		t.resetUncontrolled_unsafe()

		r.Tables[i] = t
	}
}

func init() {

	data, err := os.ReadFile("restaurants_data.json")
	println(err)

	var captured map[int][]Restaurant
	err2 := json.Unmarshal(data, &captured)
	println(err2)

	for id, list := range captured {
		if len(list) > 0 {
			if r, ok := Restaurants[id]; ok {
				// **Place restaurant data**
				r.Name = list[id].Name
				r.ID = list[id].ID
				r.Capacity = list[id].Capacity

				// **Load child elements**
				r.loadTables()
				r.loadMenu()
			}
		}
	}
}
