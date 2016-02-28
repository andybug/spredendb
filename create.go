package main

import (
	"log"
	"io/ioutil"
)

type Sport struct {
	name string
	path string
}

func listSports(root string) (sports []string) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		log.Println("file " + f.Name())
		if f.IsDir() {
			sports = append(sports, f.Name())
		}
	}

	return sports
}

func createDatabase(root string, outPath string) error {
	log.Println("Creating database with root at " + root)
	log.Println("Outputing to " + outPath)
	sports := listSports(root)

	return nil
}
