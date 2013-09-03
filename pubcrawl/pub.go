package pubcrawl

import (
	"math"
	"strconv"
	"strings"
)

const LATDEG_KM float64 = 110.54  //km
const LONDEG_KM float64 = 113.320 //km

type Pub struct {
	DisplayName string `json:"display_name"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
}

func (self *Pub) Name() string {
	address := strings.Split(self.DisplayName, ",")
	return address[0]
}

func (self *Pub) LatFloat() float64 {
	f, _ := strconv.ParseFloat(self.Lat, 64)
	return f
}
func (self *Pub) LonFloat() float64 {
	f, _ := strconv.ParseFloat(self.Lon, 64)
	return f
}

func (self *Pub) DistanceTo(other *Pub) float64 {
	avgLat := (self.LatFloat() + other.LatFloat()) / 2.0
	dx := math.Abs(self.LatFloat() - other.LatFloat())
	dx = dx * LATDEG_KM
	dy := math.Abs(self.LonFloat() - other.LonFloat())
	dy = dy * (LONDEG_KM * math.Cos(avgLat))
	z := math.Sqrt((dx * dx) + (dy * dy))
	return z
}
