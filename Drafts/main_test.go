package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestHeight(t *testing.T) {
	type args struct {
		n *big.Int
		m *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Height(tt.args.n, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Height() = %v, want %v", got, tt.want)
			}
		})
	}
}
