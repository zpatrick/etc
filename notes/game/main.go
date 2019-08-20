package main

type Adjucator func(a action.Action) error

type Action struct   {
	Type ActionType
}

func HealthAbove(r health.Reader, limit float64) Adjucator {
	return Adjucator(func(a Action) error {
		health, err := r.Read()
		if err != nil {
			return err
		}

		return health.Value() > limit
	})
}

func () {
	a := &actor.Actor{
		Name: "Zack",
		Adjucators: []Adjucator{
			HealthAbove(
		},
	}

	a.LoadAdjucators(
		HealthAbove(a, 0, func(a Action) error { if actor.Health()},
	)

	a.Re
}
