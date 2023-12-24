package main

import "testing"

func Test_gearRatiosSum(t *testing.T) {
	tests := []struct {
		name   string
		matrix []string
		want   int
	}{
		{
			name:   "only a star",
			matrix: []string{"*"},
			want:   0,
		},
		{
			name:   "only a star with other non digit chars",
			matrix: []string{"....*."},
			want:   0,
		},
		{
			name:   "minimal valid gear (one line)",
			matrix: []string{"4*5"},
			want:   20,
		},
		{
			name:   "invalid gear with one number right(one line)",
			matrix: []string{".*5"},
			want:   0,
		},
		{
			name:   "invalid gear with one number left(one line)",
			matrix: []string{"4*."},
			want:   0,
		},
		{
			name:   "valid gear (one line, len nb > 1)",
			matrix: []string{"10*523"},
			want:   5230,
		},
		{
			name:   "valid gear (one line, len nb > 1, with other chars)",
			matrix: []string{"...10*523.."},
			want:   5230,
		},
		{
			name: "valid gear (one line in matrix, len nb > 1, with other chars)",
			matrix: []string{
				".8...5...",
				".10*523.5",
				".8...5...",
			},
			want: 5230,
		},
		{
			name: "valid gear up left (multi line in matrix, len nb > 1, with other chars)",
			matrix: []string{
				".10..5...",
				"...*523.5",
				".8...5...",
			},
			want: 5230,
		},
		{
			name: "valid gear above left (multi line in matrix, len nb > 1, with other chars)",
			matrix: []string{
				".10..5...",
				"...*523.5",
				".8...5...",
			},
			want: 5230,
		},
		{
			name: "valid gear above right (multi line in matrix, len nb > 1, with other chars)",
			matrix: []string{
				".1..10...",
				"...*523.5",
				".8...5...",
			},
			want: 5230,
		},
		{
			name: "valid gear above (multi line in matrix, len nb > 1, with other chars)",
			matrix: []string{
				"..523....",
				"...*10.5",
				".8...5...",
			},
			want: 5230,
		},
		{
			name: "valid gear above (multi line in matrix, len nb > 1, with other chars)",
			matrix: []string{
				"..523....",
				"...*...5",
				".8..10...",
			},
			want: 5230,
		},
		{
			name: "valid gear below (multi line in matrix, len nb > 1, with other chars)",
			matrix: []string{
				"..523....",
				".3.*...5",
				"2.100.4.",
			},
			want: 52300,
		},
		{
			name: "invalid gear to many part numbers (multi line in matrix, len nb > 1, with other chars)",
			matrix: []string{
				"..523....",
				"..3*...5",
				"2.100.4.",
			},
			want: 0,
		},
		{
			name: "invalid gear to many part numbers (multi line in matrix, len nb > 1, with other chars)",
			matrix: []string{
				"..5.3....",
				"..3*1..5",
				"2.1.0.4.",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gearRatiosSum(tt.matrix); got != tt.want {
				t.Errorf("gearRatiosSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
