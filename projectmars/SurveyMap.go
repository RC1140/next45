package projectmars

import (
	"strconv"
	"strings"
)

type SurveyMap struct {
	Width  int
	Height int
}

func (m *SurveyMap) ParseAndSetupMap(sizeLine string) {
	gridParams := strings.Split(sizeLine, " ")
	x, err := strconv.ParseInt(gridParams[0], 10, 32)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseInt(gridParams[1], 10, 32)
	if err != nil {
		panic(err)
	}

	m.Width = int(x)
	m.Height = int(y)

}
