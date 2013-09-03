package pubcrawl

import (
//"fmt"
)

func (t *testSuite) TestDistanceToPub() {
	pub1 := &Pub{Lat: "55.001", Lon: "-1.000"}
	pub2 := &Pub{Lat: "55.002", Lon: "-1.001"}
	t.Equal(pub1.DistanceTo(pub2), 0.11057241856366766)
	t.Equal(pub2.DistanceTo(pub1), 0.11057241856366766)
}
