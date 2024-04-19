package main

// Penduduk many-to-many dengan Alamat
// Terdapat slice of struct Alamat dalam struct Penduduk
// Dibuat junction table untuk slice tersebut

type Penduduk struct {
	ID     uint     `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	Alamat []Alamat `gorm:"many2many:penduduk_alamat;"`
}

type Alamat struct {
	ID            uint   `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	AlamatLengkap string `gorm:"column:alamat_lengkap"`
}
