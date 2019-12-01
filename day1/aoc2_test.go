package main

import (
	"reflect"
	"testing"
)

func Test_reduceFuel(t *testing.T) {
	var r []float64
	type args struct {
		mass float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"successful test", args{1969}, []float64{654, 216, 70, 21, 5}},
		{"another test", args{100756}, []float64{33583, 11192, 3728, 1240, 411, 135, 43, 12, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reduceFuel(tt.args.mass, r); !reflect.DeepEqual(got[:len(got)-1], tt.want) {
				t.Errorf("reduceFuel() = %v, want %v", got[:len(got)-1], tt.want)
			}
		})
	}
}

func Test_reduceFuelSum(t *testing.T) {
	var r []float64
	type args struct {
		mass float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"successful test", args{1969}, 966},
		{"another test", args{100756}, 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(reduceFuel(tt.args.mass, r)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reduceFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_reduceFuelList(t *testing.T) {
// 	var r []float64
// 	type args []float64
// 	tests := []struct {
// 		name string
// 		args args
// 		want float64
// 	}{
// 		{"successful test", args{1969, 100756}, 966 + 50346},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := sum(reduceFuel(tt.args.mass, r)); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("reduceFuel() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
