package main

// Provinsi has one Kepala Daerah
// Terdapat struct Kepala Daerah dalam Provinsi

type Provinsi struct {
	ID           uint         `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	Nama         string       `gorm:"type:varchar(255);column:nama"`
	KepalaDaerah KepalaDaerah `gorm:"foreignKey:id"` // nama PK di tabel parent
}

type KepalaDaerah struct {
	ID   uint   `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	Nama string `gorm:"type:varchar(255);column:nama"`
}
