package model

import (
	"encoding/json"
	"io"
)

type Games []Game

// FromJSON serializes data from json
func (o *Games) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(o)
}

// ToJSON converts the collection to json
func (o *Games) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}

type Game struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	StarPoint float32    `db:"star_point"`
	PlayerNum int    `db:"player_num"`
}

// FromJSON serializes data from json
func (o *Game) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(o)
}

// ToJSON converts the collection to json
func (o *Game) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}
