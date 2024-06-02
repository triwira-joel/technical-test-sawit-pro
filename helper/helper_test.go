package helper

import (
	"reflect"
	"testing"
)

func TestAbs(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Minus number should return positive",
			args: args{-57},
			want: 57,
		},
		{
			name: "Positive should return the same",
			args: args{10},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.a); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateEstate(t *testing.T) {
	type args struct {
		width  int
		length int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "Create estate should return as intended",
			args: args{3, 3},
			want: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			name: "Create estate should return as intended 2",
			args: args{3, 4},
			want: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateEstate(tt.args.width, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEstate() = %v, want %v", got, tt.want)
			}
		})
	}
}
