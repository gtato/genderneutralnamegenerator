package genderneutralnamegenerator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameGenerator(t *testing.T) {
	ng := NewNameGenerator()
	name := ng.GenerateName()
	assert.Greater(t, len(name), 1)
	assert.NotNil(t, name)
}
