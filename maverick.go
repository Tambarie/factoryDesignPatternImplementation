package main

type Maverick struct {
	Gun
}

func NewMaverick() *Gun {
	return &Gun{
		name:  "Maverick gun",
		power: 7,
	}
}
