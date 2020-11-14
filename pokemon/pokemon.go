package pokemon

import (
	"errors"
	"pokemon/db"
)

type Pokemon struct {
	Name      string `json:"name"`
	HP        int    `json:"hp"`
	Attack    int    `json:"attack"`
	Defense   int    `json:"defense"`
	SpAttack  int    `json:"sp_attack"`
	SpDefense int    `json:"sp_defense"`
	Speed     int    `json:"speed"`
}

func GetPokemonBy(name string) (*Pokemon, error) {
	pokemonCollection := getPokemonCollection()
	for _, pokemon := range pokemonCollection.Pokemons {
		if pokemon.Name.Japanese == name {
			return getPokemonStruct(pokemon), nil
		}
	}
	return nil, errors.New("ポケモンが見つかりません")
}

func getPokemonCollection() db.PokemonCollection {
	return db.GetPokemonCollection()
}

func getPokemonStruct(pokemon db.PokemonData) *Pokemon {
	return &Pokemon{
		Name: pokemon.Name.Japanese,
		HP: pokemon.Base.HP,
		Attack: pokemon.Base.Attack,
		Defense: pokemon.Base.Defense,
		SpAttack: pokemon.Base.SpAttack,
		SpDefense: pokemon.Base.SpDefense,
		Speed: pokemon.Base.Speed}
}