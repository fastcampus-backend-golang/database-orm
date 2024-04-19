package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Penduduk struct {
	ID     uint     `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	Alamat []Alamat `gorm:"many2many:penduduk_alamat;"`
}

type Alamat struct {
	ID            uint   `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	AlamatLengkap string `gorm:"column:alamat_lengkap"`
}

func main() {
	connURI := "postgresql://postgres:password@localhost:5432/database?sslmode=disable"
	db, err := gorm.Open(postgres.Open(connURI), &gorm.Config{})
	if err != nil {
		fmt.Printf("Gagal menghubungkan ke database: %v\n", err)
		os.Exit(1)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	fmt.Println("Berhasil terhubung ke database")

	db.AutoMigrate(&Alamat{}, &Penduduk{})

	// 1. mengisi data
	pendudukInsert := Penduduk{
		Alamat: []Alamat{
			{
				AlamatLengkap: "Kota Jakarta",
			},
			{
				AlamatLengkap: "Kota Bandung",
			},
		},
	}
	if result := db.Create(&pendudukInsert); result.Error != nil {
		fmt.Printf("Gagal mengisi tabel: %v\n", result.Error)
		os.Exit(1)
	}

	// pendudukInsert otomatis diupdate dengan hasil insert
	fmt.Printf("Data berhasil ditambahkan: %v\n", pendudukInsert.ID)

	// 2. mengambil data dari dua tabel (menggunakan nama struct)
	var pendudukSelect Penduduk
	db.Preload("Alamat").First(&pendudukSelect, pendudukInsert.ID) // gunakan ID hasil insert

	fmt.Println(pendudukSelect)
}
