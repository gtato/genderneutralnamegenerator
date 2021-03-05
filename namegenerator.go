package genderneutralnamegenerator

import (
	"math/rand"
	"time"
)

type NameGenerator struct {
	random *rand.Rand
}

func (ng *NameGenerator) GenerateName() string {
	randIndex := ng.random.Intn(len(NAMES))
	return NAMES[randIndex]
}

func NewNameGenerator() *NameGenerator {
	nameGenerator := &NameGenerator{
		random: rand.New(rand.NewSource(rand.Int63())),
	}
	nameGenerator.random.Seed(time.Now().UnixNano())
	return nameGenerator
}
