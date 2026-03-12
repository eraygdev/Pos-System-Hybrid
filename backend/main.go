package main

import (
	"fmt"
)

func main() {
	db, err2 := initDB()
	if err2 != nil {
		fmt.Println(err2)
	}

	Restaurants = make(map[int]*Restaurant)
	err := initRestaurants(db)
	if err != nil {
		fmt.Println(err)
	}

	/*
		err = loadFromJSON("./restaurants_data.json")
		if err != nil {
			fmt.Println("JSON Okuma Hatası:", err)
		}
		err3 := MigrationJsonToSql(db, Restaurants)
		if err3 != nil {
			fmt.Println(err3)
		}
	*/

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
