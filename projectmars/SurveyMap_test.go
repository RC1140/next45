package projectmars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSurveyMap_ParseMapSetup(t *testing.T) {
	m := SurveyMap{}
	m.ParseAndSetupMap("8 10")

	assert.Equal(t, 8, m.Width)
	assert.Equal(t, 10, m.Height)
}
func TestSurveyMap_InvalidMapPanics(t *testing.T) {
	m := SurveyMap{}
	assert.Panics(t, func() { m.ParseAndSetupMap("A 10") }, "The Map definition should not have parsed")
	assert.Panics(t, func() { m.ParseAndSetupMap("10 B") }, "The Map definition should not have parsed")
}
