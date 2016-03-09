package main

import (
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

type Game struct {
	date       string
	uuid       string
	home_uuid  string
	home_score int
	away_uuid  string
	away_score int
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
			//round.games = readGames(round)
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
