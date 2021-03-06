package tree

import "github.com/topport/magic/pkg/structs/cappedarr"

func (t *Tree) minValue(maxNodes int) uint64 {
	c := cappedarr.New(maxNodes)
	t.iterateWithTotal(func(total uint64) bool {
		return c.Push(total)
	})
	return c.MinValue()
}
