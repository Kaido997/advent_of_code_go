package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input.txt
var example []byte

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  161,
		},
		{
			name:  "actual",
			input: input,
			want:  187825547,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1 got %v want %v", got, tt.want)
			}
		})
	}

}

//go:embed test_input_2.txt
var example2 []byte

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			name:  "example",
			input: example2,
			want:  48,
		},
		{
			name:  "actual",
			input: input,
			want:  366,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("'part2': got %v want %v", got, tt.want)
			}
		})
	}

}
