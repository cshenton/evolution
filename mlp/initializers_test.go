package mlp_test

import (
	"testing"

	"github.com/cshenton/openai-evolution/mlp"
)

func TestZeros(t *testing.T) {
	tt := []int{1, 10, 100}
	for n := range tt {
		t.Run("length "+string(n), func(t *testing.T) {
			res := mlp.Zeros(n)
			exp := make([]float64, n)
			if len(res) != len(exp) {
				t.Fatalf("expected slice length %v, but it was %v", len(exp), len(res))
			}
			for i := range res {
				if res[i] != exp[i] {
					t.Errorf(
						"expected result %v at position %v, but got %v",
						exp[i], i, res[i],
					)
				}
			}
		})
	}
}
