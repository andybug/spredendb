package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

type Season struct {
	year int
}

type Sport struct {
	name    string
	path    string
	seasons []*Season
}

func listSports(root string) (sports []*Sport) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			sport := new(Sport)
			sport.name = f.Name()
			sport.path = root + "/" + f.Name()
			sports = append(sports, sport)
		}
	}

	return sports
}

func listSeasons(sport *Sport) (seasons []*Season) {
	files, err := ioutil.ReadDir(sport.path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		match, _ := regexp.MatchString("^\\d{4}$", f.Name())
		if f.IsDir() && match {
			season := new(Season)
			season.year, _ = strconv.Atoi(f.Name())
		}
	}

	return nil
}

func createDatabase(root string, outPath string) error {
	log.Println("Creating database with root at " + root)
	log.Println("Outputing to " + outPath)
	sports := listSports(root)

	for _, sport := range sports {
		log.Println("Processing sport " + sport.name)
		sport.seasons = listSeasons(sport)
		for _, season := range sport.seasons {
			log.Printf("Processing year %d\n", season.year)
		}
	}

	return nil
}
