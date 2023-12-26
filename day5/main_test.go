package main

import (
	"reflect"
	"testing"
)

func Test_mapper_dst(t *testing.T) {
	tests := []struct {
		name string
		m    mapper
		v    int
		want int
	}{
		{
			name: "one mapping, unmatched mapping",
			m: mapper{mapping{
				dst:         50,
				src:         98,
				rangeLength: 2,
			}},
			v:    97,
			want: 97,
		},
		{
			name: "one mapping, matched mapping",
			m: mapper{mapping{
				dst:         50,
				src:         98,
				rangeLength: 2,
			}},
			v:    99,
			want: 51,
		},
		{
			name: "one mapping, matched mapping, lower limit",
			m: mapper{mapping{
				dst:         50,
				src:         98,
				rangeLength: 2,
			}},
			v:    98,
			want: 50,
		},
		{
			name: "one mapping, no match, upper limit",
			m: mapper{mapping{
				dst:         50,
				src:         98,
				rangeLength: 2,
			}},
			v:    100,
			want: 100,
		},
		{
			name: "two mapping, no match, lower limit",
			m: mapper{
				mapping{
					dst:         50,
					src:         98,
					rangeLength: 2,
				},
				mapping{
					dst:         52,
					src:         50,
					rangeLength: 48,
				},
			},
			v:    49,
			want: 49,
		},
		{
			name: "two mapping, match, lower limit",
			m: mapper{
				mapping{
					dst:         50,
					src:         98,
					rangeLength: 2,
				},
				mapping{
					dst:         52,
					src:         50,
					rangeLength: 48,
				},
			},
			v:    50,
			want: 52,
		},
		{
			name: "two mapping, match, upper limit",
			m: mapper{
				mapping{
					dst:         50,
					src:         98,
					rangeLength: 2,
				},
				mapping{
					dst:         52,
					src:         50,
					rangeLength: 48,
				},
			},
			v:    97,
			want: 99,
		},
		{
			name: "two mapping, no match, upper limit",
			m: mapper{
				mapping{
					dst:         50,
					src:         98,
					rangeLength: 2,
				},
				mapping{
					dst:         52,
					src:         50,
					rangeLength: 48,
				},
			},
			v:    98,
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.dst(tt.v); got != tt.want {
				t.Errorf("dst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSeeds(t *testing.T) {
	want := []int{79, 14, 55, 13}
	if got := getSeeds("seeds: 79 14 55 13"); !reflect.DeepEqual(got, want) {
		t.Errorf("getSeeds() = %v, want %v", got, want)
	}
}

func Test_getMappers(t *testing.T) {

	tests := []struct {
		name    string
		entries []string
		want    []mapper
	}{
		{
			name: "one mapper, two mappings",
			entries: []string{
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
			},
			want: []mapper{
				{
					{
						dst:         50,
						src:         98,
						rangeLength: 2,
					},
					{
						dst:         52,
						src:         50,
						rangeLength: 48,
					},
				},
			},
		},
		{
			name: "two mappers, two mappings",
			entries: []string{
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
			},
			want: []mapper{
				{
					{
						dst:         50,
						src:         98,
						rangeLength: 2,
					},
					{
						dst:         52,
						src:         50,
						rangeLength: 48,
					},
				},
				{
					{
						dst:         0,
						src:         15,
						rangeLength: 37,
					},
					{
						dst:         37,
						src:         52,
						rangeLength: 2,
					},
					{
						dst:         39,
						src:         0,
						rangeLength: 15,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMappers(tt.entries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMappers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLocation(t *testing.T) {
	tests := []struct {
		name    string
		mappers []mapper
		seed    int
		want    int
	}{
		{
			name: "one mapper",
			mappers: []mapper{
				{
					{
						dst:         50,
						src:         98,
						rangeLength: 2,
					},
					{
						dst:         52,
						src:         50,
						rangeLength: 48,
					},
				},
			},
			seed: 99,
			want: 51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLocation(tt.mappers, tt.seed); got != tt.want {
				t.Errorf("getLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
