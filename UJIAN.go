=========== No 1 ===========

// Membuat Perulangan Nama & NPM

package main
import (
  "fmt"
)
func main() {
  for i := 0; i <= 10; i++ {
    fmt.Println("Nama_NPM")
  }
}


=========== No 2 ===========

// Membuat Sturct yang berisi nama,npm,hoby,dan grade

package main
import "fmt"
type data struct {
  nama  string
  npm   string
  hoby  string
  umur  string
  grade string
}
func main() {
  var dataku = data {}
  fmt.Print("Masukkan Nama: ")
  fmt.Scan(&dataku.nama)
  fmt.Print("Masukkan NPM: ")
  fmt.Scan(&dataku.npm)
  fmt.Print("Masukkan Hoby: ")
  fmt.Scan(&dataku.hoby)
  fmt.Print("Masukkan Umur: ")
  fmt.Scan(&dataku.umur)
  fmt.Print("Masukkan Grade: ")
  fmt.Scan(&dataku.grade)
  fmt.Print("")
  fmt.Print("===============================================")
  fmt.Println("Nama Aku: ",dataku.nama)
  fmt.Println("NPM: ",dataku.npm)
  fmt.Println("Hoby: ",dataku.hoby)
  fmt.Println("Umur: ",dataku.umur)
  fmt.Println("Grade: ",dataku.grade)
}


