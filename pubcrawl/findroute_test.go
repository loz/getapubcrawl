package pubcrawl

func (t *testSuite) TestFinderEmptyWorks() {
	empty := make([]*Pub, 0)
	route, cost := FindRoute(empty)
	t.Equal(len(route), 0)
	t.Equal(cost, 0.0)
}

func (t *testSuite) TestFinderOneCostNothing() {
	one := make([]*Pub, 0)
	one = append(one, &Pub{})
	route, cost := FindRoute(one)
	t.Equal(route[0], one[0])
	t.Equal(cost, 0.0)
}

func (t *testSuite) TestFinderTwoAsGivenWithDistance() {
	pub1 := &Pub{Lat: "55.001", Lon: "-1.000"}
	pub2 := &Pub{Lat: "55.002", Lon: "-1.001"}
	pubs := []*Pub{pub1, pub2}

	route, cost := FindRoute(pubs)
	t.Equal(route[0], pub1)
	t.Equal(route[1], pub2)
	t.Equal(cost, pub1.DistanceTo(pub2))
}

func (t *testSuite) TestFinderPicksShortedRoute() {
	pub1 := &Pub{Lat: "55.001", Lon: "-1.001"}
	pub2 := &Pub{Lat: "55.002", Lon: "-1.002"}
	pub3 := &Pub{Lat: "55.003", Lon: "-1.003"}

	pubs := []*Pub{pub1, pub3, pub2}

	route, cost := FindRoute(pubs)
	t.Equal(route[0], pub1)
	t.Equal(route[1], pub2)
	t.Equal(route[2], pub3)

	expectedDistance := pub1.DistanceTo(pub2) +
		pub2.DistanceTo(pub3)
	t.Equal(cost, expectedDistance)
}

func (t *testSuite) TestFindNearest() {
	pub0 := &Pub{Lat: "55.000", Lon: "-1.000"}
	pub1 := &Pub{Lat: "55.001", Lon: "-1.001"}
	pub2 := &Pub{Lat: "55.002", Lon: "-1.002"}
	pub3 := &Pub{Lat: "55.003", Lon: "-1.003"}

	pubs := []*Pub{pub1, pub3, pub2}

	nearest, rest := FindNearest(pub0, pubs)
	t.Equal(nearest, pub1)
	t.Equal(rest[0], pub3)
	t.Equal(rest[1], pub2)
}
