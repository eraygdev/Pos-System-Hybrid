package main

import "fmt"

func main() {
	Restaurants = make(map[int]*Restaurant)
	err := initRestaurants()

	// Hata kontrolü (Dosya bulunamadıysa burada kalır)
	if err != nil {
		fmt.Printf("❌ Yükleme Sırasında Hata: %v\n", err)
		return
	}

	// Değişkenlerden verileri çekiyoruz
	resCount := len(Restaurants)
	tableCount := 0
	menuCount := 0

	for _, r := range Restaurants {
		tableCount += len(r.Tables)
		menuCount += len(r.Menu)
	}

	// Jilet gibi rapor
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
