package main

import "testing"

func TestNumTrees(t *testing.T) {
	testCases := []struct {
		request int
		want    int
		name    string
	}{
		{1, 1, "0"},
		{0, 1, "1"},
		{2, 2, "2"},
		{3, 5, "3"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NumTrees(tc.request)
			if got != tc.want {
				t.Errorf("NumTrees(%d) = %d; want %d", tc.request, got, tc.want)
			}
		})
	}
}
