package main

//divitiaegithub.com/barnex/cuda5 https://archive.fosdem.org/2014/schedule/event/hpc_devroom_go/
import (
	"fmt"

	D "github.com/nkozyra/domino"
)

func main() {
	d := D.NewDomino()
	d.SetModel(D.NewKNN())

	fmt.Println(d.Details)
}
