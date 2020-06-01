package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Marvel struct {
	Name string `json:"name,omitempty"`

	Character *Character `json:"character,omitempty"`
}
type Character struct {
	Cname    string `json:"cname,omitempty"`
	Maxpower string `json:"Maxpower,omitempty"`
}

var mar []Marvel
var c Character

func GetMarvelRecord(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(mar)
}
func GetMarvelOneRecord(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range mar {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Marvel{})
}

func GetMaxPower(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range mar {
		if item.Character.Cname == params["cname"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(mar)
}

func CreateMarvelDetails(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var marv Marvel
	_ = json.NewDecoder(req.Body).Decode(&marv)
	marv.Name = params["name"]
	mar = append(mar, marv)
	json.NewEncoder(w).Encode(marv)
}
func DeleteRecord(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range mar {
		if item.Name == params["name"] {
			mar = append(mar[:index], mar[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(mar)
}

func UpdateMarvelRecord(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	for index, item := range mar {
		if item.Name == params["name"] {
			mar = append(mar[:index], mar[index+1:]...)
			var marv Marvel
			_ = json.NewDecoder(req.Body).Decode(&marv)
			marv.Name = params["name"]
			mar = append(mar, marv)
			json.NewEncoder(w).Encode(marv)
		}
	}
	json.NewEncoder(w).Encode(mar)
}

func main() {
	router := mux.NewRouter()
	mar = append(mar, Marvel{Name: "mutants",
		Character: &Character{Cname: "Apocalpsy", Maxpower: "90"}})
	mar = append(mar, Marvel{Name: "Avenger",
		Character: &Character{Cname: "IronMan", Maxpower: "60"}})
	mar = append(mar, Marvel{Name: "AntiHeros",
		Character: &Character{Cname: "Madrin", Maxpower: "70"}})

	router.HandleFunc("/marvel", GetMarvelRecord).Methods("GET")
	router.HandleFunc("/marvel/{name}", GetMarvelOneRecord).Methods("GET")
	router.HandleFunc("/marvel/{cname}", GetMaxPower).Methods("GET")
	router.HandleFunc("/marvel/{name}", CreateMarvelDetails).Methods("POST")
	router.HandleFunc("/marvel/{name}", UpdateMarvelRecord).Methods("PUT")

	log.Fatal(http.ListenAndServe(":12345", router))

}
