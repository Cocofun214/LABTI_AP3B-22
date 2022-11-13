// Soal Nomor 1 Membuat Program -- Mencari Nilai Rata-Rata Dari 5 Buah Elemen Array --


----- No.1 -------

package main
import "fmt"
func main() {
  var data int
  fmt.Println("Masukkan Ukuran Array : ")
  fmr.Scan(&data)
  fmt.Println("Masukkan Elemen Array : ")
  array := make([]int, data)
  sum := 0
  for i := 0; i < data; i++ {
    fmt.Scan(&array[i])
    sum = sum + array[i]
  }
  rata := float64(sum) / float64(data)
  fmt.Printf("\nRata-Rata Array Adalah : %f", rata)
}



// Soal Nomor 2 Membuat Program -- Menampilkan Nama,Hoby, Dan Kelas Dengan Menggunakan Struct --

------- No.2 -------


package main
import "fmt"
type mhs struct {
    nama string
    hoby string
    kelas string
}
func main() {
  var data_mhs = mhs{}
  fmt.Print("Masukkan Nama : ")
  fmt.Scan(&data_mhs.nama)
  fmt.Print("Masukkan Hoby : ")
  fmt.Scan(&data_mhs.hoby)
  fmt.Print("Masukkan Kelas : ")
  fmt.Scan(&data_mhs.kelas)
  fmt.Println("")
  fmt.Println("\n------------------------------------")
  fmt.Println("NAMA :", data_mhs.nama)
  fmt.Println("HOBY :", data_mhs.hoby)
  fmt.Println("KELAS :", data_mhs.kelas)
  fmt.Println("------------------------------------\n")
}
