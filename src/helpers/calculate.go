package helpers

import "math"

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371 // Radius Bumi dalam kilometer

	// Konversi derajat ke radian
	lat1Rad := lat1 * (math.Pi / 180)
	lon1Rad := lon1 * (math.Pi / 180)
	lat2Rad := lat2 * (math.Pi / 180)
	lon2Rad := lon2 * (math.Pi / 180)

	// Hitung perbedaan
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	// Rumus Haversine
	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance
}
