package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/glebarez/go-sqlite"
)

func loadFromJSON(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Restaurants)
	if err != nil {
		return err
	}

	fmt.Println("📂 JSON verisi başarıyla belleğe yüklendi!")
	return nil
}
func MigrationJsonToSql(db *sql.DB, restaurants map[int]*Restaurant) error {
	for _, r := range Restaurants {
		res, err := db.Exec("INSERT INTO restaurants (name, capacity) VALUES (?, ?)", r.Name, r.Capacity)
		if err != nil {
			return err
		}
		resID, _ := res.LastInsertId()

		for _, t := range r.Tables {
			_, err := db.Exec("INSERT INTO tables (restaurant_id, number, waiter_id, is_busy, guest_count, total) VALUES (?, ?, ?, ?, ?, ?)",
				resID, t.Number, t.WaiterID, t.IsBusy, t.GuestCount, t.Total)
			if err != nil {
				return err
			}
		}

		for _, m := range r.Menu {
			_, err := db.Exec("INSERT INTO menus (restaurant_id, name, category, prep_time, price, stock, is_active) VALUES (?, ?, ?, ?, ?, ?, ?)",
				resID, m.Name, m.Category, m.PrepTime, m.Price, m.Stock, m.IsActive)
			if err != nil {
				return err
			}
		}
	}

	fmt.Println("✅ Json'dan SQL'e veri transferi başarılı!")
	return nil
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./pos.db")
	if err != nil {
		return nil, err
	}

	createTables := `
	CREATE TABLE IF NOT EXISTS restaurants (
		id INTEGER PRIMARY KEY,
		name TEXT,
		capacity INTEGER,
		state BOOLEAN DEFAULT 0
	);
	
	CREATE TABLE IF NOT EXISTS tables (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		restaurant_id INTEGER,
		number INTEGER,
		waiter_id INTEGER DEFAULT 1,
		is_busy BOOLEAN DEFAULT 0,
		guest_count INTEGER DEFAULT 1,
		total REAL,
		paid_value INTEGER DEFAULT 0,
		FOREIGN KEY(restaurant_id) REFERENCES restaurants(id)
	);
	
	CREATE TABLE IF NOT EXISTS menus (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		restaurant_id INTEGER,
		name TEXT,
		category INTEGER DEFAULT 1,
		prep_time INTEGER DEFAULT 15,
		price REAL,
		stock INTEGER DEFAULT 0,
		is_active BOOLEAN DEFAULT 1,
		FOREIGN KEY(restaurant_id) REFERENCES restaurants(id)
	);
	
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		table_id INTEGER,
		menu_item_id INTEGER,
		quantity INTEGER DEFAULT 1,
		order_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(table_id) REFERENCES tables(id),
		FOREIGN KEY(menu_item_id) REFERENCES menus(id)
	);`

	_, err = db.Exec(createTables)
	return db, err
}
