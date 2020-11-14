package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type PokemonCollection struct {
	Pokemons []PokemonData
}

type PokemonData struct {
	Id   int         `json:"id"`
	Name PokemonName `json:"name"`
	Type []string    `json:"type"`
	Base PokemonBase `json:"base"`
}

type PokemonName struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Chinese  string `json:"chinese"`
	French   string `json:"french"`
}

type PokemonBase struct {
	HP        int `json:"hp"`
	Attack    int `json:"attack"`
	Defense   int `json:"defense"`
	SpAttack  int `json:"spattack"`
	SpDefense int `json:"spdefense"`
	Speed     int `json:"speed"`
}

func GetPokemonCollection() PokemonCollection {
	raw, err := ioutil.ReadFile("./pokedex.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var pokemonCollection PokemonCollection
	json.Unmarshal(raw, &pokemonCollection)

	return pokemonCollection
}
