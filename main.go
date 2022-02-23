package main

import "fmt"

func PrintDetails(g Gun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()

	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}

func main() {
	ak47, _ := getGun("ak47")
	PrintDetails(*ak47)
	maverick, _ := getGun("maverick")
	PrintDetails(*maverick)
}
