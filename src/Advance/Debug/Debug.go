package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Avenger represent a single hero
type Avenger struct {
	RealName string `json:"real_name"`
	HeroName string `json:"hero_name"`
	Planet   string `json:"planet"`
	Alive    bool   `json:"alive"`
}

func (a *Avenger) isAlive() {
	a.Alive = true
}

func main() {
	avengers := []Avenger{
		{
			RealName: "Dr. Bruce Banner",
			HeroName: "Hulk",
			Planet:   "Midgar",
		},
		{
			RealName: "Tony Startk",
			HeroName: "Iron Man",
			Planet:   "Earth",
		}, {
			RealName: "Gun Gun Febrianza",
			HeroName: "Gatot Kaca",
			Planet:   "Moon",
		},
	}

	avengers[1].isAlive()
	
	jsonBytes, err := json.Marshal(avengers)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonBytes))
}
