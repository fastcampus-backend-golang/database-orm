# SQL dengan Gorm
## Cara Menjalankan
1. Jalankan postgres dengan docker
```
docker run --name postgresql -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=database -p 5432:5432 -d postgres:16
```

2. Install module yang dibutuhkan
```
go mod tidy
```

3. Jalankan program
```
go run main.go
```

## Konten
Dalam file `main.go` terdapat 8 bagian contoh:
- Menghubungkan Golang dengan Postgres
- Membuat tabel dengan auto migration
- Menambah data
- Mengambil 1 data
- Mengambil banyak data
- Memperbarui data
- Menghapus data
- Transaction

Dalam `association/main.go` terdapat 2 bagian contoh:
- Menambahkan data yang memiliki asosiasi
- Mengambil data yang memiliki asosiasi

Selain itu dalam direktori `association` terdapat contoh struct yang memiliki beberapa bentuk asosiasi