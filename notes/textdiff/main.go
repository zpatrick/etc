package main

import (
	"fmt"
	"math"
)

const (
	BodyLengthDeltaThreshold = 0.3
	BodyRuneDeltaThreshold   = 0.5
)

const (
	body1 = `
<html> 
  <head> 
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"> 
  </head> 
  <body> 
    <html> 
      <body>
        <div>
          <p>Hello, World!</p>
        </div>
      </body>
    </html>
  </body>
</html>
`

	body2 = `
<html> 
  <body>
    <div>
      <p>Hello, World!</p>
    </div>
  </body>
</html>
`

	body3 = `
<html> 
  <head> 
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"> 
  </head> 
  <body> 
    <html> 
      <body>
        <div>
          <p>Hello%2C World%21</p>
        </div>
      </body>
    </html>
  </body>
</html>
`

	body4 = `
<iunm>
  <cpez>
    <ejw>
      <q>Ifmmp, Xpsme!</q>
    </ejw>
  </cpez>
</iunm>
`
)

type RuneIndex struct {
	Count int
	Delta int
}

func bodyLengthDelta(body1, body2 string) float64 {
	switch lb1, lb2 := len(body1), len(body2); true {
	case lb1 > lb2:
		return 1.0 - float64(lb2)/float64(lb1)
	case lb2 > lb1:
		return 1.0 - float64(lb1)/float64(lb2)
	default:
		return 0.0
	}
}

func bodyRuneDelta(body1, body2 string) float64 {
	runeDeltas := map[rune]int{}
	for _, c := range body1 {
		runeDeltas[c] += 1
	}

	for _, c := range body2 {
		runeDeltas[c] -= 1
	}

	var aggregateRuneDeltaCount float64
	for _, d := range runeDeltas {
		aggregateRuneDeltaCount += math.Abs(float64(d))
	}

	averageBodyLength := float64(len(body1)+len(body2)) / 2.0
	return aggregateRuneDeltaCount / averageBodyLength
}

func diffBodies(body1, body2 string) bool {
	if bld := bodyLengthDelta(body1, body2); bld > BodyLengthDeltaThreshold {
		fmt.Printf("Body Length Delta: %v\n", bld)
		return true
	}

	if brd := bodyRuneDelta(body1, body2); brd > BodyRuneDeltaThreshold {
		fmt.Printf("Body Rune Delta: %v\n", brd)
		return true
	}

	return false
}

func main() {
	fmt.Println(diffBodies(body1, body2))
	fmt.Println(diffBodies(body1, body3))
	fmt.Println(diffBodies(body2, body4))
}
