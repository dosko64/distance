package distance

import (
	"math"
)

var r float64 = 6378100

// HaversineDistance Returns the haversine distance in meters between two points
// https://en.wikipedia.org/wiki/Haversine_formula
func HaversineDistance(lat1, lng1, lat2, lng2 float64) float64 {
	var la1, lo1, la2, lo2 float64
	la1 = rad(lat1)
	lo1 = rad(lng1)
	la2 = rad(lat2)
	lo2 = rad(lng2)
	h := haversine(la2-la1) + math.Cos(la1)*math.Cos(la2)*haversine(lo2-lo1)
	return 2 * r * math.Asin(math.Sqrt(h))
}

// haversine Return the haversine of a given theta
// https://en.wikipedia.org/wiki/Haversine_formula
func haversine(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// rad Converts degrees to radians
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// Square returns two points, that represent a square box with center the given lat, lng and d*2 as length of side, d is in meters
func Square(lat, lng, d float64) (Point, Point) {
	maxLat := lat + (d/r)*(180/math.Pi)
	minLat := lat - (d/r)*(180/math.Pi)

	maxLng := lng + (d/r)*(180/math.Pi)/math.Cos(lat*math.Pi/180)
	minLng := lng - (d/r)*(180/math.Pi)/math.Cos(lat*math.Pi/180)
	return Point{maxLat, maxLng}, Point{minLat, minLng}
}

// Point represents a Lat, Lng point
type Point struct {
	lat, lng float64
}

// New return a new point
func New(lat, lng float64) Point {
	return Point{lat, lng}
}

// DistanceFrom returns the distance between p with a given pp Point
func (p Point) DistanceFrom(pp Point) float64 {
	return HaversineDistance(p.Lat(), p.Lng(), pp.Lat(), pp.Lng())
}

// Square return two points that represent a square area with center the p Poing and side length of d*2
func (p Point) Square(d float64) (Point, Point) {
	return Square(p.lat, p.lng, d)
}

// Lat returns point's latitude
func (p Point) Lat() float64 {
	return p.lat
}

// Lng returns point's longitude
func (p Point) Lng() float64 {
	return p.lng
}
