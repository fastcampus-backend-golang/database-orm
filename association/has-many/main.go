package main

// Provinsi has many Kabupaten
// Terdapat slice of struct Kabupaten dalam Provinsi

type Kabupaten struct {
	ID   uint   `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	Nama string `gorm:"type:varchar(255);column:nama"`
}

type Provinsi struct {
	ID        uint        `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	Nama      string      `gorm:"type:varchar(255);column:nama"`
	Kabupaten []Kabupaten `gorm:"foreignKey:id"`
}
