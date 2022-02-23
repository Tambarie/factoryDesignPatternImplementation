package main

type ak47 struct {
	Gun
}

func newAK47() *Gun {
	return &Gun{
		name:  "AK47 Gun",
		power: 4,
	}
}
