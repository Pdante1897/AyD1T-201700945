package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type ArchivoJson struct {
	Numero1 int `json:"Numero1"`
	Numero2 int `json:"Numero2"`
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo xd")
}
func suma(w http.ResponseWriter, r *http.Request) {
	archivo := new(ArchivoJson)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &archivo)

	var num1, num2 int
	num1 = archivo.Numero1
	num2 = archivo.Numero2
	fmt.Println(archivo)
	var result int
	result = num1 + num2
	fmt.Fprintf(w, "La suma es %v+%v=%v", num1, num2, result)
}

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bryan Gerardo Paez Morales 201700945")
}

func resta(w http.ResponseWriter, r *http.Request) {
	archivo := new(ArchivoJson)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &archivo)

	var num1, num2 int
	num1 = archivo.Numero1
	num2 = archivo.Numero2
	fmt.Println(archivo)
	var result int
	result = num1 - num2
	fmt.Fprintf(w, "La resta es %v+%v=%v", num1, num2, result)
}

func main() {

	fmt.Println("un server papu xd")

	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/suma", suma).Methods("POST")
	router.HandleFunc("/info", info).Methods("GET")
	router.HandleFunc("/resta", resta).Methods("POST")


	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}

