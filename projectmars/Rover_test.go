package projectmars

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Tests that the rover can move in a clock wise direction by only turning right in every direction
func TestRover_MoveClockwise(t *testing.T) {

	var flagtests = []struct {
		in  string
		out []int
	}{
		{"M", []int{2, 1}},  //Move east (based on init)
		{"RM", []int{2, 0}}, //turn and move south
		{"RM", []int{1, 0}}, //turn and move west
		{"RM", []int{1, 1}}, //turn and move north
	}

	m := ProjectMars.SurveyMap{}
	m.ParseAndSetupMap("8 10")
	r := ProjectMars.Rover{}
	r.SurveyMap = m
	r.ParseAndSetupStartLocation("11 E")

	// testDirections := []string{"M", "R"}
	for _, tc := range flagtests {
		r.Commands = tc.in
		r.ExecuteCommands()

		fmt.Println(tc.in)
		assert.Equal(t, tc.out[0], r.CurrentLocation.X)
		assert.Equal(t, tc.out[1], r.CurrentLocation.Y)
	}

}
func TestRover_CantMoveOffMap(t *testing.T) {

	var flagtests = []struct {
		in  string
		out []int
	}{
		{"M", []int{0, 1}},  //Move east (based on init)
		{"M", []int{0, 1}},  //Move east (based on init)
		{"LM", []int{0, 0}}, //Move east (based on init)
		{"M", []int{0, 0}},  //Move east (based on init)
	}

	m := ProjectMars.SurveyMap{}
	m.ParseAndSetupMap("8 10")
	r := ProjectMars.Rover{}
	r.SurveyMap = m
	r.ParseAndSetupStartLocation("11 W")

	// testDirections := []string{"M", "R"}
	for _, tc := range flagtests {
		r.Commands = tc.in
		r.ExecuteCommands()

		fmt.Println(tc.in)
		assert.Equal(t, tc.out[0], r.CurrentLocation.X)
		assert.Equal(t, tc.out[1], r.CurrentLocation.Y)
	}

}

//Tests that the rover can move in a counterclock wise direction by only turning right in every direction
func TestRover_MoveCounterClockwise(t *testing.T) {

	var flagtests = []struct {
		in  string
		out []int
	}{
		{"M", []int{2, 1}},  //Move east (based on init)
		{"LM", []int{2, 2}}, //turn and move south
		{"LM", []int{1, 2}}, //turn and move west
		{"LM", []int{1, 1}}, //turn and move north to  the original postition
	}

	m := ProjectMars.SurveyMap{}
	m.ParseAndSetupMap("8 10")
	r := ProjectMars.Rover{}
	r.SurveyMap = m
	r.ParseAndSetupStartLocation("11 E")

	// testDirections := []string{"M", "R"}
	for _, tc := range flagtests {
		r.Commands = tc.in
		r.ExecuteCommands()

		fmt.Println(tc.in)
		assert.Equal(t, tc.out[0], r.CurrentLocation.X)
		assert.Equal(t, tc.out[1], r.CurrentLocation.Y)
	}

}

func TestRover_ParseValidStartLocation(t *testing.T) {
	r := ProjectMars.Rover{}
	r.ParseAndSetupStartLocation("11 W")
	assert.Equal(t, 1, r.CurrentLocation.X)
	assert.Equal(t, 1, r.CurrentLocation.Y)
	assert.Equal(t, 4, r.Heading)
}

func TestRover_ParseInValidStartHeading(t *testing.T) {
	r := ProjectMars.Rover{}
	assert.Panics(t, func() { r.ParseAndSetupStartLocation("11 X") }, "The heading should not have parsed")
}

func TestRover_ParseInValidStartLocation(t *testing.T) {
	r := ProjectMars.Rover{}
	assert.Panics(t, func() { r.ParseAndSetupStartLocation("112 X") }, "The location should not have parsed")
}

func TestRover_ParseInValidValueStartLocation(t *testing.T) {
	r := ProjectMars.Rover{}
	assert.Panics(t, func() { r.ParseAndSetupStartLocation("AA X") }, "The location should not have parsed")
	assert.Panics(t, func() { r.ParseAndSetupStartLocation("1A X") }, "The location should not have parsed")
}
func TestRover_CanPrintCurrentLocation(t *testing.T) {
	r := ProjectMars.Rover{}
	r.ParseAndSetupStartLocation("11 W")

	var res = "Current Location : X : 1 , Y 1 \n"
	res += "Current Heading : 4"
	assert.Equal(t, res, r.String())
}

func TestCoOrds_CanPrintLocation(t *testing.T) {
	c := ProjectMars.CoOrd{X: 1, Y: 1}

	var res = "X : 1 , Y 1 \n"
	assert.Equal(t, res, c.String())
}
