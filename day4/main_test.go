package main

import (
	"reflect"
	"testing"
)

func Test_score(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "no winning number",
			s:    "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			want: 0,
		},
		{
			name: "has one winning number",
			s:    "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			want: 1,
		},
		{
			name: "has two winning numbers",
			s:    "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			want: 2,
		},
		{
			name: "has 4 winning numbers",
			s:    "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nb := winingCardsNb(tt.s)
			if got := score(nb); got != tt.want {
				t.Errorf("score() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getResults(t *testing.T) {
	tests := []struct {
		name    string
		entries []string
		want    []int
	}{
		{
			name:    "one card, no winning number",
			entries: []string{"Card 1: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"},
			want:    []int{1},
		},
		{
			name:    "one card, has one winning number",
			entries: []string{"Card 1: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"},
			want:    []int{1},
		},
		{
			name:    "one card, has two winning number",
			entries: []string{"Card 1:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"},
			want:    []int{1},
		},
		{
			name: "two cards, has one and two winning number",
			entries: []string{
				"Card 1: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 2:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			},
			want: []int{1, 2},
		},
		{
			name: "three cards, has one and two winning number",
			entries: []string{
				// 2 matching
				"Card 1:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				// 2 matching
				"Card 2:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				// no matching
				"Card 2: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			},
			want: []int{1, 2, 4},
		},
		{
			name: "sample",
			entries: []string{
				// 4 matching
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				// 2 matching
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				// 2 matching
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				// 2 matching
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				// 0 matching
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				// 0 matching
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			want: []int{1, 2, 4, 8, 14, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCardsNumber(tt.entries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCardsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
