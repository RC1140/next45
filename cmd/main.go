package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	ProjectMars "github.com/rc1140/next45/projectmars"
)

//ParseDataFile reads a data file that has the rover commands and returns them in clean params
func ParseDataFile(dataFile io.Reader) (string, string, string) {
	rawbytes, err := ioutil.ReadAll(dataFile)
	if err != nil {
		panic(err)
	}
	s := string(rawbytes)
	filelines := strings.SplitN(s, "\n", 3)
	if len(filelines) != 3 {
		panic(fmt.Sprintf("Invalid number of lines [%d] passed into file", len(filelines)))
	}

	size := filelines[0]
	location := filelines[1]
	commands := filelines[2]

	return size, location, commands
}

func main() {
	d, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	size, location, commands := ParseDataFile(d)
	m := ProjectMars.SurveyMap{}
	m.ParseAndSetupMap(size)

	r := ProjectMars.Rover{}
	r.SurveyMap = m
	r.Commands = commands
	r.ParseAndSetupStartLocation(location)
	r.ExecuteCommands()

	fmt.Printf("%v", r)
}
