package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strconv"
	//	"errors"
)

type Team struct {
	uuid string
	name string
}

type TeamScore struct {
	Uuid  string `json:"uuid"`
	Score int    `json:"score"`
}

type Game struct {
	Date    string    `json:"date"`
	Uuid    string    `json:"uuid"`
	Home    TeamScore `json:"home"`
	Away    TeamScore `json:"away"`
	Neutral bool      `json:"neutral"`
}

type Round struct {
	path  string
	games []*Game
}

type Season struct {
	path   string
	year   int
	teams  []*Team
	rounds []*Round
}

type Sport struct {
	path    string
	name    string
	seasons []*Season
}

func (round *Round) readGames() error {
	log.Println("Reading " + round.path)

	file, err := ioutil.ReadFile(round.path)
	if err != nil {
		log.Fatal(err)
	}

	games := make([]Game, 0)
	err = json.Unmarshal(file, &games)
	if err != nil {
		log.Fatal(err)
	}

	for _, game := range games {
		log.Println(game.Date + "  " + game.Uuid + " " + strconv.FormatBool(game.Neutral))
	}

	return nil
}

func (season *Season) readRounds() error {
	files, err := ioutil.ReadDir(season.path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		match, _ := regexp.MatchString("^round\\d{2}.json$", f.Name())
		if match {
			round := new(Round)
			round.path = season.path + "/" + f.Name()
			round.readGames()
			season.rounds = append(season.rounds, round)
		}
	}

	return nil
}

func (sport *Sport) readSeasons() {
	files, err := ioutil.ReadDir(sport.path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		match, _ := regexp.MatchString("^\\d{4}$", f.Name())
		if f.IsDir() && match {
			season := new(Season)
			season.path = sport.path + "/" + f.Name()
			season.year, _ = strconv.Atoi(f.Name())
			season.readRounds()
			sport.seasons = append(sport.seasons, season)
		}
	}
}

func createDatabase(root string, outPath string) error {
	log.Println("Creating database with root at " + root)
	log.Println("Outputing to " + outPath)

	sport := new(Sport)
	sport.path = root
	sport.name = path.Base(root)
	sport.readSeasons()

	return nil
}
