package services

import (
	"testing"
)

func TestLookingForMidpoint(t *testing.T) {
	cases := []struct {
		in  []GPS
		out GPS
	}{
		{[]GPS{GPS{25.22, 121.25}, GPS{25.22, 121.25}}, GPS{25.22, 121.25}},
	}
	for _, c := range cases {
		result := LookingForMidpoint(c.in)
		if result.Lon != c.out.Lon || result.Lat != c.out.Lat {
			t.Errorf("Wrong %+v\n", result)
		}
	}
}
