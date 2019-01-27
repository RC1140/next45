package projectmars

import (
	"fmt"
	"strconv"
	"strings"
)

type Heading map[string]int

var RoverHeadings = Heading{
	"N": 1,
	"E": 2,
	"S": 3,
	"W": 4,
}

type CoOrd struct {
	X int
	Y int
}

func (c CoOrd) String() string {
	var res string
	res = fmt.Sprintf("X : %v , Y %v \n", c.X, c.Y)

	return res
}

type Rover struct {
	SurveyMap       SurveyMap
	Heading         int
	CurrentLocation CoOrd
	Commands        string
}

func (e *Rover) ParseAndSetupStartLocation(locationline string) {
	locationparams := strings.Split(locationline, " ")

	if len(locationparams[0]) > 2 {
		panic("Initial location too large")
	}

	x, err := strconv.ParseInt(string(locationparams[0][0]), 10, 32)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseInt(string(locationparams[0][1]), 10, 32)
	if err != nil {
		panic(err)
	}
	e.CurrentLocation.X = int(x)
	e.CurrentLocation.Y = int(y)

	val, found := RoverHeadings[string(locationparams[1][0])]
	if found {
		e.Heading = val
	} else {
		panic("Invalid heading sent")
	}
}

//IsOnMap checks that the current location of the rover is within the defined map
func (e *Rover) IsOnMap() bool {
	return e.CurrentLocation.X >= 0 &&
		e.CurrentLocation.Y >= 0 &&
		e.CurrentLocation.X <= e.SurveyMap.Width &&
		e.CurrentLocation.Y <= e.SurveyMap.Height
}

//Move executes the rovers move command which moves the rover one block in which ever location its facing
//as long as it is within the map
func (e *Rover) Move() {
	var x, y int
	x = e.CurrentLocation.X
	y = e.CurrentLocation.Y
	switch e.Heading {
	case 1:
		e.CurrentLocation.Y++
	case 2:
		e.CurrentLocation.X++
	case 3:
		e.CurrentLocation.Y--
	case 4:
		e.CurrentLocation.X--
	}
	if !e.IsOnMap() {
		e.CurrentLocation.X = x
		e.CurrentLocation.Y = y
	}
}

func (e *Rover) Turn(direction string) {
	switch direction {
	case "L":
		e.Heading--
	case "R":
		e.Heading++
	}
	if e.Heading > 4 && direction == "R" {
		e.Heading = 1
	}
	if e.Heading < 1 && direction == "L" {
		e.Heading = 4
	}
}

func (e *Rover) ExecuteCommands() {
	//TODO : Check that the commands are set first
	for _, c := range e.Commands {
		cc := string(c)
		if cc == "M" {
			e.Move()
		} else {
			e.Turn(cc)
		}
	}
}

//String prints out the rovers current location in a human readable format
func (r Rover) String() string {
	var res string
	res += fmt.Sprintf("Current Location : %v", r.CurrentLocation)
	res += fmt.Sprintf("Current Heading : %v", r.Heading)
	return res
}
