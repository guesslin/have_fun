package services

func LookingForMidpoint(position []GPS) (midpoint GPS) {
	size := float64(len(position))
	for _, p := range position {
		midpoint.Lat += p.Lat
		midpoint.Lon += p.Lon
	}
	midpoint.Lat = midpoint.Lat / size
	midpoint.Lon = midpoint.Lon / size
	return
}
