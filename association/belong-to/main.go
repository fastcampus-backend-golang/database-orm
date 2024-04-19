package main

// Kabupaten belong to one Provinsi
// Terdapat ID Provinsi & struct Provinsi di dalam struct Kabupaten

type Kabupaten struct {
	ID         uint     `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	Nama       string   `gorm:"type:varchar(255);column:nama"`
	IDProvinsi uint     `gorm:"column: id_provinsi"`
	Provinsi   Provinsi `gorm:"foreignKey:id_provinsi"`
}

type Provinsi struct {
	ID   uint   `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	Nama string `gorm:"type:varchar(255);column:nama"`
}
