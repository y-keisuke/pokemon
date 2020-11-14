package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	pokemon2 "pokemon/pokemon"
)

func pokemonToJson(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pokemon, err := pokemon2.GetPokemonBy(name)
	if err != nil {
		log.Writer()
		http.Error(w, fmt.Sprintf("{\"err\":\"%s\"}", err), 200)
		return
	}
	pokemonJson, _ := json.Marshal(pokemon)
	fmt.Fprint(w, fmt.Sprintf("%+v", string(pokemonJson)))
}

func handleRequests() {
	http.HandleFunc("/", pokemonToJson)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
