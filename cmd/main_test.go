package main

import (
	"io"
	"strings"
	"testing"
)

func TestParseDataFile(t *testing.T) {
	type args struct {
		dataFile io.Reader
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		{
			"basic file parser",
			args{strings.NewReader(`8 10
12 E
MMLMRMMRRMML`)},
			"8 10",
			"12 E",
			"MMLMRMMRRMML",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := ParseDataFile(tt.args.dataFile)
			if got != tt.want {
				t.Errorf("ParseDataFile() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseDataFile() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ParseDataFile() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
