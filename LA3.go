Soal!


Versi--1

//1) Buatlah database dan tabel dengan nama bebas untuk menyimpan data 3 
//   mahasiswa yang terdiri dari NPM, Nama, Kelas, dan Jurusan!.

* Pertama-tama membuat databasen nya terlebih dahulu menggunakan XAMPP
* Setelah Itu Membuat Folder Baru Dan Menambahakan  go.mod,Dan go.sum 
  Berikut Cara nya : 
 - go mod init_Nama Folder, Contoh : go mod init LA3

  selanjutnya membuat koneksi ke database/membuat go.sum dengan cara
 - go get -u github.com/go-sql-driver/mysql

* Selanjutnya Buka VSC dan buat file Baru dengan nama main.go



----  Codingan Main.go -----


package main
import (
     "database/sql"
     "fmt"
     
     "github.com/go-sql-driver/mysql"
)
type ok struct { // ok <-- sesuaikan dengan nama dengan nama tabel di database
    npm     int
    nama    string
    kelas   string
    jurusan string
}
func connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/iya") // iya <-- sesuaikan dengan nama database
    if err != nil {
        return nil, err
    }
    return db, nil
}
func sqlQuery() {
    db, err := connect()
    if err != nil {
      fmt.Println(err.Error())
    }
    defer db.Close()

// Menambahkan Data-1
_, err = db.Exec("Insert into ok values (?,?,?,?)", "50421692", "Jonathan Sitorus", "2IA09", "Informatika")
if err != nil {
  fmt.Println(err.Error())
  return
}
fmt.Println("Data Berhasil Ditambah")
  
// Menambahkan Data-2  
_, err = db.Exec("Insert into ok values (?,?,?,?)", "12345", "Sitorus", "2IA08", "Sastra Inggris")
if err != nil {
  fmt.Println(err.Error())
  return
}
fmt.Println("Data Berhasil Ditambah")

// Menambahkan Data-3
_, err = db.Exec("Insert into ok values (?,?,?,?)", "54321", "Jonathan Boltok", "2IA02", "Manajemen")
if err != nil {
  fmt.Println(err.Error())
  return
}
  defer rows.Close()
  var result []ok
  for rows.Next() {
    var each = ok{}
    var err = rows.Scan(&each.npm, &each.nama, &each.kelas, &each.jurusan)
    if err != nil {
      fmt.Println(err.Error())
      return
    }
    result = append(result, each)
  }
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  for _, each := range result {
    fmt.Print(each.npm)
    fmt.Print(" | ")
    fmt.Print(each.nama)
    fmt.Print(" | ")
    fmt.Print(each.kelas)
    fmt.Print(" | ")
    fmt.Print(each.jurusan, "\n")
  }
}
func main() {
  sqlQuery()
}



=============
Versi---2

=============

package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

type data_msh struct {
    id    int
    nama  string
    npm   int
    kelas string
}

func connect() (*sql.DB, error) {
    db, err := sql.Open("mysql",
    "root:@tcp(127.0.0.1:3306)/fata_54419164")
    if err != nil {
        return nil, err
    }
    return db, nil
}

func sqlQuery() {
    db, err := connect()
    if err != nil {
        fmt.Println(err.Error())
    }
    defer db.Close()

    // tambah data
    _, err = db.Exec("insert into data_msh values
    (?,?,?,?)", "", "Udin", "12345678", "2IA07")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Printf("Data berhasil ditambah")

    //Update data
    // _, err = db.Exec("update data_msh set nama= ? where id=?", "Siapa", 2)
    // if err != nil {
    //  fmt.Println(err.Error())
    //  return
    // }
    // fmt.Printf("Data berhasil diubah")

    //Hapus Data
    _, err = db.Exec("delete from data_msh where id=?", 2)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Printf("Data berhasil dihapus")

    rows, err := db.Query("select * from data_msh")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer rows.Close()

    var result []data_msh

    for rows.Next() {
        var each = data_msh{}
        var err = rows.Scan(&each.id, &each.nama, &each.npm, &each.kelas)

        if err != nil {
            fmt.Println(err.Error())
            return
        }
        result = append(result, each)
    }

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    for _, each := range result {
        fmt.Print(each.id)
        fmt.Print(" | ")
        fmt.Print(each.nama)
        fmt.Print(" | ")
        fmt.Print(each.npm)
        fmt.Print(" | ")
        fmt.Print(each.kelas, "\n")
    }
}

func main() {
    sqlQuery()
}
