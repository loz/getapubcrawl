package pubcrawl

import (
//"fmt"
)

var exampleJSON = `
[{"place_id":"1556315","licence":"Data \u00a9 OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright","osm_type":"node","osm_id":"312743305","boundingbox":["52.478443145752","52.4784469604492","-1.89822113513947","-1.89822101593018"],"lat":"52.4784465","lon":"-1.8982211","display_name":"The Newt, Stephenson Street, Gun Quarter, Birmingham, West Midlands, England, B2 4BL, United Kingdom, European Union","class":"amenity","type":"pub","importance":0.701,"icon":"http:\/\/nominatim.openstreetmap.org\/images\/mapicons\/food_pub.p.20.png","address":{"pub":"The Newt","road":"Stephenson Street","suburb":"Gun Quarter","city":"Birmingham","county":"West Midlands","state_district":"West Midlands","state":"England","postcode":"B2 4BL","country":"United Kingdom","country_code":"gb","continent":"European Union"}},{"place_id":"18350883","licence":"Data \u00a9 OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright","osm_type":"node","osm_id":"1721656777","boundingbox":["52.4749412536621","52.4749450683594","-1.89787173271179","-1.8978716135025"],"lat":"52.4749433","lon":"-1.8978717","display_name":"XXL's, Hurst Street, Gun Quarter, Birmingham, West Midlands, England, B5 4HQ, United Kingdom, European Union","class":"amenity","type":"bar","importance":0.601,"icon":"http:\/\/nominatim.openstreetmap.org\/images\/mapicons\/food_bar.p.20.png","address":{"bar":"XXL's","road":"Hurst Street","suburb":"Gun Quarter","city":"Birmingham","county":"West Midlands","state_district":"West Midlands","state":"England","postcode":"B5 4HQ","country":"United Kingdom","country_code":"gb","continent":"European Union"}},{"place_id":"18402736","licence":"Data \u00a9 OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright","osm_type":"node","osm_id":"1699331656","boundingbox":["52.4752616882324","52.4752655029297","-1.89712476730347","-1.89712464809418"],"lat":"52.4752642","lon":"-1.8971247","display_name":"Bambu, Wrottesley Street, Gun Quarter, Birmingham, West Midlands, England, B5 4HA, United Kingdom, European Union","class":"amenity","type":"bar","importance":0.601,"icon":"http:\/\/nominatim.openstreetmap.org\/images\/mapicons\/food_bar.p.20.png","address":{"bar":"Bambu","road":"Wrottesley Street","suburb":"Gun Quarter","city":"Birmingham","county":"West Midlands","state_district":"West Midlands","state":"England","postcode":"B5 4HA","country":"United Kingdom","country_code":"gb","continent":"European Union"}}]
`
var pubParser PubParser = PubParser{exampleJSON}

func (t *testSuite) TestReturnsRecordForEachParsed() {
	pubs := pubParser.Parse()
	t.Equal(3, len(pubs))
}

func (t *testSuite) TestSetsRecordAttributes() {
	pubs := pubParser.Parse()
	pub := pubs[0]
	t.Equal(pub.Name(), "The Newt")
	t.Equal(pub.Lat, "52.4784465")
	t.Equal(pub.Lon, "-1.8982211")
}

/*
func (t *testSuite) TestFinding() {
  pubs := FindPubs("Floodgate Street")
  for _, pub := range pubs {
    fmt.Println(pub)
  }
}
*/
