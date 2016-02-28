package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

type Season struct {
	year int16
}

type Sport struct {
	name string
	path string
}

func listSports(root string) (sports []*Sport) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		log.Println("file " + f.Name())
		if f.IsDir() {
			sport := new(Sport)
			sport.name = f.Name()
			sport.path = root + "/" + f.Name()
			sports = append(sports, sport)
		}
	}

	return sports
}

func createDatabase(root string, outPath string) error {
	log.Println("Creating database with root at " + root)
	log.Println("Outputing to " + outPath)
	sports := listSports(root)

	for _, sport := range sports {
		fmt.Printf("sport %s %s\n", sport.name, sport.path)
	}

	return nil
}
