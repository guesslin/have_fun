package services

import "testing"

func TestAddress2GPS(t *testing.T) {
	cases := []struct {
		addr     []rune
		lat, lon float64
	}{
		{[]rune("台北市內湖區東湖路119巷49弄28號"), 25.0700363, 121.6165096},
		// 25.0700363, 121.6165096
	}
	for _, c := range cases {
		lat, lon := Address2GPS(c.addr)
		if lat != c.lat {
			t.Errorf("%s lat == %f, want %f\n", string(c.addr), lat, c.lat)
		}
		if lon != c.lon {
			t.Errorf("%s lon == %f, want %f\n", string(c.addr), lon, c.lon)
		}
	}
}
