package services

import (
	"testing"
)

func TestFindPlaces(t *testing.T) {
	cases := []struct {
		in     GPS
		radius int32
		out    string
	}{
		{GPS{25.0700363, 121.6165096}, 100, "佳佳牛排"},
	}
	for _, c := range cases {
		results := FindPlaces(c.in, c.radius)
		count := 0
		for _, r := range results {
			if r.Name == c.out {
				count += 1
			}
		}
		if count == 0 {
			t.Errorf("Should find %s but get nothing\n")
		}
	}
}
