package simpleDirWalker

import (
	"testing"
)

// Testing the local application launch directory
func TestSDW(t *testing.T) {
	t.Run("test root", func(t *testing.T) {
		SDW(".")
	})
}
