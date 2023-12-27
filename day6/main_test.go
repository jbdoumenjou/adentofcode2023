package main

import (
	"reflect"
	"testing"
)

func Test_nbOfWaysToBeatDistance(t *testing.T) {

	tests := []struct {
		name     string
		r        race
		time     int
		distance int
		want     int
	}{
		{
			name: "no way",
			r: race{
				Time:     1,
				Distance: 1,
			},
			want: 0,
		},
		{
			name: "simplest",
			r: race{
				Time:     3,
				Distance: 1,
			},
			want: 1,
		},
		{
			name: "first line of example",
			r: race{
				Time:     7,
				Distance: 9,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.nbOfWaysToBeatDistance(); got != tt.want {
				t.Errorf("nbOfWaysToBeatDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRaces(t *testing.T) {

	entries := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	want := []race{
		{Time: 7, Distance: 9},
		{Time: 15, Distance: 40},
		{Time: 30, Distance: 200},
	}
	if got := getRaces(entries); !reflect.DeepEqual(got, want) {
		t.Errorf("getEntries() = %v, want %v", got, want)
	}

}

func Test_getValue(t *testing.T) {
	line := "Time:      7  15   30"
	want := 71530

	if got := getValue(line); got != want {
		t.Errorf("getValue() = %v, want %v", got, want)
	}
}
