package main

import (
	"fmt"
	"log"
	"net/http"
	pokemon2 "pokemon/api/pokemon"
)

func pokemonJson(w http.ResponseWriter, r *http.Request) {
	pokemon, err := pokemon2.GetPokemonBy("ヒトカゲ")
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", pokemonJson)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	pokemon, err := pokemon2.GetPokemonBy("ヒトカゲ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pokemon)
	}
	//handleRequests()
}
