package remote_package

import (
	"testing"

	cm "github.com/easierway/concurrent_map" // this is how to make an alias for a package
)

// just a demo about remote package
func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
