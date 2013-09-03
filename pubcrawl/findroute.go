package pubcrawl

import (
	//"fmt"
	"math"
)

func FindRoute(remaining []*Pub) ([]*Pub, float64) {
	if len(remaining) < 2 {
		return remaining, 0.0
	}
	if len(remaining) == 2 {
		return remaining, remaining[0].DistanceTo(remaining[1])
	}

	//work out cost for each element to the best
	//route for the rest
	bestRoute := []*Pub{}
	bestCost := math.MaxFloat64

	for i, pub := range remaining {
		others := everythingExcept(remaining, i)
		thisRoute := []*Pub{pub}
		thisCost := 0.0
		curPub := pub
		for len(others) > 0 {
			next, rest := FindNearest(curPub, others)
			thisRoute = append(thisRoute, next)
			thisCost = thisCost + curPub.DistanceTo(next)
			others = rest
			curPub = next
		}
		if thisCost < bestCost {
			bestRoute = thisRoute
			bestCost = thisCost
		}
	}
	return bestRoute, bestCost
}

func FindNearest(pub *Pub, pubs []*Pub) (*Pub, []*Pub) {
	nearest := pubs[0]
	distance := pub.DistanceTo(nearest)
	rest := []*Pub{}
	options := pubs[1:]
	for _, option := range options {
		newdistance := pub.DistanceTo(option)
		if newdistance < distance {
			//append last nearest
			rest = append(rest, nearest)
			nearest = option
			distance = newdistance
		} else {
			rest = append(rest, option)
		}
	}
	return nearest, rest
}

func AppendReversed(pub *Pub, elements []*Pub) []*Pub {
	items := []*Pub{pub}
	for i := len(elements) - 1; i >= 0; i-- {
		items = append(items, elements[i])
	}
	return items
}

func everythingExcept(elements []*Pub, index int) []*Pub {
	var others []*Pub
	for i, pub := range elements {
		if i == index {
			continue
		}
		others = append(others, pub)
	}
	return others
}
