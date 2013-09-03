package pubcrawl

import (
	"encoding/json"
	//"fmt"
)

type PubParser struct {
	source string
}

func (self *PubParser) Parse() []*Pub {
	var data []*Pub
	json.Unmarshal([]byte(self.source), &data)
	return data
}
