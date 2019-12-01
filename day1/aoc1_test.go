package main

import "testing"

func Test_calculateFuel(t *testing.T) {
	type args struct {
		mass float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test successful", args{1969}, 654},
		{"Test 2", args{100756}, 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateFuel(tt.args.mass); got != tt.want {
				t.Errorf("calculateFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
