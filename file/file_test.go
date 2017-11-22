package file

import "testing"

func TestTreeDir(t *testing.T) {
	t.Logf("%+v", treeDir("/Users/ezbuy/Projects/ezbuy/goflow/src/github.com/JodeZer/lazydog", -1))
	t.Logf("%+v", treeDir("/Users/ezbuy/Projects/ezbuy/goflow/src/github.com/JodeZer/lazydog", 2))
}

func TestListFile(t *testing.T) {
	t.Logf("%+v", listGofileByPaths(treeDir("/Users/ezbuy/Projects/ezbuy/goflow/src/github.com/JodeZer/lazydog", -1)))
}
