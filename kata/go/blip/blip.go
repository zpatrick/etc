package main

type Blip struct {
	on bool
}

func (b Blip) String() string {
	if b.on {
		return "*"
	}

	return " "
}
