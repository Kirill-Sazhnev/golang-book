package main

import (
	"testing"
)

func TestFormatDuration(t *testing.T) {
	type args struct {
		seconds int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "initial",
			args: args{478823},
			want: "5 days, 13 hours and 23 seconds",
		}, {
			name: "second",
			args: args{5091},
			want: "something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDuration(tt.args.seconds); got != tt.want {
				t.Errorf("FormatDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringTime(t *testing.T) {
	type args struct {
		frmt string
		res  string
		time map[string]int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringTime(tt.args.frmt, tt.args.res, tt.args.time); got != tt.want {
				t.Errorf("stringTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculate(t *testing.T) {
	type args struct {
		ttlSec  int64
		formSec int64
		frmt    string
		time    map[string]int64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calculate(tt.args.ttlSec, tt.args.formSec, tt.args.frmt, tt.args.time)
		})
	}
}
