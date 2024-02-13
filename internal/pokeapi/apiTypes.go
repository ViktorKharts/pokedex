package pokeapi

type Locations struct {
	Next		*string		`json:"next"`
	Previous	*string		`json:"previous"`
	Results		[]struct{
		Name	string		`json:"name"`
		Url	string		`json:"url"`
	}				`json:"results"`
}

type AreaPokemons struct {
	PokemonEncounters	[]struct{
		Pokemon		struct{
			Name	string		`json:"name"`
			Url	string		`json:"url"`
		}				`json:"pokemon"`
	}					`json:"pokemon_encounters"`
}

type Pokemon struct {
	BaseExperience		int		`json:"base_experience"`
	Height			int		`json:"height"`
	LocationAreaEncounters	string		`json:"location_area_encounters"`
	Name			string		`json:"name"`
	Stats []struct {
		BaseStat	int		`json:"base_stat"`
		Effort		int		`json:"effort"`
		Stat     struct {
			Name	string		`json:"name"`
			URL	string		`json:"url"`
		}				`json:"stat"`
	}					`json:"stats"`
	Types []struct {
		Slot		int		`json:"slot"`
		Type struct {
			Name	string		`json:"name"`
			URL	string		`json:"url"`
		}				`json:"type"`
	}					`json:"types"`
	Weight			int		`json:"weight"`
}
