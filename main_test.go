package main

import (
	"os"
	"reflect"
	"testing"
)

func TestCountDroneDistance(t *testing.T) {
	var estate1 = [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}

	var estate2 = [][]int{
		{0, 0, 0},
		{0, 0, 5},
		{10, 0, 0},
	}

	var estate3 = [][]int{
		{1, 1},
		{1, 1},
	}
	var estate4 = [][]int{
		{0, 0, 0, 0, 0},
	}
	var estate5 = [][]int{
		{0, 0, 10},
		{0, 0, 10},
		{10, 0, 10},
	}
	var estate6 = [][]int{
		{0, 5, 10, 7, 0, 1},
	}

	type args struct {
		estate [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Empty 2d plot",
			args: args{
				estate1,
			},
			want: 52,
		},
		{
			name: "Empty 1d plot",
			args: args{
				estate4,
			},
			want: 42,
		},
		{
			name: "Some trees exist",
			args: args{estate2},
			want: 112,
		},
		{
			name: "All plot has trees",
			args: args{estate3},
			want: 34,
		},
		{
			name: "Plot has couple of same tree height",
			args: args{estate5},
			want: 142,
		},
		{
			name: "Plot has couple of different tree height",
			args: args{estate6},
			want: 74,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountDroneDistance(tt.args.estate); got != tt.want {
				t.Errorf("countDroneDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadInput(t *testing.T) {
	var estate = [][]int{{10, 20, 0, 0, 0}}
	tests := []struct {
		name  string
		input string
		want  *[][]int
		want1 bool
	}{
		// TODO: Add test cases.
		{
			name:  "Valid input",
			input: "5 1 2\n1 1 10\n2 1 20\n",
			want:  &estate, // Replace with the correct expected result
			want1: true,
		},
		{
			name:  "Invalid number of inputs",
			input: "10 10\n",
			want:  nil,
			want1: false,
		},
		{
			name:  "Invalid number of inputs 2",
			input: "10 10 4 4\n",
			want:  nil,
			want1: false,
		},
		{
			name:  "Invalid width",
			input: "0 10 2\n",
			want:  nil,
			want1: false,
		},
		{
			name:  "Invalid length",
			input: "10 0 2\n",
			want:  nil,
			want1: false,
		},
		{
			name:  "Invalid count",
			input: "10 10 0\n",
			want:  nil,
			want1: false,
		},
		{
			name:  "Invalid tree x",
			input: "10 10 1\n11 1 10",
			want:  nil,
			want1: false,
		},
		{
			name:  "Invalid tree y",
			input: "10 10 1\n1 11 10",
			want:  nil,
			want1: false,
		},
		{
			name:  "Invalid tree height",
			input: "10 10 1\n1 1 40",
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a pipe to simulate stdin
			r, w, _ := os.Pipe()
			origStdin := os.Stdin
			defer func() { os.Stdin = origStdin }()
			os.Stdin = r

			// Write the test input to the pipe
			go func() {
				defer w.Close()
				w.Write([]byte(tt.input))
			}()

			got, got1 := ReadInput()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadInput() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ReadInput() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
