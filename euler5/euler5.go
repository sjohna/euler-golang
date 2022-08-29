package euler5

import (
	"euler/utilities"
	"fmt"
)

func main() {
	lcm := 1

	for i := 2; i <= 20; i++ {
		lcm = utilities.LCM(lcm, i)
	}

	fmt.Println(lcm)
}
