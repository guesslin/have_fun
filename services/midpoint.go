package service

func LookingForMidpoint( position1 GPS, position2 GPS ) GPS {
	var Position GPS
	Position.Lat = ( position1.Lat + position2.Lat ) / 2
	Position.Lon = ( position1.Lon + position2.Lon ) / 2
	return Position
}