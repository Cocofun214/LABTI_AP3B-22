package main
import (
  "fmt"
)
func main() {
  fmt.Println("Masukkan Tinggi Badan : ") //Meminta Input
  var tinggi int // Mendeklarasikan variabel tinggi dengan integer
  fmt.Scanln(&tinggi) // Membaca/Scaning input 
  if tinggi >= 165 { //Kondisi Jika tinggi Lebih Besar Sama Dengan 165
    fmt.Println("Anda LOLOS") // Jika Benar 
  } else {
    fmt.Println("Kamu Tidak Memenuhi Syarat") //Jika Tidak
  }
}
