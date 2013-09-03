package pubcrawl

import (
  "net/http"
	"io/ioutil"
	"net/url"
)

func FindPubs(name string) []*Pub {
	const OSM string = "http://nominatim.openstreetmap.org/search?format=json&addressdetails=1&q="
	res, _ := http.Get(OSM + "pubs+near+" + url.QueryEscape(name))
	jsonBytes, _ := ioutil.ReadAll(res.Body)
	parser := PubParser{string(jsonBytes)}
	return parser.Parse()
}
