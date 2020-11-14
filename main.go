package main

import (
	"fmt"
	"log"
	"net/http"
	"pokemon/db"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	pokemonCollection := db.GetPokemonCollection()
	for _, pokemon := range pokemonCollection.Pokemons {
		fmt.Println(pokemon.Name.Japanese, pokemon.Base)
	}

	//handleRequests()
}
