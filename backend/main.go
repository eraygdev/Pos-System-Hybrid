package main

import (
	"fmt"
	"pos_system/terminal_test"
)

func main() {
	terminal_test.InitCmd()

	db, err := initDB()
	if err != nil {
		fmt.Println(err)
	}

	/* Migration JSON --> sql
	    if err := loadFromJSON("./restaurants_data.json"); err != nil {
			fmt.Println("JSON Okuma Hatası:", err)
		}
		if err := MigrationJsonToSql(db, Restaurants); err != nil {
			fmt.Println(err)
		} */

	Restaurants = make(map[int]*Restaurant)
	if err := initRestaurants(db); err != nil {
		fmt.Println(err)
	}

	resCount := len(Restaurants)
	tableCount := 0
	menuCount := 0

	for _, r := range Restaurants {
		tableCount += len(r.Tables)
		menuCount += len(r.Menu)
	}

	fmt.Println("==========================================")
	fmt.Printf("🏠 Aktif Restoran : %d\n", resCount)
	fmt.Printf("🪑 Toplam Masa    : %d\n", tableCount)
	fmt.Printf("🍔 Toplam Yemek   : %d\n", menuCount)
	fmt.Println("==========================================")

	if menuCount == 0 {
		fmt.Println("⚠️ Uyarı: Restoranlar yüklendi ama menü boş!")
	} else {
		fmt.Println("🚀 Sistem kullanıma hazır!")
	}

}
