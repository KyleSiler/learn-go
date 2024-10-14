package main

import (
	"fmt"
	"math/rand"
)

func main() {
	total := 100
	ap := rand.Intn(total)
	an := total - ap
	tp := rand.Intn(ap)
	fn := ap - tp
	tn := rand.Intn(an)
	fp := an - tn
	pp := tp + fp
	pn := fn + tn

	fmt.Printf("t-%d ap-%d an-%d\n", total, ap, an)
	fmt.Printf("pp-%d tp-%d fp-%d\n", pp, tp, fp)
	fmt.Printf("pn-%d fn-%d tn-%d", pn, fn, tn)

	// TODO: Need to randomly generate a question, prompt user for the answer, check answer
}
