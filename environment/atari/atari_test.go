package atari_test

import (
	"testing"

	"github.com/cshenton/evolution/environment"
	"github.com/cshenton/evolution/environment/atari"
)

func TestAtariEnvironment(t *testing.T) {
	// This will force a compilation failure if *atari.Env doesn't implement environment.Environment
	_ = environment.Environment(&atari.Env{})
}
