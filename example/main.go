package main

import (
	"log"
	"sort"

	d "github.com/dosko64/distance"
)

func main() {
	pp := []d.Point{}
	pp = append(pp, d.New(4, 4))
	pp = append(pp, d.New(3, 3))
	pp = append(pp, d.New(4, 2))

	p := d.New(1.5, 1.5)
	log.Println(pp)
	pp = sortByDistance(p, pp, true)
	log.Println(pp)
}

func sortByDistance(p d.Point, pp []d.Point, reverse bool) []d.Point {
	sort.Slice(pp, func(i, j int) bool {
		a := d.HaversineDistance(
			pp[i].Lat(),
			pp[i].Lng(),
			p.Lat(),
			p.Lng(),
		)
		b := d.HaversineDistance(
			pp[j].Lat(),
			pp[j].Lng(),
			p.Lat(),
			p.Lng(),
		)
		if reverse {
			return a > b
		}
		return a < b
	})
	return pp
}
