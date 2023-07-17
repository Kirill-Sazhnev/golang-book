package main

import "testing"

func TestNextBigger(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextBigger(tt.args.n); got != tt.want {
				t.Errorf("NextBigger() = %v, want %v", got, tt.want)
			}
		})
	}
}
