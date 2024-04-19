package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Produk struct {
	ID       uint   `gorm:"primaryKey,autoincrement,type:serial;column:id"`
	Nama     string `gorm:"type:varchar(255);column:nama"`
	Kategori string `gorm:"type:varchar(50);column:kategori"`
	Harga    int    `gorm:"type:int;column:harga"`
}

func (Produk) TableName() string {
	return "produk"
}

func main() {
	// 1 - menghubungkan golang dengan postgres
	connURI := "postgresql://postgres:password@localhost:5432/database?sslmode=disable"
	db, err := gorm.Open(postgres.Open(connURI), &gorm.Config{
		SkipDefaultTransaction: true, // mengatur agar transaction tidak digunakan otomatis
	})
	if err != nil {
		fmt.Printf("Gagal menghubungkan ke database: %v\n", err)
		os.Exit(1)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	fmt.Println("Berhasil terhubung ke database")

	// 2 - melakukan migrasi (membuat tabel)
	db.AutoMigrate(&Produk{})

	// 3 - menambahkan data
	produk := Produk{Nama: "Kertas A4", Kategori: "Kertas", Harga: 35000}
	result := db.Create(&produk)
	if result.Error != nil {
		fmt.Printf("Gagal mengisi tabel: %v\n", result.Error)
		os.Exit(1)
	}

	fmt.Println("Data berhasil ditambahkan")

	// 4 - mengambil 1 data

	// dengan ID
	// var dataProduk Produk
	// db.First(&dataProduk, 1)

	// dengan WHERE
	dataProduk := Produk{
		ID: 1,
	}
	db.First(&dataProduk)

	fmt.Println(dataProduk)

	// 5 - mengambil banyak data

	// dengan IN
	// var produkSlice []Produk
	// db.Find(&produkSlice, []uint{1})

	// dengan WHERE
	var produkSlice []Produk
	db.Where(map[string]interface{}{"id": 1}).Find(&produkSlice, []uint{1})

	// dengan NOT
	// var produkSlice []Produk
	// db.Not(map[string]interface{}{"id": 2}).Find(&produkSlice, []uint{1})

	fmt.Println(produkSlice)

	// 6 - memperbarui data

	db.Model(Produk{ID: 1}).Updates(Produk{Nama: "New Kertas A5", Harga: 30000})
	fmt.Println("Data berhasil diupdate")

	// 7 - menghapus data

	db.Delete(Produk{ID: 1})
	fmt.Println("Data berhasil dihapus")

	// 8 - transaction

	db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Delete(&Produk{ID: 1}); result.Error != nil {
			fmt.Printf("Transaction gagal: %v\n", result.Error)

			// return error akan memanggil rollback
			return result.Error
		} else {
			fmt.Println("Transaction berhasil")

			// return nil akan memanggil commit
			return nil
		}
	})
}
