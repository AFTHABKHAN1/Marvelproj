package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Avenger struct {
	Name      string `json:"name"`
	Character []struct {
		Name     string `json:"name"`
		MaxPower int    `json:"max_power"`
	} `json:"character"`
}

type AntiHero struct {
	Name      string `json:"name"`
	Character []struct {
		Name     string `json:"name"`
		MaxPower int    `json:"max_power"`
	} `json:"character"`
}

type Mutant struct {
	Name      string `json:"name"`
	Character []struct {
		Name     string `json:"name"`
		MaxPower int    `json:"max_power"`
	} `json:"character"`
}
type DataHolder [15]struct {
	Name     string `json:"name"`
	Count    int    `json:"count"`
	MaxPower int    `json:"max_power"`
}

var data DataHolder
var avg Avenger
var mut Mutant
var anti AntiHero
var m int
var n int

func Unmarshalling(str string) {

	jsonFile, err := os.Open(str)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Avenger.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &avg)

}
func UnmarshallAnti(str string) {

	jsonAnti, err := os.Open(str)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Antihero.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonAnti.Close()
	// read our opened xmlFile as a byte array.
	byteValues, _ := ioutil.ReadAll(jsonAnti)

	json.Unmarshal(byteValues, &anti)

}
func UnmarshallMut(str string) {

	jsonMut, err := os.Open(str)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened mutant.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonMut.Close()
	// read our opened xmlFile as a byte array.
	byteValuess, _ := ioutil.ReadAll(jsonMut)

	json.Unmarshal(byteValuess, &mut)

}

func GetMarvelRecord(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(avg)
}
func GetAntiRecord(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(anti)
}
func GetMutantRecord(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(mut)
}
func GetData(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(data)
}

/*func GetMarvelsRecord(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(slice)
}*/
var o int

func GetMaxpowerAntiHero(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for i := 0; i < len(anti.Character); i++ {
		if anti.Character[i].Name == params["name"] {
			fmt.Println(anti)
			json.NewEncoder(w).Encode(anti.Character[i].MaxPower)

			for k := 0; k < len(data); k++ {
				fmt.Println("data " + data[k].Name)
				fmt.Print("params " + params["name"])

				if data[k].Name == params["name"] {
					//data[k].MaxPower = anti.Character[i].MaxPower
					//data[k].Name = anti.Character[i].Name
					data[k].Count = o + 1
					break
				} else {
					data[k].MaxPower = anti.Character[i].MaxPower
					data[k].Name = anti.Character[i].Name
					data[k].Count = o + 1
					break
				}

			}

			//	data[m].MaxPower = anti.Character[i].MaxPower
			//	data[m].Name = anti.Character[i].Name
			//	data[m].Count = n + 1
			//	m = m + 1
			return
		}

	}

	json.NewEncoder(w).Encode(anti)
}

func GetMaxpowerMutant(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for i := 0; i < len(mut.Character); i++ {
		if mut.Character[i].Name == params["name"] {
			fmt.Println(mut)
			json.NewEncoder(w).Encode(mut.Character[i].MaxPower)
			for k := 0; k < len(data); k++ {
				fmt.Println(data[k].Name)

				if params["name"] == data[k].Name {
					data[k].MaxPower = mut.Character[i].MaxPower
					data[k].Name = mut.Character[i].Name
					data[k].Count = n + 1

				} else {
					data[m].MaxPower = mut.Character[i].MaxPower
					data[m].Name = mut.Character[i].Name
					data[m].Count = n + 1

				}

			}
			m = m + 1
			return
		}
	}
	json.NewEncoder(w).Encode(mut)
}
func GetMaxpowerAvenger(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for i := 0; i < len(avg.Character); i++ {
		if avg.Character[i].Name == params["name"] {
			fmt.Println(mut)

			json.NewEncoder(w).Encode(avg.Character[i].MaxPower)

			for k := 0; k < len(data); k++ {
				fmt.Println(data[k].Name)

				if data[k].Name == params["name"] {
					//	data[k].MaxPower = avg.Character[i].MaxPower
					//	data[k].Name = avg.Character[i].Name
					data[k].Count = n + 1

				} else {
					data[m].MaxPower = avg.Character[i].MaxPower
					data[m].Name = avg.Character[i].Name
					data[m].Count = n + 1

				}

			}
			m = m + 1
			return
		}

	}
	json.NewEncoder(w).Encode(avg)
}
func DeleteRecord(w http.ResponseWriter, req *http.Request) {
	//	params := mux.Vars(req)
	m := 0

	var i int
	for i = 0; i < len(data); i++ {

		if data[i].MaxPower < m {
			m = data[i].MaxPower
			break
		}

	}
	data[i].MaxPower = 0
	data[i].Name = ""
	data[i].Count = 0
	json.NewEncoder(w).Encode(data)
}
func main() {
	router := mux.NewRouter()
	// Open our jsonFile

	Unmarshalling("C:\\Users\\user\\Desktop\\Gofile\\Avenger.json")
	UnmarshallAnti("C:\\Users\\user\\Desktop\\Gofile\\Antiheroes.json")
	UnmarshallMut("C:\\Users\\user\\Desktop\\Gofile\\Mutants.json")
	fmt.Println("%+T", anti)
	fmt.Println("%+T", mut)
	router.HandleFunc("/Avenger", GetMarvelRecord).Methods("GET")
	router.HandleFunc("/Mutant", GetMutantRecord).Methods("GET")
	router.HandleFunc("/AntiHero", GetAntiRecord).Methods("GET")
	router.HandleFunc("/Maxpower/Antihero/{name}", GetMaxpowerAntiHero).Methods("GET")
	router.HandleFunc("/Maxpower/Mutant/{name}", GetMaxpowerMutant).Methods("GET")
	router.HandleFunc("/Maxpower/Avenger/{name}", GetMaxpowerAvenger).Methods("GET")
	router.HandleFunc("/Maxpower/Data", GetData).Methods("GET")
	router.HandleFunc("/Maxpower", DeleteRecord).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))

}
