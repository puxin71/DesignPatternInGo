package headFirstPatterns_test

import (
	"log"
	"math"
	"testing"

	"github.com/puxin71/DesignPatternInGo/headFirstPatterns"
	"github.com/stretchr/testify/require"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	diff := math.Abs(a - b)
	log.Printf("%e", diff)
	return diff <= float64EqualityThreshold
}

func TestDarkRoastWithTwoWhips(t *testing.T) {
	darkRoast := headFirstPatterns.NewComponent("darkRoast", float64(4.0))
	darkRoast = headFirstPatterns.NewWhipDecorator(darkRoast)
	darkRoast = headFirstPatterns.NewWhipDecorator(darkRoast)
	require.Truef(t, almostEqual(float64(4.4), float64(darkRoast.Cost())), "error message: expect: %f, actual: %f", float64(4.4), darkRoast.Cost())
}
