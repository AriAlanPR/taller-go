package digimon

type digimon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Level          int    `json:"level"`
	Stage          string `json:"stage"`
	Attribute      string `json:"attribute"`
	PriorEvolution int    `json:"priorEvolution"`
	NextEvolution  int    `json:"nextEvolution"`
}

type DigimonList []digimon

// digimons slice to seed record digimon data.
var Digimons = DigimonList{
	{ID: 1047, Name: "War Greymon (X-Antibody)", Level: 4, Stage: "Ultimate", Attribute: "Vaccine", PriorEvolution: 76, NextEvolution: 1309},
	{ID: 771, Name: "Lucemon", Level: 7, Stage: "Child", Attribute: "Vaccine", PriorEvolution: 833, NextEvolution: 879},
	{ID: 746, Name: "Punimon", Level: 5, Stage: "Baby I", Attribute: "Data", PriorEvolution: -1, NextEvolution: 1154},
}

// Example JSON
// {
//     "id": 1334,
//     "name": "Alphamon",
//     "level": 4,
//     "stage": "Ultimate",
//     "attribute": "Vaccine",
//     "priorEvolution": 1374,
//     "nextEvolution": 560
// }

func (dl DigimonList) Get() DigimonList {
	return dl
}

func (dl DigimonList) New() digimon {
	return digimon{}
}

func (dl DigimonList) Save(d digimon) DigimonList {
	// Add the new digimon to the slice.
	Digimons = append(dl, d)
	return Digimons
}
