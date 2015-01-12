package services

import "testing"

func TestAddress2GPS(t *testing.T) {
	cases := []struct {
		addr []rune
		// lat, lon float64
		result GPS
	}{
		{[]rune("台北市內湖區東湖路119巷49弄28號"), GPS{25.0700363, 121.6165096}},
		// 25.0700363, 121.6165096
	}
	for _, c := range cases {
		r := Address2GPS(c.addr)
		if r.Lat != c.result.Lat {
			t.Errorf("%s lat == %f, want %f\n", string(c.addr), r.Lat, c.result.Lat)
		}
		if r.Lon != c.result.Lon {
			t.Errorf("%s lon == %f, want %f\n", string(c.addr), r.Lon, c.result.Lon)
		}
	}
}
