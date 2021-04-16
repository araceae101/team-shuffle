package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: main [INPUT FILE NAME]")
		os.Exit(1)
	}
	inputFilePath := flag.Args()[0]

	g, err := mkGroup(inputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	g.shuffle()

	for _, p := range g.name {
		fmt.Println(p)
	}
}

type group struct {
	name   []string
	numGen *rand.Rand
}

func mkGroup(filepath string) (*group, error) {
	person := []string{}

	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open csv: %w", err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv: %w", err)
	}
	for _, line := range csvLines {
		person = append(person, line[0])
	}
	return &group{
		name:   person,
		numGen: rand.New(rand.NewSource(time.Now().Unix())),
	}, nil
}

func (g group) shuffle() {
	if g.numGen == nil {
		g.numGen = rand.New(rand.NewSource(time.Now().Unix()))
	}
	g.numGen.Shuffle(len(g.name), func(i, j int) {
		g.name[i], g.name[j] = g.name[j], g.name[i]
	})
}
