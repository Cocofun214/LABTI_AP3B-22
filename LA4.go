Soal!

//Buatlah program API sederhana yang menerapkan method GET dan POST untuk data kode_mk,
//nama_matkul, dan sks. Uji coba program kalian dengan menggunakan Postman.


------ Method POST --------

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Kode_mk     string `json:"kode"`
	Nama_matkul string `json:"matkul"`
	Sks         string `json:"sks"`
}

var todos = []todo{
	{Kode_mk: "A", Nama_matkul: "Statistika", Sks: "2"},
	{Kode_mk: "B", Nama_matkul: "Bahasa Indonesia", Sks: "1"},
	{Kode_mk: "C", Nama_matkul: "Matematika Informatika", Sks: "2"},
}

func postTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
func main() {
	router := gin.Default()
	router.POST("/Jonathan", postTodos)
	router.Run("localhost:5050")
}


================================================================


------ Method GET -------


package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Kode_mk     string `json:"kode"`
	Nama_matkul string `json:"matkul"`
	Sks         string `json:"sks"`
}

var todos = []todo{
	{Kode_mk: "A", Nama_matkul: "Statistika", Sks: "2"},
	{Kode_mk: "B", Nama_matkul: "Bahasa Indonesia", Sks: "1"},
	{Kode_mk: "C", Nama_matkul: "Matematika Informatika", Sks: "2"},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
func main() {
	router := gin.Default()
	router.GET("/Jonathan", getTodos)
	router.Run("localhost:4040")
}
