package models

import (
	"dapi/db"
	"dapi/db/digimon"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Digimon digimon.Digimon
type DigimonList []Digimon

// digimons slice to seed record digimon data.
// var Digimons = DigimonList{
// 	{ID: 1047, Name: "War Greymon (X-Antibody)", Stage: "Ultimate", Attribute: "Vaccine", Picture: "https://digimon-api.com/images/digimon/w/War_Greymon_(X-Antibody).png"},
// 	{ID: 771, Name: "Lucemon", Stage: "Child", Attribute: "Vaccine", Picture: "https://digimon-api.com/images/digimon/w/Lucemon.png"},
// 	{ID: 746, Name: "Punimon", Stage: "Baby I", Attribute: "Data", Picture: "https://digimon-api.com/images/digimon/w/Punimon.png"},
// }

// Example JSON
// {
//     "id": 1334,
//     "name": "Alphamon",
//     "stage": "Ultimate",
//     "attribute": "Vaccine",
//     "picture": "https://digimon-api.com/images/digimon/w/Alphamon.png"
// }

func GetDigimon(id string) Digimon {
	dbconn := db.DBConn
	var digimon Digimon
	dbconn.Find(&digimon, id)

	return digimon
}

func GetDigimons() DigimonList {
	dbconn := db.DBConn
	var digimons DigimonList
	dbconn.Find(&digimons)
	return digimons
}

func NewDigimon(id int, name string, stage string, attribute string, picture string) Digimon {
	dbconn := db.DBConn
	digimon := Digimon{
		ID:        id,
		Name:      name,
		Stage:     stage,
		Attribute: attribute,
		Picture:   picture,
	}
	dbconn.Create(&digimon)
	return digimon
}

func DeleteDigimon(id string) bool {
	dbconn := db.DBConn

	var digimon Digimon
	dbconn.First(&digimon, id)
	if digimon.Name == "" {
		return false
	}
	dbconn.Delete(&digimon)
	return true
}
