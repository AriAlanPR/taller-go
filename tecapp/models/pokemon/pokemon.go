package pokemon

type pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Ability        string `json:"ability"`
	Type           string `json:"type"`
	PriorEvolution int    `json:"priorEvolution"`
	NextEvolution  int    `json:"nextEvolution"`
	Picture        string `json:"picture"`
}

type PokemonList []pokemon

// digimons slice to seed record pokemon data.
var Pokemons = PokemonList{
	{ID: 27, Name: "Sandshrew", Ability: "Sand Rush", Type: "Ground", PriorEvolution: 0, NextEvolution: 28, Picture: "https://images.wikidexcdn.net/mwuploads/wikidex/thumb/d/df/latest/20230620060651/Sandshrew.png/300px-Sandshrew.png"},
	{ID: 151, Name: "Mew", Ability: "Synchronize", Type: "Psychic", PriorEvolution: 0, NextEvolution: 0, Picture: "https://images.wikidexcdn.net/mwuploads/wikidex/thumb/b/bf/latest/20160311010530/Mew.png/300px-Mew.png"},
	{ID: 54, Name: "Psyduck", Ability: "Damp", Type: "Water", PriorEvolution: 0, NextEvolution: 55, Picture: "https://images.wikidexcdn.net/mwuploads/wikidex/thumb/3/32/latest/20230614194705/Psyduck.png/300px-Psyduck.png"},
}

// Example JSON
// {
//     "id": 999,
//     "name": "Gimmighoul",
//     "ability": "Rattled",
//     "type": "Ghost",
//     "priorEvolution": 0,
//     "nextEvolution": 1000,
//     "picture": "https://images.wikidexcdn.net/mwuploads/wikidex/thumb/7/74/latest/20221106182546/Gimmighoul.png/300px-Gimmighoul.png"
// }

func (pl PokemonList) Get() PokemonList {
	return pl
}

func (pl PokemonList) New() pokemon {
	return pokemon{}
}

func (pl PokemonList) Save(p pokemon) PokemonList {
	// Add the new pokemon to the slice.
	Pokemons = append(pl, p)
	return Pokemons
}
