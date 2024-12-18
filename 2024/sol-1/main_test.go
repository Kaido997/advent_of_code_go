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
			want:  11,
		},
		{
			name:  "actual",
			input: input,
			want:  190000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1 %v want %v", got, tt.want)
			}
		})
	}

}
func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  31,
		},
		{
			name:  "actual",
			input: input,
			want:  22014209,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2 %v want %v", got, tt.want)
			}
		})
	}

}
