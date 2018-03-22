package agent

import "math/rand"

// DiscreteAction foo
func DiscreteAction(a Agent, vals []float64) (act int, err error) {
	probs, err := a.Forward(vals)
	if err != nil {
		return 0, err
	}

	act = 0
	p := 0.0
	r := rand.Float64()
	for i := range probs {
		p += probs[i]
		if p > r {
			break
		}
		act++
	}

	return act, nil
}
