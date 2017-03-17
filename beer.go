package main

import "time"

type Beer struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	DidWeDrinkIt bool      `json:"didwedrinkit"`
	Due       time.Time `json:"due"`
	Abv			string	`json:"abv"`
}

type Beers []Beer
