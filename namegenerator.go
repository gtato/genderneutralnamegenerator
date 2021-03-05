package genderneutralnamegenerator

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type NameGenerator struct {
	random *rand.Rand
	names  []string
}

func (ng *NameGenerator) GenerateName() string {
	randIndex := ng.random.Intn(len(ng.names))
	return ng.names[randIndex]
}

func NewNameGenerator() *NameGenerator {
	dat, err := ioutil.ReadFile("names")
	check(err)
	names := strings.Split(string(dat), "\n")
	nameGenerator := &NameGenerator{
		random: rand.New(rand.NewSource(rand.Int63())),
		names:  names,
	}
	nameGenerator.random.Seed(time.Now().UnixNano())
	return nameGenerator
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
