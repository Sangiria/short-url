package decoder

import (
	"github.com/teris-io/shortid"
)

func DecodeURL() string {
	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
	result, _ := sid.Generate()
	return result
}